package evchargingpoint_test

import (
	"fmt"
	"testing"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"github.com/ev-chargingpoint/backend-evchargingpoint/charging_station"
	"github.com/ev-chargingpoint/backend-evchargingpoint/login"
	"github.com/ev-chargingpoint/backend-evchargingpoint/profile"
	"github.com/ev-chargingpoint/backend-evchargingpoint/signup"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = evcharging.MongoConnect("MONGOSTRING", "db_evchargingpoint")

func TestLogIn(t *testing.T) {
	var doc evcharging.User
	doc.Email = "dimasardnt6@gmail.com"
	doc.Password = "fghjkliow"
	user, err := login.LogIn(db, doc)
	if err != nil {
		t.Errorf("Error getting document: %v", err)
	} else {
		fmt.Println("Selamat Datang:", user)
	}
}

func TestGetUserFromEmail(t *testing.T) {
	email := "dimasardnt6@gmail.com"
	hasil, err := profile.GetUserFromEmail(email, db)
	if err != nil {
		t.Errorf("Eerror TestGetUserFromEmail: %v", err)
	} else {
		fmt.Println(hasil)
	}
}

func TestSignUpUser(t *testing.T) {
	var doc evcharging.User
	doc.NamaLengkap = "Dimas Ardianto"
	doc.NomorHp = "089647129890"
	doc.NamaKendaraan = "Hyundai ICONIC 6"
	doc.NomorPolisi = "D 1234 ABC"
	doc.Email = "dimasardnt6@gmail.com"
	doc.Password = "fghjkliow"
	doc.Confirmpassword = "fghjkliow"
	_, err := signup.SignUpUser(db, doc)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	} else {
		fmt.Println("Berhasil Sign Up dengan email :", doc.Email)
	}
}

func TestGetChargingStationFromID(t *testing.T) {
	id := "659159fd31636a48151a946c"
	objectId, _ := primitive.ObjectIDFromHex(id)
	hasil, err := charging_station.GetChargingStationFromID(objectId, db)
	if err != nil {
		t.Errorf("Error TestGetChargingStationFromID: %v", err)
	} else {
		fmt.Println(hasil)
	}
}

// func TestAddChargingStationByAdmin(t *testing.T) {
// 	var doc evcharging.ChargingSatation
// 	doc.ChargingKode = "EVCP054"
// 	doc.Nama = "EV Charger Point Bandung"
// 	doc.Alamat = "Jl. Sariasih 3"
// 	doc.NomorTelepon = "089647129890"
// 	doc.AmmountPlugs = 10
// 	doc.Daya = "100kW"
// 	doc.Connector = "CCS Combo 2"
// 	doc.Harga = "2000/kWh"
// 	doc.JamOperasional = "24 jam"
// 	doc.Longitude = "107.578"
// 	doc.Latitude = "-6.914"
// 	doc.Image = "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.mobil123.com%2Fberita%2Fmobil-listrik-terbaik-di-indonesia-5-pil"
// 	err := charging_station.PostChargingStationOlehAdmin(db, doc)
// 	if err != nil {
// 		t.Errorf("Error inserting document: %v", err)
// 	} else {
// 		fmt.Println("Data berhasil disimpan dengan nama :", doc.Nama)
// 	}
// }
