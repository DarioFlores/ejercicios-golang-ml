package main

import "fmt"

type Mascota struct {
	Nombre string
}

type Persona struct {
	Mascotas []Mascota
}

func cantidadDeMascotasDeUnaPersona(p *Persona) int {
	return len(p.Mascotas);
}

func agregarMascotaAUnaPersona(p *Persona, m Mascota)  {
	p.Mascotas = append(p.Mascotas, m);
}

func main() {
	m1 := Mascota{Nombre: "Bobby"}
	m2 := Mascota{Nombre: "Camila"}
	p := Persona{Mascotas: []Mascota{m1, m2}}
	fmt.Println(p)
	fmt.Println(cantidadDeMascotasDeUnaPersona(&p))
	m3 := Mascota{Nombre: "Firu"}
	agregarMascotaAUnaPersona(&p, m3)
	fmt.Println(p)
	fmt.Println(cantidadDeMascotasDeUnaPersona(&p))
}
