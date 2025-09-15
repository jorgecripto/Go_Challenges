package main

type Habilidad struct {
	Nombre string
	Poder  int
}

type Pokemon struct {
	PokedexID   int
	Nombre      string
	Tipo        string
	Habilidades []Habilidad
}
