package ebs

import (
	"fmt"
	"os"
	"os/exec"
)

func MountVolume(deviceName string, mountPath string) error {

	fmt.Printf("Creating mount point folder at %s\n", mountPath)

	err := os.MkdirAll(mountPath, 0755)

	if err != nil {
		fmt.Printf("Error creating mount point folder: %s\n", err)
		return err
	}

	output, err := exec.Command("mount", deviceName, mountPath).Output()

	fmt.Printf("mount command output: %s\n", output)

	if err != nil {
		fmt.Printf("Error mounting volume: %s\n", err)
		return err
	}

	fmt.Printf("Mounted %s to %s\n", deviceName, mountPath)
	return nil

}
