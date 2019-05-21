package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func FileExists(fileName string) bool {
	file, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !file.IsDir()
}

func WriteFile(destinationPath string, content []byte) {
	if err := ioutil.WriteFile(destinationPath, content, 0644); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "Generated %s. \n", destinationPath)
}

func GoFormat(path string) {
	if err := exec.Command("gofmt", "-w", path).Run(); err != nil {
		panic(err)
	}
}

func GoImports(path string) {
	if err := exec.Command("goimports", "-w", path).Run(); err != nil {
		panic(err)
	}
}
