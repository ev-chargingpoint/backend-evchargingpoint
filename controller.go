package evchargingpoint

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/go-github/v56/github"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/oauth2"
)

// Mongo Environment
func MongoConnect(MongoString, dbname string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MongoString)))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

// crud
func GetAllDocs(db *mongo.Database, col string, docs interface{}) interface{} {
	collection := db.Collection(col)
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error GetAllDocs %s: %s", col, err)
	}
	err = cursor.All(context.TODO(), &docs)
	if err != nil {
		return err
	}
	return docs
}

func InsertOneDoc(db *mongo.Database, col string, doc interface{}) (insertedID primitive.ObjectID, err error) {
	result, err := db.Collection(col).InsertOne(context.Background(), doc)
	if err != nil {
		return insertedID, fmt.Errorf("kesalahan server : insert")
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func UpdateOneDoc(id primitive.ObjectID, db *mongo.Database, col string, doc interface{}) (err error) {
	filter := bson.M{"_id": id}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		return fmt.Errorf("error update: %v", err)
	}
	if result.ModifiedCount == 0 {
		err = fmt.Errorf("tidak ada data yang diubah")
		return
	}
	return nil
}

func DeleteOneDoc(_id primitive.ObjectID, db *mongo.Database, col string) error {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// Get All User by Admin
// func GetAllUserByAdmin(db *mongo.Database) (user []User, err error) {
// 	collection := db.Collection("user")
// 	filter := bson.M{}
// 	cursor, err := collection.Find(context.Background(), filter)
// 	if err != nil {
// 		return user, fmt.Errorf("error GetAllUser mongo: %s", err)
// 	}
// 	err = cursor.All(context.Background(), &user)
// 	if err != nil {
// 		return user, fmt.Errorf("error GetAllUserByAdmin context: %s", err)
// 	}
// 	return user, nil
// }

// Get User

// Get User Login
func GetUserLogin(PASETOPUBLICKEYENV string, r *http.Request) (Payload, error) {
	tokenstring := r.Header.Get("Authorization")
	payload, err := Decode(os.Getenv(PASETOPUBLICKEYENV), tokenstring)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

// Get Id
func GetID(r *http.Request) string {
	return r.URL.Query().Get("id")
}

// Return Struct
func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

// save file to github
func SaveFileToGithub(usernameGhp, emailGhp, repoGhp, path string, r *http.Request) (string, error) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		return "", fmt.Errorf("error 1: %s", err)
	}
	defer file.Close()

	// Generate a random filename
	randomFileName, err := generateRandomFileName(handler.Filename)
	if err != nil {
		return "", fmt.Errorf("error 2: %s", err)
	}

	// Read the content of the file into a byte slice
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error 5: %s", err)
	}

	access_token := os.Getenv("GITHUB_ACCESS_TOKEN")
	if access_token == "" {
		return "", fmt.Errorf("error access token: %s", err)
	}

	// Initialize GitHub client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: access_token},
	)
	tc := oauth2.NewClient(r.Context(), ts)
	client := github.NewClient(tc)

	// Create a new repository file
	_, _, err = client.Repositories.CreateFile(r.Context(), usernameGhp, repoGhp, path+"/"+randomFileName, &github.RepositoryContentFileOptions{
		Message:   github.String("Add new file"),
		Content:   fileContent,
		Committer: &github.CommitAuthor{Name: github.String(usernameGhp), Email: github.String(emailGhp)},
	})
	if err != nil {
		return "", fmt.Errorf("error 6: %s", err)
	}

	imageUrl := "https://raw.githubusercontent.com/" + usernameGhp + "/" + repoGhp + "/main/" + path + "/" + randomFileName

	return imageUrl, nil
}

func generateRandomFileName(originalFilename string) (string, error) {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomFileName := fmt.Sprintf("%x%s", randomBytes, filepath.Ext(originalFilename))
	return randomFileName, nil
}
