package routing

// storePassWordRequest holds the UserName and EncryptedPassword of the incoming request
type storePassWordRequest struct {
	UserName          string `json:username`
	EncryptedPassword string `json:password`
}

type storeCompleteHash struct {
	HashedData string `json:hash`
}

//StoredData is the DS that contains all the data
var StoredData = make(map[string]string)
