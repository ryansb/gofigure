package gofigure

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var MongoHosts = "localhost"
var MongoDBName = "gofigure"

func Process(collection string, spec interface{}) error {
	c, connCloser := mongoConnect(collection)
	defer connCloser()

	res := map[string]interface{}{}
	err := c.Find(bson.M{}).One(&res)

	mergeMapAndStruct(res, spec)

	return err
}

func mongoConnect(collectionName string) (*mgo.Collection, func()) {
	session, err := mgo.Dial(MongoHosts)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	confCollection := session.DB(MongoDBName).C(collectionName)
	return confCollection, session.Close
}
