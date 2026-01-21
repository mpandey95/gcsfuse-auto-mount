# mount_gcs_bucket.yml

An Ansible playbook for installing and configuring gcsfuse to mount a Google Cloud Storage (GCS) bucket on remote Linux systems.

## Purpose

This Ansible playbook automates the deployment and configuration of gcsfuse across one or more servers by:
- Installing required dependencies
- Adding the gcsfuse apt repository
- Installing gcsfuse package
- Configuring FUSE filesystem permissions
- Creating the mount directory with proper permissions
- Adding the mount configuration to `/etc/fstab`
- Mounting the GCS bucket with appropriate options
- Verifying the mount was successful

## Requirements

- Ansible installed on the control machine
- Target hosts must be Debian/Ubuntu-based systems with apt
- SSH access to target hosts with sudo privileges
- Valid GCS bucket name
- Google Cloud credentials configured on target systems

## Usage

### Prerequisites

Update the variables in the playbook for your environment:

```yaml
vars:
    bucket_name: "your-gcs-bucket-name"
    mount_point: "/mnt/gcs-bucket"
```

### Running the Playbook

```bash
ansible-playbook -i inventory mount_gcs_bucket.yml
```

Or target specific hosts:

```bash
ansible-playbook -i inventory mount_gcs_bucket.yml -l hostname
```

## Configuration

Edit these variables in the playbook to customize:

- `bucket_name`: The name of your GCS bucket to mount
- `mount_point`: The local directory where the bucket will be mounted on target systems
- `gcsfuse_repo`: Automatically derived from the system's LSB codename (e.g., focal, jammy)

## Playbook Tasks

The playbook includes the following tasks:

1. **Install required packages** - Installs curl, gnupg, and lsb-release
2. **Add gcsfuse apt repository** - Configures the Google Cloud apt repository
3. **Add Google Cloud apt key** - Adds GPG key for package verification
4. **Install gcsfuse** - Installs the gcsfuse package
5. **Enable user_allow_other** - Modifies `/etc/fuse.conf` to allow non-root users
6. **Create mount directory** - Creates mount point with 755 permissions
7. **Add GCS bucket mount to /etc/fstab** - Ensures mount persists across reboots
8. **Mount GCS bucket** - Performs the actual mount operation
9. **Verify mount** - Confirms successful mounting
10. **Show gcsfuse mount** - Displays mount verification output

## Privileges

The playbook runs with `become: true`, requiring sudo access on target hosts. Ensure your Ansible user can execute sudo commands without a password, or provide the `--ask-become-pass` flag.

## Mount Options

The GCS bucket is mounted with these options:
- `rw`: Read-write access
- `implicit_dirs`: Treat prefixes as directories
- `uid=0, gid=0`: Set ownership to root
- `file_mode=777, dir_mode=777`: Set full permissions
- `allow_other`: Allow non-root users to access
- `_netdev`: Treat as network device

## Notes

- The mount is added to `/etc/fstab` for automatic remounting on system reboot
- Permissions are set to 777 for both files and directories
- The playbook is idempotent and can be run multiple times safely
