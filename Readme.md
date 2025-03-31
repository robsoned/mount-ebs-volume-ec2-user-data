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

