package config

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"github.com/mikroblog/helpers"
	"google.golang.org/api/option"
)

func InitFirebase() (*firebase.App, context.Context) {
	ctx := context.Background()
	sa := option.WithCredentialsFile(os.Getenv("CREDENTIAL_FILE"))

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		helpers.Logger.Fatalln(err)
	}

	return app, ctx
}
