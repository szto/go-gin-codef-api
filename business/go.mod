module business

go 1.16

require (
	codef v0.0.0
	github.com/gin-gonic/gin v1.7.2
)

replace (
	codef v0.0.0 => ../codef
	config v0.0.0 => ../config
)
