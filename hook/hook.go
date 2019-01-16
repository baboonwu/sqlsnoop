package hook

import (
	"context"
	"database/sql/driver"
	"log"
	"time"

	proxy "github.com/shogo82148/go-sql-proxy"
)

// NewHook ...
func NewHook() *proxy.HooksContext {
	return &proxy.HooksContext{

		PrePing: func(_ context.Context, _ *proxy.Conn) (interface{}, error) {
			log.Println("PrePing")
			return time.Now(), nil
		},

		Ping: func(c context.Context, ctx interface{}, conn *proxy.Conn) error {
			log.Println("Ping")
			return nil
		},

		PostPing: func(c context.Context, ctx interface{}, conn *proxy.Conn, err error) error {
			log.Println("PostPing")
			return nil
		},

		PreOpen: func(c context.Context, name string) (interface{}, error) {
			log.Println("PreOpen")
			return time.Now(), nil
		},

		Open: func(c context.Context, ctx interface{}, conn *proxy.Conn) error {
			log.Println("Open")
			return nil
		},

		PostOpen: func(c context.Context, ctx interface{}, conn *proxy.Conn, err error) error {
			log.Println("PostOpen")
			return nil
		},

		PreExec: func(c context.Context, stmt *proxy.Stmt, args []driver.NamedValue) (interface{}, error) {
			log.Println("PreExec")
			return time.Now(), nil
		},
		Exec: func(c context.Context, ctx interface{}, stmt *proxy.Stmt, args []driver.NamedValue, result driver.Result) error {
			log.Println("Exec")
			return nil
		},

		PostExec: func(c context.Context, ctx interface{}, stmt *proxy.Stmt, args []driver.NamedValue, result driver.Result, err error) error {
			log.Println("PostExec")
			return nil
		},

		PreQuery: func(c context.Context, stmt *proxy.Stmt, args []driver.NamedValue) (interface{}, error) {
			log.Printf("PreQuery:%v \n", time.Now().UnixNano())
			return time.Now(), nil
		},

		Query: func(c context.Context, ctx interface{}, stmt *proxy.Stmt, args []driver.NamedValue, rows driver.Rows) error {
			log.Println("Query")
			return nil
		},

		PostQuery: func(c context.Context, ctx interface{}, stmt *proxy.Stmt, args []driver.NamedValue, rows driver.Rows, err error) error {
			log.Printf("PostQuery:%v \n", time.Now().UnixNano())
			log.Printf("query time : %s; args = %v (%s)\n", stmt.QueryString, args, time.Since(ctx.(time.Time)))
			return nil
		},

		PreBegin: func(c context.Context, conn *proxy.Conn) (interface{}, error) {
			log.Println("PreBegin")
			txName := "tran01"
			return txName, nil
		},

		Begin: func(c context.Context, ctx interface{}, conn *proxy.Conn) error {
			txName := ctx.(string)
			log.Printf("Begin: trans name = %v\n", txName)
			return nil
		},

		PostBegin: func(c context.Context, ctx interface{}, conn *proxy.Conn, err error) error {
			log.Println("PostBegin")
			return nil
		},

		PreCommit: func(c context.Context, tx *proxy.Tx) (interface{}, error) {

			log.Printf("PreCommit: tx=%p \n", &tx.Tx)
			return time.Now(), nil
		},

		Commit: func(c context.Context, ctx interface{}, tx *proxy.Tx) error {
			log.Printf("Commit: tx = %p \n", tx)
			return nil
		},

		PostCommit: func(c context.Context, ctx interface{}, tx *proxy.Tx, err error) error {
			log.Println("PostCommit")
			return nil
		},

		PreRollback: func(c context.Context, tx *proxy.Tx) (interface{}, error) {
			log.Println("PreRollback")
			return time.Now(), nil
		},

		Rollback: func(c context.Context, ctx interface{}, tx *proxy.Tx) error {
			log.Printf("Rollback: trans name = %p\n", &tx.Tx)
			return nil
		},

		PostRollback: func(c context.Context, ctx interface{}, tx *proxy.Tx, err error) error {
			log.Println("PostRollback")
			return nil
		},

		PreClose: func(c context.Context, conn *proxy.Conn) (interface{}, error) {
			log.Println("PreClose")
			return time.Now(), nil
		},

		Close: func(c context.Context, ctx interface{}, conn *proxy.Conn) error {
			log.Println("Close")
			return nil
		},

		PostClose: func(c context.Context, ctx interface{}, conn *proxy.Conn, err error) error {
			log.Println("PostClose")
			return nil
		},

		PreResetSession: func(c context.Context, conn *proxy.Conn) (interface{}, error) {
			log.Println("PreResetSession")
			return time.Now(), nil
		},

		ResetSession: func(c context.Context, ctx interface{}, conn *proxy.Conn) error {
			log.Println("ResetSession")
			return nil
		},
		PostResetSession: func(c context.Context, ctx interface{}, conn *proxy.Conn, err error) error {
			log.Println("PostResetSession")
			return nil
		},
	}
}
