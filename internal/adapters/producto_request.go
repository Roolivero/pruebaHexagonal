// internal/adapters/producto_request.go
package adapters

// ProductoRequest representa la estructura de un producto en una solicitud
type ProductoRequest struct {
    ID          string  `json:"id"`
    Nombre      string  `json:"nombre"`
    Descripcion string  `json:"descripcion"` // Asegúrate de incluir este campo
    Cantidad    int     `json:"cantidad"`
    Precio      float64 `json:"precio"`
}
