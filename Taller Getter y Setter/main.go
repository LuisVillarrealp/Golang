package main

import "fmt"

type Producto struct {
	id       int
	nombre   string
	precio   float64
	cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		id:       id,
		nombre:   nombre,
		precio:   precio,
		cantidad: cantidad,
	}
}
func (p Producto) GetID() int {
	return p.id
}
func (p Producto) GetNombre() string {
	return p.nombre
}
func (p Producto) GetPrecio() float64 {
	return p.precio
}
func (p Producto) GetCantidad() int {
	return p.cantidad
}
func (p *Producto) SetPrecio(nuevoPrecio float64) {
	if nuevoPrecio > 0 {
		p.precio = nuevoPrecio
		fmt.Println("Se actualizo el producto")
	} else {
		fmt.Println("Precio invalido")
	}
}
func (p *Producto) SetCantidad(nuevaCantidad int) {
	p.cantidad = nuevaCantidad
}
func (p *Producto) MostrarInfo() {
	fmt.Println("El producto es el siguiente: ", p.GetID())
	fmt.Println("El nombre del producto es el siguiente: ", p.GetNombre())
	fmt.Println("El precio del producto es el siguiente: ", p.GetPrecio())
	fmt.Println("La cantidad del producto es el siguiente: ", p.GetCantidad())
}

type Inventario struct {
	productos map[int]Producto
}

func NuevoInventario() Inventario {
	return Inventario{
		productos: map[int]Producto{},
	}
}
func (inv *Inventario) AgregarProducto(p Producto) {
	inv.productos[p.GetID()] = p
}
func (inv *Inventario) EliminarProducto(id int) {
	producto, existe := inv.BuscarProducto(id)
	if existe {
		fmt.Println("Se elimino el producto", producto.GetID())
		fmt.Println("Nombre: ", producto.GetNombre())
		fmt.Println("Precio", producto.GetPrecio())
		fmt.Println("Cantidad", producto.GetCantidad())
		delete(inv.productos, id)
	} else {
		fmt.Println("No existe ese producto")
	}
}
func (inv *Inventario) BuscarProducto(id int) (Producto, bool) {
	producto, existe := inv.productos[id]
	return producto, existe
}
func (inv *Inventario) ListarProductos() {
	for _, datos := range inv.productos {
		datos.MostrarInfo()
	}
}

func main() {
	inv := NuevoInventario()
	var usrOpcion string
	fmt.Println("\n			Bienvenido a este sistema básico de gestión de inventario")
	for {
		fmt.Println("\n			1. Crear o Agregar un nuevo producto")
		fmt.Println("			2. Actualizar Precio")
		fmt.Println("			3. Actualizar Cantidad")
		fmt.Println("			4. Mostrar Informacion de un Producto")
		fmt.Println("			5. Eliminar un producto")
		fmt.Println("			6. Buscar producto")
		fmt.Println("			7. Listar todos los productos")
		fmt.Println("			0. Para salir ingrese 0")
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
			inv.AgregarProducto(NuevoProducto(id, nombre, precio, cantidad))
			fmt.Println("Se agrego el producto")
		case "2":
			var id int
			var nuevoPrecio float64
			fmt.Printf("Ingrese el id del producto: ")
			fmt.Scan(&id)
			fmt.Printf("Ingrese el nuevo precio: ")
			fmt.Scan(&nuevoPrecio)
			producto, existe := inv.BuscarProducto(id)
			if existe {
				producto.SetPrecio(nuevoPrecio)
				inv.AgregarProducto(producto)
				precioGuardado, _ := inv.BuscarProducto(id)
				fmt.Println("Precio guardado:", precioGuardado.GetPrecio())
			} else {
				fmt.Println("No existe ese producto")
			}

		case "3":
			var id int
			var nuevaCantidad int
			fmt.Printf("Ingrese el id del producto: ")
			fmt.Scan(&id)
			fmt.Printf("Ingrese la nuevo cantidad: ")
			fmt.Scan(&nuevaCantidad)
			producto, existe := inv.BuscarProducto(id)
			if existe {
				producto.SetCantidad(nuevaCantidad)
				inv.AgregarProducto(producto)
				fmt.Println("Se actualizo el producto")
			} else {
				fmt.Println("No existe ese producto")
			}
		case "4":
			var id int
			fmt.Printf("Ingrese el id del producto que quiere ver: ")
			fmt.Scan(&id)
			producto, existe := inv.BuscarProducto(id)
			if existe {
				producto.MostrarInfo()
			} else {
				fmt.Println("No existe ese producto")
			}

		case "5":
			var id int
			fmt.Printf("Ingrese el id del producto que quiere eliminar: ")
			fmt.Scan(&id)
			inv.EliminarProducto(id)
		case "6":
			var id int
			fmt.Printf("Ingrese el id del producto que quiere buscar: ")
			fmt.Scan(&id)
			producto, existe := inv.BuscarProducto(id)
			if existe {
				fmt.Println("El producto si existe, es el siguiente")
				producto.MostrarInfo()
			} else {
				fmt.Println("El producto no existe")
			}
		case "7":
			inv.ListarProductos()
		case "0":
			return
		default:
			fmt.Println("Opcion Invalida, ingrese otra.")
		}
	}
}
