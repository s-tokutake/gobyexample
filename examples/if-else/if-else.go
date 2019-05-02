// Goでは`if`と`else`を使った分岐は
// 直截的です。

package main

import "fmt"

func main() {

	// 基本的な例です。
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// `if`文はelseなしでも使えます。
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 文は条件に先行することができます。この文で宣言された変数は
	// すべての分岐で有効です。
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Goの場合、条件を囲む丸括弧は不要です。
// しかし、波括弧は必要です。
