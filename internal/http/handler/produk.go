// Handler Produk
package handler

import (
	"Evermos-Virtual-Intern/internal/dto"
	"Evermos-Virtual-Intern/internal/service"
	"Evermos-Virtual-Intern/internal/util"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProdukHandler struct {
	produkService service.ProdukService
}

func NewProdukHandler(produkService service.ProdukService) *ProdukHandler {
	return &ProdukHandler{produkService}
}

func (h *ProdukHandler) GetAllProduk(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	categoryID, _ := strconv.ParseInt(c.Query("category_id", "0"), 10, 64)
	tokoID, _ := strconv.ParseInt(c.Query("toko_id", "0"), 10, 64)
	minHarga, _ := strconv.Atoi(c.Query("min_harga", "0"))
	maxHarga, _ := strconv.Atoi(c.Query("max_harga", "0"))

	filter := dto.ProdukFilterParams{
		NamaProduk: c.Query("nama_produk", ""),
		Limit:      limit,
		Page:       page,
		CategoryID: categoryID,
		TokoID:     tokoID,
		MinHarga:   minHarga,
		MaxHarga:   maxHarga,
	}

	produkList, err := h.produkService.GetAllProduk(c.Context(), filter)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to fetch products", err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Succeed to GET data", nil, produkList)
}

func (h *ProdukHandler) GetProdukByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid ID", err.Error(), nil)
	}
	produk, err := h.produkService.GetProdukByID(c.Context(), id)
	if err != nil {
		return util.JSONResponse(c, http.StatusNotFound, "Product not found", err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Success", nil, produk)
}

func (h *ProdukHandler) CreateProduk(c *fiber.Ctx) error {
	// Parse form-data fields
	var req dto.CreateProdukRequest
	req.NamaProduk = c.FormValue("nama_produk")
	req.Slug = c.FormValue("slug")
	req.HargaReseller, _ = strconv.Atoi(c.FormValue("harga_reseller"))
	req.HargaKonsumen, _ = strconv.Atoi(c.FormValue("harga_konsumen"))
	req.Stok, _ = strconv.Atoi(c.FormValue("stok"))
	req.Deskripsi = c.FormValue("deskripsi")
	req.IdToko, _ = strconv.ParseInt(c.FormValue("id_toko"), 10, 64)
	req.IdCategory, _ = strconv.ParseInt(c.FormValue("id_category"), 10, 64)

	// --- Ini bagian penting: Ambil Multiple Files ---
	form, err := c.MultipartForm()
	if err != nil {
		return util.JSONResponse(c, fiber.StatusBadRequest, "Failed to parse multipart form", err.Error(), nil)
	}

	fileHeaders := form.File["photos"]
	if len(fileHeaders) == 0 {
		return util.JSONResponse(c, fiber.StatusBadRequest, "No photos uploaded", nil, nil)
	}

	// Call Service with fileHeaders
	err = h.produkService.CreateProduk(c.Context(), &req, fileHeaders)
	if err != nil {
		return util.JSONResponse(c, fiber.StatusInternalServerError, "Failed to create product", err.Error(), nil)
	}

	return util.JSONResponse(c, fiber.StatusOK, "Product created successfully", nil, nil)
}


func (h *ProdukHandler) UpdateProduk(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid ID", err.Error(), nil)
	}
	var req dto.UpdateProdukRequest
	if err := c.BodyParser(&req); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	if err := h.produkService.UpdateProduk(c.Context(), id, &req); err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to update product", err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Product updated successfully", nil, nil)
}

func (h *ProdukHandler) DeleteProduk(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid ID", err.Error(), nil)
	}
	if err := h.produkService.DeleteProduk(c.Context(), id); err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to delete product", err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Product deleted successfully", nil, nil)
}