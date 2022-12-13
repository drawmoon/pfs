package model

type File struct {
	Id   uint64 `db:"id"`
	Name string `db:"name"`
}
