// программа которая будет парсить xe.com и брать оттуда курсы валюты
// в первой версии программы только курсы доллара
// в дальнейшем можно будет указывать необходимые курсы валют
package main

import (
    "fmt"
)

const uri string = "http://www.xe.com/currency/usd-us-dollar";

func main() {
    fmt.Println(cravler())
}

func cravler() string {
    return uri;
}

func formatter() string {
    return "";
}