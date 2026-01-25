# mount-gcs.sh

**Bash script** automating GCS bucket mounting on Linux (Debian/Ubuntu/RHEL/CentOS/Fedora).

**Requirements:** Bash, sudo, curl, GCP credentials

## Configuration
```bash
BUCKET_NAME="your-bucket"
MOUNT_POINT="/mnt/gcs-bucket"
```

## Usage
```bash
chmod +x mount-gcs.sh
sudo ./mount-gcs.sh
```

**Steps:**
1. OS detection (apt vs yum)
2. Install curl, gnupg, lsb-release
3. Add gcsfuse repo
4. Install gcsfuse
5. Enable FUSE user_allow_other
6. Create mount directory
7. Mount GCS bucket
8. Add to /etc/fstab
9. Verify mount

**Features:**
- Zero dependencies (Bash only)
- Error handling via `set -e`
- Prevents duplicate fstab entries
- Auto-remount on reboot
- Environment variable override: `export BUCKET_NAME="bucket" && sudo -E bash mount-gcs.sh`

**Support:** [GitHub](https://github.com/mpandey95) | [LinkedIn](https://linkedin.com/in/manish-pandey95)

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
