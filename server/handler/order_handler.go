package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/response"
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
