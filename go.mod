module main

go 1.16

require (
	db v0.0.0
	github.com/gin-gonic/gin v1.7.2
	go.mongodb.org/mongo-driver v1.5.4
)

replace (
	config v0.0.0 => ./config
	db v0.0.0 => ./db
)
