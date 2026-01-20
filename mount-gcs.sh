#!/bin/bash
set -e

BUCKET_NAME="your-gcs-bucket-name" 
MOUNT_POINT="/mnt/gcs-bucket"

echo "=============================="
echo " Installing gcsfuse"
echo "=============================="

apt update
apt install -y curl gnupg lsb-release

export GCSFUSE_REPO=gcsfuse-$(lsb_release -c -s)
echo "deb https://packages.cloud.google.com/apt $GCSFUSE_REPO main" | tee /etc/apt/sources.list.d/gcsfuse.list
curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -

apt update
apt install -y gcsfuse

echo "=============================="
echo " Configuring fuse"
echo "=============================="

sed -i 's/#user_allow_other/user_allow_other/' /etc/fuse.conf

echo "=============================="
echo " Creating mount directory"
echo "=============================="

mkdir -p ${MOUNT_POINT}
chown root:root ${MOUNT_POINT}
chmod 755 ${MOUNT_POINT}

echo "=============================="
echo " Mounting GCS bucket"
echo "=============================="

gcsfuse \
  --implicit-dirs \
  --uid=0 \
  --gid=0 \
  --file-mode=777 \
  --dir-mode=777 \
  -o allow_other \
  ${BUCKET_NAME} \
  ${MOUNT_POINT}

echo "=============================="
echo " Persisting mount in /etc/fstab"
echo "=============================="

grep -q "${BUCKET_NAME} ${MOUNT_POINT} gcsfuse" /etc/fstab || \
echo "${BUCKET_NAME} ${MOUNT_POINT} gcsfuse rw,implicit_dirs,uid=0,gid=0,file_mode=777,dir_mode=777,allow_other,_netdev 0 0" >> /etc/fstab

echo "=============================="
echo " Verifying mount"
echo "=============================="

df -h | grep ${MOUNT_POINT}
mount | grep gcsfuse

echo "✅ GCS bucket successfully mounted at ${MOUNT_POINT}"
