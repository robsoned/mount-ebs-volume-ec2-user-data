package ebs

import (
	"fmt"
	"os"
	"time"
)

func WaitEBSDeviceToBeAttached(deviceName string, retryInterval int, maxRetries int) error {
	for i := 1; i <= maxRetries; i++ {
		if _, err := os.Stat(deviceName); err == nil {
			fmt.Printf("EBS volume %s is now available.\n", deviceName)
			return nil
		}

		fmt.Printf("Waiting for EBS volume %s to be attached. Retrying in %d seconds.\n", deviceName, retryInterval)
		time.Sleep(time.Duration(retryInterval) * time.Second)

	}

	errorMessage := fmt.Sprintf("EBS volume %s was not attached after %d retries", deviceName, maxRetries)

	fmt.Printf("%s", errorMessage)

	return fmt.Errorf("%s", errorMessage)
}
