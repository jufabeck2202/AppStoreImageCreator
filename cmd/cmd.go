package cmd

import (
	"flag"
	"fmt"
	"github.com/jufabeck2202/AppStoreImageCreator/core"
	"github.com/jufabeck2202/AppStoreImageCreator/server"
	"github.com/muesli/termenv"
	"sync"
)

func Execute() {
	startServer := flag.Bool("server", false, "start web server")
	flag.Parse()
	fmt.Println("server:", *startServer)

	if *startServer {
		server.StartServer()
	} else {
		var frameswg sync.WaitGroup
		core.AddFrame(&frameswg, "./beju.jpeg", "test", "#aabbcc", "#aabbcc" )
		frameswg.Wait()
		out := termenv.String("Hello World")
		fmt.Println(out)
	}
}
