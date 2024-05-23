package main

import (
	"fmt"
	"net/http"
)

type Course struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Id          int     `json:"id"`
	Price       float64 `json:"price"`
	Discount    float64 `json:"discount"`
	Active      bool    `json:"active"`
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf("Name: %s - Description: %s - Price: %f", c.Name, c.Description, c.Price)
}

func soma(x int, y int) (int, bool) {
	if x > 10 {
		return x + y, true
	}
	return x + y, false
}
func main() {
	a := "Olá Mundo"
	b := 10
	println(a)
	println(soma(11, 20))
	fmt.Println("Hello World", b)

	course := Course{
		Name:        "GO",
		Description: "Curso de GO",
		Id:          1,
		Price:       100,
		Discount:    10,
		Active:      true,
	}
	course.Price = 200
	fmt.Println(course.Name, course.Price)

	fmt.Println(course.GetFullInfo())
	http.HandleFunc("/", home)
	//http.HandleFunc("/", newhome)
	http.ListenAndServe(":8080", nil)

}

/* func newhome(w http.ResponseWriter, r *http.Request) {
	coursenew := Course{
		Name:        "GO",
		Description: "Curso de GO",
		Id:          1,
		Price:       100,
		Discount:    10,
		Active:      true,
	}
	json.NewDecoder(w).Encode(coursenew)
} */

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World with GO"))
}
