# mount_gcs_bucket.py

**Python 3 script** automating GCS bucket mounting on Linux (Debian/Ubuntu/RHEL/CentOS/Fedora).

## Requirements
- Python 3, Root/sudo, Valid GCS credentials

## Configuration
```python
BUCKET_NAME = "your-bucket"
MOUNT_POINT = "/mnt/gcs-bucket"
```

## Usage
```bash
sudo python3 mount_gcs_bucket.py
```

**Automates:**
1. Install curl, gnupg, lsb-release
2. Add gcsfuse repo (apt/yum)
3. Install gcsfuse
4. Enable FUSE config
5. Create mount directory
6. Mount GCS bucket
7. Persist in /etc/fstab
8. Verify mount

## Features
- **OS Detection**: Auto-detects Debian/RHEL families
- **Package Manager**: Uses apt (Debian) or yum (RHEL)
- **Error Handling**: Root privilege checks, unsupported OS detection
- **Environment Variables**: `export BUCKET_NAME="bucket" && sudo -E python3 mount_gcs_bucket.py`

**Support:** [GitHub](https://github.com/mpandey95) | [LinkedIn](https://linkedin.com/in/manish-pandey95)

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
