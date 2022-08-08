package errors

type UsernameAlreadyExistsError struct{}

func (UsernameAlreadyExistsError) Error() string {
	return "username already exists"
}

type EmailAlreadyExistsError struct{}

func (EmailAlreadyExistsError) Error() string {
	return "email already exists"
}

type InvalidUsernameOrPasswordError struct{}

func (InvalidUsernameOrPasswordError) Error() string {
	return "username or password are invalid"
}
