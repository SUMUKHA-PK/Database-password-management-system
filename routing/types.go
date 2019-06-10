package routing

// Types has all the data structures used
// 2 separate DSs are used, one to store UserID and Hash,
// other to hold all the data with key value as Userhash

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
	UserID         string `json:userID`
	UserHashString string `json:userhashstring`
}

// UserHash is the hash of a user
type UserHash struct {
	UserHashData string
}

// projectData is the map that contains data of one particular user
type projectData struct {
	ProjectPwdMap map[string]string
}

// userData is the mao of user
// primarily used to check existance of a user
var userData = make(map[string]string)

//StoredData is the DS that contains all the data of multiple users
// key is the hash of the user
// the value is the type of all the projects and their passwords
var StoredData = make(map[string]projectData)
