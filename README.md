Install
=======

> go get github.com/baboonwu/sqlsnoop

Run
===

> cd $GOPATH/src/github.com/baboonwu/sqlsnoop/hook

Monitor Query Time
------------------

> go test -test.run TestQueryTime

Monitor Mysql Transaction
-------------------------

> go test -test.run TestTransaction
