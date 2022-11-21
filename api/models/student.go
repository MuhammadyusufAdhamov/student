package models

import "time"

type Student struct{
	ID int64 	`json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	PhoneNumber     *string   `json:"phone_number"`
	CreatedAt       time.Time `json:"created_at"`
}

type CreateAllStudentRequest struct {
	FirstName       []string  `json:"first_name" binding:"required,min=2,max=30"`
	LastName        []string  `json:"last_name" binding:"required,min=2,max=30"`
	Username        []string  `json:"username" binding:"required,min=2,max=30"`
	Email           []string  `json:"email" binding:"required,email"`
	PhoneNumber     []*string `json:"phone_number"`
}

type GetAllStudentsResponse struct {
	Students []*Student `json:"categories"`
	Count int32   `json:"count"`
}
type CreateStudent struct {
	Students []*CreateAllStudentRequest`json:"student" binding:"required,dive"`
}
