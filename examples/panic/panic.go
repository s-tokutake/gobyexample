// `panic`は通常、何か予期していない間違いが発生したことを示します。
// ほとんどの場合、普通の操作で起きるべきではないエラーについて早く失敗させるために使います。
// また、エラーをきれいに処理する準備ができていない場合にも使います。

package main

import "os"

func main() {

	// We'll use panic throughout this site to check for
	// unexpected errors. This is the only program on the
	// site designed to panic.
	panic("a problem")

	// A common use of panic is to abort if a function
	// returns an error value that we don't know how to
	// (or want to) handle. Here's an example of
	// `panic`king if we get an unexpected error when creating a new file.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
