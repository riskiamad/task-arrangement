package main

import (
	"log"
	"task-scheduler/config"
	"task-scheduler/router"
)

func init() {
	if err := config.DbSetup(); err != nil {
		panic(err)
	}

	if err := config.AutoMigrateDB(); err != nil {
		panic(err)
	}
}

func main() {
	g := router.Router()

	if err := g.Run(config.Config.Port); err != nil {
		panic(err)
	}

	log.Println("server run at ", config.Config.Host)
}
