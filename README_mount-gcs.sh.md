# mount-gcs.sh

A Bash shell script for installing and configuring gcsfuse to mount a Google Cloud Storage (GCS) bucket on Linux systems.

## Purpose

This shell script automates the complete setup process for mounting a GCS bucket by:
- Installing gcsfuse and required dependencies
- Configuring FUSE filesystem permissions
- Creating a mount directory with appropriate permissions
- Mounting the GCS bucket with specified options
- Persisting the mount configuration in `/etc/fstab`
- Verifying the mount was successful

## Requirements

- Bash shell
- Root/sudo access
- Linux system with apt package manager (Debian/Ubuntu-based)
- curl and lsb-release utilities
- Valid GCS bucket name
- Google Cloud credentials configured on the system

## Usage

### Prerequisites

Before running the script, edit the configuration variables at the top:

```bash
BUCKET_NAME="your-gcs-bucket-name" 
MOUNT_POINT="/mnt/gcs-bucket"
```

### Running the Script

```bash
sudo bash mount-gcs.sh
```

or make it executable first:

```bash
chmod +x mount-gcs.sh
sudo ./mount-gcs.sh
```

The script will:
1. Install curl, gnupg, and lsb-release packages
2. Determine the Linux distribution codename
3. Add the official gcsfuse repository
4. Install the gcsfuse package
5. Enable FUSE user_allow_other configuration
6. Create and configure the mount directory
7. Mount the GCS bucket
8. Add the mount configuration to `/etc/fstab`
9. Verify the successful mount with `df` and `mount` commands

## Configuration

Edit these variables at the top of the script:

```bash
BUCKET_NAME="your-gcs-bucket-name"    # Your GCS bucket name
MOUNT_POINT="/mnt/gcs-bucket"         # Local mount directory
```

## Script Sections

The script is organized into clearly marked sections:

1. **Installing gcsfuse** - Updates apt and installs dependencies
2. **Configuring fuse** - Enables user_allow_other in /etc/fuse.conf
3. **Creating mount directory** - Sets up mount point with proper ownership and permissions
4. **Mounting GCS bucket** - Performs the actual mount with all necessary options
5. **Persisting mount in /etc/fstab** - Adds/updates the fstab entry
6. **Verifying mount** - Confirms successful mounting

## Error Handling

The script uses `set -e` to exit immediately if any command fails, ensuring the process stops before corrupting the system state.

## Mount Options Used

- `--implicit-dirs`: Treat prefixes as directories
- `--uid=0, --gid=0`: Set ownership to root user and group
- `--file-mode=777, --dir-mode=777`: Set full read/write/execute permissions
- `-o allow_other`: Allow non-root users to access the mount

## Persistence

The mount configuration is automatically added to `/etc/fstab` with:
- `_netdev` flag for network device handling
- Full permissions and options for automatic remounting on reboot

## Output

The script provides clear progress indicators:
- Section headers showing the current step
- Command execution feedback
- Mount verification output
- Success message: ✅ GCS bucket successfully mounted at /mnt/gcs-bucket

## Notes

- Must be run with sudo/root privileges
- Designed for Debian/Ubuntu-based distributions
- Mount will persist across system reboots
- Permissions allow full access to mounted bucket
- Safe to run multiple times; fstab entry check prevents duplicates
