module main

go 1.16

require (
	deposit v0.0.0
	github.com/gin-gonic/gin v1.7.2
)

replace (
	config v0.0.0 => ./config
	db v0.0.0 => ./db
	deposit v0.0.0 => ./deposit
)
