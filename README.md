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

    
