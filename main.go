package main

import "fmt"

/*
	Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle



*/

type Student struct {
	Name     string
	LastName string
	DNI      int
	Date     string
}

func (s *Student) StudentInformation() {
	fmt.Printf("Name: %s\nLastName: %s\nDNI: %d\nDate: %s\n", s.Name, s.LastName, s.DNI, s.Date)
}

/*

	Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
	La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

Y los costos adicionales son:
Pequeño: solo tiene el costo del producto
Mediano: el precio del producto + un 3% de mantenerlo en la tienda
Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

El porcentaje de mantenerlo en la tienda es en base al precio del producto.
El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.

Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Producto que tenga el método Precio.

Se debe poder ejecutar el método Precio y que el método me devuelva el precio total en base al costo del producto y los adicionales en caso que los tenga

*/

const (
	SMALL  = "SMALL"
	MEDIUM = "MEDIUM"
	LARGE  = "LARGE"
)

type Product struct {
	ProductType string
	Price       float64
}

func FactoryProduct(productType string, price float64) Product {

	return Product{
		ProductType: productType,
		Price:       price,
	}
}

func (p *Product) GetPrice() float64 {
	switch p.ProductType {
	case SMALL:
		return p.Price
	case MEDIUM:
		return p.Price + (p.Price * 3 / 100)
	case LARGE:
		return p.Price + (p.Price * 6 / 100) + 2500
	}
	return 0
}

func main() {
	// ...

	p := FactoryProduct(LARGE, 1000)
	fmt.Println(p.GetPrice())

}
