package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"tommynurwantoro/mygoweb/internal/domain"
	"tommynurwantoro/mygoweb/internal/service"

	"github.com/labstack/echo/v4"
)

type BlogHandler struct {
	Data    domain.BlogData
	Service domain.MarkdownService
}

func (b *BlogHandler) Index(c echo.Context) error {
	b.Data.Blogs = b.Service.GetAllBlogs()

	return c.Render(http.StatusOK, "blog", b.Data)
}

func (b *BlogHandler) Detail(c echo.Context) error {
	id := c.Param("id")
	b.Data.Content = template.HTML(b.Service.GetBlogByID(id))

	fmt.Println(b.Data.Content)

	return c.Render(http.StatusOK, "blog_detail", b.Data)
}

func NewBlogHandler() *BlogHandler {
	return &BlogHandler{
		Data: domain.BlogData{
			PageTitle: "Blog",
		},
		Service: service.NewMarkdownService(),
	}
}
