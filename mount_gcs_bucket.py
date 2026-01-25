#!/usr/bin/env python3
import subprocess
import sys
import os
import platform

# ==============================
# CONFIGURATION
# ==============================
BUCKET_NAME = "your-gcs-bucket-name"
MOUNT_POINT = "/mnt/gcs-bucket"

# ==============================
# HELPERS
# ==============================
def run(cmd, check=True):
    print(f"▶ {cmd}")
    subprocess.run(cmd, shell=True, check=check)

def require_root():
    if os.geteuid() != 0:
        print("❌ This script must be run as root (use sudo).")
        sys.exit(1)

def detect_os():
    """Detect OS and distribution"""
    with open('/etc/os-release', 'r') as f:
        os_info = {}
        for line in f:
            key, value = line.rstrip().split("=", 1)
            os_info[key] = value.strip('"')
    
    os_id = os_info.get('ID', '').lower()
    version = os_info.get('VERSION_ID', '')
    
    return os_id, version

def install_gcsfuse_debian():
    """Install gcsfuse on Debian/Ubuntu"""
    run("apt update")
    run("apt install -y curl gnupg lsb-release")

    distro = subprocess.check_output(
        "lsb_release -c -s", shell=True
    ).decode().strip()

    run(
        f'echo "deb https://packages.cloud.google.com/apt gcsfuse-{distro} main" '
        f'> /etc/apt/sources.list.d/gcsfuse.list'
    )

    run("curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -")
    run("apt update")
    run("apt install -y gcsfuse")

def install_gcsfuse_rhel():
    """Install gcsfuse on RHEL/CentOS/Fedora"""
    run("yum install -y curl gnupg")

    # Create gcsfuse repository configuration
    gcsfuse_repo = """[gcsfuse]
name=gcsfuse
baseurl=https://packages.cloud.google.com/yum/repos/gcsfuse-el$releasever-$basearch
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
       https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
"""
    
    with open('/etc/yum.repos.d/gcsfuse.repo', 'w') as f:
        f.write(gcsfuse_repo)
    
    run("yum install -y gcsfuse")

# ==============================
# MAIN LOGIC
# ==============================
def main():
    require_root()

    # Detect OS
    os_id, version = detect_os()
    print("==============================")
    print(f" Detected OS: {os_id} (Version: {version})")
    print("==============================")

    print("==============================")
    print(" Installing gcsfuse")
    print("==============================")

    if os_id in ['ubuntu', 'debian']:
        install_gcsfuse_debian()
    elif os_id in ['rhel', 'centos', 'fedora']:
        install_gcsfuse_rhel()
    else:
        print(f"❌ Unsupported OS: {os_id}")
        print("Supported: ubuntu, debian, rhel, centos, fedora")
        sys.exit(1)

    print("==============================")
    print(" Configuring fuse")
    print("==============================")

    run("sed -i 's/#user_allow_other/user_allow_other/' /etc/fuse.conf")

    print("==============================")
    print(" Creating mount directory")
    print("==============================")

    os.makedirs(MOUNT_POINT, exist_ok=True)
    run(f"chown root:root {MOUNT_POINT}")
    run(f"chmod 755 {MOUNT_POINT}")

    print("==============================")
    print(" Mounting GCS bucket")
    print("==============================")

    run(
        f"gcsfuse "
        f"--implicit-dirs "
        f"--uid=0 "
        f"--gid=0 "
        f"--file-mode=777 "
        f"--dir-mode=777 "
        f"-o allow_other "
        f"{BUCKET_NAME} "
        f"{MOUNT_POINT}"
    )

    print("==============================")
    print(" Persisting mount in /etc/fstab")
    print("==============================")

    fstab_entry = (
        f"{BUCKET_NAME} {MOUNT_POINT} gcsfuse "
        f"rw,implicit_dirs,uid=0,gid=0,file_mode=777,dir_mode=777,"
        f"allow_other,_netdev 0 0"
    )

    with open("/etc/fstab", "r+") as fstab:
        if fstab_entry not in fstab.read():
            fstab.write("\n" + fstab_entry + "\n")

    print("==============================")
    print(" Verifying mount")
    print("==============================")

    run(f"df -h | grep {MOUNT_POINT}", check=False)
    run("mount | grep gcsfuse", check=False)

    print(f"✅ GCS bucket successfully mounted at {MOUNT_POINT}")

# ==============================
if __name__ == "__main__":
    main()
