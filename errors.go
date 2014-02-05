package gofigure

// MIT Licensed (see README.md)- Copyright (c) 2014 Ryan S. Brown <sb@ryansb.com>

import (
	"fmt"
)

type ParseFailure struct {
	Field string
	Type  string
	Value string
}

func (p *ParseFailure) Error() string {
	return fmt.Sprintf("gofigure.Parsing: failure converting field=%s "+
		"value with='%s' to type=%s", p.Field, p.Value, p.Type)
}
