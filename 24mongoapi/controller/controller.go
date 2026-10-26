package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AKASH-KALSARIYA/mygolang/24mongoapi/model"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionstring = "mongodb://localhost:27017/"
const dbName = "Netflix"
const colName = "Watchlist"

var collection *mongo.Collection

// ================= Mongo Connection =================
func init() {

	clientOptions := options.Client().ApplyURI(connectionstring)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connection Success!!")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready!!")
}

// ================= Mongo Helpers =================

// insert
func insertonemovie(movie model.Netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted ID:", inserted.InsertedID)
}

// update
func updateOneMovie(movieId string) {

	id, err := bson.ObjectIDFromHex(movieId)
	if err != nil {
		fmt.Println("Invalid ObjectID:", err)
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Matched:", result.MatchedCount)
	fmt.Println("Modified:", result.ModifiedCount)
}

// delete one
func deleteOneMovie(movieID string) {

	id, err := bson.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	deletecount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Count:", deletecount.DeletedCount)
}

// delete all
func deleteAllMovie() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted All Count:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all
func getAllmovies() []bson.M {

	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M

	for cursor.Next(context.Background()) {

		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

// ================= Controllers =================

func GetMyAllmovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	allmovies := getAllmovies()
	json.NewEncoder(w).Encode(allmovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var movie model.Netflix

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertonemovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
