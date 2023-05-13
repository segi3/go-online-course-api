package class_room

import (
	"net/http"
	"strconv"

	useCase "online-course/internal/class_room/usecase"
	"online-course/internal/middleware"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ClassRoomHandler struct {
	useCase useCase.ClassRoomUseCase
}

func NewClassRoomHandler(useCase useCase.ClassRoomUseCase) *ClassRoomHandler {
	return &ClassRoomHandler{useCase}
}

func (handler *ClassRoomHandler) Route(r *gin.RouterGroup) {
	classRoomHandler := r.Group("/api/v1")

	classRoomHandler.Use(middleware.AuthJwt)
	{
		classRoomHandler.GET("/class_rooms", handler.FindAllByUserId)
	}
}
func (handler *ClassRoomHandler) FindAllByUserId(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	user := utils.GetCurrentUser(ctx)

	data := handler.useCase.FindAllByUserId(offset, limit, int(user.ID))

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}
