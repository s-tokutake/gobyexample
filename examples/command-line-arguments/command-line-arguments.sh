# コマンドライン引数の挙動を確かめるには
# `go build`でバイナリをビルドするのがいいでしょう。
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c

# 次はフラグを使ったコマンドラインの処理を紹介します。
