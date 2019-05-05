# 実行するとハッシュを計算し人間が読めるhexフォーマットで表示します。
$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94


# 上と同じようなパターンで異なるハッシュを計算できます。
# 例えば、MD5なら`crypto/md5`をインポートして、`md5.New()`です。

# 暗号論的ハッシュが必要なら[ハッシュの強度](http://en.wikipedia.org/wiki/Cryptographic_hash_function)ついてしっかりと調べてください。!
