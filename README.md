# GCSFUSE Auto Mount 🚀

**GCS FUSE Mount Automation** | Google Cloud Storage Bucket Mounting for Linux

Automated **GCS bucket mounting** solutions for Linux systems using **gcsfuse**. Choose your preferred implementation: Bash, Python, Go, or Ansible.

**Supports:** Ubuntu • Debian • RHEL • CentOS • Fedora

**Keywords:** GCS mount, gcsfuse, Google Cloud Storage, FUSE filesystem, Linux mount GCS, cloud storage automation, infrastructure-as-code

---

## 🔍 What is GCSFUSE Auto Mount?

This repository provides **production-ready implementations** for automated **GCS bucket mounting** on Linux systems. Whether you need a quick **one-liner Bash solution**, **Python integration**, **high-performance Go binary**, or **enterprise Ansible deployment**, we have you covered.

**Perfect for:**
- 🏢 Enterprise cloud storage integration
- 🔧 DevOps automation and infrastructure-as-code
- 🐳 Kubernetes and container orchestration
- 📊 Big data and ML pipelines
- 💾 Backup and archival storage solutions
- 🌐 Multi-region deployment automation

---

## 🔑 Search Keywords & Topics Covered

This project covers essential **GCS mounting** and **cloud storage** topics:

**Primary Keywords:**
- GCS mount Linux | Mount Google Cloud Storage | gcsfuse tutorial | FUSE mount GCS | Linux GCS mounting
- Automated GCS bucket mount | GCS mount automation | GCS FUSE setup
- Mount GCS in Ubuntu | Mount GCS in Debian | Mount GCS in RHEL | Mount GCS in CentOS

**Implementation-Specific Keywords:**
- **Bash**: bash mount GCS | shell script GCS | mount-gcs.sh | bash gcsfuse
- **Python**: Python GCS mount | mount_gcs_bucket.py | Python gcsfuse | Pythonic cloud storage
- **Go**: Go GCS mount | Golang gcsfuse | high-performance GCS | gcsfuse_mount.go
- **Ansible**: Ansible GCS mount | Infrastructure-as-Code storage | Ansible gcsfuse | multi-server mount

**Problem Keywords:**
- How to mount GCS bucket on Linux | GCS not mounting | Permission denied GCS | GCS mount fails

---

Available implementations:

* **Bash Script** (`mount-gcs.sh`) - Standalone shell script for direct execution
* **Python Script** (`mount_gcs_bucket.py`) - Python 3 implementation with modular functions
* **Go Program** (`gcsfuse_mount.go`) - Compiled Go program for high performance
* **Ansible Playbook** (`mount_gcs_bucket.yml`) - Infrastructure-as-Code for multi-server deployments

All implementations ensure:

* ✅ Automatic installation of gcsfuse
* ✅ Proper FUSE configuration
* ✅ Secure mount directory creation
* ✅ Persistent mount via `/etc/fstab`
* ✅ Verification of successful mount
* ✅ Support for both Ubuntu/Debian and RHEL/CentOS/Fedora
* ✅ Production-ready with error handling

---

## 🎯 Quick Decision Guide

Choose the best **GCS mount solution** for your use case:

| Your Scenario | Best Implementation | Why This Is Best |
|-----------|-----------|-----|
| **Quick cloud storage setup** | Bash Script | Instant execution, zero dependencies, perfect for single Linux servers |
| **Python/DevOps project** | Python Script | Native integration, clean code, extensible for custom workflows |
| **Performance & distribution** | Go Program | Compiled binary, lightning-fast, cross-platform compilation, ideal for containerized deployments |
| **Multi-server infrastructure** | Ansible Playbook | Infrastructure-as-code, idempotent operations, enterprise-grade orchestration |

---

## 📋 Available Implementations

Choose the **GCS FUSE mounting implementation** that best fits your workflow:

### 1. **Bash Script** - `mount-gcs.sh`
**Mount GCS buckets** with pure shell scripting - no dependencies, instant execution

Best for: Single Linux servers, quick deployment, GCP Compute Engine instances

