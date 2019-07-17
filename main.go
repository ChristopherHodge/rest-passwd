/*
rest-passwd/main

Read configuration and start server
*/
package main

import (
  "rest-passwd/config"
  "rest-passwd/restapi"
)

// The relative path to the config file
const CONFIG_FILE = "./config.json"

func main() {
  // read json configuration, if available,
  // or fallback to config constants
  config := config.Read(CONFIG_FILE)

	// start server
	server := restapi.Server{}
	server.Listen(config.Port) // block indefinitely
}

