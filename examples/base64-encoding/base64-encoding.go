// Goは[base64エンコーディング/デコーディング](http://en.wikipedia.org/wiki/Base64)を組み込みでサポートします。

package main

// `encoding/base64`パッケージをデフォルトの`base64`でななく
// b64という名前でインポートします。これによっていくらかスペースを節約します。
import b64 "encoding/base64"
import "fmt"

func main() {

	// エンコード/デコードをする文字列です。
	data := "abc123!?$*&()'-=@~"

	// Goは標準とURL互換の両方のbase64をサポートします。
	// ここでは標準のエンコーダを使ってエンコードします。
	// エンコーダには`[]byte`が必要なので`string`をキャストしています。
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// デコードはエラーを返す場合があります。
	// これによって入力の形式が正しいかどうかをチェックできます。
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// URL互換のbase64のエンコード/デコードです。
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
