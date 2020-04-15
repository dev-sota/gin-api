package main

import (
	"github.com/dev-sota/gin-api/infrastructure/database"
	"github.com/dev-sota/gin-api/infrastructure/router"
)

func main() {
	database.Init()
	router.Init()
	database.Close()
}
