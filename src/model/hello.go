package model

import (
	"github.com/tonyjt/tgo"
)

type Hello struct {
	tgo.ModelMysql
	Name string
	Word string
}

func NewHello() *Hello {

	data := &Hello{}
	return data
}
