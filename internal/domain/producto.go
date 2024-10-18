// internal/domain/producto.go
package domain

type Producto struct {
    ID          string
    Nombre      string
    Descripcion string // Agrega este campo
    Cantidad    int
    Precio      float64
}

// Método para verificar si un producto tiene stock disponible
func (p Producto) TieneStock() bool {
    return p.Cantidad > 0
}
