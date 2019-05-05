# `zeroval`は`main`の`i`を変えません。
# `zeroptr`は変えます。変数のメモリアドレスを参照しているからです。
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
