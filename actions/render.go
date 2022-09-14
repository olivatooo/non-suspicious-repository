package actions

import (
	"errors"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

var R *render.Engine

func init() {
	R = render.New(render.Options{
		DefaultContentType: "application/json",
	})
}

func RenderMessage(c buffalo.Context, message string) error {
	return c.Render(http.StatusAccepted, R.JSON(message))
}

func RenderCreated(c buffalo.Context, message string) error {
	return c.Render(http.StatusCreated, R.JSON(message))
}

func RenderBadRequest(c buffalo.Context, message string) error {
	return c.Error(http.StatusBadRequest, errors.New(message))
}

func RenderInternalServerError(c buffalo.Context, message string) error {
	return c.Error(http.StatusInternalServerError, errors.New(message))
}
