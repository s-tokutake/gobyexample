// 文字列から数値へのパースは基本的ですが多くのプログラムで一般的です。
// Goでの実現方法を示します。

package main

// 組み込みの`strconv`パッケージは数値のパースを提供します。
import "strconv"
import "fmt"

func main() {

	// `ParseFloat`の`64`はパースの精度をビットで指定しています。
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// `ParseInt`の`0`は文字列から基数を推論することを示します。
	// `64`は結果が64ビットに適合することを要求します。
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt`はHEXフォーマットの数値にも対応しています。
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// `ParseUint`もあります。
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi`は基数が10の便利な`int`パース関数です。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 入力がおかしい場合はエラーを返します。
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
