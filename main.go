//A simple interactive example of a rest API
//with credit and apologies for butchering their Restful API example
//to https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/
package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Tagline   string `json:"tagline,omitempty"`
}

var attendees []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range attendees {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}

func GetAttendeesEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(attendees)
}

func CreateAttendeeEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    attendees = append(attendees, person)
    json.NewEncoder(w).Encode(attendees)
}

func DeleteAttendeeEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range attendees {
        if item.ID == params["id"] {
            attendees = append(attendees[:index], attendees[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(attendees)
}

func main() {
    router := mux.NewRouter()
    attendees = append(attendees, Person{ID: "1", Firstname: "Taetem", Lastname: "Simms", Tagline: "Putting humanity into computing."})
    router.HandleFunc("/attendees", GetAttendeesEndpoint).Methods("GET")
    router.HandleFunc("/attendees/{id}", GetAttendeeEndpoint).Methods("GET")
    router.HandleFunc("/attendees/{id}", CreateAttendeeEndpoint).Methods("POST")
    router.HandleFunc("/attendees/{id}", DeleteAttendeeEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":12345", router))
}