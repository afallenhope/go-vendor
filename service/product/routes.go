package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/afallenhope/go-vendor/types"
	"github.com/afallenhope/go-vendor/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc(
		"/products",
		h.handleGetProducts,
	).Methods(http.MethodGet)

	router.HandleFunc(
		"/products",
		h.handleCreateProduct,
	).Methods(http.MethodPost)

	router.HandleFunc(
		"/products/{productID}",
		h.handleDeleteProduct,
	).Methods(http.MethodDelete)

	router.HandleFunc(
		"/products/{productID}",
		h.handleGetProduct,
	).Methods(http.MethodGet)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["productID"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing product ID"))
		return
	}

	productID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
		return
	}

	product, err := h.store.GetProductByID(productID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, product)

}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var product types.CreateProductPayload

	if err := utils.ParseJSON(r, &product); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(product); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", validationErrors))
		return
	}

	err := h.store.CreateProduct(product)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["productID"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing product id"))
		return
	}

	product, err := strconv.Atoi(str)
	h.store.DeleteProductByID(product)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
