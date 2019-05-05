# 実行してみると約40,000回操作が行われたのがわかります。
$ go run atomic-counters.go
ops: 41419

# 次はmutexです。状態を管理する別の方法を提供します。
