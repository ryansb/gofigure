package gofigure

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
