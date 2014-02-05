package gofigure

// MIT Licensed (see README.md)- Copyright (c) 2014 Ryan S. Brown <sb@ryansb.com>

import (
	"github.com/ryansb/gofigure/conf"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	. "launchpad.net/gocheck"
	"os"
	"testing"
	"time"
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
	conf.Settings.Option(conf.FileLocations("file/sample/simple.json"))
}

// TODO: make sure that config value heirarchy matches the spec
//       env var > local file > database

func (s *BasicS) TestEnvOverridesFile(c *C) {
	origdb := conf.Settings.Option(conf.UseDB(false))
	os.Setenv("GOFIGURE_PORT", "4321")
	thespec := MyAppSpec{}
	err := Process(&thespec)
	c.Check(err, Equals, nil)
	c.Check(thespec.Port, Equals, 4321)
	c.Check(thespec.AlarmLoadPct, Equals, float32(0.9))

	// Now turn off environment config and retry
	origenv := conf.Settings.Option(conf.UseEnv(false))
	thespec = MyAppSpec{}
	err = Process(&thespec)
	c.Check(err, Equals, nil)
	c.Check(thespec.Port, Equals, 1234)
	conf.Settings.Option(origenv)

	conf.Settings.Option(origdb)
}
func (s *BasicS) TestEnvOverridesDB(c *C) {
	origfile := conf.Settings.Option(conf.UseFile(false))
	session, err := mgo.DialWithTimeout(conf.Settings.MongoDBHosts, 500*time.Millisecond)
	if err != nil {
		c.Fatalf(err.Error())
	}
	defer session.Close()
	err = session.DB("gofigure").C("default").DropCollection()
	if err != nil && err.Error() != "ns not found" {
		c.Fatalf("Error dropping collection: %s", err.Error())
	}
	err = session.DB("gofigure").C("default").Create(&mgo.CollectionInfo{})
	if err != nil {
		c.Fatalf("Error creating collection: %s", err.Error())
	}
	coll := session.DB("gofigure").C("default")
	err = coll.Insert(bson.M{
		"Port":         1234,
		"AlarmLoadPct": 0.9,
		"Env":          "test",
		"Debug":        true,
	})
	if err != nil {
		c.Fatalf("Failed to insert basic config. %s", err.Error())
	}

	os.Setenv("GOFIGURE_PORT", "4321")
	thespec := MyAppSpec{}
	err = Process(&thespec)
	c.Check(err, Equals, nil)
	c.Check(thespec.Port, Equals, 4321)
	c.Check(thespec.AlarmLoadPct, Equals, float32(0.9))

	// Now try without the env
	origenv := conf.Settings.Option(conf.UseEnv(false))
	thespec = MyAppSpec{}
	err = Process(&thespec)
	c.Check(err, Equals, nil)
	c.Check(thespec.Port, Equals, 1234)
	conf.Settings.Option(origenv)

	conf.Settings.Option(origfile)
}
