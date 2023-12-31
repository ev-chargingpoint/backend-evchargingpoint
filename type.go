package evchargingpoint

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaLengkap     string             `bson:"namalengkap,omitempty" json:"namalengkap,omitempty"`
	NomorHp         string             `bson:"nomorhp,omitempty" json:"nomorhp,omitempty"`
	NamaKendaraan   string             `bson:"namakendaraan,omitempty" json:"namakendaraan,omitempty"`
	NomorPolisi     string             `bson:"nomorpolisi,omitempty" json:"nomorpolisi,omitempty"`
	Email           string             `bson:"email,omitempty" json:"email,omitempty"`
	Password        string             `bson:"password,omitempty" json:"password,omitempty"`
	Image           string             `bson:"image,omitempty" json:"image,omitempty"`
	Confirmpassword string             `bson:"confirmpass,omitempty" json:"confirmpass,omitempty"`
	Salt            string             `bson:"salt,omitempty" json:"salt,omitempty"`
}

type ChargingSatation struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ChargingKode   string             `bson:"chargingkode,omitempty" json:"chargingkode,omitempty"`
	Nama           string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Alamat         string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	AmmountPlugs   string             `bson:"ammountplugs,omitempty" json:"ammountplugs,omitempty"`
	Daya           string             `bson:"daya,omitempty" json:"daya,omitempty"`
	Connector      string             `bson:"connector,omitempty" json:"connector,omitempty"`
	Harga          string             `bson:"harga,omitempty" json:"harga,omitempty"`
	Image          string             `bson:"image,omitempty" json:"image,omitempty"`
	NomorTelepon   string             `bson:"nomortelepon,omitempty" json:"nomortelepon,omitempty"`
	JamOperasional string             `bson:"jamoperasional,omitempty" json:"jamoperasional,omitempty"`
	Longitude      string             `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude       string             `bson:"latitude,omitempty" json:"latitude,omitempty"`
}

type Transaksi struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ChargingSatation ChargingSatation   `bson:"chargingsatation,omitempty" json:"chargingsatation,omitempty"`
	User             User               `bson:"user,omitempty" json:"user,omitempty"`
	Tanggal          string             `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
	StartTime        time.Time          `bson:"starttime,omitempty" json:"starttime,omitempty"`
	EndTime          time.Time          `bson:"endtime,omitempty" json:"endtime,omitempty"`
	TotalPrice       string             `bson:"totalprice,omitempty" json:"totalprice,omitempty"`
	Status           int                `bson:"status,omitempty" json:"status,omitempty"`
}

type Credential struct {
	Status  int    `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Response struct {
	Status  int    `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Payload struct {
	Id    primitive.ObjectID `json:"id"`
	Email string             `json:"email"`
	Exp   time.Time          `json:"exp"`
	Iat   time.Time          `json:"iat"`
	Nbf   time.Time          `json:"nbf"`
}
