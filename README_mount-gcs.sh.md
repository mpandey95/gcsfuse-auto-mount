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

## Supported Distributions

The script supports both major Linux distribution families:

- **Debian/Ubuntu**: Ubuntu, Debian
- **RHEL/CentOS/Fedora**: RHEL 7+, CentOS 7+, Fedora

The script automatically detects the OS and uses the appropriate package manager (`apt` for Debian/Ubuntu, `yum` for RHEL).

## Requirements

- Bash shell
- Root/sudo access
- Linux system (Debian/Ubuntu or RHEL/CentOS/Fedora based)
- curl and package management tools
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

1. **OS Detection** - Detects the Linux distribution and version
2. **Installing gcsfuse** - Uses apt for Debian/Ubuntu or yum for RHEL/CentOS/Fedora
3. **Configuring fuse** - Enables user_allow_other in /etc/fuse.conf
4. **Creating mount directory** - Sets up mount point with proper ownership and permissions
5. **Mounting GCS bucket** - Performs the actual mount with all necessary options
6. **Persisting mount in /etc/fstab** - Adds/updates the fstab entry
7. **Verifying mount** - Confirms successful mounting

## OS-Specific Package Installation

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

## Environment Variables

Override configuration via environment variables:

```bash
export BUCKET_NAME="my-gcs-bucket"
export MOUNT_POINT="/data/gcs"
sudo -E bash mount-gcs.sh
```

## Troubleshooting

### Permission Denied

```bash
# Verify fuse.conf
grep user_allow_other /etc/fuse.conf

# If missing, add it manually
sudo sed -i 's/#user_allow_other/user_allow_other/' /etc/fuse.conf
```

### Script Exits with Error

Check the specific error message and review:

```bash
# Check system logs
sudo journalctl -xe | grep -i mount

# Verify GCS authentication
gcloud auth application-default print-access-token
```

### Mount Not Persisting After Reboot

```bash
# Verify fstab entry
grep gcsfuse /etc/fstab

# Test fstab configuration
sudo mount -a
```

## Uninstalling/Unmounting

```bash
# Unmount the bucket
sudo umount /mnt/gcs-bucket

# Remove from fstab (optional)
sudo sed -i '/your-bucket-name/d' /etc/fstab

# Uninstall gcsfuse (optional)
sudo apt remove gcsfuse  # Debian/Ubuntu
sudo yum remove gcsfuse   # RHEL/CentOS/Fedora
```

## Advanced Options

Customize the mount command in the script:

```bash
gcsfuse \
  --max-conns-per-host=100 \
  --dir-mode=755 \
  --file-mode=644 \
  ${BUCKET_NAME} \
  ${MOUNT_POINT}
```

See [gcsfuse documentation](https://github.com/GoogleCloudPlatform/gcsfuse) for all options.

## Notes

- Must be run with sudo/root privileges
- Designed for Debian/Ubuntu and RHEL/CentOS/Fedora distributions
- Automatically detects the OS and uses appropriate package manager
- Mount will persist across system reboots
- Safe to run multiple times; fstab entry check prevents duplicates
- Works on bare metal, VMs, containers, and cloud instances
