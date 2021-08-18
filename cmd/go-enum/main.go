package main

import (
	"flag"
	"fmt"
	"github.com/MikkelHJuul/go-enum/cmd/go-enum/stringer"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)
const DefaultFile = "*_gen.go"

var (
	output = flag.String("o", DefaultFile, "the file to [o]utput to, '*' is replaced with the enum name")
	target = flag.String("t", "", "the [t]arget identifier that provides the enum-set")
	pkg = flag.String("p", ".", "the [p]ackage in which the target sits (default '.')")
)

func main() {
	flag.Parse()
	if err := verifyTarget(*target); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		flag.Usage()
		os.Exit(2)
	}

	byts, err := process(*pkg, *target)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	name := *output
	if name == "" {
		baseName := strings.ReplaceAll(*output, "*", *target)
		name = filepath.Join(*pkg, strings.ToLower(baseName))
	}
	err = ioutil.WriteFile(name, byts, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

func process(pack, trgt string) ([]byte, error) {
	gen := stringer.Generator{}
	if err := gen.ParsePackage(pack, trgt); err != nil {
		return nil, err
	}
	if err := gen.Generate(); err != nil {
		return nil, err
	}
	return gen.Format(), nil
}

func verifyTarget(t string) error {
	if len(t) == 0 {
		return fmt.Errorf("no target provided")
	}
	if len(strings.Split(t, ",")) != 1 {
		return fmt.Errorf("must provide exactly one target")
	}
	return nil
}


