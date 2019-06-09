package routing

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// StorePassword is the function that stores the incoming password
func StorePassword(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var newReq storePassWordRequest
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Logic to store the password
	w.Header().Set("Success message", "Completed entry")
	w.WriteHeader(200)
	// profile := storePassWordRequest{"user", "passw"}

	// js, err := json.Marshal(profile)
	// err = errors.New("enoondu")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// log.Println(js)
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(js)
}
