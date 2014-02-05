package gofigure

import (
	"github.com/joeshaw/multierror"
	env "github.com/kelseyhightower/envconfig"
	"github.com/ryansb/gofigure/conf"
	"github.com/ryansb/gofigure/file"
	db "github.com/ryansb/gofigure/mongo"
)

func Process(spec interface{}) error {
	var el multierror.Errors
	e := func(err error) {
		if err != nil {
			el = append(el, err)
		}
	}
	if conf.Settings.UseDB {
		e(db.Process(spec))
	}
	if conf.Settings.UseFile {
		e(file.Process(spec))
	}
	if conf.Settings.UseEnv {
		e(env.Process(conf.Settings.AppName, spec))
	}
	return el.Err()
}
