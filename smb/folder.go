package smb

import "os"

func InitSMBFolder(folderPath string) error {
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Chmod(folderPath, os.ModePerm)
	return err
}
