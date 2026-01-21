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

## Requirements

- Python 3
- Root/sudo access
- Linux system with apt package manager (Debian/Ubuntu-based)
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
- ▶ prefix for each command being executed
- ❌ error indicator if running without root privileges
- Section headers showing progress through installation steps

## Notes

- The script uses `user_allow_other` in FUSE configuration to allow non-root users to access the mounted bucket
- Permissions are set to 777 for both files and directories in the mount
- The mount is configured with `_netdev` flag for network device handling
- Mount will persist across reboots via `/etc/fstab` entry
