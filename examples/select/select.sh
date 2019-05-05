# まず`"one"`を受け取り次に`"two"`を受け取ります。
$ time go run select.go 
received one
received two

# トータルの実行時間は2秒程度です。
# 1秒、2秒の`Sleeps`は平行で実行されるからです。
real	0m2.245s
