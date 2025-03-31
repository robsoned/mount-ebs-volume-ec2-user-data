package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/robsoned/mount-ebs-volume-ec2-user-data/ebs"
)

const EbsVolumeDeviceNameEnvKey = "EBS_VOLUME_DEVICE_NAME"

const ebsFolderPathEnvKey = "EBS_FOLDER_PATH"
const waitEBSVolumeFolderRetryIntervalEnvKey = "WAIT_EBS_VOLUME_FOLDER_RETRTY_INTERVAL"
const waitEBSVolumeFolderMaxRetryEnvKey = "WAIT_EBS_VOLUME_FOLDER_MAX_RETRY"

func main() {

	ebsFolderPath, err := getEBSFolterPath()

	if err != nil {
		panic(err)
	}

	ebs.CreateEBSFolder(ebsFolderPath)

	ebsVolumeDiviceName, err := getEBSVolumeDeviceName()

	if err != nil {

		panic(err)
	}

	waitEBSVolumeFolderRetryInterval, err := getWaitEBSVolumeFolderRetryInterval()

	if err != nil {

		panic(err)
	}

	waitEBSVolumeFolderMaxRetry, err := getWaitEBSVolumeFolderMaxRetry()

	if err != nil {

		panic(err)
	}

	ebs.WaitEBSDeviceToBeAttached(ebsVolumeDiviceName, waitEBSVolumeFolderRetryInterval, waitEBSVolumeFolderMaxRetry)

	isMounted, err := ebs.CheckIfVolumeIsMounted(ebsFolderPath)

	if err != nil {

		panic(err)
	}

	if isMounted {
		fmt.Printf("Volume is mounted at %s\n", ebsFolderPath)
		return
	} else {
		fmt.Printf("Volume is not mounted at %s\n", ebsFolderPath)
	}

	err = ebs.CreateFileSystem(ebsVolumeDiviceName, ebsFolderPath)

	if err != nil {
		panic(err)
	}

	err = ebs.MountVolume(ebsVolumeDiviceName, ebsFolderPath)

	if err != nil {
		panic(err)
	}

}

func getEBSFolterPath() (string, error) {

	ebsFolderPath, error := getEnvironmetVariable(ebsFolderPathEnvKey)

	if error != nil {
		return "", error
	}

	return ebsFolderPath, nil
}

func getEBSVolumeDeviceName() (string, error) {

	ebsVolumeDeviceName, error := getEnvironmetVariable(EbsVolumeDeviceNameEnvKey)

	if error != nil {
		return "", error
	}

	return ebsVolumeDeviceName, nil
}

func getWaitEBSVolumeFolderRetryInterval() (int, error) {

	waitEBSVolumeFolderRetryInterval, error := getEnvironmetVariable(waitEBSVolumeFolderRetryIntervalEnvKey)

	if error != nil {
		return 0, error
	}

	retryInterfal, error := strconv.Atoi(waitEBSVolumeFolderRetryInterval)

	if error != nil {
		return 0, error
	}

	return retryInterfal, nil
}

func getWaitEBSVolumeFolderMaxRetry() (int, error) {

	waitEBSVolumeFolderMaxRetry, error := getEnvironmetVariable(waitEBSVolumeFolderMaxRetryEnvKey)

	if error != nil {
		return 0, error
	}

	maxRetryInt, err := strconv.Atoi(waitEBSVolumeFolderMaxRetry)

	if err != nil {
		return 0, err
	}

	return maxRetryInt, nil
}

func getEnvironmetVariable(key string) (string, error) {

	value := os.Getenv(key)

	if value == "" {
		return "", errors.New(key + " is not set")
	}

	return value, nil
}
