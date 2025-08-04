package handler

import (
	"Evermos-Virtual-Intern/internal/common"
	"Evermos-Virtual-Intern/internal/dto"
	"Evermos-Virtual-Intern/internal/service"
	"Evermos-Virtual-Intern/internal/util"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TrxHandler struct {
	trxService service.TrxService
}

func NewTrxHandler(trxService service.TrxService) *TrxHandler {
	return &TrxHandler{trxService: trxService}
}

func (h *TrxHandler) GetTrxByID(c *fiber.Ctx) error {
	// Ambil ID dari path params
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid ID parameter", err.Error(), nil)
	}

	// Call Service
	trxData, err := h.trxService.GetTrxByID(c.Context(), id)
	if err != nil {
		return util.JSONResponse(c, http.StatusNotFound, "Transaction not found", err.Error(), nil)
	}

	// Success Response
	return util.JSONResponse(c, http.StatusOK, "Succeed to GET data", nil, trxData)
}

func (h *TrxHandler) GetAllTrx(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	result, err := h.trxService.GetAllTrx(c.Context(), page, limit)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to fetch transactions", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Succeed to GET data", nil, result)
}


func (h *TrxHandler) CreateTrx(c *fiber.Ctx) error {
	var req dto.CreateTrxRequest
	if err := c.BodyParser(&req); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}

	// Ambil User ID dari token / session (contoh pakai c.Locals)
	userToken, err := common.GetUserFromToken(c)
	if err != nil {
		return util.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", "Invalid token", nil)
	}
	userID := userToken.ID

	// Call Service
	if err := h.trxService.CreateTrx(c.Context(), &req, userID); err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to create transaction", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusCreated, "Transaction created successfully", nil, nil)
}
