package webhook

import (
	"net/http"

	dto "online-course/internal/webhook/dto"
	useCase "online-course/internal/webhook/usecase"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
	useCase useCase.WebhookUseCase
}

func NewWebHookHandler(useCase useCase.WebhookUseCase) *WebhookHandler {
	return &WebhookHandler{useCase}
}

func (handler *WebhookHandler) Route(r *gin.RouterGroup) {
	webhookHandler := r.Group("/api/v1")

	webhookHandler.POST("/webhooks", handler.Xendit)
}

func (handler *WebhookHandler) Xendit(ctx *gin.Context) {
	var input dto.WebhookRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	err := handler.useCase.UpdatePayment(input.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", "ok"))
}
