package output

import "os"

func fileExists(fileName string) bool {
	file, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !file.IsDir()
}
