// URLは[リソースを見つけるための統一された方法](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/)を提供します。
// ここではGoでURLをパースする方法を示します。

package main

import "fmt"
import "net"
import "net/url"

func main() {

    // このサンプルのURLをパースします。スキーム、認証情報、ホスト、
    // ポート、パス、クエリパラメータ、クエリフラグメントが含まれています。
    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    // URLをパースしエラーがないことを確かめます。
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    // スキームへのアクセスは直截的です。
    fmt.Println(u.Scheme)

    // `User`には認証情報が含まれています。
    // `Username`と`Password`が個別に取得できます。
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    // `Host`にはホスト名とポートが含まれます。
    // `SplitHostPort`を使うとホストとポートを別々に取得できます。
    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    // `path`と`#`の後のフラグメントを取り出します。
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    // `k=v`という文字列形式のクエリパラメータを取得するには
    // `RawQuery`を使います。また、マップにパースすることもできます。
    // このマップは文字列から文字列のスライスへのマップです。
    // したがって最初の値だけが欲しいなら`[0]`で取得できます。
    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}
