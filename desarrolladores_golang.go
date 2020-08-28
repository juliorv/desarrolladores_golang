package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {

    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/desarrolladores", returnAllDesarrolladores).Methods("GET")
    myRouter.HandleFunc("/desarrollador", createNuevoDesarrollador).Methods("POST")
    myRouter.HandleFunc("/desarrollador/{id}", updateDesarrollador).Methods("PUT")
    myRouter.HandleFunc("/desarrollador/{id}", deleteDesarrollador).Methods("DELETE")
    myRouter.HandleFunc("/desarrollador/{id}", returnSoloDesarrollador).Methods("GET")

    log.Fatal(http.ListenAndServe(":10001", myRouter))
}

func returnAllDesarrolladores(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllDesarrolladores")
    json.NewEncoder(w).Encode(Des)
}

func returnSoloDesarrollador(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, desarrolladores := range Des {
        if desarrolladores.Id == key {
            json.NewEncoder(w).Encode(desarrolladores)
        }
    }
}

func createNuevoDesarrollador(w http.ResponseWriter, r *http.Request) {   
    reqBody, _ := ioutil.ReadAll(r.Body)
    var des Desarrolladores 
    json.Unmarshal(reqBody, &des)
    // update our global Articles array to include
    // our new Article
    Des = append(Des, des)

    json.NewEncoder(w).Encode(des)
}


func deleteDesarrollador(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    id := vars["id"]

    for index, des := range Des {
        if des.Id == id {
            Des = append(Des[:index], Des[index+1:]...)
        }
    }

}


func updateDesarrollador(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    id := vars["id"]

    for index, des := range Des {
        if des.Id == id {
            Des = append(Des[:index], Des[index+1:]...)
            var updatedDes Desarrolladores
            json.NewDecoder(r.Body).Decode(&updatedDes)
            Des = append(Des, updatedDes)
            json.NewEncoder(w).Encode(updatedDes)
            return
        }
    }
}

type Desarrolladores struct {
    Id      string `json:"Id"`
    nombre string `json:"nombre"`
    link_github string `json:"link_github"`
    tecnologias string `json:"tecnologias"`
}

var Des []Desarrolladores

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    Des = []Desarrolladores{
        Desarrolladores{Id: "1", nombre: "JULIO NUÑEZ", link_github: "LINK", tecnologias: "GOLANG"},
        Desarrolladores{Id: "3", nombre: "JULIO NUÑEZ Z", link_github: "LINK", tecnologias: "SPRING"},
    }
    handleRequests()
}