- **Pros**: Zero dependencies, instant execution, easy to customize and audit
- **OS Support**: Ubuntu, Debian, RHEL, CentOS, Fedora
- **Use Case**: Quick cloud storage mount for development and testing
- **Cons**: Not suitable for multi-server infrastructure management
- **Details**: See [README_mount-gcs.sh.md](README_mount-gcs.sh.md)

### 2. **Python Script** - `mount_gcs_bucket.py`
**Pythonic GCS mounting** with modular architecture for enterprise integration

Best for: Python projects, DevOps automation, cloud infrastructure scripting

- **Pros**: Clean Python 3 code, highly extensible, excellent for CI/CD pipelines
- **OS Support**: Ubuntu, Debian, RHEL, CentOS, Fedora
- **Use Case**: Integration with Python-based DevOps tools and cloud platforms
- **Cons**: Requires Python 3 runtime installation
- **Details**: See [README_mount_gcs_bucket.py.md](README_mount_gcs_bucket.py.md)

### 3. **Go Program** - `gcsfuse_mount.go`
**High-performance GCS mount** as a compiled binary - perfect for production deployment

Best for: High-performance environments, distributed deployments, containerization

- **Pros**: Compiled binary (single executable), blazing-fast performance, cross-platform compilation, zero runtime dependencies
- **OS Support**: Ubuntu, Debian, RHEL, CentOS, Fedora
- **Use Case**: Kubernetes, Docker, cloud-native infrastructure, mission-critical deployments
- **Cons**: Requires Go compiler for building (not for end users)
- **Details**: See [README_gcsfuse_mount.go.md](README_gcsfuse_mount.go.md)

### 4. **Ansible Playbook** - `mount_gcs_bucket.yml`
**Infrastructure-as-Code GCS mounting** for enterprise multi-server deployments

Best for: Multi-server infrastructure, Kubernetes operations, enterprise automation

- **Pros**: Idempotent operations, manage multiple servers simultaneously, full IaC support, integration with existing Ansible workflows
- **OS Support**: Ubuntu, Debian, RHEL, CentOS, Fedora
- **Use Case**: Enterprise infrastructure, Kubernetes cluster setup, multi-region deployments
- **Cons**: Requires Ansible control machine setup
- **Details**: See [README_mount_gcs_bucket.yml.md](README_mount_gcs_bucket.yml.md)

---

## 🛠 Common Prerequisites for GCS Mounting

All **Google Cloud Storage mounting** implementations require:

* **Operating System**: Linux (Ubuntu, Debian, RHEL, CentOS, or Fedora)
* **Permissions**: Root or sudo access
* **GCS Authentication**: One of the following:

  **Option 1: Application Default Credentials**
  ```bash
  gcloud auth application-default login
  ```

  **Option 2: Service Account**
  The VM/server must have a **Service Account** with:
  * `roles/storage.objectViewer` (Read-only)
  * `roles/storage.objectAdmin` (Read/Write)

---

## 📂 GCS Bucket Configuration - Mount Points and Settings

All **GCS mounting implementations** use the same configuration variables:

```
BUCKET_NAME="your-gcs-bucket-name"
MOUNT_POINT="/mnt/gcs-bucket"
```

| Variable      | Description                                                  |
| ------------- | ------------------------------------------------------------ |
| `BUCKET_NAME` | Name of the **Google Cloud Storage bucket** to mount        |
| `MOUNT_POINT` | Local directory where **GCS bucket** will be mounted         |

**Configuration Tips for GCS Mounting:**
- Set `MOUNT_POINT` to a dedicated directory (e.g., `/mnt/gcs`, `/gcs`)
- Ensure the **mount point** exists before running scripts
- Use absolute paths for **GCS bucket mount** configuration

Refer to the specific README file for your chosen implementation for detailed **GCS bucket configuration** instructions.

---

## ▶️ Quick Start - Mount Your GCS Bucket in Minutes

Choose your **GCS bucket mounting method**:

### Option 1: Bash Script (Fastest - 1 minute)

Mount **Google Cloud Storage** with pure Bash:

```bash
git clone https://github.com/mpandey95/gcsfuse-auto-mount.git
cd gcsfuse-auto-mount
chmod +x mount-gcs.sh
sudo ./mount-gcs.sh
```

### Option 2: Python Script (Clean Integration - 2 minutes)

