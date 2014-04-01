package model

import ()

type Entity interface {
  Id() string
  Validate() bool
}
