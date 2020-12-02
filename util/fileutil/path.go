package fileutil

import "os"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}
