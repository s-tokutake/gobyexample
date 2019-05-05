# コマンドラインフラグの挙動を確かめるにはコンパイルして
# バイナリを直接実行するといいでしょう。
$ go build command-line-flags.go

# まずすべてのフラグを渡してビルドしたプログラムを実行します。
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# フラグを省略したらデフォルトの値が設定されます。
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# 末尾の位置引数はフラグの後に提供されます。
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# `flag`パッケージは位置引数の前にすべてのフラグが出現することを要求します。
# そうしないとフラグが位置引数だと解釈されるからです。
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# `-h`または`--help`フラグはコマンドラインプログラムのヘルプを
# 自動的に生成します。
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# `flag`パッケージで定義しなかったフラグを渡すと
# プログラムはエラーメッセージを表示し、ヘルプを表示します。
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...

# 次は環境変数です。これもプログラムをパラメータ化する一般的な方法です。
