package evchargingpoint_test

import (
	"fmt"
	"testing"

	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"github.com/ev-chargingpoint/backend-evchargingpoint/charge"
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

func TestGetAllChargingStation(t *testing.T) {
	hasil, err := charging_station.GetAllChargingStation(db)
	if err != nil {
		t.Errorf("Error TestGetAllChargingStation: %v", err)
	} else {
		fmt.Println(hasil)
	}
}

func TestChargeCar(t *testing.T) {
	idChargingStation, err := primitive.ObjectIDFromHex("659a39d90cb3a4271053fbd2")
	if err != nil {
		t.Fatalf("Failed to convert hex to ObjectID: %v", err)
	}

	idUser, err := primitive.ObjectIDFromHex("6593e4db369ad3e79a8d8a2d")
	if err != nil {
		t.Fatalf("Failed to convert hex to ObjectID: %v", err)
	}

	doc := evcharging.Charge{
		User: evcharging.User{
			ID: idUser,
		},
		ChargingStation: evcharging.ChargingStation{
			ID: idChargingStation,
		},
		StartTime:  "2021-08-01 12:00:00",
		EndTime:    "2021-08-01 13:00:00",
		TotalKWH:   "20",
		TotalPrice: "20000",
		Payment:    false,
		Status:     false,
	}

	hasil, err := charge.ChargeCar(doc.ChargingStation.ID, doc.User.ID, db, doc)
	if err != nil {
		t.Errorf("Error TestChargeCar: %v", err)
	}
	if hasil == nil {
		t.Errorf("Expected a result, got nil")
	}
}

func TestProcessPayment(t *testing.T) {
	idCharge, err := primitive.ObjectIDFromHex("65b4cf767c51df5030a51f79")
	if err != nil {
		t.Fatalf("Failed to convert hex to ObjectID: %v", err)
	}

	idUser, err := primitive.ObjectIDFromHex("6593e4db369ad3e79a8d8a2d")
	if err != nil {
		t.Fatalf("Failed to convert hex to ObjectID: %v", err)
	}

	doc := evcharging.Charge{
		ChargingStation: evcharging.ChargingStation{
			ID: idCharge,
		},
		User: evcharging.User{
			ID: idUser,
		},
		PaymentMethod:   "OVO",
		InputPembayaran: "20000",
		Payment:         true,
		Status:          false,
	}
	hasil, err := charge.ProcessPayment(doc.ChargingStation.ID, doc.User.ID, db, doc)
	if err != nil {
		t.Errorf("Error TestProcessPayment: %v", err)
	}
	if hasil == nil {
		t.Errorf("Expected a result, got nil")
	}
}

func TestProcessStatus(t *testing.T) {
	idCharge, err := primitive.ObjectIDFromHex("65b4cf767c51df5030a51f79")
	if err != nil {
		t.Fatalf("Failed to convert hex to ObjectID: %v", err)
	}

	idUser, err := primitive.ObjectIDFromHex("6593e4db369ad3e79a8d8a2d")
	if err != nil {
		t.Fatalf("Failed to convert hex to ObjectID: %v", err)
	}

	doc := evcharging.Charge{
		ChargingStation: evcharging.ChargingStation{
			ID: idCharge,
		},
		User: evcharging.User{
			ID: idUser,
		},
		Status: true,
	}
	hasil, err := charge.ProcessStatus(doc.ChargingStation.ID, doc.User.ID, db, doc)
	if err != nil {
		t.Errorf("Error TestProcessStatus: %v", err)
	}
	if hasil == nil {
		t.Errorf("Expected a result, got nil")
	}
}

func TestGetAllChargeByUser(t *testing.T) {
	id := "6593e4db369ad3e79a8d8a2d"
	objectId, _ := primitive.ObjectIDFromHex(id)
	hasil, err := charge.GetAllChargeByUser(objectId, db)
	if err != nil {
		t.Errorf("Error TestGetAllChargeByUser: %v", err)
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
	id := "6596b03e613816ef04ecb396"
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
// 	var doc evcharging.ChargingStation
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
