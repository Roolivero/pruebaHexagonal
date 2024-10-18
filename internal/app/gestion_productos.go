package app

import (
    "github.com/Roolivero/pruebaHexagonal/internal/domain"
    "github.com/Roolivero/pruebaHexagonal/internal/ports"
)

type ServicioGestionProductos struct {
    repositorio ports.ProductoRepository
}

// Constructor para crear el servicio
func NuevoServicioGestionProductos(repo ports.ProductoRepository) *ServicioGestionProductos {
    return &ServicioGestionProductos{repositorio: repo}
}

// Agregar un nuevo producto al inventario
func (s *ServicioGestionProductos) AgregarProducto(producto domain.Producto) error {
    return s.repositorio.AgregarProducto(producto)
}

// Obtener un producto por su ID
func (s *ServicioGestionProductos) ObtenerProducto(id string) (domain.Producto, error) {
    return s.repositorio.ObtenerProducto(id)
}

// Actualizar la información de un producto
func (s *ServicioGestionProductos) ActualizarProducto(producto domain.Producto) error {
    return s.repositorio.ActualizarProducto(producto)
}

// Eliminar un producto por su ID
func (s *ServicioGestionProductos) EliminarProducto(id string) error {
    return s.repositorio.EliminarProducto(id)
}

// Listar todos los productos
func (s *ServicioGestionProductos) ListarProductos() ([]domain.Producto, error) {
    return s.repositorio.ListarProductos()
}
