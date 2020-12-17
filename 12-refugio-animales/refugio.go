package main

import "fmt"

type Posicionable interface {
	desplazar(x,y int)
	obtenerPosicion() (x,y int)
}

type Posicion struct {
	X int
	Y int
}

func (p *Posicion) desplazar(x,y int)  {
	p.X = x
	p.Y = y
}

func (p *Posicion) obtenerPosicion() (x,y int) {
	return p.X, p.Y
}

type Cuidador struct {
	Posicion
	Nombre string
}

func (c *Cuidador) desplazar(x,y int)  {
	fmt.Println("Me desplazo caminando")
	c.X = x
	c.Y = y
}

func (c *Cuidador) pasearAnimales(x,y int, animales ...Posicionable) {
	for _, animal := range animales {
		animal.desplazar(x,y)
	}
	c.desplazar(x,y)
}

type Animal struct {
	Posicion
	Nombre string
}

type Perro struct {
	Animal
	Raza string
}

func (p *Perro) desplazar(x,y int)  {
	fmt.Println("Me desplazo corriendo")
	p.X = x
	p.Y = y
}

type Loro struct {
	Animal
	Palabras []string
}

func (l *Loro) desplazar(x,y int)  {
	fmt.Println("Me desplazo volar")
	l.X = x
	l.Y = y
}

func main() {

}
