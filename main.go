package main

import (
	"github.com/mikroblog-be/config"
	"github.com/mikroblog-be/helpers"
)

func main() {
	//Set logger
	helpers.Logger = config.InitLogger()

	//Set firebase
	app := helpers.NewFirebaseApp()
	app.App, app.Context = config.InitFirebase()

	//Add new document
	data1 := map[string]interface{}{
		"name":    "New York",
		"country": "America",
	}

	err := app.Add("cities", "1", data1)
	if err != nil {
		helpers.Logger.Fatal(err)
	}
}
