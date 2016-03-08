package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"route"
	"runtime"
)

func main() {
	ConfigRuntime()

	//StartWorkers()

	StartListening()

}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartWorkers() {
	//go statsWorker()
	//workers.StartWroker()
}

func StartListening() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	route.InitRoute(router)

	fmt.Println("start")

	router.Run(":8888")
}
