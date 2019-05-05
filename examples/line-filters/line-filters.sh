# ラインフィルタを試してみましょう。
# まず小文字の文字列を含んだのファイルを作ります。
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# Then use the line filter to get uppercase lines.
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
