package share

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

func GetRef(ctx context.Context) *db.Ref {

	conf := &firebase.Config{
		DatabaseURL: DatabaseURL,
	}
	opt := option.WithCredentialsJSON([]byte(Secret))
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("error in creating firebase DB client: ", err)
	}
	return client.NewRef("datas")
}
