// 私たちのプログラムはデータのコレクションに対して操作をする場合があります。
// 与えられた条件を満たすすベての要素を取り出したり、
// カスタムの関数ですべての要素を新しいコレクションにマップしたりします。

// これを実現するとき、慣用的に[ジェネリック](http://en.wikipedia.org/wiki/Generic_programming)な
// データ構造とアルゴリズムを使うプログラミング言語もあります。
// しかし、Goはジェネリックをサポートしません。
// Goでは必要に応じてコレクション関数を定義するのが一般的です。

// ここでは`strings`のスライスのコレクション関数を示します。
// この例を使って独自の関数を構築することも可能です。
// ヘルパ関数を定義し呼び出すよりも、インラインでコレクション操作をするコードを
// 直接書いた方がきれいな場合もあることに注意してください。

package main

import "strings"
import "fmt"

// Indexは対象の文字列`t`の最初のインデックスを返します。
// 見つからなかった場合は0を返します。
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Includeは対象文字列tがスライスに含まれている場合`true`を返します。
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Anyはスライス内のひとつ以上の文字列が術語`f`を満たす場合に`true`を返します。
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// Allはスライス内のすべての文字列が術語`f`を満たす場合に`true`を返します。
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filterは与えられたスライス内の、術語`f`を満たすすべての文字列を含む新しいスライスを返します。
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Mapはもとのスライス内の文字列に関数`f`を適用した結果を含む新しいスライスを作成します。
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	// ここでコレクション関数を実行してみます。
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	// 上の例はすべて無名関数を使っていますが、
	// 正しい型の名前付きの関数を渡すこともできます。
	fmt.Println(Map(strs, strings.ToUpper))

}
