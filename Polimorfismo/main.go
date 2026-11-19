package main

import (
	"errors"
	"fmt"
)

type Dispositivo interface {
	Encender() error
	Apagar() error
	Estatus() string
}

type Router struct {
	Nombre    string
	Encendido bool
}

type Servidor struct {
	Nombre string
	Activo bool
}

func (r *Router) Encender() error {
	if r.Encendido {
		return errors.New("El router ya está encendido")
	}
	r.Encendido = true
	return nil
}
func (r *Router) Apagar() error {
	if !r.Encendido {
		return errors.New("El router ya está apagado")
	}
	r.Encendido = false
	return nil
}
func (r *Router) Estatus() string {
	if r.Encendido {
		return r.Nombre + ":" + "encendido"
	} else {
		return r.Nombre + ":" + "apagado"
	}
}
func (s *Servidor) Encender() error {
	if s.Activo {
		return errors.New("El servidor ya está activo")
	}
	s.Activo = true
	return nil
}
func (s *Servidor) Apagar() error {
	if !s.Activo {
		return errors.New("El servidor está inactivo")
	}
	s.Activo = false
	return nil
}
func (s *Servidor) Estatus() string {
	if s.Activo {
		return s.Nombre + ":" + "encendido"
	} else {
		return s.Nombre + ":" + "apagado"
	}
}
func OperarDispositivo(d Dispositivo) {
	fmt.Println("Estatus Inicial:", d.Estatus())

	err := d.Encender()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println("Despues de haber encendido: ", d.Estatus())

	er := d.Apagar()
	if er != nil {
		fmt.Println("ERROR:", er)
	}
	fmt.Println("Despues de haber apagado: ", d.Estatus())
}

func main() {
	router := &Router{
		Nombre: "Cisco",
	}
	servidor := &Servidor{
		Nombre: "Servidor Web",
	}
	fmt.Println("************Router**************")
	OperarDispositivo(router)
	fmt.Println("************Servidor**************")
	OperarDispositivo(servidor)
}
