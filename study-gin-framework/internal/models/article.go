package models

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       uint8  `json:"age" binding:"gte=0,lte=130"`
	Email     string `json:"email" validate:"required,email"`
}

type Article struct {
	Title       string `json:"title" binding:"min=2,max=10" validate:"is-cool"`
	Description string `json:"description" binding:"max=20"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author"`
}
