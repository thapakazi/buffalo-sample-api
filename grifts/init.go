package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/thapakazi/myapi/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
