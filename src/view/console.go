package view

import "fmt"

func ShowMenu() {
	fmt.Println("Escoja una opción entre 1 y 4")
	fmt.Println("1. Lista nombres")
	fmt.Println("2. Pokémon")
	fmt.Println("3. Palabras Amigas")
	fmt.Println("4. Salir")
}

func GetNames() {
	fmt.Println("Introduzca nombres separados por coma")
	fmt.Println("Ejemplo: nombre1,nombre2,nombre3")
}

func GetPokemonId() {
	fmt.Println("Introduzca id de Pokémon")
	fmt.Println("Ejemplo: 25")
}

func GetWords() {
	fmt.Println("Introduzca 2 palabras separadas por un espacio para saber si son amigas")
	fmt.Println("Ejemplo: tokyo kyoto")
}

func ShowResponse(a ...any) {
	fmt.Println(a...)
}

func ShowSeparator() {
	fmt.Println("\n-----------------------------")
}
