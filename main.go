package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	rows.Close() //good habit to close

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Sometimes you can't use the for statement because you don't have more than one row, then you can use the if statement

/*
   if rows.Next() {
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
	}
*/

/*

  Transactions
The above example shows how you fetch data from the database, but when you want to write a web application
 then it will not only be necessary to fetch data from the db but it will also be required to write data into it.
 For that purpose, you should use transactions because for various reasons, such as having multiple go routines which access the database,
 the database might get locked. This is undesirable in your web application and the use of transactions is
  effective in ensuring your database activities either pass or fail completely depending on circumstances.
  It is clear that using transactions can prevent a lot of things from going wrong with the web app.

*/

/*
	    trashSQL, err := database.Prepare("update task set is_deleted='Y',last_modified_at=datetime() where id=?")
    if err != nil {
        fmt.Println(err)
    }
    tx, err := database.Begin()
    if err != nil {
        fmt.Println(err)
    }
    _, err = tx.Stmt(trashSQL).Exec(id)
    if err != nil {
        fmt.Println("doing rollback")
        tx.Rollback()
    } else {
        tx.Commit()
    }

*/
