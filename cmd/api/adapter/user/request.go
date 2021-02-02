package user

type addRequest struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required"`
}
