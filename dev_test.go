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
	email := "admin@gmail.com"
	hasil, err := profile.GetUserFromEmail(email, db)
	if err != nil {
		t.Errorf("Eerror TestGetUserFromEmail: %v", err)
	} else {
		fmt.Println(hasil)
	}
}

func TestSignUpUser(t *testing.T) {
	var doc evcharging.User
	doc.NamaLengkap = "Admin"
	doc.NomorHp = "089647129890"
	doc.NamaKendaraan = "admin"
	doc.NomorPolisi = "admin"
	doc.Email = "admin@gmail.com"
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

func TestCheckLatitudeLongitude(t *testing.T) {
	latitude := "-6.8763112207612425"
	longitude := "107.57108959575653"
	hasil := charging_station.CheckLatitudeLongitude(db, latitude, longitude)
	if hasil == false {
		t.Errorf("Error TestCheckLatitudeLongitude: %v", hasil)
	} else {
		fmt.Println(hasil)
	}
}

// func TestCheckLatitudelongitude(t *testing.T) {
// 	err := charging_station.CheckLatitudeLongitude(db, "-6.8763112207612425", "107.57108959575653")
// 	fmt.Println(err)
// }

// func TestChargingStation(t *testing.T) {
// 	conn := db
// 	id := "65966df243d1a9066896b75a"
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	err = charging_station.DeleteChargingStationOlehAdmin(objectId, conn)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Berhasil DeleteFishingSpot")
// 	}
// }

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
