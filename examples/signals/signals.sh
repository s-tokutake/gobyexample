# 実行するとシグナル待ちでブロックされます。
# `ctrl-C` (ターミナルは`^C`と表示する)を入力すると`SIGINT`を送信できます。
# これによってプログラムは`中断`され終了します。
$ go run signals.go
awaiting signal
^C
interrupt
exiting
