package model

type Directory struct {
	Id   uint64 `db:"id"`
	Name string `db:"name"`
}
