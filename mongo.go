package gofigure

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var MongoHosts = "localhost"
var MongoDBName = "gofigure"

func Process(collection string, spec interface{}) error {
	c, connCloser := mongoConnect(collection)
	defer connCloser()

	fmt.Println(spec)
	f := new(map[string]interface{})
	err := c.Find(bson.M{}).One(spec)
	fmt.Println(f)

	mergeMapAndStruct(*f, spec)

	if err != nil {
		fmt.Println(spec)
		return err
	}
	return nil
}

func MC(cname string) (*mgo.Collection, func()) {
	return mongoConnect(cname)
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
