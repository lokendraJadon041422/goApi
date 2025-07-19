package types

type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name" validate:"required"`
	Age    int    `json:"age" validate:"required"`
	Gender string `json:"gender" validate:"required"`
}
