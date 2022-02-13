package ports

import (
	"backend/internal/myapp/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	app app.Application
}

func NewHTTPServer(app app.Application) *HTTPServer {
	return &HTTPServer{app: app}
}

func (h *HTTPServer) Health(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *HTTPServer) Hello(c *gin.Context, params HelloParams) {
	name := ""
	if params.Name != nil {
		name = *params.Name
	}

	response := h.app.Hello(c.Request.Context(), name)

	c.String(http.StatusOK, response)
}

func (h *HTTPServer) Add(c *gin.Context, params AddParams) {
	var a, b int64
	if params.A != nil {
		a = int64(*params.A)
	}
	if params.B != nil {
		b = int64(*params.B)
	}

	result, err := h.app.Add(c.Request.Context(), a, b)
	if err != nil {
		response := struct {
			Msg string `json:"msg"`
		}{
			Msg: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := struct {
		Result int `json:"result"`
	}{
		Result: int(result),
	}
	c.JSON(http.StatusOK, response)
}
