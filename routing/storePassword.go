package routing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/Database-password-management-system/crypto"
)

// StorePassword is the function that stores the incoming password
func StorePassword(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var newReq PassWordRequest
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Logic to store the password
	// each incoming object is split to individual strings,
	// and hashed based on both of them,
	// this hashed string is the key to the hashmap,
	// the hashmap has value as the password that comes in

	hashedString := crypto.CreateSHA256Hash(newReq.UserName + newReq.EncryptedPassword)
	// the Object that is sent out as a response
	hashOutput := CompleteHash{hashedString}

	//add the password on to the hashmap
	// StoredData[hashedString] = newReq.EncryptedPassword

	//marshalling the object into JSON type
	outJSON, err := json.Marshal(hashOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Print(StoredData)

	// w.Header().Set("Success message", "Completed entry")
	// w.WriteHeader(200)
	// // profile := storePassWordRequest{"user", "passw"}

	// js, err := json.Marshal(profile)
	// err = errors.New("enoondu")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	//writing the JSON to response writer
	w.Header().Set("Content-Type", "application/json")
	w.Write(outJSON)
}
