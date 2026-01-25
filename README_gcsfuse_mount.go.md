# gcsfuse_mount.go

**Go program** for high-performance GCS bucket mounting on Linux (Debian/Ubuntu/RHEL/CentOS/Fedora).

## Requirements
- Go 1.16+ (for building only)
- sudo access
- Valid GCS credentials

## Build
```bash
# Build executable
go build -o gcsfuse_mount gcsfuse_mount.go

# Cross-compile (ARM64, 386)
GOOS=linux GOARCH=arm64 go build -o gcsfuse_mount gcsfuse_mount.go
```

## Configuration
```go
const (
  BucketName = "your-bucket"
  MountPoint = "/mnt/gcs-bucket"
)
```

## Usage
```bash
sudo ./gcsfuse_mount
# Or direct:
sudo go run gcsfuse_mount.go
```

**Automates:**
1. Verify sudo privileges
2. Detect OS (apt vs yum)
3. Install curl, gnupg
4. Add gcsfuse repo
5. Install gcsfuse
6. Enable FUSE config
7. Create mount directory
8. Mount GCS bucket
9. Add to /etc/fstab
10. Verify mount

**Advantages:**
- Single compiled executable (no runtime dependencies)
- Fast execution, production-ready
- Cross-platform compilation support
- Best for Kubernetes/containerization

