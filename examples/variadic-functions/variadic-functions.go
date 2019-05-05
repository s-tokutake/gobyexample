// [_可変長引数関数_](http://en.wikipedia.org/wiki/Variadic_function)は任意の数の末尾引数を受け取ります。
// 例えば、`fmt.Println`は可変長引数関数です。

package main

import "fmt"

// この関数は任意の数の`int`を引数に取ります。
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    // 可変長引数関数は普通に個別の引数を渡して呼び出します。
    sum(1, 2)
    sum(1, 2, 3)

    // スライスに複数の引数が入っているなら
    // このように`func(slice...)`というかたちで渡せます。
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
