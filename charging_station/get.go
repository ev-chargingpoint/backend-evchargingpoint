package charging_station

import (
	"context"
	"errors"
	"fmt"

	// "net/http"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllChargingStation(db *mongo.Database) (charging []evcharging.ChargingSatation, err error) {
	collection := db.Collection("chargingstation")
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return charging, fmt.Errorf("error GetAllChargingStation: %s", err)
	}
	err = cursor.All(context.Background(), &charging)
	if err != nil {
		return charging, err
	}
	return charging, nil
}

// func GetChargingStation(db *mongo.Database, r *http.Request) (charging []bson.M, err error) {
// 	chargingstation, err := GetAllChargingStation(db)
// 	if err != nil {
// 		return charging, err
// 	}
// 	for _, c := range chargingstation {
// 		charging = append(charging, bson.M{
// 			"_id":            c.ID,
// 			"chargingkode":   c.ChargingKode,
// 			"nama":           c.Nama,
// 			"alamat":         c.Alamat,
// 			"nomortelepon":   c.NomorTelepon,
// 			"daya":           c.Daya,
// 			"connector":      c.Connector,
// 			"harga":          c.Harga,
// 			"jamoperasional": c.JamOperasional,
// 			"longitude":      c.Longitude,
// 			"latitude":       c.Latitude,
// 		})
// 	}
// 	return charging, nil
// }

func GetChargingStationFromID(_id primitive.ObjectID, db *mongo.Database) (doc evcharging.ChargingSatation, err error) {
	collection := db.Collection("chargingstation")
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.Background(), filter).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return doc, fmt.Errorf("no data found for ID %s", _id)
		}
		return doc, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return doc, nil
}
