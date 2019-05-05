// _マップ_はGoの組み込みの[連想データ型](http://en.wikipedia.org/wiki/Associative_array)です
//(他の言語では_ハッシュ_や_ディクショナリ_と呼ばれます)。

package main

import "fmt"

func main() {

	// 空のマップを作成するには組み込みの`make`を使います。
	// `make(map[key-type]val-type)`
	m := make(map[string]int)

	// おなじみの構文`name[key] = val`でキーと値のペアを設定します。
	m["k1"] = 7
	m["k2"] = 13

	// `fmt.Println`などを使ってマップを表示するとすべての
	// キーと値のペアを表示します。
	fmt.Println("map:", m)

	// `name[key]`で値を取り出します。
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// マップに対して組み込みの`len`を呼び出すと
	// キー/値のペアの数を返します。
	fmt.Println("len:", len(m))

	// 組み込みの`delete`はマップからキー/値のペアを削除します。
	delete(m, "k2")
	fmt.Println("map:", m)

	// マップから値を取り出す時のオプショナルな2番目の戻り値は
	// そのキーがマップにあるかどうかを返します。
	// これによってキーがない場合とキーがあって値が`0`や`""`などのゼロ値の場合
	// を明確に区別できます。
	// ここでは値自体は使っていないので、_ブランク識別子_である`_`を使って無視しています。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 次のような構文を使えばひとつの行で
	// 新しいマップの宣言と初期化が行えます。
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
