# mount_gcs_bucket.py

A Python 3 script for installing and configuring gcsfuse to mount a Google Cloud Storage (GCS) bucket on Linux systems.

## Purpose

This script automates the entire process of:
- Installing gcsfuse and its dependencies
- Configuring FUSE filesystem permissions
- Creating a mount directory
- Mounting a GCS bucket
- Persisting the mount configuration in `/etc/fstab`
- Verifying the mount was successful

## Supported Distributions

The script supports both major Linux distribution families:

- **Debian/Ubuntu**: Ubuntu, Debian
- **RHEL/CentOS/Fedora**: RHEL 7+, CentOS 7+, Fedora

The script automatically detects the OS and uses the appropriate package manager (`apt` for Debian/Ubuntu, `yum` for RHEL).

## Requirements

- Python 3
- Root/sudo access
- Linux system (Debian/Ubuntu or RHEL/CentOS/Fedora based)
- Valid GCS bucket name
- Proper Google Cloud credentials configured on the system

## Usage

### Prerequisites

Before running the script, update the configuration variables at the top of the file:

```python
BUCKET_NAME = "your-gcs-bucket-name"
MOUNT_POINT = "/mnt/gcs-bucket"
```

### Running the Script

```bash
sudo python3 mount_gcs_bucket.py
```

The script requires root privileges and will:
1. Install curl, gnupg, and lsb-release packages
2. Add the official gcsfuse repository for your distribution
3. Install gcsfuse
4. Enable FUSE user_allow_other configuration
5. Create and set permissions on the mount directory
6. Mount the GCS bucket with appropriate permissions
7. Add the mount to `/etc/fstab` for persistence
8. Verify the successful mount

## Configuration

Edit these variables in the script to customize:

- `BUCKET_NAME`: The name of your GCS bucket to mount
- `MOUNT_POINT`: The local directory where the bucket will be mounted

## Output

The script provides visual feedback with:
- OS and version detection at startup
- ▶ prefix for each command being executed
- ❌ error indicator if running without root privileges
- Section headers showing progress through installation steps

## Implementation Details

### OS Detection
The script reads `/etc/os-release` to detect the operating system:
- Extracts OS ID and version information
- Supports: ubuntu, debian, rhel, centos, fedora

### Package Installation
- **Debian/Ubuntu**: Uses `apt` package manager with Google Cloud apt repository
- **RHEL/CentOS/Fedora**: Uses `yum` package manager with Google Cloud yum repository

### Error Handling
- Checks for root privileges before starting
- Exits on unsupported OS
- Fails fast on command errors

## Environment Variables

Override configuration via environment variables:

```bash
export BUCKET_NAME="my-bucket"
export MOUNT_POINT="/data/gcs"
sudo -E python3 mount_gcs_bucket.py
```

## Integration with Python Projects

```python
import subprocess
import sys

def mount_gcs_bucket(bucket_name, mount_point):
    result = subprocess.run([
        sys.executable, 'mount_gcs_bucket.py'
    ], env={'BUCKET_NAME': bucket_name, 'MOUNT_POINT': mount_point})
    return result.returncode == 0

if __name__ == "__main__":
    if mount_gcs_bucket("my-bucket", "/mnt/gcs"):
        print("Mounting successful!")
```

## Troubleshooting

### Python 3 Not Found

```bash
which python3
sudo python3 --version

# If not installed
sudo apt install python3  # Debian/Ubuntu
sudo yum install python3   # RHEL/CentOS
```

### Permission Denied

```bash
# Verify running with sudo
if [ "$EUID" -eq 0 ]; then
  echo "Running as root"
fi

# Check FUSE configuration
sudo grep user_allow_other /etc/fuse.conf
```

### Unsupported OS Error

The script only supports: ubuntu, debian, rhel, centos, fedora. Check your OS:

```bash
cat /etc/os-release | grep "^ID="
```

## Extending the Script

Add custom functions to the script:

```python
def custom_mount_options():
    """Add custom mount options"""
    return "--max-conns-per-host=100 --enable-nonseekable"

# Modify the gcsfuse command to include custom options
```

## Uninstalling

```bash
sudo umount /mnt/gcs-bucket
sudo sed -i '/your-bucket/d' /etc/fstab
```

## Notes

- The script uses `user_allow_other` in FUSE configuration to allow non-root users to access the mounted bucket
- Permissions are set to 777 for both files and directories in the mount
- The mount is configured with `_netdev` flag for network device handling
- Mount will persist across reboots via `/etc/fstab` entry
- Automatically detects and uses the appropriate package manager for your OS
- Modular design allows easy integration into larger automation workflows
