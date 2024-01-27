package charge

import (
	"context"
	"fmt"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"github.com/ev-chargingpoint/backend-evchargingpoint/charging_station"
	"github.com/ev-chargingpoint/backend-evchargingpoint/profile"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ChargeCar(idchargingstation, iduser primitive.ObjectID, db *mongo.Database, insertedDoc evcharging.Charge) (bson.M, error) {
	available, err := CheckAvailable(db, idchargingstation)
	if err != nil {
		return bson.M{}, err
	}
	fmt.Println("Available:", available)
	if available {
		return bson.M{}, fmt.Errorf("tempat pengisian daya tidak tersedia")
	}
	if insertedDoc.StartTime == "" || insertedDoc.EndTime == "" {
		return bson.M{}, fmt.Errorf("mohon untuk melengkapi data")
	}
	user, err := profile.GetUserFromID(iduser, db)
	if err != nil {
		return bson.M{}, fmt.Errorf("user tidak ditemukan")
	}
	chargingstation, err := charging_station.GetChargingStationFromID(idchargingstation, db)
	if err != nil {
		return bson.M{}, fmt.Errorf("charging station tidak ditemukan")
	}
	charge := bson.M{
		"_id": primitive.NewObjectID(),
		"chargingstation": bson.M{
			"_id": chargingstation.ID,
		},
		"user": bson.M{
			"_id": user.ID,
		},
		"starttime":  insertedDoc.StartTime,
		"endtime":    insertedDoc.EndTime,
		"totalkwh":   insertedDoc.TotalKWH,
		"totalprice": insertedDoc.TotalPrice,
		"payment":    false,
		"status":     false,
	}
	_, err = evcharging.InsertOneDoc(db, "charge_car", charge)
	if err != nil {
		return bson.M{}, err
	}
	return charge, nil
}

func CheckAvailable(db *mongo.Database, idchargingstation primitive.ObjectID) (bool, error) {
	collection := db.Collection("chargingstation")
	filter := bson.M{"_id": idchargingstation}

	var chargingStation evcharging.ChargingStation
	err := collection.FindOne(context.Background(), filter).Decode(&chargingStation)
	if err != nil {
		return false, fmt.Errorf("failed to find charging station: %s", err.Error())
	}

	print("Available: ", chargingStation.Available)
	return chargingStation.Available < 0, nil
}
