package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/mattstratton/oathkeeper/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
