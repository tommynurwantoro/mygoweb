package handler

import (
	"net/http"
	"tommynurwantoro/mygoweb/internal/domain"

	"github.com/labstack/echo/v4"
)

type PaymeHandler struct {
	Data domain.PaymeData
}

func (p *PaymeHandler) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "payme", p.Data)
}

func NewPaymeHandler() *PaymeHandler {
	return &PaymeHandler{
		Data: domain.PaymeData{
			PageTitle: "Pay Me",
		},
	}
}
