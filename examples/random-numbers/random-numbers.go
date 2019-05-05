// Goの`math/rand`パッケージは[擬似乱数](http://en.wikipedia.org/wiki/Pseudorandom_number_generator)生成を提供します。

package main

import "time"
import "fmt"
import "math/rand"

func main() {

    // 例えば、`rand.Intn`はランダムな`int`のnを返します。
    // `0 <= n < 100`です。
    fmt.Print(rand.Intn(100), ",")
    fmt.Print(rand.Intn(100))
    fmt.Println()

    // `rand.Float64`は`float64`の`f`を返します。
    // `0.0 <= f < 1.0`です。
    fmt.Println(rand.Float64())

    // これを使って範囲を変えてランダムな浮動小数点数を返すこともできます。
    // この例では`5.0 <= f' < 10.0`です。
    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()

    // デフォルトのジェネレータは決定的です。したがって、デフォルトでは、 so it'll
    // 毎回、同じ数のシーケンスを生成します。
    // 変化のあるシーケンスを作るには、シードを与えます。
    // 秘密にしたい乱数に使うのは安全ではありません。その場合は`crypto/rand`を使います。
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    // `rand`パッケージの関数と同じように`rand.Rand`を呼び出します。
    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()

    // もしシードが同じ数値なら乱数の同じシーケンスを生成します。
    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))
    fmt.Println()
    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Print(r3.Intn(100), ",")
    fmt.Print(r3.Intn(100))
}
