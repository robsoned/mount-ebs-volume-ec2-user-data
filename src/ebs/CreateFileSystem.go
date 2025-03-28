package ebs

import (
	"fmt"
	"os/exec"
)

func CreateFileSystem(deviceName string, folderPath string) error {

	isfileSystemValid := isFileSystemValid(deviceName)

	if isfileSystemValid {
		fmt.Printf("File system already exists on device %s. Not creating file system.\n", deviceName)
		return nil
	}

	fmt.Printf("Creating file system on device %s\n", deviceName)

	_, err := exec.Command("mkfs", "-t", "ext4", deviceName).Output()

	if err != nil {
		fmt.Printf("Error creating file system: %s\n", err)
		return err
	}

	fmt.Printf("File system created on device %s\n", deviceName)

	return nil
}

func isFileSystemValid(deviceName string) bool {

	fmt.Printf("Checking if file system is valid on device %s\n", deviceName)

	_, err := exec.Command("blkid", deviceName).Output()
	if err != nil {
		fmt.Printf("No file system detected on device %s\n", deviceName)
		return false
	}

	fmt.Printf("File system detected on device %s\n", deviceName)
	return true
}
