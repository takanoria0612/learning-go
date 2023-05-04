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

