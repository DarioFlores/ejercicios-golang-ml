package main

import (
	"fmt"
	"github.com/DarioFlores/10-almacen/almacen"
	"github.com/DarioFlores/10-almacen/producto"
)

func main() {
	p1 := producto.Producto{
		Nombre:             "Celular Moto G5 PLUS",
		Despripcion:        "Buena calidad precio",
		CantidadDisponible: 2,
		Categorias: 		[]string{"Informatica", "Tecnologia"},
	}

	p2 := producto.Producto{
		Nombre:             "Celular Moto G5 PLUS",
		Despripcion:        "Buena calidad precio",
		CantidadDisponible: 1,
		Categorias: 		[]string{"Informatica", "Tecnologia"},
	}

	p3 := producto.Producto{
		Nombre:             "Disco externo 1TB",
		Despripcion:        "Buena calidad",
		CantidadDisponible: 4,
		Categorias: 		nil,
	}

	p4 := producto.Producto{
		Nombre:             "Disco externo 1TB",
		Despripcion:        "Buena calidad",
		CantidadDisponible: 3,
		Categorias: 		[]string{"Informatica"},
	}

	p5 := producto.Producto{
		Nombre:             "Mouse Logitech",
		Despripcion:        "Buena calidad",
		CantidadDisponible: 5,
		Categorias: 		[]string{"Informatica"},
	}

	a := almacen.CrearAlmacenFactory();

	ok1, err1 := a.AgregarProducto(p1)
	if err1 != nil {
		fmt.Println("Error al agregar el producto 1", err1)
	}
	fmt.Println("Se agrego correctamente p1?:", ok1)

	ok2, err2 := a.AgregarProducto(p2)
	if err2 != nil {
		fmt.Println("Error al agregar el producto 2", err2)
	}
	fmt.Println("Se agrego correctamente p2?:", ok2)

	ok3, err3 := a.AgregarProducto(p3)
	if err3 != nil {
		fmt.Println("Error al agregar el producto 3", err3)
	}
	fmt.Println("Se agrego correctamente p3?:", ok3)

	ok4, err4 := a.AgregarProducto(p4)
	if err4 != nil {
		fmt.Println("Error al agregar el producto 4", err4)
	}
	fmt.Println("Se agrego correctamente p4?:", ok4)

	ok5, err5 := a.AgregarProducto(p5)
	if err5 != nil {
		fmt.Println("Error al agregar el producto 5", err5)
	}
	fmt.Println("Se agrego correctamente p5?:", ok5)


	fmt.Println(":::::::::::::::::::::::::::::::")
	a.TodosLosProductos();
	fmt.Println(":::::::::::::::::::::::::::::::")

	fmt.Println(":::::::::::::::::::::::::::::::")
	stock, errStock := a.StockDeProductoPorNombre("Disco externo 1TB")
	fmt.Println("El stock es:", stock, errStock)
	stock2, errStock2 := a.StockDeProductoPorNombre("Disco")
	fmt.Println("El stock es:", stock2, errStock2)
	fmt.Println(":::::::::::::::::::::::::::::::")


	fmt.Println(":::::::::::::::::::::::::::::::")
	okRetiro1, errRetiro1 := a.RetirarProducto("Disco externo 1TB")
	fmt.Println("El retiro se realizo?:", okRetiro1, errRetiro1)
	okRetiro2, errRetiro2 := a.RetirarProducto("Disco")
	fmt.Println("El retiro se realizo?:", okRetiro2, errRetiro2)
	fmt.Println(":::::::::::::::::::::::::::::::")
}
