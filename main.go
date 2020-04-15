package main

import "github.com/dev-sota/gin-api/infrastructure"

func main() {
	infrastructure.DBInit()
	infrastructure.RouterInit()
	infrastructure.Close()
}
