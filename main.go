package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kube-tarian/compage-template-go/pkg/controller"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sinhashubham95/go-actuator"
	log "github.com/sirupsen/logrus"
	"os"
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	router := gin.Default()
	// get the handler for actuator
	actuatorHandler := actuator.GetActuatorHandler(&actuator.Config{Endpoints: []int{
		actuator.Env,
		actuator.Info,
		actuator.Metrics,
		actuator.Ping,
		//actuator.Shutdown,
		actuator.ThreadDump,
	},
		Env:     "dev",
		Name:    "compage-template-go",
		Port:    8080,
		Version: "0.0.1",
	})
	ginActuatorHandler := func(ctx *gin.Context) {
		actuatorHandler(ctx.Writer, ctx.Request)
	}
	router.GET("/actuator/*endpoint", ginActuatorHandler)
	router.GET("/metrics", prometheusHandler())
	v1 := router.Group("/v1")
	{
		v1.GET("/users/:id", controller.GetUser)
		v1.POST("/users", controller.CreateUser)
		v1.PUT("/users/:id", controller.UpdateUser)
		v1.DELETE("/users/:id", controller.DeleteUser)
		v1.GET("/users", controller.ListUsers)
		//v1.PATCH("/:id", patching)
		//v1.HEAD("/", head)
		//v1.OPTIONS("/", options)
	}

	// call external client here if the isClient value is true

	Port := ":8080"
	if err := router.Run(Port); err != nil {
		return
	}
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}
