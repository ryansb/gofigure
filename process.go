package gofigure

import (
	"github.com/joeshaw/multierror"
	env "github.com/kelseyhightower/envconfig"
	cfg "github.com/ryansb/gofigure/conf"
	"github.com/ryansb/gofigure/file"
	db "github.com/ryansb/gofigure/mongo"
)

func Process(spec interface{}) error {
	var el multierror.Errors
	if err := db.Process(spec); err != nil {
		el = append(el, err)
	}
	if err := file.Process(spec); err != nil {
		el = append(el, err)
	}
	if err := env.Process(cfg.Settings.AppName, spec); err != nil {
		el = append(el, err)
	}
	return el.Err()
}
