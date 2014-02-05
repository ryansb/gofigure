package file

// MIT Licensed (see README.md)- Copyright (c) 2014 Ryan S. Brown <sb@ryansb.com>

import (
	"encoding/json"
	"github.com/joeshaw/multierror"
	"github.com/ryansb/gofigure/conf"
	"github.com/ryansb/gofigure/merger"
	"io/ioutil"
)

// Process takes a spec and a list of files. The list of files is in
// **reverse** order of precedence, meaning that a value in the first file can
// be overridden by any subsequent file.
//
// Example:
// Process(mySpec, "/etc/myapp.json", "/usr/local/etc/myapp.json", "/home/me/.myapp.json")
// values in /home/me/.myapp.json will override any values set in /etc/myapp.json
func Process(spec interface{}) error {
	var el multierror.Errors
	res := map[string]interface{}{}
	for _, fname := range conf.Settings.FileLocations {
		contents, err := ioutil.ReadFile(fname)
		if err != nil {
			el = append(el, err)
			continue
		}
		json.Unmarshal(contents, &res)
	}

	if err := merger.MapAndStruct(res, spec); err != nil {
		el = append(el, err)
	}

	return el.Err()
}
