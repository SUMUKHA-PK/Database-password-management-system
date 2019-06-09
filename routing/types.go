package routing

// PassWordRequest holds the UserName and EncryptedPassword of the incoming request
type PassWordRequest struct {
	UserName          string `json:username`
	EncryptedPassword string `json:password`
}

// CompleteHash hash the hashedData
type CompleteHash struct {
	HashedData string `json:hash`
}

// User has each user
type User struct {
	UserID string
}

// UserHash is the hash of a user
type UserHash struct {
	UserHashData string
}

// projectData is the map that contains data of one particular user
type projectData struct {
	ProjectPwdMap map[string]string
}

//StoredData is the DS that contains all the data of multiple users
var StoredData = make(map[string]projectData)
