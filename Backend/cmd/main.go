package main

import (
	"Backend/api"
)

func main() {
	server := api.SetupRouter()

	server.Run(":8085")
}
