package Routes

type LoginRequestForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Different structures in case more data needs to be collected
// on sign up.
type RegisterRequestForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PasswordChangeForm struct {
	Original string `json:"password"`
	New      string `json:"new-password"`
}
