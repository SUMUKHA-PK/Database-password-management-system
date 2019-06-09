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

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RegisterUser implements /registerUser end point
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var newReq User
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// log.Println(util.StringWithCharset(random.Intn(20)+10, charset))
	preHashString := newReq.UserID + util.StringWithCharset(random.Intn(20)+10, charset)
	hashedString := crypto.CreateSHA256Hash(preHashString)
	hashOutput := UserHash{hashedString}
	log.Println(hashedString)
	outJSON, err := json.Marshal(hashOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(outJSON)
}
