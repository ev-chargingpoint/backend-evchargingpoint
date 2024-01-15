package profile

import (
	"net/http"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	response evcharging.Response
)

func Get(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := evcharging.MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	payload, err := evcharging.GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	user, err := GetUserFromID(payload.Id, conn)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Get Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data": bson.M{
			"_id":           user.ID,
			"nama_lengkap":  user.NamaLengkap,
			"email":         user.Email,
			"nomorhp":       user.NomorHp,
			"namakendaraan": user.NamaKendaraan,
			"nomorpolisi":   user.NomorPolisi,
		},
	}
	return evcharging.GCFReturnStruct(responData)
}

func Put(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := evcharging.MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	user, err := evcharging.GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = "Gagal Decode Token : " + err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	// err = json.NewDecoder(r.Body).Decode(&user)
	// if err != nil {
	// 	response.Message = "error parsing application/json: " + err.Error()
	// 	return evcharging.GCFReturnStruct(response)
	// }
	data, err := PutProfile(user.Id, conn, r)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	response.Status = 200
	response.Message = "Berhasil mengubah profile"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}
	return evcharging.GCFReturnStruct(responData)
}
