package hook

import (
	"database/sql"
	"log"
	"testing"

	mysql "github.com/go-sql-driver/mysql"
	proxy "github.com/shogo82148/go-sql-proxy"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	sql.Register("mysql:proxy", proxy.NewProxyContext(&mysql.MySQLDriver{}, NewHook()))
	log.Println("Drivers:", sql.Drivers())
}

func TestQueryMysql(t *testing.T) {

	db, err := sql.Open("mysql:proxy", "root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalf("Open filed: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("select * from event limit 1")
	defer rows.Close()

	for rows.Next() {
		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		//将数据保存到 record 字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		log.Println(record)
		log.Println("exit")
	}

}

func TestTransaction(t *testing.T) {

	db, err := sql.Open("mysql:proxy", "root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalf("Open filed: %v", err)
	}
	defer db.Close()

	// ctx := context.WithCancel()
	// tx, err := db.BeginTx(ctx, nil) // 创建tx对象
	tx, err := db.Begin() // 创建tx对象
	log.Printf("New tx: tx=%p \n", tx)
	rows, err := tx.Query("select * from event limit 1")
	defer rows.Close()

	for rows.Next() {
		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		//将数据保存到 record 字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		log.Println(record)
	}

	tx.Commit()
	log.Printf("Post Commit tx: tx=%p \n", tx)

}
