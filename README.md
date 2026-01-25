# GCSFUSE Auto Mount - Complete Overview

Automatically install, configure, and mount Google Cloud Storage (GCS) buckets on Linux systems using **gcsfuse**. Choose from four different implementations based on your needs.

---

## 📦 Available Implementations

### 1. **Bash Script** - `mount-gcs.sh`
**Best for:** Single servers, direct execution, lightweight automation

- **Pros:** No dependencies (only bash), fast execution, easy to customize
- **Requirements:** Bash, root/sudo, Debian/Ubuntu, curl, lsb-release
- **Installation:** `chmod +x mount-gcs.sh && sudo ./mount-gcs.sh`
- **Details:** [README_mount-gcs.sh.md](README_mount-gcs.sh.md)

### 2. **Python Script** - `mount_gcs_bucket.py`
**Best for:** Python environments, programmatic integration, modular code

- **Pros:** Clean Python code, easier to extend, integrates with Python projects
- **Requirements:** Python 3, root/sudo, Debian/Ubuntu
- **Installation:** `sudo python3 mount_gcs_bucket.py`
- **Details:** [README_mount_gcs_bucket.py.md](README_mount_gcs_bucket.py.md)

### 3. **Ansible Playbook** - `mount_gcs_bucket.yml`
**Best for:** Multi-server deployments, infrastructure-as-code, repeatable automation

- **Pros:** Manage multiple servers, idempotent, part of IaC ecosystem
- **Requirements:** Ansible, Debian/Ubuntu targets, SSH with sudo access
- **Installation:** `ansible-playbook -i inventory mount_gcs_bucket.yml`
- **Details:** [README_mount_gcs_bucket.yml.md](README_mount_gcs_bucket.yml.md)

### 4. **Go Program** - `gcsfuse_mount.go`
**Best for:** Production systems, compiled binaries, performance-critical deployments

- **Pros:** Single binary, no runtime dependencies, fast execution
- **Requirements:** Go compiler, Debian/Ubuntu
- **Compilation & Installation:** See [README_gcsfuse_mount.go.md](README_gcsfuse_mount.go.md)
- **Details:** [README_gcsfuse_mount.go.md](README_gcsfuse_mount.go.md)

---

## 🛠 Common Prerequisites (All Implementations)

- **Linux OS:** Ubuntu or Debian-based distributions
- **Root/Sudo Access:** Required for installation and mounting
- **Google Cloud Authentication:** One of the following:
  ```bash
  gcloud auth application-default login
  ```
  OR a Service Account with `Storage Object Viewer` or `Storage Object Admin` permissions

---

## ⚙️ Configuration

All implementations use the same configuration variables:

```bash
BUCKET_NAME="your-gcs-bucket-name"      # Your GCS bucket name
MOUNT_POINT="/mnt/gcs-bucket"           # Where to mount locally
```

**Where to set:**
- **Bash/Python/Go:** Edit the variables at the top of the script
- **Ansible:** Edit the `vars` section in the playbook

---

## 🚀 Quick Start Comparison

| Task | Bash | Python | Ansible | Go |
|------|------|--------|---------|-----|
| **Single Server** | ✅ Best | ✅ Good | ⚠️ Overkill | ✅ Best |
| **Multiple Servers** | ❌ No | ❌ No | ✅ Best | ❌ No |
| **No Dependencies** | ✅ Yes | ❌ Needs Python | ❌ Needs Ansible | ✅ Yes |
| **Customizable** | ✅ Easy | ✅ Easy | ✅ Easy | ⚠️ Needs recompile |
| **Production Ready** | ✅ Yes | ✅ Yes | ✅ Yes | ✅ Yes |

---

## ✅ What All Implementations Do

Regardless of which you choose, each implementation performs:

1. **Install Dependencies** - curl, gnupg, lsb-release
2. **Add GCSFuse Repository** - Official Google Cloud apt repository
3. **Install GCSFuse** - The main mounting tool
4. **Configure FUSE** - Enable `user_allow_other` for non-root access
5. **Create Mount Directory** - `/mnt/gcs-bucket` (or your custom path)
6. **Mount GCS Bucket** - With full read/write permissions (777)
7. **Persist Mount** - Add entry to `/etc/fstab` for automatic remount on reboot
8. **Verify Mount** - Test that the mount was successful

---

## 📋 Mount Configuration Details

All implementations mount with these options:

```
rw              # Read-write access
implicit_dirs   # Treat prefixes as directories
uid=0, gid=0    # Root ownership
file_mode=777   # Full file permissions
dir_mode=777    # Full directory permissions
allow_other     # Allow non-root user access
_netdev         # Treat as network device
```

---

## 🔍 How to Choose

1. **Single server, no dependencies needed?** → **Bash Script** (`mount-gcs.sh`)
2. **Part of Python application?** → **Python Script** (`mount_gcs_bucket.py`)
3. **Managing many servers?** → **Ansible Playbook** (`mount_gcs_bucket.yml`)
4. **Production binary, maximum performance?** → **Go Program** (`gcsfuse_mount.go`)

---

## 📂 Project Structure

```
gcsfuse-auto-mount/
├── README.md                          # Main project README
├── README_overview.md                 # This file (overview of all)
├── mount-gcs.sh                       # Bash implementation
├── README_mount-gcs.sh.md             # Bash documentation
├── mount_gcs_bucket.py                # Python implementation
├── README_mount_gcs_bucket.py.md      # Python documentation
├── mount_gcs_bucket.yml               # Ansible implementation
├── README_mount_gcs_bucket.yml.md     # Ansible documentation
├── gcsfuse_mount.go                   # Go implementation
├── README_gcsfuse_mount.go.md         # Go documentation
└── LICENSE                            # Project license
```

---

## 🎯 Getting Started

1. **Choose your implementation** based on your use case
2. **Read the specific README** for that implementation
3. **Update configuration variables** (`BUCKET_NAME` and `MOUNT_POINT`)
4. **Run the script/playbook** with appropriate permissions
5. **Verify the mount** using:
   ```bash
   df -h | grep gcsfuse
   mount | grep gcsfuse
   ```

---

## 📚 Implementation-Specific Documentation

- [Bash Script Guide](README_mount-gcs.sh.md)
- [Python Script Guide](README_mount_gcs_bucket.py.md)
- [Ansible Playbook Guide](README_mount_gcs_bucket.yml.md)
- [Go Program Guide](README_gcsfuse_mount.go.md)

---

## ✨ Features (All Implementations)

- ✅ Automatic gcsfuse installation
- ✅ FUSE filesystem configuration
- ✅ Persistent mount across reboots
- ✅ Automatic mount verification
- ✅ Works on GCP Compute Engine, on-prem, and cloud VMs
- ✅ Idempotent (safe to run multiple times)
- ✅ Clear error handling and feedback

---

## 📝 Notes

- All implementations are production-ready
- Choose based on your infrastructure and preferences
- Each can be run safely multiple times
- Mount will persist across system reboots
- Full read/write access is granted to the mounted bucket

---

## 📞 Support

For issues specific to:
- **GCSFuse:** See [official documentation](https://cloud.google.com/storage/docs/gcs-fuse)
- **Google Cloud:** Check [Cloud Storage documentation](https://cloud.google.com/storage/docs)
- **This project:** Review the implementation-specific README files
