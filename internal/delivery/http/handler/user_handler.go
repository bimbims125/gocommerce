package handler

import (
	"encoding/json"
	"gocommerce/internal/entity"
	"gocommerce/internal/infra"
	"gocommerce/internal/usecase"
	"gocommerce/internal/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	usecase *usecase.UserUseCase
}

func NewUserHandler(usecase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.usecase.Create(r.Context(), &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload entity.User
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.usecase.GetUserByEmail(r.Context(), payload.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if !utils.ComparePasswords(user.Password, []byte(payload.Password)) {
		utils.JSONResponse(w, http.StatusUnauthorized, map[string]interface{}{"message": "Invalid email or password!"})
		return
	}

	token, err := infra.GenerateJWT([]byte(os.Getenv("JWT_SECRET")), user.ID, user.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.usecase.GetUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.JSONResponse(w, http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	strId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.usecase.GetUserByID(r.Context(), strId)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, map[string]interface{}{"message": "User not found!"})
		return
	}
	utils.JSONResponse(w, http.StatusOK, user)
}
