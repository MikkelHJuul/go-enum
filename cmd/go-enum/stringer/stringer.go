// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// from golang.org/x/tools/cmd/stringer with some changes
package stringer

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
)

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf bytes.Buffer // Accumulated output.
	pkg *Package     // Package we are scanning.

}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

// File holds a single parsed file and associated data.
type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	typeName string  // Name of the constant type.
	value   Value // Accumulator for constant values of that type.
}

type Package struct {
	name  string
	defs  map[*ast.Ident]types.Object
	files []*File
}

// ParsePackage analyzes the single package constructed from the patterns and tags.
// ParsePackage exits if there is an error.
func (g *Generator) ParsePackage(pkg, target string) error {
	cfg := &packages.Config{
		Mode: packages.LoadAllSyntax,
		// TODO: Need to think about constants in test files. Maybe write type_string_test.go
		// in a separate pass? For later.
		Tests:      false,
	}
	pkgs, err := packages.Load(cfg, pkg)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		return fmt.Errorf("error: %d packages found", len(pkgs))
	}
	g.addPackage(pkgs[0], target)
	return nil
}

// addPackage adds a type checked Package and its syntax files to the generator.
func (g *Generator) addPackage(pkg *packages.Package, target string) {
	g.pkg = &Package{
		name:  pkg.Name,
		defs:  pkg.TypesInfo.Defs,
		files: make([]*File, len(pkg.Syntax)),
	}

	for i, file := range pkg.Syntax {
		g.pkg.files[i] = &File{
			file:        file,
			pkg:         g.pkg,
			typeName: 	 target,
		}
	}
}



// Format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) Format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}