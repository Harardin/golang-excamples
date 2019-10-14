package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	// This is mongoDB drivers
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config configuration file structure
type Config struct {
	MongoConnectionStr string `json:"mongoConnString"`
}

// Document this is mongoDB Document model
type Document struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	GoodName    string             `json:"name" bson:"name"`
	Price       string             `json:"price" bson:"price"`
	Parameters  []string           `json:"parameters" bson:"parameters"`
	ImagesList  []string           `json:"imgs" bson:"imgs"`
	Description string             `json:"description" bson:"description"`
}

// RequestProduct This is JSON post to recieve products from MongoDb
type RequestProduct struct {
	CategoryName  string `json:"category"`
	HowMuchToTake int    `json:"count"`
}

// ConfigData contains data of the configuration file
var ConfigData Config

func main() {
	// Geting config file
	configFile, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Fatal("Configuration file is lost")
		return
	}
	json.Unmarshal(configFile, &ConfigData)

	http.HandleFunc("/api/imgs", GetImage)
	http.HandleFunc("/api/get-data", GetGoodsInfo)

	err = http.ListenAndServe(":11", nil)
	if err != nil {
		log.Fatal("ListenServer: ", err)
		return
	}
}

// GetImage geting an image from mongo and gives it as a file via http
func GetImage(w http.ResponseWriter, r *http.Request) {
	//var ImgStringTemplate string = `<!DOCTYPE html><html><html lang="en"><head></head><body><img src="data:image/jpg;base64,{{.Image}}"></body></html>`

	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		fmt.Fprintf(w, `{"err":"Link param is missing"}`)
		return
	}
	if len(keys) > 1 {
		fmt.Fprintf(w, `{"err":"To many img params"}`)
		return
	}
	// Setting Up the connection to MongoDB
	clientOption := options.Client().ApplyURI(ConfigData.MongoConnectionStr)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal("Failed to connect to mongoDB")
		return
	}
	// Setting Database to get imgs from
	db := client.Database("ImgsDataBase")
	// Geting the image from MongoDB by it's name
	buc, err := gridfs.NewBucket(db)
	if err != nil {
		log.Fatal("Failed to create new Bucket for Img")
		client.Disconnect(context.TODO())
		return
	}

	// This is the format "auqcup62.jpg"
	key := keys[0]
	imgBuf := new(bytes.Buffer)
	buc.DownloadToStreamByName(key, imgBuf)

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(imgBuf.Bytes())))
	if _, err := w.Write(imgBuf.Bytes()); err != nil {
		log.Println("Unable to write an image")
	}

	/*imgStr := base64.StdEncoding.EncodeToString(imgBuf.Bytes())

	if tmpl, err := template.New("image").Parse(ImgStringTemplate); err != nil {
		log.Println("Unable to parse image template, Error: ", err)
	} else {
		data := map[string]interface{}{"Image": imgStr}
		if err = tmpl.Execute(w, data); err != nil {
			log.Println("Unable to execute template, Error: ", err)
			client.Disconnect(context.TODO())
			return
		}
	}*/
	client.Disconnect(context.TODO())
}

// GetGoodsInfo Gets the JsonDocument with info about the good from mongoDB
func GetGoodsInfo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/get-data" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		fmt.Fprint(w, `{"err":"Should be a POST request with JSON"}`)
	case "POST":
		clientOptions := options.Client().ApplyURI(ConfigData.MongoConnectionStr)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal("Failed to connect to mongoDB")
			return
		}
		// Geting JSON Post
		var JSONRequest RequestProduct
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&JSONRequest)
		if err != nil {
			fmt.Fprint(w, `{"err":"Bad request"}`)
			client.Disconnect(context.TODO())
			return
		}

		// Geting the Document
		var DocumentsList []*Document
		findOption := options.Find()
		findOption.SetLimit(int64(JSONRequest.HowMuchToTake))
		collection := client.Database("GoodsInfo").Collection(JSONRequest.CategoryName)
		cur, err := collection.Find(context.TODO(), bson.D{{}}, findOption)
		if err != nil {
			log.Fatal("Failed to get the collection from MongoDB, Err: ", err)
			client.Disconnect(context.TODO())
			return
		}
		for cur.Next(context.TODO()) {
			var elem Document
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
			DocumentsList = append(DocumentsList, &elem)
		}
		cur.Close(context.TODO())
		jsonString, err := json.Marshal(DocumentsList)
		check(err)
		fmt.Fprintf(w, string(jsonString))
	default:
		fmt.Fprintf(w, `{"err":"Acceptig only POST request with approptiate JSON"}`)
	}
}

// check is a basic check for error
func check(e error) {
	if e != nil {
		log.Fatal("Unexpected error: ", e)
	}
}
