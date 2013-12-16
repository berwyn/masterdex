package controller

import (
	"github.com/codegangsta/martini"
)

type Controller interface {
	Register(*martini.ClassicMartini)
}
