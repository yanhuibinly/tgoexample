package mongo

import (
	"github.com/tonyjt/tgo"
	"model"
)

type Record struct {
	tgo.DaoMongo
}

func NewRecord() *Record {

	return &Record{tgo.DaoMongo{CollectionName: "record", AutoIncrementId: true, PrimaryKey: "_id"}}
}

func (r *Record) Insert(data *model.Record) error {

	return r.DaoMongo.Insert(data)
}
