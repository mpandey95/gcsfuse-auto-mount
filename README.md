# GCSFUSE Auto Mount 🚀

This repository contains **multiple implementations** to **install, configure, mount, and persist a Google Cloud Storage (GCS) bucket** on Linux systems using **gcsfuse**.

Available implementations:

* **Bash Script** (`mount-gcs.sh`) - Standalone shell script for direct execution
* **Python Script** (`mount_gcs_bucket.py`) - Python 3 implementation with modular functions
* **Ansible Playbook** (`mount_gcs_bucket.yml`) - Infrastructure-as-Code for multi-server deployments

All implementations ensure:

* Automatic installation of gcsfuse
* Proper FUSE configuration
* Secure mount directory creation
* Persistent mount via `/etc/fstab`
* Verification of successful mount

---

## � Available Implementations

Choose the implementation that best fits your workflow:

### 1. **Bash Script** - `mount-gcs.sh`
Best for: Single servers, direct execution, lightweight automation

- **Pros**: No dependencies (only bash), fast execution, easy to customize
- **Cons**: Not suitable for multi-server management
- **Details**: See [README_mount-gcs.sh.md](README_mount-gcs.sh.md)

### 2. **Python Script** - `mount_gcs_bucket.py`
Best for: Python-based environments, programmatic integration, modular code

- **Pros**: Clean Python code, easier to extend, good for Python projects
- **Cons**: Requires Python 3 installation
- **Details**: See [README_mount_gcs_bucket.py.md](README_mount_gcs_bucket.py.md)

### 3. **Ansible Playbook** - `mount_gcs_bucket.yml`
Best for: Multi-server deployments, infrastructure management, repeatable automation

- **Pros**: Manage multiple servers, idempotent, part of IaC ecosystem
- **Cons**: Requires Ansible installation and setup
- **Details**: See [README_mount_gcs_bucket.yml.md](README_mount_gcs_bucket.yml.md)

---

## 🛠 Common Prerequisites

All implementations require:

* Linux OS (Ubuntu/Debian based)
* Root or sudo privileges
* Google Cloud SDK authentication already configured:

  ```bash
  gcloud auth application-default login
  ```

  OR the VM/server has a **Service Account** with:

  * `Storage Object Viewer` / `Storage Object Admin`

---

## 📂 Configuration

All implementations use the same configuration variables:

```
BUCKET_NAME="your-gcs-bucket-name"
MOUNT_POINT="/mnt/gcs-bucket"
```

| Variable      | Description                                  |
| ------------- | -------------------------------------------- |
| `BUCKET_NAME` | Name of the GCS bucket                       |
| `MOUNT_POINT` | Local directory where bucket will be mounted |

Refer to the specific README file for your chosen implementation for detailed configuration instructions.

---

## ▶️ Quick Start

Choose one of the following based on your needs:

### Option 1: Bash Script (Simplest)

```bash
git clone https://github.com/mpandey95/gcsfuse-auto-mount.git
cd gcsfuse-auto-mount
chmod +x mount-gcs.sh
sudo ./mount-gcs.sh
```

### Option 2: Python Script

```bash
git clone https://github.com/mpandey95/gcsfuse-auto-mount.git
cd gcsfuse-auto-mount
sudo python3 mount_gcs_bucket.py
```

### Option 3: Ansible Playbook

```bash
git clone https://github.com/mpandey95/gcsfuse-auto-mount.git
cd gcsfuse-auto-mount
ansible-playbook -i inventory mount_gcs_bucket.yml
```

For detailed usage instructions, see the specific README for your chosen implementation.

---

## �📌 Features

* ✅ Installs **gcsfuse** automatically
* ✅ Configures **FUSE** to allow non-root access
* ✅ Mounts GCS bucket with correct permissions
* ✅ Persists mount across VM reboots
* ✅ Verifies mount status
* ✅ Suitable for **GCP Compute Engine**, **on-prem Linux**, and **cloud VMs**

---

## 🛠 Prerequisites

Before running any implementation, ensure:

* Linux OS (Ubuntu/Debian based)
* Root or sudo privileges
* Google Cloud SDK authentication already configured:

  ```bash
  gcloud auth application-default login
  ```

  OR the VM has a **Service Account** with:

  * `Storage Object Viewer` / `Storage Object Admin`

---

## 📂 Script Variables

Update these variables in your chosen implementation:

```
BUCKET_NAME="your-gcs-bucket-name"
MOUNT_POINT="/mnt/gcs-bucket"
```

| Variable      | Description                                  |
| ------------- | -------------------------------------------- |
| `BUCKET_NAME` | Name of the GCS bucket                       |
| `MOUNT_POINT` | Local directory where bucket will be mounted |

For specific configuration instructions, refer to the README file for your chosen implementation.

---

## ▶️ How to Use

1. **Choose your implementation**: Bash, Python, or Ansible
2. **Read the specific README**: Each has detailed instructions
3. **Update configuration**: Set `BUCKET_NAME` and `MOUNT_POINT`
4. **Run the implementation**: Follow your chosen implementation's guide

See the implementation-specific README files for detailed step-by-step instructions:
- [README_mount-gcs.sh.md](README_mount-gcs.sh.md) for Bash
- [README_mount_gcs_bucket.py.md](README_mount_gcs_bucket.py.md) for Python
- [README_mount_gcs_bucket.yml.md](README_mount_gcs_bucket.yml.md) for Ansible

---

## 🔁 Persistent Mount (Auto Mount on Reboot)

The script automatically adds the following entry to `/etc/fstab`:

```text
bucket-name /mount/dir gcsfuse rw,implicit_dirs,uid=0,gid=0,file_mode=777,dir_mode=777,allow_other,_netdev 0 0
```

This ensures the bucket remains mounted after reboot.

---

## 🔍 Verification

The script validates the mount using:

```bash
df -h
mount | grep gcsfuse
```

Expected output:

```
✅ GCS bucket successfully mounted at /mount/dir
```

---

## 💡 Use Cases

* 📦 Centralized file storage for applications
* 🧠 ML/AI model storage
* 🗄 Log aggregation
* 🧪 CI/CD artifact storage
* 🌐 Shared storage across VMs
* 🔄 Backup & archival access

---

## 🔐 Security Notes

* Avoid using `777` permissions in production unless required
* Prefer IAM-based access via **Service Accounts**
* Use least-privilege IAM roles

---

## 🧩 Troubleshooting

**Bucket not mounting?**

* Check service account permissions
* Verify bucket name
* Ensure `/etc/fuse.conf` has:

  ```
  user_allow_other
  ```

**Mount fails after reboot?**

* Confirm `_netdev` is present in `/etc/fstab`
* Ensure network is available

---

## 👨‍💻 Author

**Manish Pandey**
Cloud | DevOps | Platform Engineer

* 🔗 GitHub: [https://github.com/mpandey95](https://github.com/mpandey95)
* 💼 LinkedIn: [https://www.linkedin.com/in/manish-pandey95](https://www.linkedin.com/in/manish-pandey95)

---

## ⭐ Contributing

Contributions are welcome!
Feel free to open issues or submit PRs to improve the script.

---

## 📜 License

MIT License – free to use, modify, and distribute.
