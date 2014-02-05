package conf

// MIT Licensed (see README.md)- Copyright (c) 2014 Ryan S. Brown <sb@ryansb.com>

// conf is actually how you configure gofigure through the magic of
// self-referential functions. An implemention that is a shameless duplicate of
// [post by Rob Pike][pike-blogpost]
//
// Example:
// figOpt := new(GoFigOpt)
// figOpt.Option(conf.MongoHosts("localhost"))
// // load a your local test configuration
// prevMongos := figOpt.Option(conf.MongoHosts("p-mongo-1.go.com,p-mongo-2.go.com"))
// // process a prod config for sillies
// figOpt.Option(prevMongos) // now your mongo server will be set to "localhost" again
//
// [pike-blogpost]: http://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html

var Settings GoFigOpt

func init() {
	Settings = GoFigOpt{
		UseDB:             true,
		UseFile:           true,
		UseEnv:            true,
		AppName:           "gofigure",
		MongoDBHosts:      "localhost",
		MongoDBName:       "gofigure",
		MongoDBCollection: "default",
		FileLocations:     []string{"./gofigure.json", "./.gofigure.json"},
	}
}

type GoFigOpt struct {
	UseDB             bool     // default true
	UseFile           bool     // default true
	UseEnv            bool     // default true
	AppName           string   // default "gofigure"
	MongoDBHosts      string   // default "localhost"
	MongoDBName       string   // default "gofigure"
	MongoDBCollection string   // default "default"
	FileLocations     []string // default []string{"./gofigure.json", "./.gofigure.json"}
}

type option func(*GoFigOpt) option

func (g *GoFigOpt) Option(opts ...option) (previous option) {
	for _, opt := range opts {
		previous = opt(g)
	}
	return
}

func AppName(h string) option {
	return func(g *GoFigOpt) option {
		prev := g.AppName
		g.AppName = h
		return AppName(prev)
	}
}

func MongoDBHosts(h string) option {
	return func(g *GoFigOpt) option {
		prev := g.MongoDBHosts
		g.MongoDBHosts = h
		return MongoDBHosts(prev)
	}
}

func MongoDBName(h string) option {
	return func(g *GoFigOpt) option {
		prev := g.MongoDBName
		g.MongoDBName = h
		return MongoDBName(prev)
	}
}

func FileLocations(h ...string) option {
	return func(g *GoFigOpt) option {
		prev := g.FileLocations
		g.FileLocations = h
		return FileLocations(prev...)
	}
}

func UseDB(h bool) option {
	return func(g *GoFigOpt) option {
		prev := g.UseDB
		g.UseDB = h
		return UseDB(prev)
	}
}

func UseFile(h bool) option {
	return func(g *GoFigOpt) option {
		prev := g.UseFile
		g.UseFile = h
		return UseFile(prev)
	}
}

func UseEnv(h bool) option {
	return func(g *GoFigOpt) option {
		prev := g.UseEnv
		g.UseEnv = h
		return UseEnv(prev)
	}
}
