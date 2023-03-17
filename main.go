package main

import (
	"github.com/gaurishhs/uptimo/config"
	"github.com/gaurishhs/uptimo/server"
)
 
var Version = "(unknown)"

func main() {
  // if len(os.Args) > 0 && os.Args[1] == "version" {
	// 	fmt.Printf("Version: %s", Version)
	// 	os.Exit(0)
	// }

  apiServer := server.NewServer(nil, &config.Config{
		Debug: true,
	})

	apiServer.Start()
} 
