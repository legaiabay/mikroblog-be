package helpers

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

type firebaseAppInterface interface {
	NewFirebaseApp()
	Add(collection string, doc string, data map[string]interface{}) error
	Read(collection string, doc string)
	ReadAll(collection string) (map[int]map[string]interface{}, error)
	ReadWhere(collection string, field string, op string, value interface{}) (map[int]map[string]interface{}, error)
	Delete(collection string, doc string) error
	DeleteColumn(collection string, doc string, column string) error
}

type firebaseApp struct {
	App     *firebase.App
	Context context.Context
}

func NewFirebaseApp() firebaseApp {
	app := firebaseApp{}

	return app
}

func (this firebaseApp) Add(collection string, doc string, data map[string]interface{}) {
	client, clientErr := this.App.Firestore(this.Context)
	defer client.Close()
	if clientErr != nil {
		log.Fatalln(clientErr)
		return
	}

	_, err := client.Collection(collection).Doc(doc).Set(this.Context, data)

	if err != nil {
		Logger.Fatalf("Failed adding alovelace: %v", err)
		return
	}

	Logger.Printf("inserted : %s", ToJSON(data))
}

func (this firebaseApp) Read(collection string, doc string) map[string]map[string]interface{} {
	client, clientErr := this.App.Firestore(this.Context)
	defer client.Close()
	if clientErr != nil {
		log.Fatalln(clientErr)
		return nil
	}

	dsnap, err := client.Collection(collection).Doc(doc).Get(this.Context)
	if err != nil {
		return nil
	}

	var temp map[string]interface{}
	dsnap.DataTo(&temp)

	data := make(map[string]map[string]interface{})
	data[dsnap.Ref.ID] = temp

	Logger.Printf("read : %s", ToJSON(data))
	return data
}

func (this firebaseApp) ReadAll(collection string) map[string]map[string]interface{} {
	client, clientErr := this.App.Firestore(this.Context)
	defer client.Close()
	if clientErr != nil {
		log.Fatalln(clientErr)
		return nil
	}

	docs := make(map[string]map[string]interface{})
	iter := client.Collection(collection).Documents(this.Context)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			Logger.Fatalf("Failed to iterate: %v", err)
			return nil
		}

		docs[doc.Ref.ID] = doc.Data()
	}

	Logger.Printf("read all : %s", ToJSON(docs))
	return docs
}

func (this firebaseApp) ReadWhere(collection string, field string, op string, value interface{}) map[string]map[string]interface{} {
	client, clientErr := this.App.Firestore(this.Context)
	defer client.Close()
	if clientErr != nil {
		Logger.Fatalln(clientErr)
		return nil
	}

	docs := make(map[string]map[string]interface{})
	iter := client.Collection(collection).Where(field, op, value).Documents(this.Context)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil
		}

		docs[doc.Ref.ID] = doc.Data()
	}

	Logger.Printf("read where : %s", ToJSON(docs))
	return docs
}

func (this firebaseApp) Delete(collection string, doc string) {
	client, clientErr := this.App.Firestore(this.Context)
	defer client.Close()
	if clientErr != nil {
		Logger.Fatalln(clientErr)
		return
	}

	_, err := client.Collection(collection).Doc(doc).Delete(this.Context)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		Logger.Errorf("An error has occurred: %s", err)
		return
	}

	Logger.Printf("deleted : %s -> %s", collection, doc)
}

func (this firebaseApp) DeleteColumn(collection string, doc string, column string) {
	client, clientErr := this.App.Firestore(this.Context)
	defer client.Close()
	if clientErr != nil {
		Logger.Fatalln(clientErr)
		return
	}

	_, err := client.Collection(collection).Doc(doc).Update(this.Context, []firestore.Update{
		{
			Path:  column,
			Value: firestore.Delete,
		},
	})

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		Logger.Errorf("An error has occurred: %s", err)
		return
	}

	Logger.Printf("deleted column : %s -> %s -> %s", collection, doc, column)
}
