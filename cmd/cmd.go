package cmd

import (
	"flag"
	"fmt"
	"image/color"

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

		// retrieve color profile supported by terminal
		p := termenv.ColorProfile()

		// supports hex values
		// will automatically degrade colors on terminals not supporting RGB
		out = out.Foreground(p.Color("#abcdef"))
		// but also supports ANSI colors (0-255)
		out = out.Background(p.Color("69"))
		// ...or the color.Color interface
		out = out.Foreground(p.FromColor(color.RGBA{255, 128, 0, 255}))

		fmt.Println(out)
	}
}
