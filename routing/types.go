package routing

// storePassWordRequest holds the UserName and EncryptedPassword of the incoming request
type storePassWordRequest struct {
	UserName          string `json:username`
	EncryptedPassword string `json:password`
}
