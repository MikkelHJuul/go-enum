// this file is partially adapted from golang.org/x/tools/cmd/stringer
package stringer

import (
	"fmt"
	enum2 "github.com/MikkelHJuul/go-enum/enum"
	"go/ast"
	"go/types"
	"strings"
)

type EnumDecl interface {
	enum2.Named
	String() string
}

func (k keyValueDecl) Name() string {
	return ""
}

type keyValueDecl struct {
	EnumDecl
	kvs []struct{key, value fmt.Stringer}
}

func (k keyValueDecl) String() string {
	var listOfKVs []string
	for _, kv := range k.kvs {
		listOfKVs = append(listOfKVs, fmt.Sprintf("%s: %s", kv.key, kv.value))
	}
	return fmt.Sprintf("{%s}", strings.Join(listOfKVs,", "))
}

type basicDecl struct {

}

func (b basicDecl) String() string {
	return ""
}

func (b basicDecl) Name() string {
	return ""
}


func (f *File) genDecl(node ast.Node) bool {
	//Imports??
	valueNode, ok := node.(*ast.Ident)
	if !ok {
		return true
	}
	if valueNode.Name != f.typeName {
		return true
	}
	decl := valueNode.Obj.Decl
	valueDecl, ok := decl.(*ast.ValueSpec)
	if !ok {
		return true
	}
	obj, ok := f.pkg.defs[valueNode]
	if !ok {
		return true
	}
	if !hasNameMethod() {

	}
	values := valueDecl.Values[0]
	comp, ok := values.(*ast.CompositeLit)
	if !ok {
		return true
	}
	_ = comp.Type.(*ast.Ident).Name
	enums := []EnumDecl{}
	for _, elt := range comp.Elts {
		var val EnumDecl
		if compLit, ok := elt.(*ast.CompositeLit); ok {
			for _, vals := range compLit.Elts {
				switch vals.(type) {
				case *ast.BasicLit:
					val = basicDecl{}
				case *ast.KeyValueExpr:
					val = keyValueDecl{}
				default:
				}
			}
		}
		enums = append(enums, val)
	}



	return true
}

func (g *Generator) Generate() error {
	for _, file := range g.pkg.files {
		// Set the state for this run of the walker.
		if file.file != nil {
			ast.Inspect(file.file, file.genDecl)
		}
	}

	return nil
}



// Value represents a declared constant.
type Value struct {
	originalName string // The name of the constant.
	name         string // The name with trimmed prefix.
	// The value is stored as a bit pattern alone. The boolean tells us
	// whether to interpret it as an int64 or a uint64; the only place
	// this matters is when sorting.
	// Much of the time the str field is all we need; it is printed
	// by Value.String.
	value  uint64 // Will be converted to int64 when needed.
	signed bool   // Whether the constant is a signed type.
	str    string // The string representation given by the "go/constant" package.
}

func (v *Value) String() string {
	return v.str
}


func hasNameMethod(o types.Object) bool {
	a := []types.Type{o.Type(), types.NewPointer(o.Type())}
	for i := range a {
		ms := types.NewMethodSet(a[i])
		for j := 0; j < ms.Len(); j++ {
			object := ms.At(j).Obj()
			if m, ok := object.(*types.Func); ok {

				if strings.Contains(m.String(), "Name()") {
					return true
				}
			}
		}
	}
	return false
}