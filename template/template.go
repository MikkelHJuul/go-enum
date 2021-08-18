package template

type SerializationObj struct {
	PackageName string
	EnumType    string
	Enums       []struct {
		Name, Type string
		Ordinal    int
	}
}

const enumGeneratorTmpl = `
package {{ .PackageName }}


{{range $enum := .Enums}}
func {{ $enum.Name }}() {{ $enum.Type }} {
	return Values()[{{ $enum.Ordinal }}]
}


{{end}}

func EnumMap() map[string]{{ .EnumTypeName }} {
	return map[string]{{ .EnumTypeName }}{
	{{range $enum := .Enums}}
		"{{ $enum.Name }}": Values()[{{ $enum.ordinal }}],
	{{end}}
	}
}


func ValueOf(enumName string) (val {{ .EnumTypeName }}, ok bool) {
	val, ok = EnumMap()[enumName]
	return
}
`
