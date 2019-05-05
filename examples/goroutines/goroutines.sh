# このプログラムを実行するとブロッキングの呼び出しの出力が先に表示されます。
# そしてその後にゴルーチンの出力です。
# ひとつゴルーチンの出力に別のゴルーチンの出力が挟み込まれるのは
# ゴルーチンが平行に動いていることを示しています。
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
<enter>
done

# 次はGoの平行処理でゴルーチンを補完しているチャンネルです。
