module deposit

go 1.16

require (
	db v0.0.0
	github.com/gin-gonic/gin v1.7.2
	go.mongodb.org/mongo-driver v1.6.0
)

replace (
	config v0.0.0 => ../config
	db v0.0.0 => ../db
)
