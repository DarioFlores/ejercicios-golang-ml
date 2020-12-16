package main

import (
	"fmt"
	"github.com/DarioFlores/09-twitter/datos"
	"github.com/DarioFlores/09-twitter/persona"
)

func main() {

	t1 := datos.Tweet{Texto: "Hola Mundo"}
	t2 := datos.Tweet{Texto: "Hola Mundo"}
	t3 := datos.Tweet{Texto: "Hola Mundo"}
	t4 := datos.Tweet{Texto: "Hola Mundo"}
	t5 := datos.Tweet{Texto: "Hola Mundo"}
	t6 := datos.Tweet{Texto: "Hola Mundo"}
	t7 := datos.Tweet{Texto: "Hola Mundo"}

	p1 := persona.Persona{Nombre: "Dario", Tweets: []datos.Tweet{t1,t2,t3}}
	p2 := persona.Persona{Nombre: "Gisella", Tweets: []datos.Tweet{t4,t5,t6,t7}}
	fmt.Println(p1)
	fmt.Println(p2)
	personaQueMasTwitteo(p1,p2)
	fmt.Println(p1.Nombre, " - ", p1.TodosLosTweets())
	fmt.Println(p1.Nombre, " - ", p1.CantidadDeTweets())
}

func personaQueMasTwitteo(personas ...persona.Persona)  {
	var auxPersona persona.Persona;
	mayor := 0;
	for _, pers := range personas{
		if mayor < pers.CantidadDeTweets() {
			auxPersona = pers;
		}
	}
	fmt.Println(auxPersona.Nombre);
}