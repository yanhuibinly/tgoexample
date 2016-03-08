package service

import (
	"dao/mongo"
	"dao/mysql"
	"dao/redis"
	"github.com/tonyjt/tgo"
	"model"
	"pconst"
)

func HelloSay(name string) (string, int) {

	if name == "xxxx" {
		return "", pconst.ERROR_HELLO_DIRTY
	}

	daoRedis := redis.NewHello()

	var data *model.Hello

	result := daoRedis.Get(name, data)

	if !result {
		daoMysql := mysql.NewHello()
		var err error
		data, err = daoMysql.GetByName(name)
		if err != nil {
			tgo.UtilLogErrorf("get hello error, name:%s,error:%s", name, err.Error())
			return "", pconst.ERROR_MYSQL
		}
		//写入redis,懒的写了
	}
	//写记录到mongo
	if data != nil {
		record := model.NewRecord()
		record.Name = name

		daoMongo := mongo.NewRecord()

		err := daoMongo.Insert(record)

		if err != nil {
			tgo.UtilLogErrorf("insert hello error, name:%s,error:%s", name, err.Error())
			return "", pconst.ERROR_MONGO
		}
		return data.Word, 0
	} else {
		return "", pconst.ERROR_HELLO_NOTFOUND
	}

}
