package signup

import (
	"encoding/json"
	"net/http"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	response evcharging.Response
	user     evcharging.User
)

func Post(MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := evcharging.MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	email, err := SignUpUser(conn, user)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Berhasil daftar akun dengan email :" + user.Email
	responseData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data": bson.M{
			"email": email,
		},
	}
	return evcharging.GCFReturnStruct(responseData)
}
