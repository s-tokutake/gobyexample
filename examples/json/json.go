// Goは組み込みでJSONのエンコーディングとデコーディングをサポートします。
// 組み込みのデータ型とカスタムのデータ型のエンコーディングとデコーディングを提供します。

package main

import "encoding/json"
import "fmt"
import "os"

// ふたつの構造体を使ってカスタムの型の
// エンコーディングとデコーディングをを実演します。
type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// まず、基本データ型をJSON文字列にします。
	// いくつかのアトミックな値の例を示します。
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// スライスとマップは予想通りの
	// JSON配列とJSONオブジェクトにエンコーディングされます。
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSONパッケージを使うとカスタムのデータ型を自動的にエンコードできます。
	// 出力に含まれるのは公開されているフィールドだけであり、
	// デフォルトではフィールド名はJSONのキーとして使われます。
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 構造体のフィールド宣言にタグを付けることで
	// JSONのキー名をカスタマイズできます。
	// 上の`response2`の定義のタグを見てください。
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// JSONのデータをGoの値にデコードします。
	// この例は一般的なデータ構造です。
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// JSONパッケージがデコードしたデータを設定するための
	// 変数を宣言する必要があります。この`map[string]interface{}`は
	// 文字列から任意のデータ型へのマップを保持します。
	var dat map[string]interface{}

	// 実際にデコードします。関連するエラーをチェックしています。
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// デコードされたマップの値を使うには
	// 適切な型にキャストする必要があります。
	// この例では`num`を`float64`にキャストしています。
	num := dat["num"].(float64)
	fmt.Println(num)

	// ネストしたデータにはひと続きのキャストが必要です。
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// JSONをカスタムのデータ型にデコードすることもできます。
	// こうするとプログラムはよりタイプセーフになり、
	// デコードしたデータにアクセスするときに、
	// 型アサーションが必要なくなります。
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 上の例ではバイトと文字列をデータとJSONの間の媒介として使いました。
	// JSONを直接`os.Stdout`のような`os.Writer`にストリームすることができます。
	// HTTPレスポンスのボディにもストリームできます。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
