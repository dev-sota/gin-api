package main

import (
	router "github.com/dev-sota/gin-api/infrastructure"
	"github.com/dev-sota/gin-api/infrastructure/database"
)

func main() {
	database.Init()
	router.Init()
	database.Close()
}
