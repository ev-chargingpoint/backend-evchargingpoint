package charging_station

import (
	"fmt"
	"net/http"
	"strconv"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var imageUrl string

func PutChargingStationOlehAdmin(_id primitive.ObjectID, db *mongo.Database, r *http.Request) (bson.M, error) {
	chargingkode := r.FormValue("chargingkode")
	nama := r.FormValue("nama")
	alamat := r.FormValue("alamat")
	nomortelepon := r.FormValue("nomortelepon")
	daya := r.FormValue("daya")
	connector := r.FormValue("connector")
	image := r.FormValue("file")
	harga := r.FormValue("harga")
	jamoperasional := r.FormValue("jamoperasional")
	longitude := r.FormValue("longitude")
	latitude := r.FormValue("latitude")
	ammountplugsStr := r.FormValue("ammountplugs")
	ammountplugs, err := strconv.Atoi(ammountplugsStr)
	if err != nil {
		return bson.M{}, fmt.Errorf("ammountplugs should be an integer: %s", err)
	}
	available := ammountplugs

	if chargingkode == "" || nama == "" || alamat == "" || nomortelepon == "" || ammountplugs == 0 || daya == "" || connector == "" || harga == "" || jamoperasional == "" || latitude == "" || longitude == "" || available == 0 {
		return bson.M{}, fmt.Errorf("mohon untuk melengkapi data")
	}
	if image != "" {
		imageUrl = image
	} else {
		imageUrl, err := evcharging.SaveFileToGithub("dimasardnt6", "dimasardnt6@gmail.com", "push-image", "evchargingpoint", r)
		if err != nil {
			return bson.M{}, fmt.Errorf("error save file: %s", err)
		}
		image = imageUrl
	}

	chargingstation := bson.M{
		"chargingkode":   chargingkode,
		"nama":           nama,
		"alamat":         alamat,
		"nomortelepon":   nomortelepon,
		"ammountplugs":   ammountplugs,
		"available":      available,
		"daya":           daya,
		"connector":      connector,
		"image":          image,
		"harga":          harga,
		"jamoperasional": jamoperasional,
		"longitude":      longitude,
		"latitude":       latitude,
	}
	err = evcharging.UpdateOneDoc(_id, db, "chargingstation", chargingstation)
	if err != nil {
		return bson.M{}, err
	}
	return chargingstation, nil
}
