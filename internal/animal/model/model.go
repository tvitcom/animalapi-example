package model

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/tvitcom/animalapi-example/pkg/util"
)


// The DOB - as integer like yyyymmdd: 19991231
type Animal struct {
	Id    int64  `json:"id"`
	Kind  string `json:"kind" form:"kind" binding:"required"`
	Name  string `json:"name" form:"name" binding:"required"`
	Dob   int    `json:"dob" form:"dob" binding:"gte=19991231"`
	Owner string `json:"owner" form:"owner" binding:"required"`
	Error error
}

func IndexWithPage(limit int64, offset int64) []Animal {
	db := util.GetDbConn()
	defer db.Close()

	query := "SELECT id, kind, name, dob, owner FROM animal LIMIT $1 OFFSET $2"
	rows, err := db.Query(query, limit, offset)
	defer rows.Close()
	util.PanicError(err)

	var items []Animal
	for rows.Next() {
		var sl Animal

		err = rows.Scan(&sl.Id, &sl.Name, &sl.Dob, &sl.Owner)
		items = append(items, sl)
	}

	return items
}

func Count() int64 {
	db := util.GetDbConn()
	defer db.Close()

	var count int64
	query := "SELECT COUNT(*) FROM animal"
	row := db.QueryRow(query)
	row.Scan(&count)
	return count
}

func Create(item Animal) (int64, error) {
	db := util.GetDbConn()
	defer db.Close()

	query := "INSERT INTO animal (kind, name, dob, owner) VALUES($1, $2, $3, $4);"
	stmt, stmtErr := db.Prepare(query)
	util.PanicError(stmtErr)

	res, queryErr := stmt.Exec(item.Kind, item.Name, item.Dob, item.Owner)
	util.PanicError(queryErr)

	id, getLastInsertIdErr := res.LastInsertId()
	util.PanicError(getLastInsertIdErr)

	return id, queryErr
}

func FindById(id int64) Animal {
	var item Animal
	db := util.GetDbConn()
	defer db.Close()

	query := "SELECT id, kind, name, dob, owner FROM animal WHERE id = $1;"

	row := db.QueryRow(query, id)
	row.Scan(&item.Id, &item.Kind, &item.Name, &item.Dob, &item.Owner)

	return item
}

func Put(id int64, item Animal) (Animal, error) {
	db := util.GetDbConn()
	defer db.Close()

	query := "UPDATE animal SET kind = $1, name = $1, dob = $2, owner = $3 WHERE id = $4"
	stmt, stmtErr := db.Prepare(query)
	util.PanicError(stmtErr)

	_, queryErr := stmt.Exec(item.Kind, item.Name, item.Dob, item.Owner, id)
	util.PanicError(queryErr)

	item.Id = id
	return item, queryErr
}

func Delete(item Animal) error {
	db := util.GetDbConn()
	defer db.Close()

	query := "DELETE FROM animal WHERE id = $1"
	stmt, stmtErr := db.Prepare(query)
	util.PanicError(stmtErr)

	_, queryErr := stmt.Exec(item.Id)
	util.PanicError(queryErr)

	return queryErr
}
