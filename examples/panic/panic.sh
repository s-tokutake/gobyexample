# これを実行するとpanicが発生します。
# エラーメッセージとゴルーチンのトレースが表示され
# 0でないステータスが返ります。
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# エラー処理で例外を使う他の言語と違い、
# Goは可能な場合はエラーを示す戻り値を返すのが普通です。
