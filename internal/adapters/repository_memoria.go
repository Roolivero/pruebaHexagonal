// internal/adapters/repository_memoria.go
package adapters

import (
    "errors"
    "github.com/Roolivero/pruebaHexagonal/internal/domain"
    "github.com/Roolivero/pruebaHexagonal/internal/ports" // Importa el paquete ports
)

// RepositorioMemoria es la implementación de ProductoRepository
type RepositorioMemoria struct {
    productos map[string]domain.Producto
}

// Constructor para crear un nuevo repositorio en memoria
func NuevoRepositorioMemoria() *RepositorioMemoria {
    return &RepositorioMemoria{
        productos: make(map[string]domain.Producto),
    }
}

// Implementación de la interfaz ProductoRepository
var _ ports.ProductoRepository = (*RepositorioMemoria)(nil) // Verifica que RepositorioMemoria implementa la interfaz

// AgregarProducto agrega un nuevo producto al repositorio.
func (r *RepositorioMemoria) AgregarProducto(producto domain.Producto) error {
    if _, existe := r.productos[producto.ID]; existe {
        return errors.New("el producto ya existe")
    }
    r.productos[producto.ID] = producto
    return nil
}

// ObtenerProducto obtiene un producto por su ID.
func (r *RepositorioMemoria) ObtenerProducto(id string) (domain.Producto, error) {
    producto, existe := r.productos[id]
    if !existe {
        return domain.Producto{}, errors.New("producto no encontrado")
    }
    return producto, nil
}

// ActualizarProducto actualiza un producto existente.
func (r *RepositorioMemoria) ActualizarProducto(producto domain.Producto) error {
    if _, existe := r.productos[producto.ID]; !existe {
        return errors.New("producto no encontrado")
    }
    r.productos[producto.ID] = producto
    return nil
}

// EliminarProducto elimina un producto por su ID.
func (r *RepositorioMemoria) EliminarProducto(id string) error {
    if _, existe := r.productos[id]; !existe {
        return errors.New("producto no encontrado")
    }
    delete(r.productos, id)
    return nil
}

// ListarProductos lista todos los productos.
func (r *RepositorioMemoria) ListarProductos() ([]domain.Producto, error) {
    productos := make([]domain.Producto, 0, len(r.productos))
    for _, producto := range r.productos {
        productos = append(productos, producto)
    }
    return productos, nil
}
