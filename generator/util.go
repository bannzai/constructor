package generator

import (
	"fmt"
	"io/ioutil"
	"os"
)

func fileExists(fileName string) bool {
	file, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !file.IsDir()
}

func writeFile(destinationPath string, content []byte) {
	if err := ioutil.WriteFile(destinationPath, content, 0644); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "Generated %s. \n", destinationPath)
}
