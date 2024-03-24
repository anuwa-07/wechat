package main
import (
	"fmt"
	"time"

	"net/http"
	"github.com/gorilla/mux"
	"github.com/anuwa-07/wechat/internal/handlers"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

// main function and run the server
func main() {
	r := mux.NewRouter();
	r.HandleFunc("/employee/create", handlers.CreateUser).Methods("POST");
	http.Handle("/", r);
	//
	// configure the server and run it...
	server := &http.Server{
		Addr:    ":8080",
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	//
	fmt.Println("Server is running on port 8080...")
	server.ListenAndServe();
}