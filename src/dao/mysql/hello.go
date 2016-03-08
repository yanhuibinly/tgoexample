package mysql

import (
	"fmt"
	"github.com/tonyjt/tgo"
	"model"
)

type Hello struct {
	tgo.DaoMysql
}

func NewHello() *Hello {

	return &Hello{tgo.DaoMysql{TableName: "Hello"}}
}

func (c *Hello) GetByName(name string) (*model.Hello, error) {

	data := model.NewHello()

	condition := fmt.Sprintf("name=%s", name)

	err := c.DaoMysql.Select(condition, &data)

	return data, err
}
