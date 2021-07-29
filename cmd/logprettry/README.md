## Log pretty output

This tool is for [logger](https://github.com/momaek/toolbox/tree/master/logger). 

You can pretty print json log output:

REQ + normal log
``` 
2021-07-30 00:07:20.922128 [Okw8Pygshc6xTpYW] [info] [REQ] [Started] |  GET  | / ::1
2021-07-30 00:07:20.922137 [Okw8Pygshc6xTpYW] [info] [/toolbox/logger/route_test.go:15] hello
2021-07-30 00:07:20.922162 [Okw8Pygshc6xTpYW] [info] [REQ] [Completed] |  GET  |  200  | [12.899µs] | / ::1
```

SQL log:

```
2021-07-30 00:02:30 [THISISDATABASETEST] [info] [/Users/wentx7n/momaek/src/toolbox/db/db_test.go:16] 2.256815ms select * from bill_record RowsAffected: 0
2021-07-30 00:02:30 [THISISDATABASETEST] [error] [/Users/wentx7n/momaek/src/toolbox/db/db_test.go:19] 385.108µs SELECT * FROM `bill_record` WHERE id  1 RowsAffected: 0 Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near '1' at line 1
```

### Usage

```
go run /path/to/your/main.go | logpretty

```
