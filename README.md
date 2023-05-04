# Goのポインタについて学ぶ

これはGo言語の基本を学ぶためのGoプロジェクトです。このプロジェクトには、`pointer`パッケージを含むいくつかのパッケージがあります。`pointer`パッケージには2つのファイルが含まれています。

- `pointerCase1.go`: Goでポインタを使用する方法を示すプログラム。
- `pointerCase2.go`: Goでポインタレシーバを使用する方法を示すプログラム。

## Pointerパッケージ

`pointer`パッケージには、Goでポインタを使用する方法を示すプログラムが含まれています。以下のファイルが含まれています。

- `pointerCase1.go`: メモリアドレスを格納するためにポインタを使用する方法を示すプログラム。
- `pointerCase2.go`: 値を変更するためにポインタレシーバを使用する方法を示すプログラム。

### 'pointerCase1.go'
    
    package main

    import "fmt"

    func main() {
        var x = 100
        fmt.Println("x address:\t", &x)

        var y *int
        fmt.Println("y value:\t", y)
        fmt.Println("y address:\t", &y)
        y = &x
        fmt.Println("y value:\t", y)
        fmt.Println("y address:\t", &y)
    }

    
#### 説明
1. main関数が呼び出され、スタックフレームが作成される。xとyというローカル変数が含まれる。
2. xがスタック上に確保され、値100が割り当てられる。
3. xのアドレスが出力される。
4. yがスタック上に確保されるが、まだどのアドレスも指していない（nil）。
5. yの値（nil）とy自体のアドレスが出力される。
6. yにxのアドレスが割り当てられる。yはxを指す。
7. yの値（xのアドレス）とy自体のアドレスが再度出力される。

このプログラムでは、ヒープは使用されていない。すべての変数とポインタはスタック上に配置されている。
大切なポイントは、`y`はポインタで`y`自体にもアドレスがあるということ。

## struct パッケージ

### ___概要___
このGoプログラムは、簡単なUser構造体を定義し、それを使用してインスタンスを作成する方法を示しています。このプログラムでは、logパッケージを使用して出力をコンソールに表示します。

### プログラムの説明
#### パッケージのインポート
```Go
package main

import (
    "log"
    "time"
)
```
このプログラムでは、`main`パッケージを使用しています。また、`log`と`time`パッケージをインポートしています。これらのパッケージは、ログの出力や時間の扱い方に関する機能を提供しています。

### __User構造体の定義___
```Go
type User struct {
    FirstName   string
    LastName    string
    PhoneNumber string
    Age         int
    BirthDate   time.Time
}
```
このプログラムでは、`User`という名前の構造体を定義しています。User構造体には、5つのフィールドがあります。`FirstName`、`LastName`、`PhoneNumber`は`string`型、`Ag`eは`int`型、`BirthDate`は`time.Time`型です。これらのフィールドは、ユーザーの名前、電話番号、年齢、誕生日に関する情報を格納します。

### ___main関数___
```Go
func main() {
    user := User {
        FirstName: "Trevor",
        LastName: "Sawler",
        PhoneNumber: "1 555 555 1212",
    }

    log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}
```
main関数は、このプログラムのエントリーポイントです。この関数では、User構造体のインスタンスを作成しています。
<br>
`FirstName`、`LastName`、`PhoneNumber`フィールドには、それぞれ値が設定されていますが、`BirthDate`フィールドには値が設定されていません。`BirthDate`の値が初期化されていないため、ゼロ値である`0001-01-01 00:00:00 +0000 UTC`が表示されます。

最後に、`log.Println()`関数を使用して、User構造体の`FirstName`、`LastName`、`BirthDate`フィールドを出力しています。

### ___実行方法___
このプログラムを実行するには、以下のコマンドを使用します。

go run main.go
### ___結果___
上記のコマンドを実行すると、以下のような結果が表示されます。

yaml
Copy code
Trevor Sawler Birthdate: 0001-01-01 00:00:00 +0000 UTC
### ___結論___
このプログラムは、Go言語で簡単な構造体を定義する方法と、その構造体のインスタンスを作成する方法を示しています。また、ログ出力としてlogパッケージを使用する方法を示しています。


このプログラムを通じて、Go言語で構造体とログ出力を扱う方法を学ぶことができました。