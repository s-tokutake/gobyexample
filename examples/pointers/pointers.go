// Goは<em><a href="http://en.wikipedia.org/wiki/Pointer_(computer_programming)">ポインタ</a></em>をサポートします。
// プログラム内で値とレコードの参照を渡すことができます。

package main

import "fmt"

// ふたつの関数`zeroval`と`zeroptr`を比較することで
// ポインタがどのように動作するかを示します。`zeroval`は
// `int`のパラメータを取ります。よって引数は値として渡されます。
// `zeroval`は`ival`のコピーを受け取ります。
// これは関数の呼び出し使われた値とは区別されます。
func zeroval(ival int) {
	ival = 0
}

// 対して、`zeroptr`のパラメータは`*int`です。
// つまり`int`のポインタを受け取ります。関数の中身の`*iptr`は
// pointerのメモリアドレスからそのアドレスの現在の値を_デリファレンス_します。
// デリファレンスされたポインタに値を設定すると参照されたアドレスの値が変わります。
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// `&i`という構文は`i`のメモリアドレス、
	// つまり、`i`のポインタを返します。
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// ポインタもプリントすることができます。
	fmt.Println("pointer:", &i)
}
