package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tonyjt/tgo"
	"pconst"
	"service"
)

func Hello(c *gin.Context) {

	var name string

	name = tgo.UtilRequestGetParam(c, "name")

	var code int

	var word string

	if tgo.UtilIsEmpty(name) {
		code = pconst.ERROR_PARAMETER
	} else {
		word, code = service.HelloSay(name)

	}
	//msg := tgo.GetCodeMessage(code)

	tgo.UtilResponseReturnJson(c, code, word)

}
