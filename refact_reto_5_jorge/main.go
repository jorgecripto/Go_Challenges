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
			PokedexID: 150,
			Nombre:    "Mewtwo",
			Tipo:      "Psíquico",
			Habilidades: []Habilidad{
				{Nombre: "Psíquico", Poder: 90},
				{Nombre: "Psicorayo", Poder: 65},
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
			Tipo:      "Fuego",
			Habilidades: []Habilidad{
				{Nombre: "Lanzallamas", Poder: 90},
				{Nombre: "Garra Dragón", Poder: 80},
			},
		},
		{
			PokedexID: 95,
			Nombre:    "Onix",
			Tipo:      "Roca",
			Habilidades: []Habilidad{
				{Nombre: "Avalancha", Poder: 75},
				{Nombre: "Lanzarrocas", Poder: 50},
			},
		},
		{
			PokedexID: 144,
			Nombre:    "Articuno",
			Tipo:      "Hielo",
			Habilidades: []Habilidad{
				{Nombre: "Rayo Hielo", Poder: 90},
				{Nombre: "Ventisca", Poder: 110},
			},
		},
		{
			PokedexID: 94,
			Nombre:    "Gengar",
			Tipo:      "Fantasma",
			Habilidades: []Habilidad{
				{Nombre: "Bola Sombra", Poder: 80},
				{Nombre: "Puño Sombra", Poder: 60},
			},
		},
		{
			PokedexID: 143,
			Nombre:    "Snorlax",
			Tipo:      "Normal",
			Habilidades: []Habilidad{
				{Nombre: "Golpe Cuerpo", Poder: 85},
				{Nombre: "Descanso", Poder: 0},
			},
		},
		{
			PokedexID: 18,
			Nombre:    "Pidgeot",
			Tipo:      "Volador",
			Habilidades: []Habilidad{
				{Nombre: "Ataque Aéreo", Poder: 60},
				{Nombre: "Huracán", Poder: 110},
			},
		},
		{
			PokedexID: 131,
			Nombre:    "Lapras",
			Tipo:      "Agua",
			Habilidades: []Habilidad{
				{Nombre: "Hidrobomba", Poder: 110},
				{Nombre: "Canto Helado", Poder: 55},
			},
		},
		{
			PokedexID: 68,
			Nombre:    "Machamp",
			Tipo:      "Lucha",
			Habilidades: []Habilidad{
				{Nombre: "Puño Dinámico", Poder: 100},
				{Nombre: "Golpe Cruzado", Poder: 100},
			},
		},
		{
			PokedexID: 26,
			Nombre:    "Raichu",
			Tipo:      "Eléctrico",
			Habilidades: []Habilidad{
				{Nombre: "Trueno", Poder: 110},
				{Nombre: "Puño Trueno", Poder: 75},
			},
		},
		{
			PokedexID: 89,
			Nombre:    "Muk",
			Tipo:      "Veneno",
			Habilidades: []Habilidad{
				{Nombre: "Bomba Lodo", Poder: 90},
				{Nombre: "Gas Venenoso", Poder: 40},
			},
		},
		{
			PokedexID: 149,
			Nombre:    "Dragonite",
			Tipo:      "Dragón",
			Habilidades: []Habilidad{
				{Nombre: "Furia Dragón", Poder: 120},
				{Nombre: "Danza Dragón", Poder: 0},
			},
		},
		{
			PokedexID: 3,
			Nombre:    "Venusaur",
			Tipo:      "Planta",
			Habilidades: []Habilidad{
				{Nombre: "Rayo Solar", Poder: 120},
				{Nombre: "Bomba Germen", Poder: 80},
			},
		},
		{
			PokedexID: 59,
			Nombre:    "Arcanine",
			Tipo:      "Fuego",
			Habilidades: []Habilidad{
				{Nombre: "Llamarada", Poder: 150},
				{Nombre: "Rueda Fuego", Poder: 60},
			},
		},
		{
			PokedexID: 87,
			Nombre:    "Dewgong",
			Tipo:      "Hielo",
			Habilidades: []Habilidad{
				{Nombre: "Rayo Aurora", Poder: 65},
				{Nombre: "Granizo", Poder: 0},
			},
		},
		{
			PokedexID: 65,
			Nombre:    "Alakazam",
			Tipo:      "Psíquico",
			Habilidades: []Habilidad{
				{Nombre: "Psicorrayo", Poder: 65},
				{Nombre: "Futuro Visto", Poder: 120},
			},
		},
		{
			PokedexID: 76,
			Nombre:    "Golem",
			Tipo:      "Roca",
			Habilidades: []Habilidad{
				{Nombre: "Terremoto", Poder: 100},
				{Nombre: "Explosión", Poder: 250},
			},
		},
		{
			PokedexID: 9,
			Nombre:    "Blastoise",
			Tipo:      "Agua",
			Habilidades: []Habilidad{
				{Nombre: "Hidrobomba", Poder: 110},
				{Nombre: "Pulso Agua", Poder: 60},
			},
		},
		{
			PokedexID: 22,
			Nombre:    "Fearow",
			Tipo:      "Volador",
			Habilidades: []Habilidad{
				{Nombre: "Pico Taladro", Poder: 80},
				{Nombre: "Ataque Aéreo", Poder: 60},
			},
		},
		{
			PokedexID: 57,
			Nombre:    "Primeape",
			Tipo:      "Lucha",
			Habilidades: []Habilidad{
				{Nombre: "Puño Certero", Poder: 150},
				{Nombre: "Enfado", Poder: 20},
			},
		},
		{
			PokedexID: 38,
			Nombre:    "Ninetales",
			Tipo:      "Fuego",
			Habilidades: []Habilidad{
				{Nombre: "Pulso Fuego", Poder: 70},
				{Nombre: "Colmillo Ígneo", Poder: 65},
			},
		},
		{
			PokedexID: 102,
			Nombre:    "Exeggcute",
			Tipo:      "Planta",
			Habilidades: []Habilidad{
				{Nombre: "Semilla Bomba", Poder: 80},
				{Nombre: "Confusión", Poder: 50},
			},
		},
		{
			PokedexID: 124,
			Nombre:    "Jynx",
			Tipo:      "Hielo",
			Habilidades: []Habilidad{
				{Nombre: "Puño Hielo", Poder: 75},
				{Nombre: "Beso Amoroso", Poder: 0},
			},
		},
		{
			PokedexID: 145,
			Nombre:    "Zapdos",
			Tipo:      "Eléctrico",
			Habilidades: []Habilidad{
				{Nombre: "Rayo", Poder: 90},
				{Nombre: "Chispazo", Poder: 80},
			},
		},
		{
			PokedexID: 28,
			Nombre:    "Sandslash",
			Tipo:      "Tierra",
			Habilidades: []Habilidad{
				{Nombre: "Excavar", Poder: 80},
				{Nombre: "Garra", Poder: 40},
			},
		},
		{
			PokedexID: 73,
			Nombre:    "Tentacruel",
			Tipo:      "Agua",
			Habilidades: []Habilidad{
				{Nombre: "Hidropulso", Poder: 60},
				{Nombre: "Ácido", Poder: 40},
			},
		},
		{
			PokedexID: 24,
			Nombre:    "Arbok",
			Tipo:      "Veneno",
			Habilidades: []Habilidad{
				{Nombre: "Mordisco", Poder: 60},
				{Nombre: "Colmillo Veneno", Poder: 50},
			},
		},
		{
			PokedexID: 104,
			Nombre:    "Cubone",
			Tipo:      "Tierra",
			Habilidades: []Habilidad{
				{Nombre: "Huesomerang", Poder: 50},
				{Nombre: "Cabezazo", Poder: 70},
			},
		},
		{
			PokedexID: 101,
			Nombre:    "Electrode",
			Tipo:      "Eléctrico",
			Habilidades: []Habilidad{
				{Nombre: "Autodestrucción", Poder: 200},
				{Nombre: "Chispa", Poder: 65},
			},
		},
		{
			PokedexID: 106,
			Nombre:    "Hitmonlee",
			Tipo:      "Lucha",
			Habilidades: []Habilidad{
				{Nombre: "Patada Salto", Poder: 130},
				{Nombre: "Megapatada", Poder: 120},
			},
		},
		{
			PokedexID: 36,
			Nombre:    "Clefable",
			Tipo:      "Normal",
			Habilidades: []Habilidad{
				{Nombre: "Metrónomo", Poder: 0},
				{Nombre: "Vozarrón", Poder: 90},
			},
		},
		{
			PokedexID: 78,
			Nombre:    "Rapidash",
			Tipo:      "Fuego",
			Habilidades: []Habilidad{
				{Nombre: "Rueda Fuego", Poder: 60},
				{Nombre: "Llamarada", Poder: 150},
			},
		},
		{
			PokedexID: 103,
			Nombre:    "Exeggutor",
			Tipo:      "Planta",
			Habilidades: []Habilidad{
				{Nombre: "Cabezazo Zen", Poder: 80},
				{Nombre: "Psíquico", Poder: 90},
			},
		},
		{
			PokedexID: 91,
			Nombre:    "Cloyster",
			Tipo:      "Agua",
			Habilidades: []Habilidad{
				{Nombre: "Carámbano", Poder: 85},
				{Nombre: "Hidrobomba", Poder: 110},
			},
		},
		{
			PokedexID: 142,
			Nombre:    "Aerodactyl",
			Tipo:      "Roca",
			Habilidades: []Habilidad{
				{Nombre: "Avalancha", Poder: 75},
				{Nombre: "Colmillo Roca", Poder: 65},
			},
		},
		{
			PokedexID: 107,
			Nombre:    "Hitmonchan",
			Tipo:      "Lucha",
			Habilidades: []Habilidad{
				{Nombre: "Puño Fuego", Poder: 75},
				{Nombre: "Puño Hielo", Poder: 75},
			},
		},
		{
			PokedexID: 93,
			Nombre:    "Haunter",
			Tipo:      "Fantasma",
			Habilidades: []Habilidad{
				{Nombre: "Lengüetazo", Poder: 30},
				{Nombre: "Bola Sombra", Poder: 80},
			},
		},
		{
			PokedexID: 85,
			Nombre:    "Dodrio",
			Tipo:      "Volador",
			Habilidades: []Habilidad{
				{Nombre: "Pico Taladro", Poder: 80},
				{Nombre: "Tri Ataque", Poder: 80},
			},
		},
		{
			PokedexID: 105,
			Nombre:    "Marowak",
			Tipo:      "Tierra",
			Habilidades: []Habilidad{
				{Nombre: "Porra", Poder: 65},
				{Nombre: "Terremoto", Poder: 100},
			},
		},
		{
			PokedexID: 123,
			Nombre:    "Scyther",
			Tipo:      "Bicho",
			Habilidades: []Habilidad{
				{Nombre: "Cuchillada", Poder: 70},
				{Nombre: "Tajo Aéreo", Poder: 75},
			},
		},
		{
			PokedexID: 82,
			Nombre:    "Magnezone",
			Tipo:      "Eléctrico",
			Habilidades: []Habilidad{
				{Nombre: "Rayo Carga", Poder: 50},
				{Nombre: "Electrocañón", Poder: 120},
			},
		},
		{
			PokedexID: 146,
			Nombre:    "Moltres",
			Tipo:      "Fuego",
			Habilidades: []Habilidad{
				{Nombre: "Tormenta Ígnea", Poder: 110},
				{Nombre: "Envite Ígneo", Poder: 120},
			},
		},
		{
			PokedexID: 51,
			Nombre:    "Dugtrio",
			Tipo:      "Tierra",
			Habilidades: []Habilidad{
				{Nombre: "Excavar", Poder: 80},
				{Nombre: "Magnitud", Poder: 70},
			},
		},
		{
			PokedexID: 121,
			Nombre:    "Starmie",
			Tipo:      "Agua",
			Habilidades: []Habilidad{
				{Nombre: "Hidropulso", Poder: 60},
				{Nombre: "Psicorrayo", Poder: 65},
			},
		},
		{
			PokedexID: 112,
			Nombre:    "Rhydon",
			Tipo:      "Roca",
			Habilidades: []Habilidad{
				{Nombre: "Perforador", Poder: 80},
				{Nombre: "Megacuerno", Poder: 120},
			},
		},
		{
			PokedexID: 125,
			Nombre:    "Electabuzz",
			Tipo:      "Eléctrico",
			Habilidades: []Habilidad{
				{Nombre: "Puño Trueno", Poder: 75},
				{Nombre: "Rayo", Poder: 90},
			},
		},
		{
			PokedexID: 44,
			Nombre:    "Gloom",
			Tipo:      "Planta",
			Habilidades: []Habilidad{
				{Nombre: "Polvo Veneno", Poder: 0},
				{Nombre: "Absorber", Poder: 20},
			},
		},
		{
			PokedexID: 113,
			Nombre:    "Chansey",
			Tipo:      "Normal",
			Habilidades: []Habilidad{
				{Nombre: "Huevo Bomba", Poder: 100},
				{Nombre: "Suavidad", Poder: 0},
			},
		},
		{
			PokedexID: 126,
			Nombre:    "Magmar",
			Tipo:      "Fuego",
			Habilidades: []Habilidad{
				{Nombre: "Puño Fuego", Poder: 75},
				{Nombre: "Llamarada", Poder: 150},
			},
		},
		{
			PokedexID: 96,
			Nombre:    "Drowzee",
			Tipo:      "Psíquico",
			Habilidades: []Habilidad{
				{Nombre: "Confusión", Poder: 50},
				{Nombre: "Psicorrayo", Poder: 65},
			},
		},
		{
			PokedexID: 40,
			Nombre:    "Wigglytuff",
			Tipo:      "Normal",
			Habilidades: []Habilidad{
				{Nombre: "Doble Bofetón", Poder: 15},
				{Nombre: "Vozarrón", Poder: 90},
			},
		},
		{
			PokedexID: 71,
			Nombre:    "Victreebel",
			Tipo:      "Planta",
			Habilidades: []Habilidad{
				{Nombre: "Hoja Afilada", Poder: 55},
				{Nombre: "Ácido", Poder: 40},
			},
		},
		{
			PokedexID: 127,
			Nombre:    "Pinsir",
			Tipo:      "Bicho",
			Habilidades: []Habilidad{
				{Nombre: "Tijera X", Poder: 80},
				{Nombre: "Agarre", Poder: 55},
			},
		},
		{
			PokedexID: 84,
			Nombre:    "Doduo",
			Tipo:      "Volador",
			Habilidades: []Habilidad{
				{Nombre: "Picotazo", Poder: 35},
				{Nombre: "Ataque Furia", Poder: 15},
			},
		},
		{
			PokedexID: 97,
			Nombre:    "Hypno",
			Tipo:      "Psíquico",
			Habilidades: []Habilidad{
				{Nombre: "Psíquico", Poder: 90},
				{Nombre: "Hipnosis", Poder: 0},
			},
		},
		{
			PokedexID: 115,
			Nombre:    "Kangaskhan",
			Tipo:      "Normal",
			Habilidades: []Habilidad{
				{Nombre: "Megapuño", Poder: 80},
				{Nombre: "Golpe", Poder: 40},
			},
		},
		{
			PokedexID: 34,
			Nombre:    "Nidoking",
			Tipo:      "Veneno",
			Habilidades: []Habilidad{
				{Nombre: "Megacuerno", Poder: 120},
				{Nombre: "Doble Patada", Poder: 30},
			},
		},
		{
			PokedexID: 53,
			Nombre:    "Persian",
			Tipo:      "Normal",
			Habilidades: []Habilidad{
				{Nombre: "Cuchillada", Poder: 70},
				{Nombre: "Arañazo", Poder: 40},
			},
		},
		{
			PokedexID: 117,
			Nombre:    "Seadra",
			Tipo:      "Agua",
			Habilidades: []Habilidad{
				{Nombre: "Hidropulso", Poder: 60},
				{Nombre: "Rayo Aurora", Poder: 65},
			},
		},
		{
			PokedexID: 148,
			Nombre:    "Dragonair",
			Tipo:      "Dragón",
			Habilidades: []Habilidad{
				{Nombre: "Pulso Dragón", Poder: 85},
				{Nombre: "Ciclón", Poder: 40},
			},
		},
		{
			PokedexID: 55,
			Nombre:    "Golduck",
			Tipo:      "Agua",
			Habilidades: []Habilidad{
				{Nombre: "Psicorrayo", Poder: 65},
				{Nombre: "Confusión", Poder: 50},
			},
		},
		{
			PokedexID: 74,
			Nombre:    "Geodude",
			Tipo:      "Roca",
			Habilidades: []Habilidad{
				{Nombre: "Lanzarrocas", Poder: 50},
				{Nombre: "Autorrodar", Poder: 30},
			},
		},
		{
			PokedexID: 83,
			Nombre:    "Farfetch'd",
			Tipo:      "Volador",
			Habilidades: []Habilidad{
				{Nombre: "Corte", Poder: 50},
				{Nombre: "Ataque Aéreo", Poder: 60},
			},
		},
		{
			PokedexID: 137,
			Nombre:    "Porygon",
			Tipo:      "Normal",
			Habilidades: []Habilidad{
				{Nombre: "Psicorrayo", Poder: 65},
				{Nombre: "Conversión", Poder: 0},
			},
		},
		{
			PokedexID: 110,
			Nombre:    "Weezing",
			Tipo:      "Veneno",
			Habilidades: []Habilidad{
				{Nombre: "Bomba Lodo", Poder: 90},
				{Nombre: "Explosión", Poder: 250},
			},
		},
		{
			PokedexID: 147,
			Nombre:    "Dratini",
			Tipo:      "Dragón",
			Habilidades: []Habilidad{
				{Nombre: "Pulso Dragón", Poder: 85},
				{Nombre: "Ciclón", Poder: 40},
			},
		},
		{
			PokedexID: 92,
			Nombre:    "Gastly",
			Tipo:      "Fantasma",
			Habilidades: []Habilidad{
				{Nombre: "Lengüetazo", Poder: 30},
				{Nombre: "Maldición", Poder: 0},
			},
		},
		{
			PokedexID: 75,
			Nombre:    "Graveler",
			Tipo:      "Roca",
			Habilidades: []Habilidad{
				{Nombre: "Avalancha", Poder: 75},
				{Nombre: "Autodestrucción", Poder: 200},
			},
		},
		{
			PokedexID: 120,
			Nombre:    "Staryu",
			Tipo:      "Agua",
			Habilidades: []Habilidad{
				{Nombre: "Pistola Agua", Poder: 40},
				{Nombre: "Rapidez", Poder: 60},
			},
		},
		{
			PokedexID: 128,
			Nombre:    "Tauros",
			Tipo:      "Normal",
			Habilidades: []Habilidad{
				{Nombre: "Cornada", Poder: 65},
				{Nombre: "Golpe", Poder: 40},
			},
		},
		{
			PokedexID: 31,
			Nombre:    "Nidoqueen",
			Tipo:      "Veneno",
			Habilidades: []Habilidad{
				{Nombre: "Megacuerno", Poder: 120},
				{Nombre: "Doble Patada", Poder: 30},
			},
		},
		{
			PokedexID: 108,
			Nombre:    "Lickitung",
			Tipo:      "Normal",
			Habilidades: []Habilidad{
				{Nombre: "Lengüetazo", Poder: 30},
				{Nombre: "Golpe Cuerpo", Poder: 85},
			},
		},
		{
			PokedexID: 122,
			Nombre:    "Mr. Mime",
			Tipo:      "Psíquico",
			Habilidades: []Habilidad{
				{Nombre: "Psicorrayo", Poder: 65},
				{Nombre: "Barrera", Poder: 0},
			},
		},
		{
			PokedexID: 12,
			Nombre:    "Butterfree",
			Tipo:      "Bicho",
			Habilidades: []Habilidad{
				{Nombre: "Polvo Veneno", Poder: 0},
				{Nombre: "Psicorrayo", Poder: 65},
			},
		},
		{
			PokedexID: 20,
			Nombre:    "Raticate",
			Tipo:      "Normal",
			Habilidades: []Habilidad{
				{Nombre: "Superdiente", Poder: 80},
				{Nombre: "Ataque Rápido", Poder: 40},
			},
		},
		{
			PokedexID: 86,
			Nombre:    "Seel",
			Tipo:      "Agua",
			Habilidades: []Habilidad{
				{Nombre: "Rayo Aurora", Poder: 65},
				{Nombre: "Cabezazo", Poder: 70},
			},
		},
		{
			PokedexID: 15,
			Nombre:    "Beedrill",
			Tipo:      "Bicho",
			Habilidades: []Habilidad{
				{Nombre: "Picadura", Poder: 60},
				{Nombre: "Pin Misil", Poder: 25},
			},
		},
		{
			PokedexID: 80,
			Nombre:    "Slowbro",
			Tipo:      "Psíquico",
			Habilidades: []Habilidad{
				{Nombre: "Psíquico", Poder: 90},
				{Nombre: "Confusión", Poder: 50},
			},
		},
		{
			PokedexID: 139,
			Nombre:    "Omastar",
			Tipo:      "Roca",
			Habilidades: []Habilidad{
				{Nombre: "Hidropulso", Poder: 60},
				{Nombre: "Avalancha", Poder: 75},
			},
		},
		{
			PokedexID: 62,
			Nombre:    "Poliwrath",
			Tipo:      "Agua",
			Habilidades: []Habilidad{
				{Nombre: "Puño Dinámico", Poder: 100},
				{Nombre: "Hidropulso", Poder: 60},
			},
		},
		{
			PokedexID: 130,
			Nombre:    "Gyarados",
			Tipo:      "Agua",
			Habilidades: []Habilidad{
				{Nombre: "Hidrobomba", Poder: 110},
				{Nombre: "Furia Dragón", Poder: 120},
			},
		},
		{
			PokedexID: 79,
			Nombre:    "Slowpoke",
			Tipo:      "Psíquico",
			Habilidades: []Habilidad{
				{Nombre: "Confusión", Poder: 50},
				{Nombre: "Psicorrayo", Poder: 65},
			},
		},
		{
			PokedexID: 81,
			Nombre:    "Magnemite",
			Tipo:      "Eléctrico",
			Habilidades: []Habilidad{
				{Nombre: "Chispa", Poder: 65},
				{Nombre: "Onda Trueno", Poder: 0},
			},
		},
		{
			PokedexID: 46,
			Nombre:    "Paras",
			Tipo:      "Bicho",
			Habilidades: []Habilidad{
				{Nombre: "Arañazo", Poder: 40},
				{Nombre: "Espora", Poder: 0},
			},
		},
		{
			PokedexID: 134,
			Nombre:    "Vaporeon",
			Tipo:      "Agua",
			Habilidades: []Habilidad{
				{Nombre: "Hidropulso", Poder: 60},
				{Nombre: "Rayo Aurora", Poder: 65},
			},
		},
		{
			PokedexID: 135,
			Nombre:    "Jolteon",
			Tipo:      "Eléctrico",
			Habilidades: []Habilidad{
				{Nombre: "Impactrueno", Poder: 40},
				{Nombre: "Pin Misil", Poder: 25},
			},
		},
		{
			PokedexID: 136,
			Nombre:    "Flareon",
			Tipo:      "Fuego",
			Habilidades: []Habilidad{
				{Nombre: "Ascuas", Poder: 40},
				{Nombre: "Llamarada", Poder: 150},
			},
		},
	}

	//fmt.Println(pokemons)

	myTeamPokemon := CreateMyPokemonTeam(pokemons)
	fmt.Println(myTeamPokemon)

}
