package main

import (
	"strings"
)

type myEnum interface {
	hello()
	Name() string
}

type base struct {
	PrettyName string
	number     int
}

type eOne struct {
	base
}

func (o eOne) legs() int {
	return o.number
}
func (o eOne) hello() {}

type eTwo struct {
	base
}

func (o eTwo) wings() int {
	return o.number * o.number
}
func (o eTwo) hello() {}

func (a base) Name() string {
	return strings.ToUpper(a.PrettyName)
}

//go:generate go-enum
func Values() [5]myEnum {
	return [...]myEnum{
		eOne{base{"hello0", 1}},
		eTwo{base{"hello2", 3}},
		eOne{base{"hello3", 4}},
		eOne{base{"hello4", 5}},
		eTwo{base{"hello5", 13}},
	}
}
