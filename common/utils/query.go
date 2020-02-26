package utils

import (
	"fmt"
	"github.com/NielJun/go-logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 查询通过gin框架传过来的参数
func GetQueryInt64(c *gin.Context, key string) (value int64, err error) {

	str, ok := c.GetQuery(key)
	if !ok {
		logger.Error("question id is not availd,not found question_id")
		err = fmt.Errorf(" invalid params,not found key : %s", key)
		return
	}

	value, err = strconv.ParseInt(str, 10, 64)
	if err != nil {
		logger.Error("string conv to int faild,err : %#v ", err)
		err = fmt.Errorf(" str convert to int64 faild , err : %#v", err)
		return
	}
	return
}
