module db

go 1.16

require (
	config v0.0.0
	go.mongodb.org/mongo-driver v1.5.4
)

replace config v0.0.0 => ../config
