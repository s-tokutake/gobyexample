// Goの_構造体_は型付けられたフィールドのコレクションです。
// データをまとめてレコードを形作るに便利です。

package main

import "fmt"

// `person`構造体には`name`と`age`フィールドがあります。
type person struct {
	name string
	age  int
}

func main() {

	// この構文で新しい構造体を作ります。
	fmt.Println(person{"Bob", 20})

	// 初期化でフィールドを指定することもできます。
	fmt.Println(person{name: "Alice", age: 30})

	// 省略されたフィールドはゼロ値になります。
	fmt.Println(person{name: "Fred"})

	// `&`プレフィクスを使うと構造体のポインタが作れます。
	fmt.Println(&person{name: "Ann", age: 40})

	// ドットを使ってフィールドにアクセスします。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 構造体のポインタにもドットが使えます。
	// ポインタは自動的に参照先を見ます。
	sp := &s
	fmt.Println(sp.age)

	// 構造体は可変です。
	sp.age = 51
	fmt.Println(sp.age)
}
