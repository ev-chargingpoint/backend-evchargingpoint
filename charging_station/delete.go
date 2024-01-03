package charging_station

import (
	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteChargingStationOlehAdmin(_id primitive.ObjectID, db *mongo.Database) error {
	err := evcharging.DeleteOneDoc(_id, db, "chargingstation")
	if err != nil {
		return err
	}
	return nil
}
