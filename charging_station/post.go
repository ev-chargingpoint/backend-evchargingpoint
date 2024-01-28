package charging_station

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckLatitudeLongitude(db *mongo.Database, latitude, longitude string) bool {
	collection := db.Collection("chargingstation")
	filter := bson.M{"latitude": latitude, "longitude": longitude}
	err := collection.FindOne(context.Background(), filter).Decode(&evcharging.ChargingStation{})
	return err == nil
}

func CheckChargingKode(db *mongo.Database, chargingkode string) bool {
	collection := db.Collection("chargingstation")
	filter := bson.M{"chargingkode": chargingkode}
	err := collection.FindOne(context.Background(), filter).Decode(&evcharging.ChargingStation{})
	return err == nil
}

// By Admin
func PostChargingStationOlehAdmin(db *mongo.Database, r *http.Request) (bson.M, error) {
	chargingkode := r.FormValue("chargingkode")
	nama := r.FormValue("nama")
	alamat := r.FormValue("alamat")
	nomortelepon := r.FormValue("nomortelepon")
	daya := r.FormValue("daya")
	connector := r.FormValue("connector")
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
	if CheckChargingKode(db, chargingkode) {
		return bson.M{}, fmt.Errorf("kode sudah ada")
	}
	if CheckLatitudeLongitude(db, latitude, longitude) {
		return bson.M{}, fmt.Errorf("charging station sudah terdaftar")
	}

	imageUrl, err := evcharging.SaveFileToGithub("dimasardnt6", "dimasardnt6@gmail.com", "push-image", "evchargingpoint", r)
	if err != nil {
		return bson.M{}, fmt.Errorf("error save file: %s", err)
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
		"image":          imageUrl,
		"harga":          harga,
		"jamoperasional": jamoperasional,
		"longitude":      longitude,
		"latitude":       latitude,
	}
	_, err = evcharging.InsertOneDoc(db, "chargingstation", chargingstation)
	if err != nil {
		return bson.M{}, err
	}
	return chargingstation, nil
}
