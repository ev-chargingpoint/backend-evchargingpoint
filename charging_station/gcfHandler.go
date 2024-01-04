package charging_station

import (
	"net/http"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	response evcharging.Response
)

func Post(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := evcharging.MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	user, err := evcharging.GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	if user.Email != "admin@gmail.com" {
		response.Message = "Anda tidak memiliki akses"
		return evcharging.GCFReturnStruct(response)
	}
	data, err := PostChargingStationOlehAdmin(conn, r)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	//
	response.Status = 201
	response.Message = "Berhasil tambah charging station"
	responseData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}
	return evcharging.GCFReturnStruct(responseData)
}

func Put(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := evcharging.MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	user, err := evcharging.GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	if user.Email != "admin@gmail.com" {
		response.Message = "Anda tidak memiliki akses"
		return evcharging.GCFReturnStruct(response)
	}
	id := evcharging.GetID(r)
	if id == "" {
		response.Message = "Wrong parameter"
		return evcharging.GCFReturnStruct(response)
	}
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid id parameter"
		return evcharging.GCFReturnStruct(response)
	}

	data, err := PutChargingStationOlehAdmin(idparam, conn, r)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Berhasil update charging station"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}
	return evcharging.GCFReturnStruct(responData)
}

func GetChargingStationHandler(MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := evcharging.MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	id := evcharging.GetID(r)
	if id == "" {
		chargingstation, err := GetAllChargingStation(conn)
		if err != nil {
			response.Message = err.Error()
			return evcharging.GCFReturnStruct(response)
		}
		//
		return evcharging.GCFReturnStruct(chargingstation)
	}
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	chargingstation, err := GetChargingStationFromID(idparam, conn)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Berhasil get charging station"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    chargingstation,
	}
	return evcharging.GCFReturnStruct(responData)
}

func HapusChargingStationHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := evcharging.MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	user, err := evcharging.GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	if user.Email != "admin@gmail.com" {
		response.Message = "Anda tidak memiliki akses"
		return evcharging.GCFReturnStruct(response)
	}
	id := evcharging.GetID(r)
	if id == "" {
		response.Message = "Wrong parameter"
		return evcharging.GCFReturnStruct(response)
	}
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid id parameter"
		return evcharging.GCFReturnStruct(response)
	}
	err = DeleteChargingStationOlehAdmin(idparam, conn)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	//
	response.Status = 204
	response.Message = "Hapus Charging Station Berhasil"
	return evcharging.GCFReturnStruct(response)
}
