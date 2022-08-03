package rest

import (
	"GriffinBackend/rdb"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type GriffinWS struct {
	Conn     *gin.Engine
	Database *redis.Client
}

func (g GriffinWS) StartService() GriffinWS {
	// start web server connection
	g.Conn = gin.Default()
	// add database initialization
	d := rdb.Connect()
	g.Database = d
	return g
}

func (g GriffinWS) PingTest() GriffinWS {
	g.Conn.GET("/ping", pingPong)
	return g
}

func (g GriffinWS) Version() GriffinWS {
	g.Conn.GET("/version", version)
	return g
}

func (g GriffinWS) AddEmployee() GriffinWS {
	g.Conn.POST("/employee", func(c *gin.Context) {
		postEmployeePermanent(c, g.Database)
	})
	return g
}

func (g GriffinWS) GetEmployee() GriffinWS {
	g.Conn.GET("/employee", func(c *gin.Context) {
		getEmployee(c, g.Database)
	})
	return g
}
