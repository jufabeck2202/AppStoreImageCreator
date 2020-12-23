package cmd

import (
	"flag"
	"fmt"
	"github.com/jufabeck2202/AppStoreImageCreator/core"
	"github.com/jufabeck2202/AppStoreImageCreator/server"
	"github.com/muesli/termenv"
)

func Execute() {
	startServer := flag.Bool("server", false, "start web server")
	flag.Parse()
	fmt.Println("server:", *startServer)

	if *startServer {
		server.StartServer()
	} else {
		//core.CreateGradient(80000, 80000)
		core.StartConcat(false)
		out := termenv.String("Hello World")
		fmt.Println(out)
	}
}
