# 標準とURLのbase64エンコーダではエンコード結果がわずかに違います(末尾の `+`と`-`)。
# しかし、デコードは望み通りの元の文字列になります。
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
