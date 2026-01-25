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

## Implementation Docs
- [Bash Script](README_mount-gcs.sh.md)
- [Python Script](README_mount_gcs_bucket.py.md)
- [Ansible Playbook](README_mount_gcs_bucket.yml.md)
- [Go Program](README_gcsfuse_mount.go.md)

## Author & Skills
**Manish Pandey** — Senior DevOps/Platform Engineer

### 🛠️ Technology Stack

#### ☁️ Cloud & Platforms
![GCP](https://img.shields.io/badge/GCP-Expert-blue?style=flat-square&logo=google-cloud)
![AWS](https://img.shields.io/badge/AWS-Advanced-orange?style=flat-square&logo=amazon-aws)

#### ⚙️ Platform & DevOps
![Kubernetes](https://img.shields.io/badge/Kubernetes-Expert-blue?style=flat-square&logo=kubernetes)
![Docker](https://img.shields.io/badge/Docker-Expert-2496ED?style=flat-square&logo=docker)
![Terraform](https://img.shields.io/badge/Terraform-Advanced-7B42BC?style=flat-square&logo=terraform)
![Helm](https://img.shields.io/badge/Helm-Advanced-0F1689?style=flat-square&logo=helm)
![Ansible](https://img.shields.io/badge/Ansible-Expert-EE0000?style=flat-square&logo=ansible)
![CI/CD](https://img.shields.io/badge/CI%2FCD-Expert-green?style=flat-square&logo=github-actions)

#### 🔐 Security & Ops
![IAM](https://img.shields.io/badge/IAM-Expert-red?style=flat-square)
![Networking](https://img.shields.io/badge/Networking-Advanced-blue?style=flat-square)
![Monitoring](https://img.shields.io/badge/Monitoring%20%26%20Logging-Expert-green?style=flat-square&logo=grafana)
![Secrets Management](https://img.shields.io/badge/Secrets%20Mgmt-Advanced-yellow?style=flat-square&logo=vault)

#### 🧑‍💻 Programming
![Python](https://img.shields.io/badge/Python-Expert-3776AB?style=flat-square&logo=python)
![Bash](https://img.shields.io/badge/Bash-Expert-4EAA25?style=flat-square&logo=gnu-bash)
![YAML](https://img.shields.io/badge/YAML-Advanced-CB171E?style=flat-square)
![Go](https://img.shields.io/badge/Go-Advanced-00ADD8?style=flat-square&logo=go)

#### 💾 Database
![SQL](https://img.shields.io/badge/SQL-Expert-CC2927?style=flat-square&logo=mysql)
![MongoDB](https://img.shields.io/badge/MongoDB-Advanced-13AA52?style=flat-square&logo=mongodb)

## Connect With Me
- **GitHub:** [mpandey95](https://github.com/mpandey95)
- **LinkedIn:** [manish-pandey95](https://www.linkedin.com/in/manish-pandey95)
- **Email:** mnshkmrpnd@gmail.com

## License
See [LICENSE](LICENSE)
**Support:** [GitHub](https://github.com/mpandey95) • [LinkedIn](https://www.linkedin.com/in/manish-pandey95)
