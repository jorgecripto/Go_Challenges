package main

import "fmt"

const type1 = "Fuego"
const type2 = "Agua"
const type3 = "Planta"

func MyPokemonTeam(message string, pokemons []Pokemon) string {
	var (
	name_1 string
	name_2 string
	name_3 string

	clase_1 string
	clase_2 string
	clase_3 string

	hability_1 string
	hability_2 string
	hability_3 string
	)
	var counterWater int 
	var counterFire int
	var counterPlant int
	var totalCounter int
	for _, pokemon := range pokemons {

		if pokemon.Tipo == type1 ||
			pokemon.Tipo == type2 ||
			pokemon.Tipo == type3 {
			if counterWater + counterFire + counterPlant == 3 {break}

			if totalCounter < 3 {
				if totalCounter == 0{
				name_1 = pokemon.Nombre
					if pokemon.Tipo == type2 && counterWater < 1{
					clase_1 = pokemon.Tipo
					CheckType(clase_1, &counterWater, &counterFire, &counterPlant)
				}
					if pokemon.Tipo == type1 && counterFire < 1{
					clase_1 = pokemon.Tipo
					CheckType(clase_1, &counterWater, &counterFire, &counterPlant)
				}
					if pokemon.Tipo == type3 && counterPlant < 1{
					clase_1 = pokemon.Tipo
					CheckType(clase_1, &counterWater, &counterFire, &counterPlant)
				}
				hability_1 = pokemon.Habilidades[0].Nombre
				}
				if totalCounter == 1{
				name_2 = pokemon.Nombre
				if pokemon.Tipo == type2 && counterWater < 1{
					clase_2 = pokemon.Tipo
					CheckType(clase_2, &counterWater, &counterFire, &counterPlant)
				}
					if pokemon.Tipo == type1 && counterFire < 1{
					clase_2 = pokemon.Tipo
					CheckType(clase_2, &counterWater, &counterFire, &counterPlant)
				}
					if pokemon.Tipo == type3 && counterPlant < 1{
					clase_2 = pokemon.Tipo
					CheckType(clase_2, &counterWater, &counterFire, &counterPlant)
				}
				hability_2 = pokemon.Habilidades[0].Nombre
				
				}
				if totalCounter == 2{
				name_3 = pokemon.Nombre
				if pokemon.Tipo == type2 && counterWater < 1{
					clase_3 = pokemon.Tipo
					CheckType(clase_3, &counterWater, &counterFire, &counterPlant)
				}
					if pokemon.Tipo == type1 && counterFire < 1{
					clase_3 = pokemon.Tipo
					CheckType(clase_3, &counterWater, &counterFire, &counterPlant)
				}
					if pokemon.Tipo == type3 && counterPlant < 1{
					clase_3 = pokemon.Tipo
					CheckType(clase_3, &counterWater, &counterFire, &counterPlant)
				}
				hability_3 = pokemon.Habilidades[0].Nombre
							
				}
				totalCounter = counterWater + counterFire + counterPlant
			}
			
		}

	}
	finalMessage := fmt.Sprintf(message, name_1, clase_1, hability_1, name_2, clase_2, hability_2, name_3, clase_3, hability_3)
	return finalMessage
}


func CheckType(clase string, counterWater *int, counterFire *int, counterPlant *int) {
	switch clase {
	case "Agua": 
		*counterWater = *counterWater + 1
	case "Fuego":
		*counterFire = *counterFire + 1
	case "Planta":
		*counterPlant = *counterPlant + 1		
	}
}
