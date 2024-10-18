// main.go
package main

import (
    "github.com/Roolivero/pruebaHexagonal/internal/adapters"
    "github.com/Roolivero/pruebaHexagonal/internal/domain"
    "github.com/Roolivero/pruebaHexagonal/internal/app"
    "github.com/labstack/echo/v4"
    "net/http"
)

func main() {
    // Crear el repositorio en memoria
    repositorio := adapters.NuevoRepositorioMemoria()

    // Crear el servicio de gestión de productos
    servicio := app.NuevoServicioGestionProductos(repositorio)// Asegúrate de que este método exista

    // Configurar el servidor Echo
    e := echo.New()

    // Definir las rutas
    e.POST("/productos", func(c echo.Context) error {
        var producto adapters.ProductoRequest
        if err := c.Bind(&producto); err != nil {
            return c.String(http.StatusBadRequest, "Error al procesar la solicitud")
        }

        nuevoProducto := domain.Producto{
            ID:          producto.ID,
            Nombre:      producto.Nombre,
            Descripcion: producto.Descripcion,
            Cantidad:    producto.Cantidad,
            Precio:      producto.Precio, // Incluye este campo si lo necesitas
        }

        if err := servicio.AgregarProducto(nuevoProducto); err != nil {
            return c.String(http.StatusInternalServerError, err.Error())
        }
        return c.String(http.StatusCreated, "Producto agregado")
    })

    e.GET("/productos/:id", func(c echo.Context) error {
        id := c.Param("id")
        producto, err := servicio.ObtenerProducto(id)
        if err != nil {
            return c.String(http.StatusNotFound, err.Error())
        }
        return c.JSON(http.StatusOK, producto)
    })

    e.PUT("/productos/:id", func(c echo.Context) error {
        id := c.Param("id")
        var producto adapters.ProductoRequest
        if err := c.Bind(&producto); err != nil {
            return c.String(http.StatusBadRequest, "Error al procesar la solicitud")
        }

        actualizado := domain.Producto{
            ID:          id,
            Nombre:      producto.Nombre,
            Descripcion: producto.Descripcion,
            Cantidad:    producto.Cantidad,
            Precio:      producto.Precio, // Incluye este campo si lo necesitas
        }

        if err := servicio.ActualizarProducto(actualizado); err != nil {
            return c.String(http.StatusInternalServerError, err.Error())
        }
        return c.String(http.StatusOK, "Producto actualizado")
    })

    e.DELETE("/productos/:id", func(c echo.Context) error {
        id := c.Param("id")
        if err := servicio.EliminarProducto(id); err != nil {
            return c.String(http.StatusNotFound, err.Error())
        }
        return c.String(http.StatusOK, "Producto eliminado")
    })

    e.GET("/productos", func(c echo.Context) error {
        productos, err := servicio.ListarProductos()
        if err != nil {
            return c.String(http.StatusInternalServerError, err.Error())
        }
        return c.JSON(http.StatusOK, productos)
    })

    // Iniciar el servidor en el puerto 8080
    e.Logger.Fatal(e.Start(":8080"))
}
