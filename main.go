package main

import (
	"bufio"
	"bytes"
	"embed"
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

const DefaultFile = "*_gen.go"

var (
	output = flag.String("o", DefaultFile, "the file to [o]utput to, '*' is replaced with the enum name")
	pkg    = flag.String("p", ".", "the [p]ackage(dir) in which the target sits (default '.')")
)

const testName = "TestNot_EnumBuild"

//go:embed template
var templates embed.FS

func main() {
	flag.Parse()

	pkgName, err := packageName(*pkg)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	err = prepareAndExecuteGoTest(templateNames{
		Dir:         *pkg,
		PackageName: pkgName,
		FileName:    *output,
		TestName:    testName,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func prepareAndExecuteGoTest(names templateNames) error {
	templ, err := template.ParseFS(templates, "template/go-enum.inject_test.tmpl")
	if err != nil {
		return err
	}
	path := filepath.Join(names.Dir, "go-enum.inject_test.go")
	_ = os.Remove(path)
	f, err := os.Create(path)
	defer func(p string) { os.Remove(p) }(path)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	err = templ.Execute(w, names)
	if err != nil {
		return err
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	byt, err := templates.ReadFile("template/go-enum.enum.tmpl")
	if err != nil {
		return err
	}
	path = filepath.Join(names.Dir, "go-enum.enum.tmpl")
	_ = os.Remove(path)
	err = ioutil.WriteFile(path, byt, 0644)
	defer func(p string) { os.Remove(p) }(path)
	if err != nil {
		return err
	}
	err = executeGoTest(names)
	return err
}

func executeGoTest(names templateNames) error {
	//this is probably flimsy
	cmd := exec.Command("go", "test", "-run", names.TestName)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error executing command: %s ,%v", out.String(), err)
	}
	return nil
}

type templateNames struct {
	Dir, PackageName, FileName, TestName string
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
