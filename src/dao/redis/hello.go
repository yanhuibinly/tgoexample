package redis

import (
	"github.com/tonyjt/tgo"
	"model"
)

type Hello struct {
	tgo.DaoRedis
}

func NewHello() *Hello {

	dao := &Hello{tgo.DaoRedis{KeyName: "hello"}}

	return dao
}

func (h *Hello) Get(name string, data *model.Hello) bool {
	return h.DaoRedis.Get(name, data)
}
