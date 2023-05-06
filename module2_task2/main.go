package main

import (
  "fmt"
  "io"
  "log"
  "net/http"
  "os"

  "github.com/gorilla/mux"
)

func main() {
  httpAddr := "0.0.0.0:9999"
  if port := os.Getenv("PORT"); port != "" {
    httpAddr = "0.0.0.0:" + port
  }
  fmt.Println("HTTP Server listening on", httpAddr)

  // Start an HTTP server using the custom router
  log.Fatal(http.ListenAndServe(httpAddr, setupRouter()))
}

func setupRouter() *mux.Router {
  // Create a new empty HTTP Router
  r := mux.NewRouter()

  // When an HTTP GET request is received on the path /health, delegates to the function "HealthCheckHandler()"
  r.HandleFunc("/health", HealthCheckHandler).Methods("GET")

  // when an HTTP GET request is received on the path /hello
  r.HandleFunc("/hello", HelloHandler).Methods("GET")

  return r
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
  // Print a line in the logs
  fmt.Println("HIT: healthcheck")

  // Write the string "ALIVE" into the response's body
  _, _ = io.WriteString(w, "ALIVE")

  // End of the function: return HTTP 200 by default
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
  // Extract the query parameters from the GET request
  queryParams := r.URL.Query()

  // Retrieve the query parameters with the key "name"
  nameParams := queryParams["name"]

  var name string
  switch len(nameParams) {
  case 0:
    // Set the name variable to an empty string when there is no parameter "name" in the request
    name = ""
  case 1:
    // Set the name variable to the unique parameter "name" in the request
    name = nameParams[0]
  default:
    // Set the name variable to the last occurence of the parameters "name" in the request
    name = nameParams[len(nameParams)-1]
  }

  // Set a default value if the name is empty
  if name == "" {
    name = "there"
  }

  // Write the string "Hello <name>" into the response's body
  _, _ = io.WriteString(w, fmt.Sprintf("Hello %s!", name))

  // Print a line in the ACCESS log
  fmt.Printf("HIT: hello handler with name %s \n", name)
}