**Support:** [GitHub](https://github.com/mpandey95) | [LinkedIn](https://linkedin.com/in/manish-pandey95)

```go
const (
	BucketName = "your-gcs-bucket-name"    // Your GCS bucket name
	MountPoint = "/mnt/gcs-bucket"         // Local mount directory
)
```

Then rebuild:

```bash
go build -o gcsfuse_mount gcsfuse_mount.go
```

## Program Structure

### Key Functions

- **`requireRoot()`** - Validates that the program is running with root privileges
- **`detectOS()`** - Detects the Linux distribution by reading `/etc/os-release`
- **`run(cmd string, check bool)`** - Executes shell commands with optional error checking
- **`installGCSFuseDebian()`** - Handles installation for Debian/Ubuntu systems
- **`installGCSFuseRHEL()`** - Handles installation for RHEL/CentOS/Fedora systems
- **`main()`** - Orchestrates the entire mounting process

### Error Handling

The program uses Go's error handling patterns:
- Returns errors from function calls
- Uses `panic()` for fatal errors
- Validates root privileges at startup
- Checks for unsupported OS and exits gracefully

## Output

The program provides clear progress indicators:
- ▶ prefix for each command being executed
- ❌ error indicator if running without root privileges
- Section headers showing progress through installation steps
- OS and version detection at startup
- Success message: ✅ GCS bucket successfully mounted at /mnt/gcs-bucket

## OS-Specific Installation

### Ubuntu/Debian
Uses `apt` package manager to install:
- curl, gnupg, lsb-release
- Adds Google Cloud apt repository
- Installs gcsfuse

### RHEL/CentOS/Fedora
Uses `yum` package manager to install:
- curl, gnupg
- Adds Google Cloud yum repository configuration
- Installs gcsfuse

## Mount Options Used

- `--implicit-dirs`: Treat prefixes as directories
- `--uid=0, --gid=0`: Set ownership to root user and group
- `--file-mode=777, --dir-mode=777`: Set full read/write/execute permissions
- `-o allow_other`: Allow non-root users to access the mount

## Persistence

The mount configuration is automatically added to `/etc/fstab` with:
- `_netdev` flag for network device handling
- Full permissions and options for automatic remounting on reboot

## Advantages of Go Implementation

- **Single Binary**: No dependencies (except Go for building) - can be distributed as a single executable
- **Performance**: Compiled language, faster execution than scripts
- **Type Safety**: Go's strong type system catches errors at compile time
- **Cross-Platform Compilation**: Easy to build for different architectures
- **Error Handling**: Explicit error handling with Go's error pattern
- **Maintainability**: Structured code is easier to extend and maintain

## Environment Variables

The program uses hardcoded constants, but can be modified to accept environment variables:

```go
package main

import "os"

var (
	BucketName = getEnv("BUCKET_NAME", "your-gcs-bucket-name")
	MountPoint = getEnv("MOUNT_POINT", "/mnt/gcs-bucket")
)

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}
```

## Extending the Program

### Add Command-Line Flags

```go
import "flag"

var (
	bucketName = flag.String("bucket", "your-gcs-bucket-name", "GCS bucket name")
	mountPoint = flag.String("mount", "/mnt/gcs-bucket", "Mount directory")
)

func init() {
	flag.Parse()
}
```

### Add Logging

```go
import "log"

func main() {
	log.Println("Starting gcsfuse mount...")
	// Rest of the code
}
```

### Add Configuration File Support

```go
import "encoding/json"

type Config struct {
	BucketName string `json:"bucket_name"`
	MountPoint string `json:"mount_point"`
}
```

## Troubleshooting

### Permission Denied

```bash
# Verify running with sudo
sudo ./gcsfuse_mount

# Check FUSE configuration
sudo grep user_allow_other /etc/fuse.conf
```

### Unsupported OS Error

The program only supports: ubuntu, debian, rhel, centos, fedora. Check your OS:

```bash
cat /etc/os-release | grep "^ID="
```

### Build Errors

Ensure Go is installed:

```bash
go version
which go
```

If Go is not installed, download from [golang.org](https://golang.org/dl/)

### Mount Not Persisting After Reboot

```bash
# Verify fstab entry
grep gcsfuse /etc/fstab

# Test fstab configuration
sudo mount -a
```

## Uninstalling

```bash
# Unmount the bucket
sudo umount /mnt/gcs-bucket

# Remove from fstab (optional)
sudo sed -i '/your-bucket-name/d' /etc/fstab

# Uninstall gcsfuse (optional)
sudo apt remove gcsfuse  # Debian/Ubuntu
sudo yum remove gcsfuse   # RHEL/CentOS/Fedora
```

## Comparison with Other Implementations

| Feature | Bash | Python | Go | Ansible |
|---------|------|--------|----|---------| 
| **Build Required** | ❌ | ❌ | ✅ | ❌ |
| **Performance** | ⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐⭐ |
| **Type Safety** | ❌ | ⭐⭐ | ⭐⭐⭐⭐⭐ |
| **Single Binary** | ❌ | ❌ | ✅ | ❌ |
| **Cross-Compile** | ❌ | ⚠️ | ✅ | ❌ |
| **Dependencies** | None | Python 3 | None | Ansible |
| **Ease of Use** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ |

## Advanced Options

Customize the mount command in the program:

```go
run(fmt.Sprintf(
	`gcsfuse --max-conns-per-host=100 --implicit-dirs --uid=0 --gid=0 %s %s`,
	BucketName,
	MountPoint,
), true)
```

See [gcsfuse documentation](https://github.com/GoogleCloudPlatform/gcsfuse) for all options.

## Building for Docker/Container

```dockerfile
FROM golang:1.20-alpine as builder
WORKDIR /build
COPY gcsfuse_mount.go .
RUN go build -o gcsfuse_mount gcsfuse_mount.go

FROM ubuntu:22.04
COPY --from=builder /build/gcsfuse_mount /usr/local/bin/
RUN chmod +x /usr/local/bin/gcsfuse_mount
CMD ["sudo", "gcsfuse_mount"]
```

## Notes

- Must be run with sudo/root privileges
- Designed for Debian/Ubuntu and RHEL/CentOS/Fedora distributions
- Automatically detects the OS and uses appropriate package manager
- Mount will persist across system reboots
- Compiled Go binary can be distributed without Go runtime
- Works on bare metal, VMs, containers, and cloud instances
- High performance implementation suitable for production use
