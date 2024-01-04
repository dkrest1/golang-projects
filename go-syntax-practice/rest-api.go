package main


type Application struct {
	ID string 	`json:"id,omitempty"`
	Name string	`json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Developer string  `json`
	Details *Details
} 

type Details struct {
	Language string
	Type string
}

var app []Application

const (
	PORT = ":1909"
)

func main() {
	
}