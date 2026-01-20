# GCSFUSE Auto Mount Script 🚀

This repository contains a **Bash script** to **install, configure, mount, and persist a Google Cloud Storage (GCS) bucket** on a Linux VM using **gcsfuse**.

The script ensures:

* Automatic installation of gcsfuse
* Proper FUSE configuration
* Secure mount directory creation
* Persistent mount via `/etc/fstab`
* Verification of successful mount

---

## 📌 Features

* ✅ Installs **gcsfuse** automatically
* ✅ Configures **FUSE** to allow non-root access
* ✅ Mounts GCS bucket with correct permissions
* ✅ Persists mount across VM reboots
* ✅ Verifies mount status
* ✅ Suitable for **GCP Compute Engine**, **on-prem Linux**, and **cloud VMs**

---

## 🛠 Prerequisites

Before running this script, ensure:

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

Update these variables at the top of the script:

```bash
BUCKET_NAME="your-gcs-bucket-name"
MOUNT_POINT="/mnt/gcs-bucket"
```

| Variable      | Description                                  |
| ------------- | -------------------------------------------- |
| `BUCKET_NAME` | Name of the GCS bucket                       |
| `MOUNT_POINT` | Local directory where bucket will be mounted |

---

## ▶️ How to Use

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/mpandey95/gcsfuse-auto-mount.git
cd gcsfuse-auto-mount
```

### 2️⃣ Make Script Executable

```bash
chmod +x mount-gcs.sh
```

### 3️⃣ Run the Script

```bash
sudo ./mount-gcs.sh
```

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
