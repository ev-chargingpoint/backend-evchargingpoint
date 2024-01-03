package login

import (
	"encoding/hex"
	"fmt"

	"github.com/badoux/checkmail"
	evcharging "github.com/ev-chargingpoint/backend-evchargingpoint"
	"github.com/ev-chargingpoint/backend-evchargingpoint/profile"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/argon2"
)

// Login Admin, User
func LogIn(db *mongo.Database, insertedDoc evcharging.User) (user evcharging.User, err error) {
	if insertedDoc.Email == "" || insertedDoc.Password == "" {
		return user, fmt.Errorf("mohon untuk melengkapi data")
	}
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return user, fmt.Errorf("email tidak valid")
	}
	existsDoc, err := profile.GetUserFromEmail(insertedDoc.Email, db)
	if err != nil {
		return
	}
	salt, err := hex.DecodeString(existsDoc.Salt)
	if err != nil {
		return user, fmt.Errorf("kesalahan server : salt")
	}
	hash := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	if hex.EncodeToString(hash) != existsDoc.Password {
		return user, fmt.Errorf("password salah")
	}
	return existsDoc, nil
}
