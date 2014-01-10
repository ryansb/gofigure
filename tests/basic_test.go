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

func resetDB(c *C) {
	session := connectMongo(c)
	defer session.Close()
	err := session.DB("gofigure").C("default").DropCollection()
	if err != nil {
		c.Fatalf("Error dropping collection: %s", err.Error())
	}
	err = session.DB("gofigure").C("default").Create(&mgo.CollectionInfo{})
	if err != nil {
		c.Fatalf("Error creating collection: %s", err.Error())
	}
}

func (s *BasicS) TestMongoConfig(c *C) {
	resetDB(c)
	session := connectMongo(c)
	defer session.Close()
	coll := session.DB("gofigure").C("default")

	err := coll.Insert(bson.M{
		"Port":  1234,
		"Env":   "test",
		"Debug": true,
	})
	if err != nil {
		c.Errorf("Failed to insert basic config. %s", err.Error())
	}
	var i map[string]interface{}
	err = coll.Find(bson.M{}).One(&i)
	if err != nil {
		c.Errorf("Failed to query. %s", err.Error())
	}
	c.Log("Got back: ", i)
	c.Check(i["Port"], Equals, 1234)
	c.Check(i["Env"], Equals, "test")
	c.Check(i["Debug"], Equals, true)

	gofigure.MongoHosts = testMongoHost

	conf := MyAppSpec{}
	err = gofigure.Process("default", &conf)
	c.Check(err, Equals, nil)
	c.Check(conf.Debug, Equals, true)
	c.Check(conf.Port, Equals, 1234)
	c.Check(conf.Env, Equals, "test")
}
