package util

import (
	"log"
	"os"
)

func MakeBinDir() string {
	folderPath := GetRootPath() + "/bin"
	_, err := os.Stat(folderPath)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(folderPath, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	return folderPath
}
