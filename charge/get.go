package charge

import (
	"context"
	"errors"
	"fmt"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"github.com/ev-chargingpoint/backend-evchargingpoint/charging_station"
	"github.com/ev-chargingpoint/backend-evchargingpoint/profile"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetChargeFromID(_id primitive.ObjectID, db *mongo.Database) (charge evcharging.Charge, err error) {
	collection := db.Collection("charge_car")
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.Background(), filter).Decode(&charge)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return charge, fmt.Errorf("charge dengan ID %s tidak ditemukan", _id)
		}
		return charge, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	chargingstation, err := charging_station.GetChargingStationFromID(charge.ChargingStation.ID, db)
	if err != nil {
		return charge, fmt.Errorf("charging station tidak ditemukan")
	}
	user, err := profile.GetUserFromID(charge.User.ID, db)
	if err != nil {
		return charge, fmt.Errorf("user tidak ditemukan")
	}
	charge.ChargingStation = chargingstation
	dataUser := evcharging.User{
		ID:            user.ID,
		NamaLengkap:   user.NamaLengkap,
		NomorHp:       user.NomorHp,
		NamaKendaraan: user.NamaKendaraan,
		NomorPolisi:   user.NomorPolisi,
		Email:         user.Email,
	}
	charge.User = dataUser
	return charge, nil
}

func GetAllChargeByUser(iduser primitive.ObjectID, db *mongo.Database) (charge []evcharging.Charge, err error) {
	collection := db.Collection("charge_car")
	filter := bson.M{"user._id": iduser}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return charge, fmt.Errorf("error GetAllChargeByUser: %s", err)
	}
	err = cursor.All(context.Background(), &charge)
	if err != nil {
		return charge, err
	}

	for i, data := range charge {
		chargingstation, err := charging_station.GetChargingStationFromID(data.ChargingStation.ID, db)
		if err != nil {
			return charge, fmt.Errorf("charging station tidak ditemukan")
		}
		user, err := profile.GetUserFromID(data.User.ID, db)
		if err != nil {
			return charge, fmt.Errorf("user tidak ditemukan")
		}
		data.ChargingStation = chargingstation
		dataUser := evcharging.User{
			ID:            user.ID,
			NamaLengkap:   user.NamaLengkap,
			NomorHp:       user.NomorHp,
			NamaKendaraan: user.NamaKendaraan,
			NomorPolisi:   user.NomorPolisi,
			Email:         user.Email,
		}
		data.User = dataUser
		charge[i] = data
	}

	return charge, nil
}
