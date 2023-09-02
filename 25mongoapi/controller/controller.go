package controller

import (
	"25mongoapi/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const connectionString = "mongodb+srv://<root>:<root>@cluster0.9jw7kca.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success.")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection is ready.")
}

// HELPERS

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	fatalLog(err)

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	fatalLog(err)

	fmt.Println("Modified count: ", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	fatalLog(err)

	fmt.Println("One delete with count: ", result.DeletedCount)
}

func deleteAllMovies() {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	fatalLog(err)

	fmt.Println("Number of movies deleted: ", result.DeletedCount)
}

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	fatalLog(err)
	defer cursor.Close(context.Background())

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		fatalLog(err)

		movies = append(movies, movie)
	}

	return movies
}

func fatalLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
