package solidGate

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var listNotifications [] interface{}

func SaveSolidGateProd(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)

		body, err := ioutil.ReadAll(r.Body)

		errorHandeler(err)

		if len(listNotifications) > 100 {
			listNotifications = nil
		}

		var notific interface{}

		err = json.Unmarshal(body, &notific)

		errorHandeler(err)

		listNotifications = append(listNotifications, notific)

		w.Write([]byte("{\"status\":\"ok\"}"))
	}
	defer r.Body.Close()
}

func BackSolidGateProd(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		list, _ := json.Marshal(listNotifications)

		w.Write([]byte(list))
		listNotifications = nil
	}
	defer r.Body.Close()
}



func errorHandeler(err error) {
	if err != nil {
		log.Print(err)
	}
}
