package main

import (
	"fmt"
)

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		ID:       id,
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}

func (p *Producto) ActualizarPrecio(nuevoPrecio float64) float64 {
	p.Precio = nuevoPrecio
	return p.Precio
}
func (p *Producto) ActualizarCantidad(nuevoCantidad int) int {
	p.Cantidad = nuevoCantidad
	return p.Cantidad
}
func (p *Producto) MostrarInfo(existe bool) {
	if existe {
		fmt.Println("El producto es el siguiente: ", p.ID)
		fmt.Println("El nombre del producto es el siguiente: ", p.Nombre)
		fmt.Println("El precio del producto es el siguiente: ", p.Precio)
		fmt.Println("El precio del producto es el siguiente: ", p.Cantidad)
	} else {
		fmt.Println("No existe ese producto")
	}
}

func AgregarProducto(p Producto, inventario map[int]Producto) map[int]Producto {
	inventario[p.ID] = Producto{ID: p.ID, Nombre: p.Nombre, Precio: p.Precio, Cantidad: p.Cantidad}
	return inventario
}
func EliminarProducto(id int, inventario map[int]Producto) {
	producto, existe := inventario[id]
	if existe {
		fmt.Println("Se elimino el producto", producto)
		delete(inventario, id)
	} else {
		fmt.Println("No existe ese producto")
	}
}
func BuscarProducto(id int, inventario map[int]Producto) (Producto, bool) {
	producto, existe := inventario[id]
	return producto, existe
}

func ListarProductos(inventario map[int]Producto) {
	for id, datos := range inventario {
		fmt.Println("ID: ", id, "Datos del producto: ", datos)
	}
}

func main() {
	inventario := make(map[int]Producto)
	var usrOpcion string
	fmt.Println("Bienvenido a este sistema básico de gestión de inventario")
	for {
		fmt.Println("1. Crear o Agregar un nuevo producto")
		fmt.Println("2. Actualizar Precio")
		fmt.Println("3. Actualizar Cantidad")
		fmt.Println("4. Mostar Informacion de un Producto")
		fmt.Println("5. Eliminar un producto")
		fmt.Println("6. Buscar producto")
		fmt.Println("7. Listar todos los productos")
		fmt.Println("0. Para salir ingrese 0")
		fmt.Printf("Que desea hacer: ")
		fmt.Scan(&usrOpcion)

		switch usrOpcion {
		case "1":
			var id int
			var nombre string
			var precio float64
			var cantidad int
			fmt.Printf("Ingrese el id: ")
			fmt.Scan(&id)
			fmt.Printf("Ingrese el nombre: ")
			fmt.Scan(&nombre)
			fmt.Printf("Ingrese el precio: ")
			fmt.Scan(&precio)
			fmt.Printf("Ingrese la cantidad: ")
			fmt.Scan(&cantidad)
			inventario = AgregarProducto(NuevoProducto(id, nombre, precio, cantidad), inventario)
			fmt.Println("Se agrego el producto")
		case "2":
			var id int
			var nuevoPrecio float64
			fmt.Printf("Ingrese el id del producto: ")
			fmt.Scan(&id)
			fmt.Printf("Ingrese el nuevo precio: ")
			fmt.Scan(&nuevoPrecio)
			producto, existe := inventario[id]
			if existe {
				producto.ActualizarPrecio(nuevoPrecio)
				inventario[id] = producto
				fmt.Println("Se actualizo el producto")
			} else {
				fmt.Println("No existe ese producto")
			}
		case "3":
			var id int
			var nuevoCantidad int
			fmt.Printf("Ingrese el id del producto: ")
			fmt.Scan(&id)
			fmt.Printf("Ingrese la nuevo cantidad: ")
			fmt.Scan(&nuevoCantidad)
			producto, existe := inventario[id]
			if existe {
				producto.ActualizarCantidad(nuevoCantidad)
				inventario[id] = producto
				fmt.Println("Se actualizo el producto")
			} else {
				fmt.Println("No existe ese producto")
			}
		case "4":
			var id int
			fmt.Printf("Ingrese el id del producto que quiere ver: ")
			fmt.Scan(&id)
			producto, existe := inventario[id]
			producto.MostrarInfo(existe)

		case "5":
			var id int
			fmt.Printf("Ingrese el id del producto que quiere eliminar: ")
			fmt.Scan(&id)
			EliminarProducto(id, inventario)
		case "6":
			var id int
			fmt.Printf("Ingrese el id del producto que quiere buscar: ")
			fmt.Scan(&id)
			producto, existe := BuscarProducto(id, inventario)
			if existe {
				fmt.Println("El producto si existe, es el siguiente", producto)
			} else {
				fmt.Println("El producto no existe")
			}
		case "7":
			ListarProductos(inventario)
		case "0":
			return
		default:
			fmt.Println("Opcion Invalida, ingrese otra.")
		}
	}
}
