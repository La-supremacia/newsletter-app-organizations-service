package database

import (
	"fmt"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() error {
	mongoDBUri := os.Getenv("MONGODB_URI") //env.GoDotEnvVariable("MONGODB_URI")
	err := mgm.SetDefaultConfig(nil, "apps", options.Client().ApplyURI(mongoDBUri))
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to mongo on URI", mongoDBUri)
	return err
}
