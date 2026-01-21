package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ==============================
// CONFIGURATION
// ==============================
const (
	BucketName = "your-gcs-bucket-name"
	MountPoint = "/mnt/gcs-bucket"
)

// ==============================
// HELPERS
// ==============================
func run(cmd string, check bool) error {
	fmt.Println("▶", cmd)
	command := exec.Command("bash", "-c", cmd)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil && check {
		return err
	}
	return nil
}

func requireRoot() {
	if os.Geteuid() != 0 {
		fmt.Println("❌ This program must be run as root (use sudo).")
		os.Exit(1)
	}
}

func detectOS() (string, string) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		fmt.Println("❌ Unable to detect OS")
		os.Exit(1)
	}
	defer file.Close()

	osInfo := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			osInfo[parts[0]] = strings.Trim(parts[1], `"`)
		}
	}

	return strings.ToLower(osInfo["ID"]), osInfo["VERSION_ID"]
}

// ==============================
// INSTALLERS
// ==============================
func installGCSFuseDebian() error {
	if err := run("apt update", true); err != nil {
		return err
	}
	if err := run("apt install -y curl gnupg lsb-release", true); err != nil {
		return err
	}

	out, err := exec.Command("lsb_release", "-c", "-s").Output()
	if err != nil {
		return err
	}
	distro := strings.TrimSpace(string(out))

	if err := run(fmt.Sprintf(
		`echo "deb https://packages.cloud.google.com/apt gcsfuse-%s main" > /etc/apt/sources.list.d/gcsfuse.list`,
		distro,
	), true); err != nil {
		return err
	}

	if err := run("curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -", true); err != nil {
		return err
	}

	if err := run("apt update", true); err != nil {
		return err
	}
	return run("apt install -y gcsfuse", true)
}

func installGCSFuseRHEL() error {
	if err := run("yum install -y curl gnupg", true); err != nil {
		return err
	}

	repo := `[gcsfuse]
name=gcsfuse
baseurl=https://packages.cloud.google.com/yum/repos/gcsfuse-el$releasever-$basearch
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
       https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
`
	if err := os.WriteFile("/etc/yum.repos.d/gcsfuse.repo", []byte(repo), 0644); err != nil {
		return err
	}

	return run("yum install -y gcsfuse", true)
}

// ==============================
// MAIN
// ==============================
func main() {
	requireRoot()

	osID, version := detectOS()
	fmt.Println("==============================")
	fmt.Printf(" Detected OS: %s (Version: %s)\n", osID, version)
	fmt.Println("==============================")

	fmt.Println(" Installing gcsfuse")
	fmt.Println("==============================")

	switch osID {
	case "ubuntu", "debian":
		if err := installGCSFuseDebian(); err != nil {
			panic(err)
		}
	case "rhel", "centos", "fedora":
		if err := installGCSFuseRHEL(); err != nil {
			panic(err)
		}
	default:
		fmt.Println("❌ Unsupported OS:", osID)
		os.Exit(1)
	}

	fmt.Println("==============================")
	fmt.Println(" Configuring fuse")
	fmt.Println("==============================")
	run("sed -i 's/#user_allow_other/user_allow_other/' /etc/fuse.conf", true)

	fmt.Println("==============================")
	fmt.Println(" Creating mount directory")
	fmt.Println("==============================")
	os.MkdirAll(MountPoint, 0755)
	run(fmt.Sprintf("chown root:root %s", MountPoint), true)

	fmt.Println("==============================")
	fmt.Println(" Mounting GCS bucket")
	fmt.Println("==============================")

	run(fmt.Sprintf(
		`gcsfuse --implicit-dirs --uid=0 --gid=0 --file-mode=777 --dir-mode=777 -o allow_other %s %s`,
		BucketName,
		MountPoint,
	), true)

	fmt.Println("==============================")
	fmt.Println(" Persisting mount in /etc/fstab")
	fmt.Println("==============================")

	fstabEntry := fmt.Sprintf(
		"%s %s gcsfuse rw,implicit_dirs,uid=0,gid=0,file_mode=777,dir_mode=777,allow_other,_netdev 0 0",
		BucketName,
		MountPoint,
	)

	data, _ := os.ReadFile("/etc/fstab")
	if !strings.Contains(string(data), fstabEntry) {
		f, _ := os.OpenFile("/etc/fstab", os.O_APPEND|os.O_WRONLY, 0644)
		defer f.Close()
		f.WriteString("\n" + fstabEntry + "\n")
	}

	fmt.Println("==============================")
	fmt.Println(" Verifying mount")
	fmt.Println("==============================")
	run("df -h | grep "+MountPoint, false)
	run("mount | grep gcsfuse", false)

	fmt.Println("✅ GCS bucket successfully mounted at", MountPoint)
}

