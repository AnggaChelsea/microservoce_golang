package models
import(
		"simplelogic/services/auth/repository/response/errors"
		"strings"
)
type User struct{
	Id int64 `json:"id,omitempty"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
}

func(user *User) Validate() *errors.ErrorResponse {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
	user.Role = strings.TrimSpace(user.Role)
	if user.Name == "" {
		return errors.NewBadRequestError("Name is valid")
	}
	if user.Email == "" {
		return errors.NewBadRequestError("Email is valid")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("Password is valid")
	}
	if user.Role == "" {
		return errors.NewBadRequestError("Role is valid")
	}
	return nil

}