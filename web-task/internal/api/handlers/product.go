package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "web-task/internal/models"
    "web-task/internal/service"
    "web-task/pkg/utils/response"
)

func ListProducts(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
    
    svc := c.MustGet("productService").(*service.ProductService)
    products, total, err := svc.ListProducts(page, pageSize)
    if err != nil {
        c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
        return
    }

    c.JSON(http.StatusOK, response.Success(gin.H{
        "items": products,
        "total": total,
        "page":  page,
        "size":  pageSize,
    }))
}

func GetProduct(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, response.Error(400, "Invalid product ID"))
        return
    }

    svc := c.MustGet("productService").(*service.ProductService)
    product, err := svc.GetProduct(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, response.Error(404, "Product not found"))
        return
    }

    c.JSON(http.StatusOK, response.Success(product))
}

func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, response.Error(400, err.Error()))
        return
    }

    svc := c.MustGet("productService").(*service.ProductService)
    if err := svc.CreateProduct(&product); err != nil {
        c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
        return
    }

    c.JSON(http.StatusOK, response.Success(product))
}

func UpdateProduct(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, response.Error(400, "Invalid product ID"))
        return
    }

    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, response.Error(400, err.Error()))
        return
    }
    product.ID = uint(id)

    svc := c.MustGet("productService").(*service.ProductService)
    if err := svc.UpdateProduct(&product); err != nil {
        c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
        return
    }

    c.JSON(http.StatusOK, response.Success(product))
}

func DeleteProduct(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, response.Error(400, "Invalid product ID"))
        return
    }

    svc := c.MustGet("productService").(*service.ProductService)
    if err := svc.DeleteProduct(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, response.Error(500, err.Error()))
        return
    }

    c.JSON(http.StatusOK, response.Success(nil))
} 