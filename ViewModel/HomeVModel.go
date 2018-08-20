package ViewModel

import (
	"../Model"
)

type Home struct {
	Model.HomeModel
}

func (m Home) IsValid() bool {
	/* do some checks and return true if it's valid... */
	return m.ID > 0
}
