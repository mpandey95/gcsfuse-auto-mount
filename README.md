# GCSFUSE Auto Mount 🚀

**Production-ready GCS bucket mounting for Linux** — Bash • Python • Go • Ansible

Automated solutions to install, configure, and mount Google Cloud Storage buckets on Linux using **gcsfuse**.

| Implementation | Best For | Cmd |
|---|---|---|
| **Bash** | Single servers | `chmod +x mount-gcs.sh && sudo ./mount-gcs.sh` |
| **Python** | Python projects | `sudo python3 mount_gcs_bucket.py` |
| **Ansible** | Multi-server | `ansible-playbook -i inventory mount_gcs_bucket.yml` |
| **Go** | High performance | See [README_gcsfuse_mount.go.md](README_gcsfuse_mount.go.md) |

**Prerequisites:** Linux (Ubuntu/Debian/RHEL), root/sudo, GCP auth: `gcloud auth application-default login`

**Configuration:** Edit `BUCKET_NAME="your-bucket"` and `MOUNT_POINT="/mnt/gcs-bucket"`

**What it does:** Install deps → Add gcsfuse repo → Install gcsfuse → Enable FUSE → Create mount dir → Mount bucket → Add to fstab → Verify

**Features:** Auto-install, persistent mounts, production-ready, works on GCP/on-prem/cloud VMs

**Support:** [GitHub](https://github.com/mpandey95) • [LinkedIn](https://www.linkedin.com/in/manish-pandey95)

[Docs](README_mount-gcs.sh.md) | [License](LICENSE)
