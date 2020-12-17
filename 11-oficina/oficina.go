package main

import "fmt"

type Empleado struct {
	Nombre string
	Edad int
	Sueldo float64
	LlamadasAtendidas int
}

type Administrativo struct {
	Empleado
}

type Duenio struct {
	Empleado
}

type Supervisor struct {
	Empleado
	AdministrativosACargo []Administrativo
}

func (a *Empleado) atenderLlamada()  {
	a.LlamadasAtendidas += 1
	fmt.Println("Hola habla el Administrativo")
}

func (d *Duenio) felicitarSupervisor(e *Empleado) {
	fmt.Println("el dueño felicita a", e.Nombre)
}

func (s *Supervisor) darAumentoAdministrativo(a *Administrativo)  {
	a.Sueldo += 200
}

func main() {
	//fmt.Printf("%+v %+v %+v \n", Administrativo, Supervisor, Duenio)
}
