package template



const valueOfMethod = `
package {{.Package}}

import(
	{{range $import := .Imports}}
	"{{ printf $import }}"
	{{end}}
)

{{range $enum := .Enums}}
func {{ $enum.Name }}() {{ .EnumType }} {
	return {{ printf $val }}
}


{{end}}

func Values() [{{len .Enums}}]{{ .EnumType }} {
	return [{{len .Enums}}]{{ .EnumType }} {
	{{range $val := .Enums}}
		{{ printf $val }},
	{{end}}
	}
}

func EnumMap() map[string]{{ .EnumType }} {
	return map[string]{{ .EnumType }}{
	{{range $val := .Enums}}
		"{{$val.Name}}": {{ printf $val }},
	{{end}}
	}
}


func ValueOf(enumName string) (val {{ .EnumType }}, ok bool) {
	val, ok = EnumMap()[enumName]
	return
}
`