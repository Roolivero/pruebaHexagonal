package ports

import "github.com/Roolivero/pruebaHexagonal/internal/domain"

// ProductoRepository define la interfaz para manejar operaciones de productos
type ProductoRepository interface {
    AgregarProducto(producto domain.Producto) error
    ObtenerProducto(id string) (domain.Producto, error)
    ActualizarProducto(producto domain.Producto) error
    EliminarProducto(id string) error
    ListarProductos() ([]domain.Producto, error)
}