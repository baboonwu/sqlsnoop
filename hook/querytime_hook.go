package hook

import (
	"context"
	"database/sql/driver"
	"log"
	"time"

	proxy "github.com/shogo82148/go-sql-proxy"
)

// NewQueryTimeHook is used to monitor query time
func NewQueryTimeHook() *proxy.HooksContext {
	return &proxy.HooksContext{
		PreQuery: func(c context.Context, stmt *proxy.Stmt, args []driver.NamedValue) (interface{}, error) {
			return time.Now(), nil
		},

		Query: func(c context.Context, ctx interface{}, stmt *proxy.Stmt, args []driver.NamedValue, rows driver.Rows) error {
			return nil
		},

		PostQuery: func(c context.Context, ctx interface{}, stmt *proxy.Stmt, args []driver.NamedValue, rows driver.Rows, err error) error {
			log.Printf("query time [%s] : %s \n", time.Since(ctx.(time.Time)), stmt.QueryString)
			return nil
		},
	}
}
