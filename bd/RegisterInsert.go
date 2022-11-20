package bd

import (
	"context"
	"time"

	"github.com/oscaralcalde/twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterInsert(usuario models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	usuario.Password, _ = EncryptPassword(usuario.Password)

	result, err := col.InsertOne(ctx, usuario)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
