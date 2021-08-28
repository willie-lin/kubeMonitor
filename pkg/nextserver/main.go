package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {

	log, _ := zap.NewDevelopment()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	client, err := NewClient()
	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))
		return
	}

	fmt.Println("eeee")

	defer client.Close()
	ctx := context.Background()

	//autoMigration := database.AutoMigration
	autoMigration := AutoMigration
	autoMigration(client, ctx)

	//debugMode := database.DebugMode
	debugMode := DebugMode

	debugMode(err, client, ctx)

	e.GET("/", Hello(client))

	e.Logger.Fatal(e.Start(":2022"))

}
