package ebs

import (
	"fmt"
	"os"
)

func CreateEBSFolder(folderPath string) error {

	if exists, err := checkIfFolderExists(folderPath); exists {
		if err != nil {
			fmt.Printf("Error checking if folder exists: %s\n", err)
			return err
		}
		fmt.Printf("Folder %s already exists. Not creating folder \n", folderPath)
		return nil
	}

	fmt.Printf("Foder %s does not exist. Creating folder.\n", folderPath)

	err := os.MkdirAll(folderPath, 0755)

	if err != nil {
		fmt.Printf("Error creating folder: %s\n", err)
	}

	return err

}

func checkIfFolderExists(folderPath string) (bool, error) {

	fmt.Printf("Checking if folder exists at %s\n", folderPath)

	if _, err := os.Stat(folderPath); err == nil {
		return true, nil
	}

	return false, nil

}
