package charge

import (
	"fmt"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"github.com/ev-chargingpoint/backend-evchargingpoint/profile"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProcessPayment(_id, iduser primitive.ObjectID, db *mongo.Database, insertedDoc evcharging.Charge) (bson.M, error) {
	charge, err := GetChargeFromID(_id, db)
	if err != nil {
		return bson.M{}, fmt.Errorf("charge tidak ditemukan")
	}

	user, err := profile.GetUserFromID(iduser, db)
	if err != nil {
		return bson.M{}, fmt.Errorf("user tidak ditemukan")
	}

	if charge.User.ID != user.ID {
		return bson.M{}, fmt.Errorf("user tidak sesuai")
	}

	if insertedDoc.InputPembayaran != charge.TotalPrice {
		return bson.M{}, fmt.Errorf("pembayaran tidak sesuai")
	}

	if insertedDoc.Payment {
		charge.Payment = true
	}

	data := bson.M{
		"starttime":       charge.StartTime,
		"endtime":         charge.EndTime,
		"totalkwh":        charge.TotalKWH,
		"totalprice":      charge.TotalPrice,
		"inputpembayaran": insertedDoc.InputPembayaran,
		"paymentmethod":   insertedDoc.PaymentMethod,
		"payment":         charge.Payment,
		"status":          charge.Status,
	}

	err = evcharging.UpdateOneDoc(_id, db, "charge_car", data)
	if err != nil {
		return bson.M{}, err
	}
	return data, nil
}

func ProcessStatus(_id, iduser primitive.ObjectID, db *mongo.Database, insertedDoc evcharging.Charge) (bson.M, error) {
	charge, err := GetChargeFromID(_id, db)
	if err != nil {
		return bson.M{}, fmt.Errorf("charge tidak ditemukan")
	}

	user, err := profile.GetUserFromID(iduser, db)
	if err != nil {
		return bson.M{}, fmt.Errorf("user tidak ditemukan")
	}

	if charge.User.ID != user.ID {
		return bson.M{}, fmt.Errorf("user tidak sesuai")
	}

	if insertedDoc.Status {
		charge.Status = true
	}

	data := bson.M{
		"starttime":       charge.StartTime,
		"endtime":         charge.EndTime,
		"totalkwh":        charge.TotalKWH,
		"totalprice":      charge.TotalPrice,
		"inputpembayaran": charge.InputPembayaran,
		"paymentmethod":   charge.PaymentMethod,
		"payment":         charge.Payment,
		"status":          charge.Status,
	}

	err = evcharging.UpdateOneDoc(_id, db, "charge_car", data)
	if err != nil {
		return bson.M{}, err
	}
	return data, nil
}
