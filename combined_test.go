package gofigure

// MIT Licensed (see README.md)- Copyright (c) 2014 Ryan S. Brown <sb@ryansb.com>

import (
	"github.com/ryansb/gofigure/conf"
	. "launchpad.net/gocheck"
	"testing"
)

type BasicS struct{}

var basicS = Suite(&BasicS{})

func Test(t *testing.T) {
	TestingT(t)
}

type MyAppSpec struct {
	Env          string
	Port         int
	Debug        bool
	AlarmLoadPct float32
}

func init() {
	conf.Settings.Option(conf.MongoDBHosts("172.16.0.101"))
	conf.Settings.Option(conf.FileLocations("sample/simple.json"))
}

// TODO: make sure that config value heirarchy matches the spec
//       env var > local file > database
