package login

import (
	"encoding/json"
	"net/http"
	"os"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	credential evcharging.Credential
	response   evcharging.Response
	user       evcharging.User
)

func Post(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := evcharging.MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	user, err := LogIn(conn, user)
	if err != nil {
		response.Message = err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	tokenstring, err := evcharging.Encode(user.ID, user.Email, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		response.Message = "Gagal Encode Token : " + err.Error()
		return evcharging.GCFReturnStruct(response)
	}
	//
	credential.Message = "Selamat Datang " + user.Email
	credential.Token = tokenstring
	credential.Status = 200
	responData := bson.M{
		"status":  credential.Status,
		"message": credential.Message,
		"data": bson.M{
			"token": credential.Token,
			"email": user.Email,
		},
	}
	return evcharging.GCFReturnStruct(responData)
}
