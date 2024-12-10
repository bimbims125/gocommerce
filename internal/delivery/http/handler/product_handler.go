package handler

import (
	"context"
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

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	usecase *usecase.ProductUseCase
}

func NewProductHandler(usecase *usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{usecase: usecase}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{"message": "Error parsing form: " + err.Error()})
		return
	}

	// Extract and validate form data
	name := r.FormValue("name")
	price := r.FormValue("price")
	categoryID := r.FormValue("category_id")
	stock := r.FormValue("stock")

	floatPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{"message": "Invalid price format"})
		return
	}

	intCategoryID, err := strconv.Atoi(categoryID)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{"message": "Invalid category ID format"})
		return
	}

	intStock, err := strconv.Atoi(stock)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{"message": "Invalid stock format"})
		return
	}

	// Retrieve the uploaded file
	file, handler, err := r.FormFile("image")
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{"message": "Error retrieving file: " + err.Error()})
		return
	}
	defer file.Close()

	// Validate image format
	if !utils.ValidateImageExt(handler.Filename) {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{"message": "Invalid image format. Only jpg, jpeg, png allowed"})
		return
	}

	// Generate a unique filename
	uniqueFilename := fmt.Sprintf("%d-%s", time.Now().Unix(), handler.Filename)
	filepath := fmt.Sprintf("../internal/static/images/%s", uniqueFilename)

	// Respond immediately to the client
	utils.JSONResponse(w, http.StatusAccepted, map[string]interface{}{"message": "File upload in progress"})

	// Process file upload and product creation asynchronously
	go func() {
		// Save the file to the server
		dst, err := os.Create(filepath)
		if err != nil {
			log.Printf("Error creating file: %v", err)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			log.Printf("Error saving file: %v", err)
			return
		}

		// Optionally upload the file to a cloud service
		imageURL, err := utils.UploadImageImageKit(filepath)
		if err != nil {
			log.Printf("Error uploading to ImageKit: %v", err)
			return
		}

		// Create the product entity
		product := entity.Product{
			Name:       name,
			Price:      floatPrice,
			CategoryID: intCategoryID,
			Stock:      intStock,
			ImageURL:   imageURL,
		}

		// Save the product in the database
		_, err = h.usecase.CreateProduct(context.Background(), &product)
		if err != nil {
			log.Printf("Error saving product: %v", err)
			return
		}

		log.Printf("Product created successfully with image: %s", imageURL)
	}()
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.usecase.GetProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.JSONResponse(w, http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	strId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := h.usecase.GetProductByID(r.Context(), strId)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, map[string]interface{}{"message": "Product not found!"})
		return
	}
	utils.JSONResponse(w, http.StatusOK, product)
}
