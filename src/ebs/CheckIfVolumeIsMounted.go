package ebs

import (
	"fmt"
	"os/exec"
)

// CheckIfVolumeIsMounted checks if the given volume device name is mounted by reading /proc/mounts.
func CheckIfVolumeIsMounted(folderPath string) (bool, error) {

	fmt.Printf("Checking if volume is mounted at %s\n", folderPath)

	cmd := exec.Command("mountpoint", folderPath)
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("mountpoint cmd Output: %s", output)
		return false, nil
	}

	fmt.Printf("mountpoint cmd Output: %s", output)

	return true, nil

}
