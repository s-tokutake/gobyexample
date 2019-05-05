// 自然な順ではない順でコレクションをソートしたい場合があります。
// 例えば、アルファベット順ではなく長さで文字列をソートしたい場合です。
// ここではGoでの実現方法を示します。

package main

import "sort"
import "fmt"

// Goでカスタムの関数でソートをするためには
// 対応する型が必要です。ここでは`byLength`型を定義します。
// これは単に`[]string`のエイリアスです。
type byLength []string

// `sort.Interface`を実装します。`Len`、`Less`、`Swap`を持ちます。
// これによって `sort`パッケージのジェネリックな`Sort`関数を使えます。
// `Len`と`Swap`は多くの型で同じような実装になります。
// `Less`は実際のカスタムの並べ替えのロジックを持ちます。
// この例では、文字列の長さの昇順に並べます。
// したがって、`len(s[i])`と`len(s[j])`を使います。
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// 準備が整いました。`fruits`スライスを`byLength`にキャストし、
// `sort.Sort`を使ってカスタムのソートを実行します。
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
