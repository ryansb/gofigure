package file

import (
	"github.com/ryansb/gofigure"
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
	gofigure.Settings.Option(gofigure.FileLocations("sample/simple.json"))
}

func (s *BasicS) TestFileString(c *C) {
	conf := MyAppSpec{}
	err := Process(&conf)
	c.Check(err, Equals, nil)
	c.Check(conf.Env, Equals, "test")
}

func (s *BasicS) TestFileBool(c *C) {
	conf := MyAppSpec{}
	err := Process(&conf)
	c.Check(err, Equals, nil)
	c.Check(conf.Debug, Equals, true)
}

func (s *BasicS) TestFileFloat(c *C) {
	conf := MyAppSpec{}
	err := Process(&conf)
	c.Check(err, Equals, nil)
	c.Check(conf.AlarmLoadPct, Equals, float32(0.9))
}

func (s *BasicS) TestFileInt(c *C) {
	conf := MyAppSpec{}
	err := Process(&conf)
	c.Check(err, Equals, nil)
	c.Check(conf.Port, Equals, 1234)
}
