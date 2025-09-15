package main

import (
	"fmt"
)

func CreatePokemonTeam(pokemons []Pokemon) string {

	const typeFire = "Fuego"
	const typeWater = "Agua"
	const typePlant = "Planta"

	message := `Hola soy Jorge Iván, este es mi equipo pokemon:`
	message2 := ``
	baseMessage := ` %v. %s, su tipo es: %s y su habilidad inicial es: %s;`
	newMessage := ``

	counter := 0
	waterCounter := 0
	fireCounter := 0
	plantCounter := 0

	for _, pokemon := range pokemons {

		if counter > 3 {
			break
		}

		if pokemon.Tipo == typeFire && fireCounter < 1 {

			fireCounter++
			message2 = fmt.Sprintf(baseMessage, counter+1, pokemon.Nombre, pokemon.Tipo, pokemon.Habilidades[0].Nombre)
			if counter < 1 {
				newMessage = message + message2
			} else {
				newMessage = newMessage + message2
			}
			counter++

		}
		if pokemon.Tipo == typeWater && waterCounter < 1 {

			waterCounter++
			message2 = fmt.Sprintf(baseMessage, counter+1, pokemon.Nombre, pokemon.Tipo, pokemon.Habilidades[0].Nombre)
			if counter < 1 {
				newMessage = message + message2
			} else {
				newMessage = newMessage + message2
			}
			counter++

		}
		if pokemon.Tipo == typePlant && plantCounter < 1 {

			plantCounter++
			message2 = fmt.Sprintf(baseMessage, counter+1, pokemon.Nombre, pokemon.Tipo, pokemon.Habilidades[0].Nombre)
			if counter < 1 {
				newMessage = message + message2
			} else {
				newMessage = newMessage + message2
			}
			counter++

		}

	}

	return newMessage

}
