package http

import (
	"net/http"

	"github.com/ckwhen/go-todo/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TodoHandler struct {
	TodoUsecase domain.TodoUsecase
}

func NewTodoHandler(rg *gin.RouterGroup, todoUsecase domain.TodoUsecase) {
	handler := &TodoHandler{
		TodoUsecase: todoUsecase,
	}

	rg.GET("/", handler.getTodos)
	rg.POST("/", handler.createTodo)
}

func (t *TodoHandler) getTodos(ctx *gin.Context) {
	todos, err := t.TodoUsecase.GetAll(ctx)

	if err != nil {
		logrus.Error(err)

		ctx.JSON(http.StatusInternalServerError, "Internal error")
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

func (t *TodoHandler) createTodo(ctx *gin.Context) {
	var todo domain.Todo

	if err := ctx.BindJSON(&todo); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "Internal error")
		return
	}

	if err := t.TodoUsecase.Store(ctx, &todo); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "Internal error")
		return
	}

	ctx.JSON(http.StatusOK, todo)
}
