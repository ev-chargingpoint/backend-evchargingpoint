package profile

import (
	"fmt"
	"net/http"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var imageUrl string

func PutProfile(idparam primitive.ObjectID, db *mongo.Database, r *http.Request) (bson.M, error) {
	dataUser, err := GetUserFromID(idparam, db)
	if err != nil {
		return bson.M{}, err
	}
	namalengkap := r.FormValue("namalengkap")
	nomorhp := r.FormValue("nomorhp")
	namakendaraan := r.FormValue("namakendaraan")
	nomorpolisi := r.FormValue("nomorpolisi")

	image := r.FormValue("file")

	if namalengkap == "" || nomorhp == "" || namakendaraan == "" || nomorpolisi == "" {
		return bson.M{}, fmt.Errorf("mohon untuk melengkapi data")
	}
	if image != "" {
		imageUrl = image
	} else {
		imageUrl, err = evcharging.SaveFileToGithub("dimasardnt6", "dimasardnt6@gmail.com", "push-image", "evchargingpoint", r)
		if err != nil {
			return bson.M{}, fmt.Errorf("error save file: %s", err)
		}
	}

	profile := bson.M{
		"namalengkap":   namalengkap,
		"nomorhp":       nomorhp,
		"namakendaraan": namakendaraan,
		"nomorpolisi":   nomorpolisi,
		"email":         dataUser.Email,
		"password":      dataUser.Password,
		"image":         imageUrl,
	}
	err = evcharging.UpdateOneDoc(idparam, db, "user", profile)
	if err != nil {
		return bson.M{}, err
	}
	return nil, err
}
