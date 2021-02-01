package user

type addUserRequest struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required"`
}
