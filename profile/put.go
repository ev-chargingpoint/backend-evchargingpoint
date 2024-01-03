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

func PutUser(idparam primitive.ObjectID, db *mongo.Database, r *http.Request) (err error) {

	namalengkap := r.FormValue("namalengkap")
	nomorhp := r.FormValue("nomorhp")
	namakendaraan := r.FormValue("namakendaraan")
	nomorpolisi := r.FormValue("nomorpolisi")

	image := r.FormValue("file")

	if namalengkap == "" || nomorhp == "" || namakendaraan == "" || nomorpolisi == "" {
		return fmt.Errorf("mohon untuk melengkapi data")
	}
	if image != "" {
		imageUrl = image
	} else {
		imageUrl, err = evcharging.SaveFileToGithub("dimasardnt6", "dimasardnt6@gmail.com", "push-image", "evchargingpoint", r)
		if err != nil {
			return fmt.Errorf("error save file: %s", err)
		}
		image = imageUrl
	}

	profile := bson.M{
		"namalengkap":   namalengkap,
		"nomorhp":       nomorhp,
		"namakendaraan": namakendaraan,
		"nomorpolisi":   nomorpolisi,
		"image":         image,
	}
	err = evcharging.UpdateOneDoc(idparam, db, "user", profile)
	if err != nil {
		return err
	}
	return nil
}
