package controller

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

const default_image = "https://res.cloudinary.com/dwd3oedmz/image/upload/v1735407083/default.png.png"

func GetProduct(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		productId := c.Param("id")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		product, err := srv.GetProduct(ctx, productId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func GetProducts(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		query := c.Request.URL.Query()
		skip, take := 0, 0
		var err error
		var search string
		if query.Has("search") {
			search = query.Get("search")
		}
		if query.Has("skip") {
			skip, err = strconv.Atoi(query.Get("skip"))
		}
		if query.Has("take") {
			take, err = strconv.Atoi(query.Get("take"))
		}
		if err != nil {
			log.Println(err)
		}
		products, err := srv.GetProducts(ctx, search, skip, take)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if len(products.Products) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No products found"})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

func PostProduct(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.ProductDocument
		err := c.ShouldBind(&product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		priceStr := c.DefaultPostForm("price", "0")
		stockStr := c.DefaultPostForm("stock", "0")
		product.Price, err = strconv.ParseFloat(priceStr, 64)
		product.Stock, err = strconv.ParseUint(stockStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price or stock"})
			return
		}

		product.SellerName = c.GetString("name")
		product.SellerID = c.GetString("id")

		file, err := c.FormFile("image")
		if err != nil {
			product.ImageUrl = default_image
		}
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open image file"})
			return
		}
		defer src.Close()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		uploadResult, err := srv.CloudinaryStorage.Upload.Upload(ctx, src, uploader.UploadParams{
			Folder: product.Category,
		})
		product.ImageUrl = uploadResult.SecureURL
		defer cancel()

		createdProduct, err := srv.PostProduct(ctx, &product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, createdProduct)
	}
}
