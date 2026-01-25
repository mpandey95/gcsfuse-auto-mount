# mount_gcs_bucket.yml

An Ansible playbook for installing and configuring gcsfuse to mount a Google Cloud Storage (GCS) bucket on remote Linux systems.

## Purpose

This Ansible playbook automates the deployment and configuration of gcsfuse across one or more servers by:
- Installing required dependencies
- Adding the gcsfuse apt/yum repository
- Installing gcsfuse package
- Configuring FUSE filesystem permissions
- Creating the mount directory with proper permissions
- Adding the mount configuration to `/etc/fstab`
- Mounting the GCS bucket with appropriate options
- Verifying the mount was successful

## Supported Distributions

The playbook supports both major Linux distribution families:

- **Debian/Ubuntu**: Ubuntu, Debian
- **RHEL/CentOS/Fedora**: RHEL 7+, CentOS 7+, Fedora

The playbook automatically detects the OS family (`ansible_os_family`) and uses conditional tasks to apply the correct package manager and configuration.

## Requirements

- Ansible installed on the control machine
- Target hosts must be Debian/Ubuntu or RHEL/CentOS/Fedora based systems
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

1. **Display detected OS family** - Shows the detected OS for verification
2. **Install required packages (Debian/Ubuntu)** - Installs curl, gnupg, lsb-release
3. **Install required packages (RHEL/CentOS/Fedora)** - Installs curl, gnupg
4. **Add gcsfuse apt repository (Debian/Ubuntu)** - Configures Google Cloud apt repository
5. **Add Google Cloud apt key (Debian/Ubuntu)** - Adds GPG key for package verification
6. **Add gcsfuse yum repository (RHEL/CentOS/Fedora)** - Configures Google Cloud yum repository
7. **Install gcsfuse (Debian/Ubuntu)** - Installs gcsfuse package via apt
8. **Install gcsfuse (RHEL/CentOS/Fedora)** - Installs gcsfuse package via yum
9. **Enable user_allow_other** - Modifies `/etc/fuse.conf` to allow non-root users
10. **Create mount directory** - Creates mount point with 755 permissions
11. **Add GCS bucket mount to /etc/fstab** - Ensures mount persists across reboots
12. **Mount GCS bucket** - Performs the actual mount operation
13. **Verify mount** - Confirms successful mounting
14. **Show gcsfuse mount** - Displays mount verification output

## Conditional Execution

The playbook uses Ansible's `when` conditions to execute OS-specific tasks:

- Tasks with `when: ansible_os_family == "Debian"` run on Ubuntu/Debian systems
- Tasks with `when: ansible_os_family == "RedHat"` run on RHEL/CentOS/Fedora systems

This ensures only the appropriate commands are executed for each system type.

## Package Manager Support

- **Debian/Ubuntu**: Uses `apt` module for package management
- **RHEL/CentOS/Fedora**: Uses `yum` module for package management and `yum_repository` for repository configuration

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

## Inventory Setup

Create an `inventory.ini` file:

```ini
[gcs-servers]
server1.example.com
server2.example.com
server3.example.com

[gcs-servers:vars]
ansible_user=ubuntu
ansible_ssh_private_key_file=~/.ssh/id_rsa
```

## Advanced Usage

### Variable Overrides

```bash
ansible-playbook -i inventory mount_gcs_bucket.yml \
  -e bucket_name="production-bucket" \
  -e mount_point="/data/production"
```

### With Password Prompt

```bash
ansible-playbook -i inventory mount_gcs_bucket.yml \
  --ask-become-pass
```

### Dry Run (Check Mode)

```bash
ansible-playbook -i inventory mount_gcs_bucket.yml \
  --check --diff
```

## Scaling to Multiple Environments

Create separate playbooks for different environments:

```yaml
# site.yml
- import_playbook: mount_gcs_bucket.yml
  vars:
    bucket_name: "{{ environment }}-bucket"
    mount_point: "/mnt/{{ environment }}-gcs"
```

## Health Check Playbook

```yaml
- name: Verify GCS Mounts
  hosts: gcs-servers
  tasks:
    - name: Check mount status
      command: df -h /mnt/gcs-bucket
      register: mount_status
    
    - name: Display mount status
      debug:
        msg: "{{ mount_status.stdout_lines }}"
```

## Troubleshooting

### SSH Connection Issues

```bash
# Test SSH connectivity
ansible gcs-servers -i inventory -m ping

# Enable verbose output
ansible-playbook -i inventory mount_gcs_bucket.yml -vvv
```

### Sudoers Configuration

Ensure passwordless sudo on target hosts:

```bash
# On target hosts
sudo visudo
# Add line: ubuntu ALL=(ALL) NOPASSWD:ALL
```

### Repository Key Errors

```bash
# Manually add GPG keys on target
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys KEY_ID
```

## Uninstalling

Create an uninstall playbook:

```yaml
- name: Unmount GCS bucket
  hosts: gcs-servers
  become: true
  tasks:
    - name: Unmount GCS bucket
      mount:
        path: /mnt/gcs-bucket
        state: absent
    
    - name: Remove gcsfuse
      package:
        name: gcsfuse
        state: absent
```

## Notes

- The mount is added to `/etc/fstab` for automatic remounting on system reboot
- Permissions are set to 777 for both files and directories
- The playbook is idempotent and can be run multiple times safely
- Automatically detects OS family and uses appropriate package manager
- Supports both Debian/Ubuntu and RHEL/CentOS/Fedora distributions
- Requires Ansible 2.0+ with knowledge of target systems' OS family
- Works seamlessly with Ansible Galaxy and Tower/AWX
