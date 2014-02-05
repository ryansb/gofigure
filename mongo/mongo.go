package mongo

// MIT Licensed (see README.md)- Copyright (c) 2014 Ryan S. Brown <sb@ryansb.com>

import (
	"github.com/ryansb/gofigure/conf"
	"github.com/ryansb/gofigure/merger"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func Process(spec interface{}) error {
	c, connCloser := mongoConnect(conf.Settings.MongoDBCollection)
	defer connCloser()

	res := map[string]interface{}{}
	err := c.Find(bson.M{}).One(&res)

	merger.MapAndStruct(res, spec)

	return err
}

func mongoConnect(collectionName string) (*mgo.Collection, func()) {
	session, err := mgo.Dial(conf.Settings.MongoDBHosts)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	confCollection := session.DB(conf.Settings.MongoDBName).C(collectionName)
	return confCollection, session.Close
}
