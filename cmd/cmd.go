package cmd

import (
	"flag"
	"fmt"
	"github.com/jufabeck2202/AppStoreImageCreator/server"
	"github.com/jufabeck2202/AppStoreImageCreator/core"
)

func Execute() {
	startServer := flag.Bool("server", false, "start web server")
	generate := flag.Bool("generate", false, "start web server")
	frames := flag.Bool("frames", false, "start web server")

	flag.Parse()
	fmt.Println("server:", *startServer)

	if *startServer {
		server.StartServer()
	} else if *generate {
		core.GenerateImages()
	} else if *frames {
		core.GenerateTestFrames()
	}

}
