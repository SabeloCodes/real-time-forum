package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
		"github.com/gorilla/websocket"
	// uuid "githb.com/satori/go.uuid"
	// "golang.org/x/crypto/bcrypt"
)

// Struct used to upgrade HTTP connections to WebSocket connections.
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

//The DBase struct represents the database connection.
type DBase struct{
	DB *sql.DB
}


// The Init function to initialize the database tables.
func Init (database *sql.DB){
	database.Exec(`
		CREATE TABLE IF NOT EXISTS "Users" (
			"userId" INTEGER PRIMARY KEY AUTOINCREMENT,
			"nickName" 	VARCHAR(64) NOT NULL UNIQUE,
			"email"  TEXT NOT NULL UNIQUE,
			"FName" TEXT NOT NULL,
			"LName" TEXT NOT NULL,
			"password" TEXT NOT NULL,
			"gender" TEXT NOT NULL,
			"age" INTEGER NOT NULL,
			"sessionId" TEXT,
			"loggedIn" TEXT,
			FOREIGN KEY (sessionId) REFERENCES "Session" ("sessionId")
			
		);
	`)
	database.Exec(`
	CREATE TABLE IF NOT EXISTS "Session" (
		"sessionId"	TEXT PRIMARY KEY,
		"userId"	INTEGER NOT NULL,
		FOREIGN KEY (userId)
			REFERENCES "Users" ("userId")
	);
	`)

	// todos check for the null options.
	database.Exec(`
	CREATE TABLE IF NOT EXISTS "Post" (
		"postId"	TEXT UNIQUE NOT NULL PRIMARY KEY,
		"userId"	TEXT NOT NULL,
		"title"     TEXT NOT NULL,
		"category"	TEXT NOT NULL,
		"category2"	TEXT,
		"datePosted" TEXT NOT NULL,
		"body"	TEXT , 
		FOREIGN KEY ("userId")
			REFERENCES "Users" ("userId")
	);
	`)

	database.Exec(`
	CREATE TABLE IF NOT EXISTS "Comment" (
		"commentId" 	TEXT UNIQUE NOT NULL PRIMARY KEY,
		"postId"		TEXT NOT NULL,
		"userId"		TEXT NOT NULL,
		"dateCommented" 	TEXT NOT NULL,
		"body"			TEXT,
		FOREIGN KEY ("postId")
			REFERENCES "Post" ("postId")
		FOREIGN KEY ("userID")
			REFERENCES "Users" ("userId")
	);
	`)

	database.Exec(`
	CREATE TABLE IF NOT EXISTS "Reaction" (
		"reactionId" TEXT NOT NULL PRIMARY KEY,
		"postId"	TEXT NOT NULL,
		"commentId" TEXT NOT NULL,
		"userId"	TEXT NOT NULL,
		"total_reactions"		int,
		"likes"     int,
		"Dislikes"  int,
		FOREIGN KEY ("postId")
			REFERENCES "Post" ("postId")
		FOREIGN KEY (commentId)
			REFERENCES "Comment" ("commentId")
		FOREIGN KEY ("userId")
			REFERENCES "Users" ("userId")
	);
	`)

}


// Basic HTTP handler function that writes "Home Page!!" as the response when accessing the root path "/".
func homePage (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page!!")
}


// For reading messages from the WebSocket connection and logging them. Also echoes back the received message to the client by writing it back to the WebSocket connection.
func reader(conn *websocket.Conn){
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil{
			log.Println(err)
			return
		}
	}

}


// Handler function for the WebSocket endpoint
func wsEndpoint (w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {return true}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// Logs message when a client successfully connects
	log.Println("Clinet Successfully Connected...")

	// Calls the reader function to handle the WebSocket messages.
	reader(ws)
}


// Setting up the HTTP routes. "/" with homePage handler function and "/ws" with wsEndpoint handler function
func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}


// Open SQLite database using sql.Open function.
// If any error occurs during opening of database, program logs error and terminates. 
func main() {
	database, err := sql.Open("sqlite3", "real-time-forum.db")
	if err != nil{
		log.Fatal(err)
	}
	
	// The Init function is called to initialize the database tables.
	Init(database)
	defer database.Close()
	
	// The setupRoutes function is called to set up the HTTP routes.
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
	

	
}