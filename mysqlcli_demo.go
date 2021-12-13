package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	sqlBuilder "github.com/smartwalle/dbs"
)

type User struct {
	Id   int64  `sql:"id"`
	Name string `sql:"name"`
	Age  int    `sql:"age"`
}

/* CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `human_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8
*/

func mysqlcli_demo() {
	// try_connect()
	// create_table()
	// mysqlcli_sqlBuilder_demo()

	// mysqlcli_update_demo()
	// mysqlcli_delete_demo()

	//select
	// mysqlcli_select_demo()
	// mysqlcli_select_prepare_demo()
	// mysqlcli_update_prepare_demo()

	//insert
	// mysqlcli_insert_demo()
	// mysqlcli_insert_prepare_demo()
	// mysqlcli_insert_demo_with_sqlBuilder_sugar()
	// mysqlcli_insert_demo_with_sqlBuilder_sugar1()
	// mysqlcli_insert_prepare_demo_with_sqlBuilder_sugar()
}

/*
user@unix(/path/to/socket)/dbname?charset=utf8  // unix系统的，我一般不用
user:password@tcp(localhost:5555)/dbname?charset=utf8 // 涉及到远程链接调用的
user:password@/dbname  // 本地使用的
user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname // 涉及到IPV6的
*/
func try_connect() {
	db, err := sql.Open("mysql", "root:123456@tcp(10.24.14.37:3306)/global_dev_test1?parseTime=true")
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/global_dev_test1?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func create_table() {
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	db, err := sql.Open("mysql", "root:123456@tcp(10.24.14.37:3306)/global_dev_test1?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS global_dev_test1.hello(world varchar(50))")
	_, err = db.Exec("CREATE TABLE `user` ( `id` int(11) NOT NULL AUTO_INCREMENT, `name` varchar(32) DEFAULT NULL, `age` int(11) DEFAULT NULL, PRIMARY KEY (`id`), UNIQUE KEY `human_id_uindex` (`id`)) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8")

	if err != nil {
		log.Fatalln(err)
	}
}

func mysqlcli_sqlBuilder_demo() {
	var sb = sqlBuilder.NewSelectBuilder()
	sb.Selects("u.id", "u.name", "u.age")
	sb.From("user", "AS u")
	sb.Where("u.id = ?", 1)
	sb.Limit(1)

	sqlStr, args, _ := sb.ToSQL()
	fmt.Println("sqlStr:", sqlStr)
	fmt.Println("args:", args)
}

func mysqlcli_insert_demo() {
	sqlStr := fmt.Sprintf("insert into %s (name, age) values(?,?), (?,?)", "user")
	var sqlVal = []interface{}{"name1", 18, "name2", 19}
	db, err := sql.Open("mysql", "root:123456@tcp(10.24.14.37:3306)/global_dev_test1?parseTime=true")
	defer db.Close()
	res, err := db.Exec(sqlStr, sqlVal...)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	fmt.Printf("LastInsertId id: %d\n", id)
}

func mysqlcli_insert_prepare_demo() {
	sqlStr := fmt.Sprintf("insert into %s (name, age) values(?,?), (?,?)", "user")
	var sqlVal = []interface{}{"name1", 18, "name2", 19}
	db, err := sql.Open("mysql", "root:123456@tcp(10.24.14.37:3306)/global_dev_test1?parseTime=true")
	defer db.Close()

	stmt, err := db.Prepare(sqlStr)
	res, err := stmt.Exec(sqlVal...)
	// res, err := db.Exec(sqlStr, sqlVal...)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	fmt.Printf("LastInsertId id: %d\n", id)
}

func mysqlcli_insert_demo_with_sqlBuilder_sugar() {
	var ib = sqlBuilder.NewInsertBuilder()
	ib.Table("user")
	ib.Columns("name", "age")
	ib.Values("用户1", 18)
	ib.Values("用户2", 20)

	fmt.Println(ib.ToSQL())
	sqlStr, sqlVal, err := ib.ToSQL()
	db, err := sql.Open("mysql", "root:123456@tcp(10.24.14.37:3306)/global_dev_test1?parseTime=true")
	defer db.Close()
	res, err := db.Exec(sqlStr, sqlVal...)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	fmt.Printf("LastInsertId id: %d\n", id)
}

func mysqlcli_insert_demo_with_sqlBuilder_sugar1() {
	var ib = sqlBuilder.NewInsertBuilder()
	ib.Table("user")
	ib.Columns("name", "age")
	ib.Values("用户1", 18)
	ib.Values("用户2", 20)

	fmt.Println(ib.ToSQL())
	db, err := sql.Open("mysql", "root:123456@tcp(10.24.14.37:3306)/global_dev_test1?parseTime=true")
	defer db.Close()
	ib.Exec(db)
	if err != nil {
		log.Fatal(err)
	}
}

func mysqlcli_insert_prepare_demo_with_sqlBuilder_sugar() {
	var ib = sqlBuilder.NewInsertBuilder()
	ib.Table("user")
	ib.Columns("name", "age")
	ib.Values("用户1", 18)
	ib.Values("用户2", 20)
	fmt.Println(ib.ToSQL())

	sqlStr, sqlVal, err := ib.ToSQL()
	// db, err := sql.Open("mysql", "数据库连接信息")
	db, err := sql.Open("mysql", "root:123456@tcp(10.24.14.37:3306)/global_dev_test1?parseTime=true")
	defer db.Close()
	stmt, err := db.Prepare(sqlStr)
	res, err := stmt.Exec(sqlVal...)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	fmt.Printf("LastInsertId id: %d\n", id)
}

func mysqlcli_update_demo() {
	var ub = sqlBuilder.NewUpdateBuilder()
	ub.Table("user")
	ub.SET("name", "name1")
	ub.SET("age", 10)
	ub.Where("id = ? ", 2)
	ub.Limit(1)
	fmt.Println(ub.ToSQL())

	// db, err := sql.Open("mysql", "数据库连接信息")
	db, err := sql.Open("mysql", "root:123456@tcp(10.24.14.37:3306)/global_dev_test1?parseTime=true")
	defer db.Close()
	res, err := ub.Exec(db)
	if err != nil {
		log.Fatal(err)
	}
	affect, err := res.RowsAffected()
	fmt.Printf("affect: %d\n", affect)
}

func mysqlcli_update_prepare_demo() {
	var ub = sqlBuilder.NewUpdateBuilder()
	ub.Table("user")
	ub.SET("name", "name1")
	ub.SET("age", 10)
	ub.Where("id = ? ", 1)
	ub.Limit(1)
	fmt.Println(ub.ToSQL())

	updateSql, updateVal, _ := ub.ToSQL()
	fmt.Println("selectPlayerData selectSQL:", updateSql)

	// db, err := sql.Open("mysql", "数据库连接信息")
	db, err := sql.Open("mysql", "root:123456@tcp(10.24.14.37:3306)/global_dev_test1?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// sb.Scan(db, &user)
	// rb.Exec(db)

	stmt, err := db.Prepare(updateSql)
	if err != nil {
		// fmt.Println("======prepare failed, err:", err)
		fmt.Errorf("prepare failed, err: %v", err)
		return
	}

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	fmt.Println("======exec=========1")
	res, err := stmt.Exec(updateVal...)
	// res, err := db.Exec(updateSql, updateVal...)
	affect, err := res.RowsAffected()
	fmt.Printf("affect: %d\n", affect)
}

func mysqlcli_delete_demo() {
	var rb = sqlBuilder.NewDeleteBuilder()
	rb.Table("user")
	rb.Where("id = ?", 1)
	rb.Limit(1)
	fmt.Println(rb.ToSQL())
	// db, err := sql.Open("mysql", "数据库连接信息")
	db, err := sql.Open("mysql", "root:123456@tcp(10.24.14.37:3306)/global_dev_test1?parseTime=true")
	defer db.Close()
	rb.Exec(db)
	if err != nil {
		log.Fatal(err)
	}
}

func mysqlcli_select_demo() {
	var sb = sqlBuilder.NewSelectBuilder()
	sb.Selects("u.id", "u.name", "u.age")
	sb.From("user", "AS u")
	sb.Where("u.id = ?", 1)
	sb.Limit(1)

	// db, err := sql.Open("mysql", "数据库连接信息")
	db, err := sql.Open("mysql", "root:123456@tcp(10.24.14.37:3306)/global_dev_test1?parseTime=true")
	if err != nil {
		fmt.Println("连接数据库出错：", err)
		return
	}
	defer db.Close()

	selSql, selVal, _ := sb.ToSQL()
	var user User

	rows, err := db.Query(selSql, selVal...)
	if err != nil {
		fmt.Println("Query 出错：", err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	fmt.Printf("user: %v\n", user)
	fmt.Println(user.Id, user.Name, user.Age)
}

func mysqlcli_select_scan_demo() {
	var sb = sqlBuilder.NewSelectBuilder()
	sb.Selects("u.id", "u.name", "u.age")
	sb.From("user", "AS u")
	sb.Where("u.id = ?", 1)
	sb.Limit(1)

	// db, err := sql.Open("mysql", "数据库连接信息")
	db, err := sql.Open("mysql", "root:123456@tcp(10.24.14.37:3306)/global_dev_test1?parseTime=true")
	if err != nil {
		fmt.Println("连接数据库出错：", err)
		return
	}
	defer db.Close()

	var user User
	if err := sb.Scan(db, &user); err != nil {
		fmt.Println("Query 出错：", err)
		return
	}

	fmt.Println(user.Id, user.Name, user.Age)
}

func mysqlcli_select_prepare_demo() {
	var sb = sqlBuilder.NewSelectBuilder()
	sb.Selects("u.id", "u.name AS username", "u.age")
	// sb.Select(sqlBuilder.Alias("b.amount", "user_amount"))
	sb.From("user", "AS u")
	// sb.LeftJoin("bank", "AS b ON b.user_id = u.id")
	sb.Where("u.id = ?", 1)
	fmt.Println(sb.ToSQL())

	selSQL, selVal, _ := sb.ToSQL()
	fmt.Println("selectPlayerData selectSQL:", selSQL)

	// db, err := sql.Open("mysql", "数据库连接信息")
	db, err := sql.Open("mysql", "root:123456@tcp(10.24.14.37:3306)/global_dev_test1?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// sb.Scan(db, &user)
	// rb.Exec(db)

	stmt, err := db.Prepare(selSQL)
	if err != nil {
		// fmt.Println("======prepare failed, err:", err)
		fmt.Errorf("prepare failed, err: %v", err)
		return
	}

	fmt.Println("======exec=========1")

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	rows, err := stmt.Query(selVal...)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	// id, err := rows.LastInsertId()
	// affect, err := rows.RowsAffected()

	if err != nil {
		// lu.Log().Error("exec select data from db failed. roleID:", roleID, err)
		fmt.Errorf("exec select data from db failed. err:", err)
		return
	}
	fmt.Println("======exec=========2")

	// var playerDBData PlayerDataDB
	var user User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	fmt.Printf("user: %v\n", user)
}
