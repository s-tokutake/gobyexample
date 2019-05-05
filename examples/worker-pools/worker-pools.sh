# さまざまなワーカーから実行された5つのジョブを示します。
# トータルの処理で約5秒かかっていますがプログラム自体は2秒程度で完了しています。
# これは3つのワーカーが平行して動いているからです。
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
