package main

import (
	"github.com/mikroblog/config"
	"github.com/mikroblog/helpers"
)

func main() {
	//Set logger
	helpers.Logger = config.InitLogger()

	//Set firebase
	app := helpers.NewFirebaseApp()
	app.App, app.Context = config.InitFirebase()

	//Add new document
	data := map[string]interface{}{
		"name":    "New York",
		"country": "America",
	}

	app.Add("cities", "1", data)

	//Read all documents
	app.ReadAll("cities")

	//Read document by key
	app.Read("cities", "1")

	//Read document with condition
	app.ReadWhere("cities", "country", "==", interface{}("America"))

	//Delete document column
	for _, docId := range helpers.ArrayKeys(app.ReadAll("cities")) {
		app.DeleteColumn("cities", docId, "country")
	}

	//Delete document
	for _, docId := range helpers.ArrayKeys(app.ReadAll("cities")) {
		app.Delete("cities", docId)
	}
}
