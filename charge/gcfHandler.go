package charge

import (
	"encoding/json"
	"net/http"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var response evcharging.Response

func GetChargeHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := evcharging.MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	user, err := evcharging.GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}

	id := evcharging.GetID(r)
	if id == "" {
		charge, err := GetAllChargeByUser(user.Id, conn)
		if err != nil {
			response.Message = err.Error()
			return evcharging.GCFReturnStruct(response)
		}
		responData := bson.M{
			"status":  200,
			"message": "Berhasil mendapatkan data",
			"data":    charge,
		}
		return evcharging.GCFReturnStruct(responData)
	}
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid id parameter"
		return evcharging.GCFReturnStruct(response)
	}

	charge, err := GetChargeFromID(idparam, conn)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Berhasil mendapatkan data"
	responseData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    charge,
	}
	return evcharging.GCFReturnStruct(responseData)
}

func ChargeHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := evcharging.MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	err := json.NewDecoder(r.Body).Decode(&evcharging.Charge{})
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return evcharging.GCFReturnStruct(response)
	}

	user, err := evcharging.GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	id := evcharging.GetID(r)
	if id == "" {
		response.Message = "Invalid id parameter"
		return evcharging.GCFReturnStruct(response)
	}
	idchargingstation, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid id parameter"
		return evcharging.GCFReturnStruct(response)
	}
	data, err := ChargeCar(idchargingstation, user.Id, conn, evcharging.Charge{})
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	response.Status = 201
	response.Message = "Berhasil melakukan charge"
	responseData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}
	return evcharging.GCFReturnStruct(responseData)
}

func PaymentAndStatusChargeHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := evcharging.MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	var chargeData evcharging.Charge
	err := json.NewDecoder(r.Body).Decode(&chargeData)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return evcharging.GCFReturnStruct(response)
	}

	user, err := evcharging.GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}

	id := evcharging.GetID(r)
	if id == "" {
		response.Message = "Invalid id parameter"
		return evcharging.GCFReturnStruct(response)
	}

	idcharge, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid id parameter"
		return evcharging.GCFReturnStruct(response)
	}

	var data bson.M
	if chargeData.Payment {
		data, err = ProcessPayment(idcharge, user.Id, conn, chargeData)
		if err != nil {
			response.Message = err.Error()
			return evcharging.GCFReturnStruct(response)
		}
		response.Message = "Proses pembayaran berhasil"
	} else if chargeData.Status {
		data, err = ProcessStatus(idcharge, user.Id, conn, chargeData)
		if err != nil {
			response.Message = err.Error()
			return evcharging.GCFReturnStruct(response)
		}
		response.Message = "Pengecasan selesai"
	} else {
		response.Message = "Invalid condition: Payment or Status should be true"
		return evcharging.GCFReturnStruct(response)
	}

	response.Status = 201
	responseData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}
	return evcharging.GCFReturnStruct(responseData)
}
