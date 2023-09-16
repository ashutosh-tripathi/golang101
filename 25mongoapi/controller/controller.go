package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/ashutosh-tripathi/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://127.0.0.1:27017"
const dbname = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB")
	collection = client.Database(dbname).Collection(collectionName)

	fmt.Println("collection is populated")
}
func InsertMovie(movie model.Netflix) string {
	fmt.Println("movie is", movie)
	res, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		fmt.Println("got error while inserting")
	}
	insertedID := res.InsertedID.(primitive.ObjectID)
	insertedIDString := insertedID.Hex()

	fmt.Println("inserted movie with id", insertedIDString)
	return insertedIDString
}
func FindAllMovie() []bson.M {
	cursor, _ := collection.Find(context.Background(), bson.M{})
	defer cursor.Close(context.Background())
	var movies []bson.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		var err error
		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		fmt.Println("movie is:", movie)
		movies = append(movies, movie)

	}
	return movies

}
func FindOneMovie(id string) bson.M {
	primId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	var movie bson.M
	if err = collection.FindOne(context.Background(), bson.M{"_id": primId}).Decode(&movie); err != nil {
		log.Fatal(err)
	}
	return movie

}

func UpdateMovie(movieId string) {
	primObjId, _ := primitive.ObjectIDFromHex(movieId)
	id := bson.M{"_id": primObjId}
	update := bson.M{"$set": bson.M{"watched": true}}
	res, err := collection.UpdateOne(context.Background(), id, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("updated %d record", res.UpsertedCount)

}

func DeleteMovie(movieId string) {
	primObjId, _ := primitive.ObjectIDFromHex(movieId)
	idstring := bson.M{"id": primObjId}
	res, err := collection.DeleteOne(context.Background(), idstring)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %d record", res.DeletedCount)

}
func DeleteAllMovie() {
	deleteCount, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %d records", deleteCount.DeletedCount)
}
