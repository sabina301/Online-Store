package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/response"
)

func (h *Handler) MakeOrder(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		response.NewError(ctx, "Error with get user", http.StatusBadRequest)
		return
	}
	orderId, err := h.serv.MakeOrder(userId)
	if err != nil {
		response.NewError(ctx, "Error with make order", http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"orderId": orderId,
	})
}

func (h *Handler) GetAllOrdersUser(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		response.NewError(ctx, "cant get user", http.StatusBadRequest)
		return
	}
	orders, err := h.serv.GetAllOrdersUser(userId)
	if err != nil {
		response.NewError(ctx, "cant get orders", http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

func (h *Handler) GetAllOrdersAdmin(ctx *gin.Context) {
	orders, err := h.serv.GetAllOrdersAdmin()
	if err != nil {
		response.NewError(ctx, "cant get orders", http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, orders)
}
