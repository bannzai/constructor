package file

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/tools/imports"
)

func FileExists(fileName string) bool {
	file, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !file.IsDir()
}

func WriteFile(destinationPath string, content string) {
	if err := ioutil.WriteFile(destinationPath, []byte(content), 0644); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "Generated %s. \n", destinationPath)
}

func GoFormat(path string) {
	// NONE:
}

func GoImports(path string) {
	// reference: https://github.com/golang/tools/blob/master/cmd/goimports/goimports.go#L41
	options := &imports.Options{
		TabWidth:   8,
		TabIndent:  true,
		Comments:   true,
		Fragment:   true,
		FormatOnly: false,
	}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	src, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	res, err := imports.Process(path, src, options)
	if err != nil {
		panic(err)
	}
	WriteFile(path, string(res))
}
