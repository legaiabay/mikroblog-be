package testing

import (
	"testing"

	"github.com/mikroblog-be/config"
	"github.com/mikroblog-be/helpers"
)

func TestFirestoreAdd(t *testing.T) {
	//Set firebase
	app := helpers.NewFirebaseApp()
	app.App, app.Context = config.InitFirebase()

	//Add new document
	data := map[string]interface{}{
		"name":    "New York",
		"country": "America",
	}

	err := app.Add("cities", "1", data)
	if err != nil {
		t.Errorf("Add FAILED, expected %v, got %v", nil, err)
	} else {
		t.Log("Add PASSED")
	}
}

func TestFirestoreReadAll(t *testing.T) {
	//Set firebase
	app := helpers.NewFirebaseApp()
	app.App, app.Context = config.InitFirebase()

	//Read all documents
	_, err := app.ReadAll("cities")
	if err != nil {
		t.Errorf("ReadAll FAILED, expected %v, got %v", nil, err)
	} else {
		t.Log("ReadAll PASSED")
	}
}

func TestFirestoreRead(t *testing.T) {
	//Set firebase
	app := helpers.NewFirebaseApp()
	app.App, app.Context = config.InitFirebase()

	//Read document by key
	_, err := app.Read("cities", "1")
	if err != nil {
		t.Errorf("Read FAILED, expected %v, got %v", nil, err)
	} else {
		t.Log("Read PASSED")
	}
}

func TestFirestoreReadWhere(t *testing.T) {
	//Set firebase
	app := helpers.NewFirebaseApp()
	app.App, app.Context = config.InitFirebase()

	//Read document with condition
	_, err := app.ReadWhere("cities", "country", "==", interface{}("America"))
	if err != nil {
		t.Errorf("ReadWhere FAILED, expected %v, got %v", nil, err)
	} else {
		t.Log("ReadWhere PASSED")
	}
}

func TestFirestoreDeleteColumn(t *testing.T) {
	//Set firebase
	app := helpers.NewFirebaseApp()
	app.App, app.Context = config.InitFirebase()

	//Delete document column
	data, err := app.Read("cities", "1")
	if err != nil {
		t.Errorf("ReadAll() FAILED, expected %v, got %v", nil, err)
	}
	for _, docId := range helpers.ArrayKeys(data) {
		err = app.DeleteColumn("cities", docId, "country")
		if err != nil {
			t.Errorf("DeleteColumn %v FAILED, expected %v, got %v", docId, nil, err)
			break
		} else {
			t.Logf("DeleteColumn %v PASSED", docId)
		}
	}
}

func TestFirestoreDelete(t *testing.T) {
	//Set firebase
	app := helpers.NewFirebaseApp()
	app.App, app.Context = config.InitFirebase()

	//Delete document
	data, err := app.Read("cities", "1")
	if err != nil {
		t.Errorf("ReadAll() FAILED, expected %v, got %v", nil, err)
	}
	for _, docId := range helpers.ArrayKeys(data) {
		err = app.Delete("cities", docId)
		if err != nil {
			t.Errorf("DeleteColumn() FAILED, expected %v, got %v", nil, err)
			break
		} else {
			t.Logf("DeleteColumn %v PASSED", docId)
		}
	}
}
