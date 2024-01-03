package signup

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"github.com/ev-chargingpoint/backend-evchargingpoint/profile"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/argon2"
)

func SignUpUser(db *mongo.Database, insertedDoc evcharging.User) (string, error) {
	if insertedDoc.NamaLengkap == "" || insertedDoc.NomorHp == "" || insertedDoc.NamaKendaraan == "" || insertedDoc.NomorPolisi == "" || insertedDoc.Email == "" || insertedDoc.Password == "" {
		return "", fmt.Errorf("mohon untuk melengkapi data")
	}
	if err := checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return "", fmt.Errorf("email tidak valid")
	}
	userExists, _ := profile.GetUserFromEmail(insertedDoc.Email, db)
	if insertedDoc.Email == userExists.Email {
		return "", fmt.Errorf("email sudah terdaftar")
	}
	if insertedDoc.Confirmpassword != insertedDoc.Password {
		return "", fmt.Errorf("konfirmasi password salah")
	}
	if strings.Contains(insertedDoc.Password, " ") {
		return "", fmt.Errorf("password tidak boleh mengandung spasi")
	}
	if len(insertedDoc.Password) < 8 {
		return "", fmt.Errorf("password terlalu pendek")
	}
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf("kesalahan server : salt")
	}
	hashedPassword := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	user := bson.M{
		"namalengkap":   insertedDoc.NamaLengkap,
		"nomorhp":       insertedDoc.NomorHp,
		"namakendaraan": insertedDoc.NamaKendaraan,
		"nomorpolisi":   insertedDoc.NomorPolisi,
		"email":         insertedDoc.Email,
		"password":      hex.EncodeToString(hashedPassword),
		"salt":          hex.EncodeToString(salt),
	}
	_, err = evcharging.InsertOneDoc(db, "user", user)
	if err != nil {
		return "", err
	}
	return insertedDoc.Email, nil
}
