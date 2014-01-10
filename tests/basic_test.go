package gofigure_test

import (
	"github.com/ryansb/gofigure"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	. "launchpad.net/gocheck"
	"testing"
	"time"
)

type BasicS struct{}

var basicS = Suite(&BasicS{})

func Test(t *testing.T) {
	TestingT(t)
}

type MyAppSpec struct {
	Env   string
	Port  int
	Debug bool
}

const testMongoHost = "172.16.0.101"

func connectMongo(c *C) *mgo.Session {
	session, err := mgo.DialWithTimeout(testMongoHost, 500*time.Millisecond)
	if err != nil {
		c.Fatalf(err.Error())
	}
	session.SetMode(mgo.Monotonic, true)
	session.SetSyncTimeout(200 * time.Millisecond)
	session.SetSocketTimeout(200 * time.Millisecond)
	return session
}

func resetDB(c *C) error {
	session := connectMongo(c)
	err := session.DB("gofigure").C("default").DropCollection()
	if err != nil {
		return err
	}
	err = session.DB("gofigure").C("default").Create(&mgo.CollectionInfo{})
	return err
}

func (s *BasicS) TestMongoConfig(c *C) {
	resetDB(c)
	session := connectMongo(c)
	coll := session.DB("gofigure").C("default")

	err := coll.Insert(bson.M{
		"Port":  1234,
		"Env":   "test",
		"Debug": true,
	})
	if err != nil {
		c.Errorf("Failed to insert basic config. %s", err.Error())
	}
	var i interface{}
	err = coll.Find(new(interface{})).One(&i)
	if err != nil {
		c.Errorf("Failed to query. %s", err.Error())
	}

	gofigure.MongoHosts = testMongoHost

	conf := MyAppSpec{}
	err = gofigure.Process("default", &conf)
	c.Check(err, Equals, nil)
	c.Check(conf.Debug, Equals, true)
	c.Check(conf.Port, Equals, 1234)
	c.Check(conf.Env, Equals, "test")
}
