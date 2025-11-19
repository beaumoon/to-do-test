package main

<<<<<<< HEAD
import "net/http"

func main() {
	// Your code here
=======
import (
	"encoding/json"
	"fmt"
	"net/http"
)

// initialise storage structure for passed todo items
type To_Do struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var To_Do_List []To_Do

func main() {

	// test items
	newTodo := To_Do{
		Title:       "itemtest1",
		Description: "walking my dog",
	}
	To_Do_List = append(To_Do_List, newTodo)

	newTodo2 := To_Do{
		Title:       "itemtest2",
		Description: "swimming my seal",
	}
	To_Do_List = append(To_Do_List, newTodo2)

	// set up when handler is triggered & intialise listener
	http.HandleFunc("/", ToDoListHandler)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)

>>>>>>> b6e6ab2 (Completed test)
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

<<<<<<< HEAD
	// Your code here
=======
	// sets permissions for allowed HTTP requests & header types, and clarifies return type
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// test for method being passed
	fmt.Println("method:", r.Method)

	// because frontend checks for permitted HTTP options when trying to post, return that POST is allowed
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(To_Do_List)
		return
	}

	if r.Method == "POST" {
		var todo To_Do

		// break down the request's body into the To_Do structure
		json.NewDecoder(r.Body).Decode(&todo)

		// NOT IMPLEMENTED: add error (400) handling if request passed in an invalid format

		// add new to-do to list in memory
		To_Do_List = append(To_Do_List, todo)

		// return
		json.NewEncoder(w).Encode(todo)
		return
	}

>>>>>>> b6e6ab2 (Completed test)
}
