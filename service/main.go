package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
)

func initSecrets() error {
    // .secrets.env will contain any api access tokens
    return godotenv.Load(".secrets.env")
}

func main() {

	// Load secrets
	err := initSecrets()
    if err != nil {
        fmt.Println("Cannot initialize the service - %w", err)
        return
    }

    // initialise gofr object
    app := gofr.New()
	
    // register route greet
    app.GET("/chat", func(ctx *gofr.Context) (interface{}, error) {

        return "Hello World!", nil
    })

	app.GET("/googleGroup", func(ctx *gofr.Context) (interface{}, error) {
		return GoogleGroupHandler( ctx )
	} )

	app.POST("/googleGroupReplay", func(ctx *gofr.Context) (interface{}, error) {
		return GoogleGroupReplay( ctx )
	} )

	app.POST("/tweet", func(ctx *gofr.Context) (interface{}, error) {
		return TweetHandler( ctx )
	} )

    // Runs the server, it will listen on the default port 8000.
    // it can be over-ridden through configs
   app.Run()
}