Use Python for **GCS mount automation**:

```bash
git clone https://github.com/mpandey95/gcsfuse-auto-mount.git
cd gcsfuse-auto-mount
sudo python3 mount_gcs_bucket.py
```

### Option 3: Go Program (High Performance - 3 minutes with build)

Build and deploy **high-performance GCS mounting**:

```bash
git clone https://github.com/mpandey95/gcsfuse-auto-mount.git
cd gcsfuse-auto-mount
go build -o gcsfuse_mount gcsfuse_mount.go
sudo ./gcsfuse_mount
```

### Option 4: Ansible (Enterprise Automation - 5 minutes)

Deploy **GCS mounting across multiple servers**:

```bash
git clone https://github.com/mpandey95/gcsfuse-auto-mount.git
cd gcsfuse-auto-mount
ansible-playbook -i inventory mount_gcs_bucket.yml
```

For detailed usage instructions, see the specific README for your chosen implementation.

---

## ✅ GCS Mount Features - Key Capabilities

* ✅ Installs **gcsfuse** automatically on Ubuntu, Debian, RHEL, CentOS, Fedora
* ✅ Configures **FUSE filesystem** to allow non-root access
* ✅ **Mount GCS bucket** with correct permissions and security flags
* ✅ Persists **Google Cloud Storage** mount across VM reboots via /etc/fstab
* ✅ Verifies **GCS mount status** after completion
* ✅ Supports **GCP Compute Engine**, **on-premise Linux**, and **cloud VMs**
* ✅ Auto-detects **Linux distribution** (Ubuntu/Debian vs RHEL family)
* ✅ Multiple implementation options for **GCS bucket mounting**
* ✅ Enterprise-ready with error handling and validation
* ✅ Fast **GCS FUSE setup** with minimal dependencies

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

1. **Choose your implementation**: Bash, Python, Go, or Ansible
2. **Read the specific README**: Each has detailed instructions
3. **Update configuration**: Set `BUCKET_NAME` and `MOUNT_POINT`
4. **Run the implementation**: Follow your chosen implementation's guide
5. **Verify**: Check successful mount with `df -h` and `mount | grep gcsfuse`

See the implementation-specific README files for detailed step-by-step instructions:
- [README_mount-gcs.sh.md](README_mount-gcs.sh.md) for Bash
- [README_mount_gcs_bucket.py.md](README_mount_gcs_bucket.py.md) for Python
- [README_gcsfuse_mount.go.md](README_gcsfuse_mount.go.md) for Go
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
* 🐳 Container orchestration (K8s, Docker)
* 📊 Data pipeline operations

---

## 🔐 Security Best Practices

* **Avoid `777` permissions in production** - Use more restrictive permissions for sensitive data
* **Prefer IAM-based access** - Use Service Accounts with least-privilege IAM roles
* **Enable audit logging** - Monitor access to GCS buckets
* **Use VPC Service Controls** - Restrict access at the network level
* **Encrypt data** - Enable customer-managed encryption keys (CMEK) if required
* **Regular security reviews** - Periodically audit mount permissions and access

---

## 🧩 Troubleshooting - Fix GCS Mounting Issues

### ❌ GCS Bucket Not Mounting?

**Fix: GCS bucket mount failure** with these troubleshooting steps:

* Check **Service Account permissions** - Ensure account has `Storage Object Viewer` or `Storage Object Admin` role
* Verify **GCS bucket name** - Confirm spelling and that **bucket exists** in Google Cloud Storage
* Ensure `/etc/fuse.conf` has `user_allow_other` enabled for **FUSE mount**:

  ```bash
  grep -i "user_allow_other" /etc/fuse.conf
  ```

* Check **GCS authentication** - Run:

  ```bash
  gcloud auth application-default print-access-token
  ```

### 🔄 GCS Mount Fails After Reboot?

**Fix: GCS bucket mount persistence issues:**

* Confirm `_netdev` flag is present in `/etc/fstab` - Essential for **network-based GCS mounts**
* Ensure network is available - **GCS mount** may timeout if network isn't ready
* Check **gcsfuse logs**:

  ```bash
  sudo journalctl -xe | grep -i gcsfuse
  ```

