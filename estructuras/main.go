package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[int]string
}

func main() {
	Go := Course{
		Name:    "Go desde cero",
		Price:   12.24,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Clases: map[int]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	css := Course{
		Name:   "CSS desde cero",
		IsFree: true,
	}

	js := Course{}
	js.Name = "Curso JS"
	js.UserIDs = []uint{12, 67}

	fmt.Println(Go.Name)
	fmt.Printf("% + v\n", css)
	fmt.Printf("% + v", js)
}
