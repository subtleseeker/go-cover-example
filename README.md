nikhil-mundra@nikhil-mundra:~/tmp/del/cover-main/go/src/test$ go test -c -coverpkg=../...
nikhil-mundra@nikhil-mundra:~/tmp/del/cover-main/go/src/test$ ./test.test -test.v -test.coverprofile=coverage2.out -test.run "Pkg2"
=== RUN   TestPkg2_Bye
bye
--- PASS: TestPkg2_Bye (0.00s)
    bye_test.go:10: bye
PASS
coverage: 22.2% of statements in ../...
nikhil-mundra@nikhil-mundra:~/tmp/del/cover-main/go/src/test$ ./test.test -test.v -test.coverprofile=coverage1.out -test.run "Pkg1"
=== RUN   TestPkg1_Hello
hello
--- PASS: TestPkg1_Hello (0.00s)
    hello_test.go:13: hello
PASS
coverage: 44.4% of statements in ../...
nikhil-mundra@nikhil-mundra:~/tmp/del/cover-main/go/src/test$ ./test.test -test.v -test.coverprofile=coverage.out -test.run "Pkg1|Pkg2"
=== RUN   TestPkg2_Bye
bye
--- PASS: TestPkg2_Bye (0.00s)
    bye_test.go:10: bye
=== RUN   TestPkg1_Hello
hello
--- PASS: TestPkg1_Hello (0.00s)
    hello_test.go:13: hello
PASS
coverage: 66.7% of statements in ../...
nikhil-mundra@nikhil-mundra:~/tmp/del/cover-main/go/src/test$ ../../bin/gocovermerge coverage1.out coverage2.out > covermerge.out
nikhil-mundra@nikhil-mundra:~/tmp/del/cover-main/go/src/test$ go tool cover -func=covermerge.out
pkg1/hello.go:5:        Hello           80.0%
pkg1/hello2.go:3:       Hello2          0.0%
pkg2/bye.go:5:          Bye             100.0%
test/test_utils.go:3:   Util_HelloWorld 0.0%
total:                  (statements)    66.7%
