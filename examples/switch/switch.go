// _スイッチ文_は多くの分岐の条件を示します。

package main

import "fmt"
import "time"

func main() {

	// 基本的な`switch`です。
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// カンマを使って同じ`case`文の中で複数の式を書けます。
	// `default`ケースは任意です。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// 式のない`switch`はif/elseの別表現です。
	// またここでは、`case`式が非定数でどのように書くかも示しています。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// 型の`switch`は値ではなく型を比較します。
	// インターフェース値の型を見つけるのに使えます。
	// 例えば、ここでは、変数`t`がcase文に書いてある型かどうかを判定しています。
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
