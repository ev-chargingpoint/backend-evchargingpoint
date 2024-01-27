package charging_station

import (
	"context"
	"errors"
	"fmt"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllChargingStation(db *mongo.Database) (charging []evcharging.ChargingStation, err error) {
	collection := db.Collection("chargingstation")
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return charging, fmt.Errorf("error GetAllChargingStation: %s", err)
	}

	for cursor.Next(context.Background()) {
		var station evcharging.ChargingStation
		err := cursor.Decode(&station)
		if err != nil {
			return charging, fmt.Errorf("error decoding charging station: %s", err)
		}

		count, err := db.Collection("charge_car").CountDocuments(context.TODO(), bson.M{
			"chargingstation._id": station.ID,
			"payment":             true,
			"status":              false,
		})
		if err != nil {
			return charging, fmt.Errorf("failed to count charge transactions: %s", err.Error())
		}

		station.Available = station.AmmountPlugs - int(count)

		_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": station.ID}, bson.M{"$set": bson.M{"available": station.Available}})
		if err != nil {
			return charging, fmt.Errorf("failed to update available count for station ID %s: %s", station.ID, err.Error())
		}

		charging = append(charging, station)
	}

	return charging, nil
}

func GetChargingStationFromID(_id primitive.ObjectID, db *mongo.Database) (doc evcharging.ChargingStation, err error) {
	collection := db.Collection("chargingstation")
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.Background(), filter).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return doc, fmt.Errorf("no data found for ID %s", _id)
		}
		return doc, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}

	count, err := db.Collection("charge_car").CountDocuments(context.TODO(), bson.M{
		"chargingstation._id": doc.ID,
		"payment":             true,
		"status":              false,
	})
	if err != nil {
		return doc, fmt.Errorf("failed to count charge transactions: %s", err.Error())
	}

	doc.Available = doc.AmmountPlugs - int(count)

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": doc.ID}, bson.M{"$set": bson.M{"available": doc.Available}})
	if err != nil {
		return doc, fmt.Errorf("failed to update available count for station ID %s: %s", doc.ID, err.Error())
	}

	return doc, nil
}
