package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

type Location struct {
	X     int    `json:"valor_x"`
	Y     int    `json:"valor_y"`
	Name  string `json:"locate_name"`
	valor int    // campo que inicia com letra minuscula é "private"
}

func (p *Location) Sum(l Location) {
	p.X += l.X
	p.Y += l.Y
}

func main() {
	// Declarando e atribuindo valores usando os nomes dos campos
	student := Student{
		Name: "Pery",
		Age:  27,
	}

	fmt.Println("Aluno:", student)

	// Declarando e atribuindo valores diretamente
	myHouse := Location{3, 4, "casa", 500}

	// Atribuindo uma estrutura vazia
	emptyStruct := Student{}

	// Mais uma vez declarando e atribuindo usando o nomes dos campos
	schoolLocation := Location{
		Y:    100,
		X:    200,
		Name: "escola",
	}

	// Alocando em seguida atribuindo valores
	var anotherLocation Location
	anotherLocation.X = 10
	anotherLocation.Y = 20
	anotherLocation.Name = "trabalho"

	fmt.Println("Minha casa:", myHouse)
	fmt.Println("Outra localidade:", anotherLocation)
	fmt.Println("Escola:", schoolLocation)
	fmt.Println("Aluno:", student)
	fmt.Printf("Estrutura vazia: %q\r\n", emptyStruct)

	myHouse.Sum(anotherLocation)

	fmt.Println("Localidade da minha casa somada com outra localidade", myHouse)

	fmt.Printf("Minha casa: %v\r\n", myHouse)
	fmt.Printf("Minha casa: %+v\r\n", myHouse)

	// Brincando com JSON
	j, err := json.Marshal(myHouse)
	if err != nil {
		panic(err)
	}

	fmt.Println("Minha casa json:", string(j))

	newLocation := Location{}
	err = json.Unmarshal(j, &newLocation)
	if err != nil {
		panic(err)
	}

	fmt.Println("Pondo depois do Unmarshal:", newLocation)

}
