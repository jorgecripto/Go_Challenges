package main

import "fmt"

const typeFire = "Eléctrico"
const typeWater = "Bicho"
const typePlant = "Psíquico"

func CreateMyPokemonTeam(pokemons []Pokemon) string {
	var finalMessage string

	pokemonsName := []string{"", "", ""}
	pokemonsType := []string{"", "", ""}
	pokemonsHabilities := []string{"", "", ""}
	counter := 0

	for _, pokemon := range pokemons {

		if pokemon.Tipo == typeFire || pokemon.Tipo == typeWater || pokemon.Tipo == typePlant {
			if pokemonsType[0] != pokemon.Tipo && pokemonsType[1] != pokemon.Tipo && pokemonsType[2] != pokemon.Tipo {
				pokemonsName[counter] = pokemon.Nombre
				pokemonsType[counter] = pokemon.Tipo
				pokemonsHabilities[counter] = pokemon.Habilidades[0].Nombre
				counter++
			}
			if counter == 3 {
				break
			}
		}
	}
	finalMessage = fmt.Sprintf(`Hola soy Jorge Iván, este es mi equipo Pokémon:
	1. %s, su tipo es: %s y su habilidad inicial es: %s.
	2. %s, su tipo es: %s y su habilidad inicial es: %s.
	3. %s, su tipo es: %s y su habilidad inicial es: %s.`,
	pokemonsName[0], pokemonsType[0], pokemonsHabilities[0],
	pokemonsName[1], pokemonsType[1], pokemonsHabilities[1],
	pokemonsName[2], pokemonsType[2], pokemonsHabilities[2])

	return finalMessage 
}