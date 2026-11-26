package main

import "fmt"

type ProductoBase struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

type Electronico struct {
	ProductoBase
	GarantiaMeses int
}
type Alimento struct {
	ProductoBase
	FechaExpiracion string
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) ProductoBase {
	return ProductoBase{
		ID:       id,
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}
func NuevoProductoElectronico(id int, nombre string, precio float64, cantidad int, garantia int) Electronico {
	return Electronico{
		ProductoBase:  NuevoProducto(id, nombre, precio, cantidad),
		GarantiaMeses: garantia,
	}
}
func NuevoProductoAlimento(id int, nombre string, precio float64, cantidad int, fechaExp string) Alimento {
	return Alimento{
		ProductoBase:    NuevoProducto(id, nombre, precio, cantidad),
		FechaExpiracion: fechaExp,
	}
}

type InformacionProducto interface {
	MostrarInfo() string
}

type Inventario struct {
	Productos []InformacionProducto
}

func (i *Inventario) AgregarProducto(p InformacionProducto) {
	i.Productos = append(i.Productos, p)
}

func (i *Inventario) ListarProductos() {
	for _, p := range i.Productos {
		fmt.Println(p.MostrarInfo())
	}
}

func (p *ProductoBase) MostrarInfo() string {
	return fmt.Sprintf("El producto es el siguiente: %d\n"+
		"El nombre del producto es el siguiente: %s\n"+
		"El precio del producto es el siguiente: %.2f\n"+
		"La cantidad del producto es el siguiente: %d", p.ID, p.Nombre, p.Precio, p.Cantidad)
}
func (e *Electronico) MostrarInfo() string {
	return fmt.Sprintf("%s\nLa garantia del producto en meses es el siguiente: %d",
		e.ProductoBase.MostrarInfo(), e.GarantiaMeses)
}
func (a *Alimento) MostrarInfo() string {
	return fmt.Sprintf("%s\nLa fecha de expiracion del producto es el siguiente: %s",
		a.ProductoBase.MostrarInfo(), a.FechaExpiracion)
}

func main() {
	inventario := Inventario{}
	ps5 := NuevoProductoElectronico(1, "ps5", 850, 5, 12)
	agua := NuevoProductoAlimento(2, "Agua", 1.10, 30, "15/12/2026")

	inventario.AgregarProducto(&ps5)
	inventario.AgregarProducto(&agua)
	inventario.ListarProductos()
}
