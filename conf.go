package gofigure

// conf is actually how you configure gofigure through the magic of
// self-referential functions. An implemention that is a shameless duplicate of
// [post by Rob Pike][pike-blogpost]
//
// Example:
// figOpt := new(GoFigOpt)
// figOpt.Option(gofigure.MongoHosts("localhost"))
// // load a your local test configuration
// prevMongos := figOpt.Option(gofigure.MongoHosts("p-mongo-1.go.com,p-mongo-2.go.com"))
// // process a prod config for sillies
// figOpt.Option(prevMongos) // now your mongo server will be set to "localhost" again
//
// [pike-blogpost]: http://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html

var Settings GoFigOpt

func init() {
	Settings = GoFigOpt{
		MongoDBHosts:      "localhost",
		MongoDBName:       "gofigure",
		MongoDBCollection: "default",
		FileLocations:     []string{"./gofigure.json", "./.gofigure.json"},
	}
}

type GoFigOpt struct {
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
