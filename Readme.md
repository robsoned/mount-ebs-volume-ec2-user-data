# Mount EBS Volume EC2 User Data

This repository provides a Go-based application to automate the process of mounting an EBS volume on an EC2 instance.

## Features

- Waits for the EBS volume to be attached.
- Creates a folder for mounting the EBS volume if it doesn't exist.
- Checks if the volume is already mounted.
- Creates a file system on the EBS volume if necessary.
- Mounts the EBS volume to the specified folder.

## Environment Variables

The application uses the following environment variables:

- `EBS_VOLUME_DEVICE_NAME`: The device name of the EBS volume (e.g., sdh).
- `EBS_FOLDER_PATH`: The folder path where the EBS volume will be mounted.
- `WAIT_EBS_VOLUME_FOLDER_RETRTY_INTERVAL`: Retry interval (in seconds) for checking if the EBS volume is attached.
- `WAIT_EBS_VOLUME_FOLDER_MAX_RETRY`: Maximum number of retries for checking if the EBS volume is attached.

...
## Example Usage in EC2 User-Data

Add the following script to your EC2 instance's user-data script:

```bash


MOUNT_VOLUME_BINARY_URL=https://github.com/robsoned/mount-ebs-volume-ec2-user-data/releases/download/0.0.2/mount-ebs-volume-ec2-user-data-0.0.2-linux-amd64.tar.gz

curl -L $MOUNT_VOLUME_BINARY_URL -o /tmp/mount-ebs-volume-ec2-user-data.tar.gz && \
tar -xzf /tmp/mount-ebs-volume-ec2-user-data.tar.gz -C /tmp

EBS_VOLUME_DEVICE_NAME=/dev/sdh \
WAIT_EBS_VOLUME_FOLDER_RETRTY_INTERVAL=10 \
WAIT_EBS_VOLUME_FOLDER_MAX_RETRY=10 \
EBS_FOLDER_PATH=/home/ec2-user/ebs-volume \
/tmp/mount-ebs-volume-ec2-user-data

``` 

> **Note:** Remember to update the environment variables, such as `EBS_VOLUME_DEVICE_NAME`, to match your specific configuration as needed.

> **Important:** This script must be run as the root user, or with `sudo` privileges.