package user

import (
	"fmt"
	"net/http"

	"github.com/afallenhope/go-vendor/config"
	"github.com/afallenhope/go-vendor/service/auth"
	"github.com/afallenhope/go-vendor/types"
	"github.com/afallenhope/go-vendor/utils"
	"github.com/go-playground/validator/v10" // Right now, we're not validating....
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc(
		"/login",
		h.handleLogin,
	).Methods("POST")

	router.HandleFunc(
		"/register",
		h.handleRegister,
	).Methods("POST")

	router.HandleFunc(
		"/logout",
		h.handleLogout,
	).Methods("GET")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		validationError := err.(validator.ValidationErrors).Error()
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", validationError))
		return
	}

	u, err := h.store.GetUserByUsername(payload.Username)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user or password invalid"))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user or password invalid"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		// I thought about adding this, however, from a security standpoint, I'm not sure if I should.
		// Specifying what is required in the payload, leaves to guessing the columns in which need to be sent.
		// For now, I'm commenting this out, however, I may add it back in.
		validationError := err.(validator.ValidationErrors).Error()
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", validationError))
		return
	}

	_, err := h.store.GetUserByUsername(payload.Username)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("use already exists"))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		UUID:     payload.UUID,
		Username: payload.Username,
		Password: hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleLogout(w http.ResponseWriter, r *http.Request) {}