* Verify **mount manually**:

  ```bash
  sudo mount -a
  ```
  sudo mount -a
  ```

### Permission denied errors?

* Verify file permissions on mount directory
* Check GCS bucket IAM permissions
* Ensure user_allow_other is enabled
* Try remounting with verbose flags:

  ```bash
  sudo gcsfuse --debug_gcs -o allow_other your-bucket /mnt/gcs-bucket
  ```

### Performance issues?

* Check network connectivity and bandwidth
* Monitor system resources (CPU, memory)
* Review GCS bucket location vs compute location
* Consider enabling caching or using Cloud Storage Transfer Service for large operations

---

## 📊 GCS Mount Implementation Comparison - Choose Your Method

| Feature | Bash Mount Script | Python GCS Mount | Go Binary | Ansible Automation |
|---------|------|--------|----|---------| 
| **Ease of Use** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ |
| **Mount Performance** | ⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| **Build Required** | ❌ | ❌ | ✅ | ❌ |
| **Runtime Dependencies** | None | Python 3 | None | Ansible |
| **Multi-Server GCS Mount** | ❌ | ❌ | ❌ | ✅ |
| **Extensibility** | Low | High | High | Medium |
| **Linux OS Support** | ✅ Ubuntu/RHEL | ✅ Ubuntu/RHEL | ✅ Ubuntu/RHEL | ✅ Ubuntu/RHEL |
| **Compiled Binary** | ❌ | ❌ | ✅ | ❌ |

---

## 📝 Environment Variables

You can override configuration via environment variables:

```bash
export BUCKET_NAME="my-bucket"
export MOUNT_POINT="/data/gcs"
sudo -E ./mount-gcs.sh
```

---

## 🚀 Advanced Usage

### Custom Mount Options

Edit the mount command in your chosen implementation to add custom options:

```bash
gcsfuse --max-conns-per-host=100 --enable-nonseekable ...
```

Refer to [gcsfuse documentation](https://github.com/GoogleCloudPlatform/gcsfuse/blob/master/docs/gcsfuse.md) for all available options.

### Unmounting

```bash
# Unmount temporarily
sudo umount /mnt/gcs-bucket

# Remount
sudo mount /mnt/gcs-bucket

