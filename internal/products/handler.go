package products

import (
	"net/http"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/products/web"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		svc: s,
	}
}

func (h *Handler) GetProducts(ctx *gin.Context) {
	sellerID := ctx.Query("seller_id")
	if sellerID == "" {
		web.ResponseErr(ctx, http.StatusBadRequest, "Invalid Parameter")
		return
	}
	products, err := h.svc.GetAllBySeller(sellerID)
	if err != nil {
		web.ResponseErr(ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	web.ResponseOk(ctx, http.StatusOK, "Ok", products)
}
