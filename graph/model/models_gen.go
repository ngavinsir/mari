// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateUserInput struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Name     *string `json:"name"`
}

type User struct {
	ID    int     `json:"id"`
	Email string  `json:"email"`
	Name  *string `json:"name"`
}
