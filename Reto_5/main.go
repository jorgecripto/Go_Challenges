package main

import "fmt"

func main() {

	pokemons := []Pokemon{
		{
			PokedexID: 25,
			Nombre:    "Pikachu",
			Tipo:      "Eléctrico",
			Habilidades: []Habilidad{
				{Nombre: "Impactrueno", Poder: 40},
				{Nombre: "Ataque Rápido", Poder: 40},
			},
		},
		{
			PokedexID: 1,
			Nombre:    "Bulbasaur",
			Tipo:      "Planta",
			Habilidades: []Habilidad{
				{Nombre: "Látigo Cepa", Poder: 45},
				{Nombre: "Drenadoras", Poder: 20},
			},
		},
		{
			PokedexID: 7,
			Nombre:    "Squirtle",
			Tipo:      "Agua",
			Habilidades: []Habilidad{
				{Nombre: "Pistola Agua", Poder: 40},
				{Nombre: "Burbuja", Poder: 20},
			},
		},
		{
			PokedexID: 6,
			Nombre:    "Charizard",
			Tipo:      "Fuego/Volador",
			Habilidades: []Habilidad{
				{Nombre: "Lanzallamas", Poder: 90},
				{Nombre: "Garra Dragón", Poder: 80},
			},
		},
		{
			PokedexID: 9,
			Nombre:    "Blastoise",
			Tipo:      "Agua",
			Habilidades: []Habilidad{
				{Nombre: "Hidrobomba", Poder: 110},
				{Nombre: "Rayo Hielo", Poder: 90},
			},
		},
		{
			PokedexID: 4,
			Nombre:    "Charmander",
			Tipo:      "Fuego",
			Habilidades: []Habilidad{
				{Nombre: "Ascuas", Poder: 40},
				{Nombre: "Lanzallamas", Poder: 90},
			},
		},
	}

	//!Al iniciar el reto esta linea se debe comentar
	//fmt.Println(pokemons)

	var message string = `Hola soy Jorge Iván, este es mi equipo pokemón:
	1. %s, su tipo es %s y su habilidad inicial es: %s
	2. %s, su tipo es %s y su habilidad inicial es: %s
	3. %s, su tipo es %s y su habilidad inicial es: %s`

	newMessage := MyPokemonTeam(message, pokemons)
	fmt.Println(newMessage)

}

//Para este reto usted sera un entrenador pokemon novato, por tal motivo
//debera formar su equipo pokemon inicial, el cual debe estar compuesto por:
// - 1 pokemón de tipo fuego
// - 1 pokemón de tipo agua
// - 1 pokemón de tipo planta
//
//Al ejecutar su programa debera mostrar un texto compuesto de la siguiente
//manera:
//
//Hola soy ${nombre del entrenador}, este es mi equipo pokemón:
//
// 1. ${nombre del primer pokemón}, su tipo es: ${tipo del pokemón} y su habilidad inicial es:
//${nombre de la primera habilidad}.
//
// 2. ${nombre del segundo pokemón}, su tipo es: ${tipo del pokemón} y su habilidad inicial es:
//${nombre de la primera habilidad}.
//
// 3. ${nombre del tercer pokemón}, su tipo es: ${tipo del pokemón} y su habilidad inicial es:
//${nombre de la primera habilidad}.
