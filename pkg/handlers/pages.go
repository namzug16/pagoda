package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/pager"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/ui/basecoat"
	"github.com/mikestefanello/pagoda/pkg/ui/models"
	"github.com/mikestefanello/pagoda/pkg/ui/pages"
)

type Pages struct{}

func init() {
	Register(new(Pages))
}

func (h *Pages) Init(c *services.Container) error {
	return nil
}

func (h *Pages) Routes(g *echo.Group) {
	g.GET("/", h.Home).Name = routenames.Home
	g.GET("/kitchen-sink", h.KitchenSink).Name = routenames.KitchenSink
	g.GET("/about", h.About).Name = routenames.About
	g.GET("/fragments/toast/:category", h.Toast)
}

func (h *Pages) Home(ctx echo.Context) error {
	pgr := pager.NewPager(ctx, 4)

	return pages.Home(ctx, &models.Posts{
		Posts: h.fetchPosts(&pgr),
		Pager: pgr,
	})
}

// fetchPosts is a mock example of fetching posts to illustrate how paging works.
func (h *Pages) fetchPosts(pager *pager.Pager) []models.Post {
	pager.SetItems(20)
	posts := make([]models.Post, 20)

	for k := range posts {
		posts[k] = models.Post{
			ID:    k + 1,
			Title: fmt.Sprintf("Post example #%d", k+1),
			Body:  fmt.Sprintf("Lorem ipsum example #%d ddolor sit amet, consectetur adipiscing elit. Nam elementum vulputate tristique.", k+1),
		}
	}
	return posts[pager.GetOffset() : pager.GetOffset()+pager.ItemsPerPage]
}

func (h *Pages) About(ctx echo.Context) error {
	return pages.About(ctx)
}

func (h *Pages) KitchenSink(ctx echo.Context) error {
	return pages.KitchenSink(ctx)
}

func (h *Pages) Toast(ctx echo.Context) error {
	category := ctx.Param("category")

	titles := map[string]string{
		"success": "Success",
		"error":   "Error",
		"info":    "Info",
		"warning": "Warning",
	}

	descriptions := map[string]string{
		"success": "Your changes have been saved.",
		"error":   "Something went wrong. Please try again.",
		"info":    "Here is some useful information.",
		"warning": "Please review this before continuing.",
	}

	title, ok := titles[category]
	if !ok {
		return ctx.NoContent(http.StatusNotFound)
	}

	toast := basecoat.Toast(basecoat.ToastParams{
		Category:    category,
		Title:       title,
		Description: descriptions[category],
	})

	return ctx.HTML(http.StatusOK, toast.String())
}
