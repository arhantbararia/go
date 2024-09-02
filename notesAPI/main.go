package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const help_data = map[string]interface{}{
// 	"/create":    "POST :Create a note"	}
// 	// "boolValue":   true,
// 	// "stringValue": "hello!",
// 	// "objectValue": map[string]interface{}{
// 	"arrayValue": []int{1, 2, 3, 4},
// },
// }

var mdbClient *mongo.Client

type Scope struct {
	Project string
	Area    string
}

type Note struct {
	Title string
	Text  string
	Tags  []string
	Scope Scope
}

func main() {

	const serverAddr string = "127.0.0.1:8081"
	fmt.Println("Hola Caracola")
	mdb_connection_string := "mongodb://localhost:27017/"

	ctxBg := context.Background()
	var err error
	mdbClient, err = mongo.Connect(ctxBg, options.Client().ApplyURI(mdb_connection_string))

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = mdbClient.Disconnect(ctxBg); err != nil {
			panic(err)
		}

	}()

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Notes Taking API GET /help to get help :)"))
	})
	http.HandleFunc("POST /create", createNotes)
	http.HandleFunc("GET /view", viewAllNotes)

	http.HandleFunc("GET /help", showHelp)
	log.Fatal(http.ListenAndServe(serverAddr, nil))

}

func viewAllNotes(w http.ResponseWriter, r *http.Request) {
	notesCollection := mdbClient.Database("Notesbook").Collection("Notes")
	result, err := notesCollection.Find(r.Context(), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer result.Close(r.Context())
	for result.Next(r.Context()) {
		var note Note
		err := result.Decode(&note)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Note: %+v", note)
	}

}

func createNotes(w http.ResponseWriter, r *http.Request) {
	var note Note
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//write to mongodb
	notesCollection := mdbClient.Database("Notesbook").Collection("Notes")
	result, err := notesCollection.InsertOne(r.Context(), note)
	//The returned context is always non-nil; it defaults to the background context. (if a request does not have context)
	// you must be thinking what does inserting in the db has to do with request apart from data,
	// but you will notice that every function related to mongodb requires context background, and we are passing
	// the context (TTL etc. ) of the request as a context background object to the mongodb connection function

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Id: %v", result.InsertedID)

	fmt.Fprintf(w, "Note: %+v", note)

}

func showHelp(w http.ResponseWriter, r *http.Request) {
	help_data := map[string]interface{}{
		"/create": "POST :Create a note",
		"/view":   "GET :View all notes"}

	helpjsonData, err := json.Marshal(help_data)
	if err != nil {
		fmt.Printf("could not marshal help data to json: %s\n", err)
		return
	}

	fmt.Fprintf(w, "Help: %s", helpjsonData)
}
