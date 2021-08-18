package enum


type Named interface {
	Name() string
}

type Enum string

func (e Enum) Name() string {
	return string(e)
}