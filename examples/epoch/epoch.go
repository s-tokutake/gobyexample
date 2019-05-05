// [Unixエポック](http://en.wikipedia.org/wiki/Unix_time)からの秒、ミリ秒、ナノ秒を取得する
// のはプログラムの一般的な要件です。
// ここではGoでの実現方法を示します。

package main

import "fmt"
import "time"

func main() {

    // `time.Now`の`Unix`と`UnixNano`でUnixエポックからの経過時間を
    // 秒とナノ秒で取得できます。
    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)

    // `UnixMillis`はありません。ミリ秒で取得するにはナノ秒から計算します。
    millis := nanos / 1000000
    fmt.Println(secs)
    fmt.Println(millis)
    fmt.Println(nanos)

    // 整数値の秒やナノ秒を対応する`time`に変換できます。
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}
