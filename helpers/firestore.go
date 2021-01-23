package helpers

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

type firebaseAppInterface interface {
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

func (this firebaseApp) Add(collection string, doc string, data map[string]interface{}) error {
	client, clientErr := this.App.Firestore(this.Context)
	defer client.Close()
	if clientErr != nil {
		return clientErr
	}

	_, err := client.Collection(collection).Doc(doc).Set(this.Context, data)
	if err != nil {
		return err
	}

	Logger.Printf("inserted : %s", ToJSON(data))

	return nil
}

func (this firebaseApp) Read(collection string, doc string) (map[string]map[string]interface{}, error) {
	client, clientErr := this.App.Firestore(this.Context)
	defer client.Close()
	if clientErr != nil {
		log.Fatalln(clientErr)
	}

	dsnap, err := client.Collection(collection).Doc(doc).Get(this.Context)
	if err != nil {
		return nil, err
	}

	var temp map[string]interface{}
	dsnap.DataTo(&temp)

	data := make(map[string]map[string]interface{})
	data[dsnap.Ref.ID] = temp

	Logger.Printf("read : %s", ToJSON(data))

	return data, nil
}

func (this firebaseApp) ReadAll(collection string) (map[string]map[string]interface{}, error) {
	client, clientErr := this.App.Firestore(this.Context)
	defer client.Close()
	if clientErr != nil {
		return nil, clientErr
	}

	docs := make(map[string]map[string]interface{})
	iter := client.Collection(collection).Documents(this.Context)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		docs[doc.Ref.ID] = doc.Data()
	}

	Logger.Printf("read all : %s", ToJSON(docs))

	return docs, nil
}

func (this firebaseApp) ReadWhere(collection string, field string, op string, value interface{}) (map[string]map[string]interface{}, error) {
	client, clientErr := this.App.Firestore(this.Context)
	defer client.Close()
	if clientErr != nil {
		return nil, clientErr
	}

	docs := make(map[string]map[string]interface{})
	iter := client.Collection(collection).Where(field, op, value).Documents(this.Context)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		docs[doc.Ref.ID] = doc.Data()
	}

	Logger.Printf("read where : %s", ToJSON(docs))
	return docs, nil
}

func (this firebaseApp) Delete(collection string, doc string) error {
	client, clientErr := this.App.Firestore(this.Context)
	defer client.Close()
	if clientErr != nil {
		return clientErr
	}

	_, err := client.Collection(collection).Doc(doc).Delete(this.Context)
	if err != nil {
		return err
	}

	Logger.Printf("deleted : %s -> %s", collection, doc)

	return nil
}

func (this firebaseApp) DeleteColumn(collection string, doc string, column string) error {
	client, clientErr := this.App.Firestore(this.Context)
	defer client.Close()
	if clientErr != nil {
		return clientErr
	}

	_, err := client.Collection(collection).Doc(doc).Update(this.Context, []firestore.Update{
		{
			Path:  column,
			Value: firestore.Delete,
		},
	})
	if err != nil {
		return err
	}

	Logger.Printf("deleted column : %s -> %s -> %s", collection, doc, column)

	return nil
}
