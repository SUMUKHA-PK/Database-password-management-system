package routing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	random "math/rand"
	"net/http"

	"github.com/SUMUKHA-PK/Database-password-management-system/crypto"
	"github.com/SUMUKHA-PK/Database-password-management-system/util"
)

// RegisterUser implements /registerUser end point
// this adds the user to the userData map.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "No input found!", http.StatusBadRequest)
		return
	}
	var newReq User
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var username = newReq.UserID
	if _, ok := userData[username]; ok {
		http.Error(w, "User already exists!", http.StatusBadRequest)
		return
	}

	// log.Println(util.StringWithCharset(random.Intn(20)+10, charset))
	preHashString := newReq.UserID + util.StringWithCharset(random.Intn(20)+10, util.Charset)
	hashedString := crypto.CreateSHA256Hash(preHashString)
	userData[username] = hashedString
	hashOutput := UserHash{hashedString}
	log.Println(userData)
	outJSON, err := json.Marshal(hashOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(outJSON)
}
