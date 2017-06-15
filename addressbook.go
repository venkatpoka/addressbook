package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "fmt"
    "strconv"

    "github.com/gorilla/mux"
   _ "github.com/mattn/go-sqlite3"
)

type AddressBook struct {
    ID        string
    Firstname string
    Lastname  string
    Email     string
    Phone     string
}

var addressBook []AddressBook

func GetAddBookEndpoint(w http.ResponseWriter, req *http.Request) {
    fmt.Println("I am from GetPersonEndpoint")
    params := mux.Vars(req)
    rows, err := database.Query("SELECT id, firstname, lastname, email, phone FROM addressBook WHERE id = " + params["id"])

    if err != nil {
        fmt.Println("DB error", err)
    }
    addrbook := new(AddressBook)
    for rows.Next() {
        rows.Scan(&addrbook.ID, &addrbook.Firstname, &addrbook.Lastname, &addrbook.Email, &addrbook.Phone)
        fmt.Println(addrbook)
        json.NewEncoder(w).Encode(addrbook)
        return
    }
    json.NewEncoder(w).Encode(&AddressBook{})
}

func GetEntireAddBookEndpoint(w http.ResponseWriter, req *http.Request) {
    fmt.Println("I am from GetPeopleEndpoint")
    json.NewEncoder(w).Encode(addressBook)
}

func CreateAddBookEndpoint(w http.ResponseWriter, req *http.Request) {
    fmt.Println("I am from CreatePersonEndpoint")
    params := mux.Vars(req)
    var addBook AddressBook
    _ = json.NewDecoder(req.Body).Decode(&addBook)
    addBook.ID = params["id"]

    statement, _ := database.Prepare("INSERT INTO addressBook (id, firstname, lastname, email, phone) VALUES (?, ?, ?, ?, ?)")
    statement.Exec(addBook.ID, addBook.Firstname, addBook.Lastname, addBook.Email, addBook.Phone)
    json.NewEncoder(w).Encode(addBook)
}

func DeleteAddBookEndpoint(w http.ResponseWriter, req *http.Request) {
    fmt.Println("I am from DeletePersonEndpoint")
    params := mux.Vars(req)
    for index, item := range addressBook {
        if item.ID == params["id"] {
            addressBook = append(addressBook[:index], addressBook[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(addressBook)
}

var database *sql.DB

func main() {
    database, _ = sql.Open("sqlite3", "./addressBook.db")
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS addressBook (id TEXT PRIMARY KEY, firstname TEXT, lastname TEXT, email TEXT, phone TEXT)")
    statement.Exec()
    statement, _ = database.Prepare("INSERT INTO addressBook (id, firstname, lastname) VALUES (?, ?, ?)")
    statement.Exec("1", "Nic", "Raboy")
    rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
    var id int
    var firstname string
    var lastname string
    for rows.Next() {
      rows.Scan(&id, &firstname, &lastname)
      fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
    }

    router := mux.NewRouter()
    addressBook = append(addressBook, AddressBook{ID: "1", Firstname: "Venkat", Lastname: "Poka", Email: "venkat.poka@gmail.com", Phone: "847-532-5055"})
    addressBook = append(addressBook, AddressBook{ID: "2", Firstname: "Moksha", Lastname: "Poka", Email: "venkat.ju@gmail.com", Phone: "847-532-5055"})
    router.HandleFunc("/addressBook", GetEntireAddBookEndpoint).Methods("GET")
    router.HandleFunc("/addressBook/{id}", GetAddBookEndpoint).Methods("GET")
    router.HandleFunc("/addressBook/{id}", CreateAddBookEndpoint).Methods("POST")
    router.HandleFunc("/addressBook/{id}", DeleteAddBookEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":12345", router))
}
