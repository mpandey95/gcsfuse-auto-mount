# mount_gcs_bucket.yml

**Ansible playbook** for multi-server GCS bucket mounting on Debian/Ubuntu/RHEL/CentOS/Fedora.

## Requirements
- Ansible control machine
- SSH access to targets with sudo
- Debian/Ubuntu or RHEL systems
- Valid GCS credentials

## Configuration
```yaml
vars:
  bucket_name: "your-bucket"
  mount_point: "/mnt/gcs-bucket"
```

## Usage
```bash
ansible-playbook -i inventory mount_gcs_bucket.yml
# Or specific hosts:
ansible-playbook -i inventory mount_gcs_bucket.yml -l hostname
```

## What it does
1. Detect OS family (Debian/RedHat)
2. Install curl, gnupg, lsb-release (Debian) or curl, gnupg (RedHat)
3. Add gcsfuse apt/yum repo
4. Install gcsfuse
5. Enable FUSE user_allow_other
6. Create mount directory
7. Add mount to /etc/fstab
8. Mount GCS bucket
9. Verify mount

## Features
- **Conditional tasks**: OS-family detection with Ansible `when` conditions
- **Package managers**: Apt for Debian/Ubuntu, yum for RHEL
- **Multi-server**: Deploy to multiple hosts simultaneously
- **Idempotent**: Safely re-run without issues
- **Repo auto-derivation**: gcsfuse_repo from LSB codename

**Support:** [GitHub](https://github.com/mpandey95) | [LinkedIn](https://linkedin.com/in/manish-pandey95)

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
