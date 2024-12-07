package handler

import (
	"fmt"
	"gocommerce/internal/entity"
	"gocommerce/internal/usecase"
	"gocommerce/internal/utils"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ProductHandler struct {
	usecase *usecase.ProductUseCase
}

func NewProductHandler(usecase *usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{usecase: usecase}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse the form and handle any errors
	err := r.ParseMultipartForm(10 << 20) // Limit file size to 10MB
	if err != nil {
		log.Println("Error parsing form:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	price := r.FormValue("price")
	categoryID := r.FormValue("category_id")
	stock := r.FormValue("stock")

	floatPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stringCategoryID, err := strconv.Atoi(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	intStock, err := strconv.Atoi(stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get the uploaded file from the form
	file, handler, err := r.FormFile("image")
	if err != nil {
		log.Println("Error retrieving file:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate a unique filename for the image
	uniqueFilename := fmt.Sprintf("%d-%s", time.Now().Unix(), handler.Filename)
	filepath := fmt.Sprintf("../internal/static/images/%s", uniqueFilename)

	// Create the destination file
	dst, err := os.Create(filepath)
	if err != nil {
		log.Println("Error creating file:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	if _, err := io.Copy(dst, file); err != nil {
		log.Println("Error writing file:", err)
		http.Error(w, "Error writing file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Upload to ImageKit
	imageURL := utils.UploadImageImageKit(filepath)

	product := entity.Product{
		Name:       name,
		Price:      floatPrice,
		CategoryID: stringCategoryID,
		Stock:      intStock,
		ImageURL:   imageURL,
	}

	id, err := h.usecase.CreateProduct(r.Context(), &product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"id": id,
	})

}
