package almacen

import (
	"errors"
	"fmt"
	"github.com/DarioFlores/10-almacen/producto"
)

type Almacen struct {
	Productos []producto.Producto
}

func CrearAlmacenFactory() Almacen {
	return Almacen{Productos: nil};
}


func (a *Almacen) AgregarProducto(p producto.Producto) (ok bool, err error) {
	for _, p2 := range a.Productos {
		if p2.Nombre == p.Nombre {
			return false, errors.New("Ya existe ese producto")
		}
	}
	if len(p.Categorias) == 0 {
		return false, errors.New("Debe tener al menos una cartegoria")
	}
	a.Productos = append(a.Productos, p)
	return true, nil
}

func (a *Almacen) TodosLosProductos() {
	fmt.Println("Lista de todos los productos")
	for i, p := range a.Productos {
		fmt.Println(i+1, "-")
		fmt.Println("\tNombre:", p.Nombre)
		fmt.Println("\tCantidadDisponible:", p.CantidadDisponible)
	}
}

func (a *Almacen) StockDeProductoPorNombre(nombre string) (stock int, err error) {
	for _, p := range a.Productos {
		if nombre == p.Nombre {
			return p.CantidadDisponible, nil;
		}
	}
	return 0, errors.New("No se encontro el producto")
}

func (a *Almacen) RetirarProducto(nombre string) (ok bool, err error) {
	var pro producto.Producto;
	for _, p := range a.Productos {
		if nombre == p.Nombre {
			pro = p
		}
	}
	if pro.CantidadDisponible != 0 {
		pro.CantidadDisponible = pro.CantidadDisponible - 1;
		return true,nil
	}
	return false, errors.New("No tenemos stock de este producto")
}