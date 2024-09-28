package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anurag0006/apiwithdb/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Uri = "mongodb+srv://anurag:anurag@cluster0.1lq8k5p.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const DbName = "netflixgolang"
const ColName = "watchlist"

// most important:
var collection *mongo.Collection

// connect with MongoDB
func init() {
	// client option
	clientOption := options.Client().ApplyURI(Uri)
	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MongoDB connection successful!")

	collection = client.Database(DbName).Collection(ColName)

	// collection instance
	log.Println("Collection instance is ready")
}

// MONGODB Helpers:

// inserting 1 record in the db
func insertOneMovie(movie models.Netflix)  {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal( err)
	}
	fmt.Println("Inserted 1 movie with id", inserted.InsertedID)
}

func updateOneMovie(movieId string)  {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal( err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal( err)
	}

	fmt.Println("Modified count", result.ModifiedCount)
}



func deleteOneMovie(movieId string){
	id,_ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id":id}

	deleteCnt, err:=collection.DeleteOne(context.Background(),filter)

	if err != nil {
		log.Fatal( err)
	}
	fmt.Println("Movie got deleted with delete count :",deleteCnt)
}

// deleting all record from mongdb

func deleteAll(){

	deletedResult,err := collection.DeleteMany(context.Background(),bson.D{{}},nil)

	if err != nil {
		log.Fatal( err)
	}
	fmt.Println("Movie got deleted with delete count :",deletedResult.DeletedCount)
}


func getAllMovies() []primitive.M{
	cur,err := collection.Find(context.Background(),bson.D{{}})

	if err != nil {
		log.Fatal( err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M
		err := cur.Decode(&movie)
		if(err != nil){
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())

    return movies
}




// Actual controllers - file:

func GetAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}


func CreateNewMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-control-Allow-Methods","POST")
	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
    insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}


func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-control-Allow-Methods","POST")
    
	params := mux.Vars(r)

	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}


func DeleteOneMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-control-Allow-Methods","POST")
    
	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Movie  with the given id deleted!")
}


func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-control-Allow-Methods","POST")
	deleteAll()
	json.NewEncoder(w).Encode("All Movies Deleted!")
}