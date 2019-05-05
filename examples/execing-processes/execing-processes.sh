# このプログラムは実行すると`ls`に置き換わります。
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Goは従来のUnixの`fork`は提供しません。
# しかし、ゴルーチンの開始、プロセスの起動、プロセスの実行が
# `fork`を使う場合のほとんどのケースをカバーするため問題になりません。
