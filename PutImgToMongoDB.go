package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	// This is mongoDB libs
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/gridfs"
)

var reader io.Reader = (*os.File)(nil)

// This is an excample of storing image in mongoDB and retrieving it from mongoDB
func main() {
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	//client, err := mongo.Connect(context.TODO(), clientOptions)
	//check(err)
	//db := client.Database("ThisDB")

	// Reading an image file
	//reader, err = os.Open("Free-wallpaper-bioshock-rapture.jpg")
	//check(err)

	//buc, err := gridfs.NewBucket(db)
	//check(err)
	//buc.UploadFromStream("walpaperImg.jpg", reader)
	//stream, err := buc.OpenDownloadStreamByName("walpaperImg.jpg")

	//img, err := os.Create("img.jpg")
	//check(err)

	// Какой io.Writer туда нужно передать
	//buc.DownloadToStreamByName("walpaperImg.jpg", img)

	fmt.Println("This is some text")

	mainroute := "/api"
	http.HandleFunc(mainroute+"/img", GetImage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

// GetImage is geting images from mongoDB
// Тип должно работать но надо дофиксить
func GetImage(w http.ResponseWriter, r *http.Request) {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	check(err)

	db := client.Database("ThisDB")

	var gridfs db.GridFS

	name := "walpaperImg.jpg"

	f, err := gridfs.Open(name)
	check(err)

	defer f.Close()

	http.ServeContent(w, r, name, time.Now(), f)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
