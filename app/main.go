package main

import (
	router "github.com/dev-sota/gin-api/app/infrastructure"
	"github.com/dev-sota/gin-api/app/infrastructure/database"
)

func main() {
	database.Init()
	router.Init()
	database.Close()
}
