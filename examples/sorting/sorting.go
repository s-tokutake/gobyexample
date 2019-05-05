// Goの`sort`パッケージは組み込みの型とユーザー定義の型に対する
// ソート処理を実装します。まずは、組み込み型に対するソートを見てみましょう。

package main

import "fmt"
import "sort"

func main() {

    // Sortメソッドは組み込みの型に使えます。
    // ここでは文字列をソートしています。
    // ソートはin-placeで行われます。
    // したがって与えられたスライスをソートし、新しいスライスを返しません。
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    // これは`int`のソートです。
    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    // `sort`パッケージを使って
    // スライスがソート済みかどうかをチェックできます。
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}
