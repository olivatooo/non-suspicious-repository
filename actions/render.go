package actions

import (
	"github.com/gobuffalo/buffalo/render"
)

var R *render.Engine

func init() {
	R = render.New(render.Options{
		DefaultContentType: "application/json",
	})
}
