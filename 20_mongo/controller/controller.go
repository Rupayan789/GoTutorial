package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	models "netflix_api_mock/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = "mongodb+srv://routrupayan99:tgVZhmuSgJLUNZxZ@cluster0.vemaddx.mongodb.net/?retryWrites=true&w=majority"
const dbName string = "Netflix"
const colName string = "Watchlist"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println("Some error happened while Mongo connection")
		log.Fatal(err)
	}

	fmt.Println("Mongo Db Connection established successfully")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection Instance created", collection)
}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		if err := cur.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return movies
}

func insertOneMovie(movie models.Netflix) {
	insertResult, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Element _id is :", insertResult.InsertedID)
}

func updateOneMovie(movieId string) {
	objectID, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": objectID}
	update := bson.M{"watched": true}
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified Count is :", updateResult.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	objectID, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"id": objectID}
	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count is :", deleteResult.DeletedCount)
}

func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted count is :", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func GetAllMoviesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	movies := getAllMovies()
	json.NewEncoder(w).Encode(movies)
}

func CreateMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie models.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode("Successfully created a movie")
}

func MarkMovieWatchedController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode("Movie successfully updated")
}

func DeleteOneMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Movie successfully deleted")
}

func DeleteAllMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	movieCountDeleted := deleteAllMovie()
	json.NewEncoder(w).Encode(fmt.Sprintf("Count of movies deleted is %v", movieCountDeleted))
}
