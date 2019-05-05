# `go run`を使って`exit.go`を実行すると
# `go`が終了を取り出し結果を表示します。
$ go run exit.go
exit status 3

# ビルドしてバイナリを実行するとステータスはターミナルで確認できます。
$ go build exit.go
$ ./exit
$ echo $?
3

# `!`は決して表示されないことに注意してください。