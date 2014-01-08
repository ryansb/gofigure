package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/ryansb/gofigure"
)

type MyAppSpec struct {
	Debug      bool
	Port       int
	MongoHost  string
	DBName     string
	Collection string
}

func main() {
	fmt.Println("main")
	s := new(MyAppSpec)
	envconfig.Process("GOFIGURE", s)
	fmt.Println(s.Debug)

	//c, closer := gofigure.MC("default")
	//defer closer()
	//c.Insert(bson.M{"Port": 1235, "DBName": "gofigure"})

	//	gofigure.MongoHosts = "172.16.0.101"

	err := gofigure.Process("default", s)
	if err != nil {
		println("process failure")
		println(err.Error())
		println("process failure")
	}
	fmt.Println(s.Debug)
	fmt.Println(s.Port)
}
