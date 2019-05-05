# file-writingを実行してみましょう。
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# 次にファイルの中身を確認します。
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# 次は今まで学んできたI/Oのアイディアを
# `stdin`と`stdout`のストリームに応用してみます。