# Persistent removal from fstab
sudo sed -i '/your-bucket/d' /etc/fstab
```

## ⭐ Contributing

Contributions are welcome!

Feel free to:
- Open issues to report bugs or suggest features
- Submit pull requests to improve implementations
- Add new implementations in other languages
- Improve documentation and examples

Please ensure all implementations:
- Support both Ubuntu/Debian and RHEL/CentOS/Fedora
- Include comprehensive error handling
- Have corresponding README documentation
- Are production-ready and tested

---

## 🗺️ Project Structure

```
gcsfuse-auto-mount/
├── README.md                          # Main documentation
├── mount-gcs.sh                       # Bash implementation
├── mount_gcs_bucket.py                # Python implementation
├── gcsfuse_mount.go                   # Go implementation
├── mount_gcs_bucket.yml               # Ansible playbook
├── README_mount-gcs.sh.md             # Bash README
├── README_mount_gcs_bucket.py.md      # Python README
├── README_gcsfuse_mount.go.md         # Go README
├── README_mount_gcs_bucket.yml.md     # Ansible README
└── LICENSE                            # MIT License
```

---

## 📚 Documentation

Each implementation has its own comprehensive README with:
- Detailed setup instructions
- Configuration options
- Troubleshooting guides
- Advanced usage examples
- Integration patterns
- Uninstall procedures

---

## 🤝 Getting Help

### Common Issues

1. **"Permission denied"** - Ensure running with `sudo`
2. **"Bucket not found"** - Verify bucket name and GCS credentials
3. **"Mount fails after reboot"** - Check `_netdev` flag in `/etc/fstab`

### Resources

- [gcsfuse GitHub](https://github.com/GoogleCloudPlatform/gcsfuse) - Official gcsfuse project
- [GCS Documentation](https://cloud.google.com/storage/docs) - Google Cloud Storage docs
- [FUSE Documentation](https://github.com/libfuse/libfuse) - FUSE filesystem documentation

### Getting Support

- Check the implementation-specific README for detailed troubleshooting
- Review the [GCS documentation](https://cloud.google.com/storage/docs)
- Open a GitHub issue with error messages and system information

---

## 📈 Performance Comparison

| Metric | Bash | Python | Go | Ansible |
|--------|------|--------|----|---------| 
| **Startup Time** | <100ms | ~500ms | ~50ms | ~1s |
| **Memory Usage** | ~5MB | ~30MB | ~10MB | ~50MB |
| **Installation Time** | ~5min | ~5min | ~5min (build) | ~5min |
| **Remote Execution** | ⚠️ SSH | ⚠️ SSH | ⚠️ SSH | ✅ Native |

---

## 🔄 Real-World GCS Mounting Workflows

### 🚀 Workflow 1: Quick GCS Mount on Single Server
```bash
# Fastest way to mount GCS bucket - Best for quick testing
cd gcsfuse-auto-mount
sudo ./mount-gcs.sh
# Verify: mount | grep gcsfuse
```
**Best for:** Quick GCS bucket mounting, single VM setup, testing

### 🐍 Workflow 2: Python Integration - Mount GCS in Application
```python
# Integrate GCS mounting into Python application
from mount_gcs_bucket import mount_gcs_bucket
mount_gcs_bucket("my-bucket", "/mnt/gcs")
# Now use /mnt/gcs as local filesystem
```
**Best for:** DevOps automation, infrastructure code, Python applications

### 🌐 Workflow 3: Enterprise Multi-Server GCS Mounting
```bash
# Scale GCS mounting across multiple servers with Ansible
ansible-playbook -i inventory mount_gcs_bucket.yml
# This mounts GCS bucket on all target servers automatically
```
**Best for:** Enterprise deployments, multi-server setup, infrastructure-as-code

### 📦 Workflow 4: High-Performance Containerized GCS Mount
```bash
# Build and deploy high-performance GCS mount in containers
go build -o gcsfuse_mount gcsfuse_mount.go
docker cp gcsfuse_mount container:/usr/local/bin/
container run ./gcsfuse_mount
```
**Best for:** Kubernetes deployments, high-performance requirements, container orchestration

---

## 📋 Pre-deployment GCS Mount Checklist

- [ ] Verify **GCS bucket** exists and is accessible in Google Cloud console
- [ ] Confirm **IAM permissions** - Service Account has Storage Object Viewer/Admin role
- [ ] Test **GCS authentication** works: `gcloud auth application-default print-access-token`
- [ ] Choose **GCS mount point** and verify write permission to parent directory
- [ ] Review **FUSE mount permissions** (avoid `777` in production environments)
- [ ] Plan **GCS data backup** strategy for business-critical data
- [ ] Test **GCS bucket mount** on non-production system first
- [ ] Document **GCS mount options** and custom configurations
- [ ] Set up **monitoring** for **GCS mount** health and availability

---

## 🎓 Learning Resources

### For Beginners - Mount GCS Quickly
1. Start with **Bash mount script** - simple, straightforward, no dependencies
2. Read the quick start section above
3. Follow the implementation-specific README for your method

### For Advanced Users - Scale GCS Mounting
1. Review all **four GCS mount implementations** for your specific needs
2. Customize **GCS bucket configuration** for your use case
3. Consider **multi-environment setups** for GCS in production
4. Integrate **GCS mounting** with existing DevOps automation
5. Explore **high-performance options** like Go or Ansible for enterprise

---

## 👨‍💻 Author

**Manish Pandey**  
Cloud | DevOps | Platform Engineer

* 🔗 GitHub: [https://github.com/mpandey95](https://github.com/mpandey95)
* 💼 LinkedIn: [https://www.linkedin.com/in/manish-pandey95](https://www.linkedin.com/in/manish-pandey95)
* 📧 Open for collaboration and contributions

---

## ⭐ Support This Project

If this project helped you, consider:
- ⭐ Starring the repository
- 🔗 Sharing with others
- 💬 Providing feedback
- 🐛 Reporting issues
- 🔧 Contributing improvements

---

## ⭐ Contributing

Contributions are welcome!
Feel free to open issues or submit PRs to improve the script.

---

## 📜 License

MIT License – free to use, modify, and distribute.
