// Goは<a href="http://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>再帰関数</em></a>をサポートします。
// おなじみの階乗の例を示します。

package main

import "fmt"

// この`fact`関数は`fact(0)`にたどり着くまで自分自身を呼びます。
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
