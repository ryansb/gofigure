package gofigure

// MIT Licensed (see README.md)- Copyright (c) 2014 Ryan S. Brown <sb@ryansb.com>

import (
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
	Settings.Option(MongoDBHosts("172.16.0.101"))
	Settings.Option(FileLocations("sample/simple.json"))
}

// TODO: make sure that config value heirarchy matches the spec
//       env var > local file > database
