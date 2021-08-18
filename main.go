package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const DefaultFile = "*_gen.go"

var (
	output = flag.String("o", DefaultFile, "the file to [o]utput to, '*' is replaced with the enum name")
	pkg    = flag.String("p", ".", "the [p]ackage(dir) in which the target sits (default '.')")
)

func main() {
	flag.Parse()

	var enumName string
	pkgName, err := packageName(*pkg)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	name := *output
	if name == "" {
		baseName := strings.ReplaceAll(*output, "*", enumName)
		name = filepath.Join(pkgName, strings.ToLower(baseName))
	}
	err = ioutil.WriteFile(name, []byte{}, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

//packageName returns the package name given a path
//from: https://stackoverflow.com/questions/25262754/how-to-get-name-of-current-package-in-go
func packageName(path string) (string, error) {
	fset := token.NewFileSet()

	// parse the go soure file, but only the package clause
	astFile, err := parser.ParseDir(fset, path, nil, parser.PackageClauseOnly)
	if err != nil {
		return "", err
	}
	packageName := ""
	for _, pack := range astFile {
		if pack.Name == "" {
			continue
		}
		packageName = pack.Name
		break
	}
	if packageName == "" {
		return "", fmt.Errorf("could not read package name from given dir")
	}

	return packageName, nil
}
