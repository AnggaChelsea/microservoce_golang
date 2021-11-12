package response

import(
	"simplelogic/services/users/repository/usecase"
)

type userresponse struct {
	Id int64 `json:"id"`
	Message string `json:"message"`
	Data []usecase.User `json:"data"`
}