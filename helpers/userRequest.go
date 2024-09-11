package helpers

type Signup_Body struct {
	UserName string
	Email    string
	Password string
}

type Login struct {
	Email    string
	Password string
}