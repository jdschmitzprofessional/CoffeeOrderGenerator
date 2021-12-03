package main

import (
	"CoffeeOrderGenerator/Structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var coffee Structs.Coffee
	jsonInput, err := ioutil.ReadFile("resources/coffee.json")
	checkError(err)
	json.Unmarshal(jsonInput, &coffee)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		responseString := generateCoffeeOrder(coffee)
		fmt.Fprint(writer, responseString)
	})
	server := http.Server{
		Addr: ":8082",
	}
	fmt.Println("Starting server")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getRandomElement(array []string) string {
	return array[rand.Intn(len(array))]
}

func generateCoffeeOrder(coffee Structs.Coffee) string {
	i := rand.Intn(2)
	switch i {
	case 1:
		return fmt.Sprintf(
			"I'd like a %s %s %s with %s %s and %s",
			getRandomElement(coffee.Roasts),
			getRandomElement(coffee.Sizes),
			getRandomElement(coffee.Types),
			getRandomElement(coffee.Modifiers),
			getRandomElement(coffee.Condiments),
			getRandomElement(coffee.Condiments),
		)
	default:
		return fmt.Sprintf("I'd like %s %s %s %s, one with %s %s and the others with %s, %s on the side for each of them with %s %s.",
			getRandomElement(coffee.Quantities),
			getRandomElement(coffee.SizesPlural),
			getRandomElement(coffee.Roasts),
			getRandomElement(coffee.TypesPlural),
			getRandomElement(coffee.Modifiers),
			getRandomElement(coffee.Condiments),
			getRandomElement(coffee.Modifiers),
			getRandomElement(coffee.Condiments),
			getRandomElement(coffee.Modifiers),
			getRandomElement(coffee.Condiments),
		)
	}
}
