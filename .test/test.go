package main

import (
	"github.com/MikkelHJuul/go-enum/enum"
)

type ASD struct {
	enum.Enum
	Asd int
}

var vals = []ASD{
	{"hello0", 1},
	{Enum: "hello2", Asd: 3},
	{"hello3", 4},
	{Enum: "hello4", Asd: 5},
	{"hello5", 13},
}
