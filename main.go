package main

import (
	"go-gin-codef-api/src/router"
)

func main() {
	r := router.Router()
	r.Run(":9999")
}
