package cmd

import (
	"flag"
	"fmt"

	"github.com/jufabeck2202/AppStoreImageCreator/core"
	"github.com/jufabeck2202/AppStoreImageCreator/server"
)

func Execute() {
	startServer := flag.Bool("server", false, "start web server")
	flag.Parse()
	fmt.Println("server:", *startServer)

	if *startServer {
		server.StartServer()
	} else {
		core.CreateGradient(80000, 80000)
	}
}
