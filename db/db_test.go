package db

import "testing"

func TestDB(t *testing.T) {
	conf := Config{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Database: "unit_test",
	}

	Init(nil, &conf)

	db := Get("THISISDATABASETEST")
	db.Exec("select * from bill_record")

	r := BillRecord{}
	db.Find(&r, "id  1")
}

type BillRecord struct {
	ID int64
}

func (b *BillRecord) TableName() string {
	return "bill_record"
}
