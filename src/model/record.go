package model

import (
	"github.com/tonyjt/tgo"
)

type Record struct {
	tgo.ModelMongo
	Name string
}

func NewRecord() *Record {
	data := &Record{}

	return data
}
