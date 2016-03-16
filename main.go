// программа которая будет парсить xe.com и брать оттуда курсы валюты
// в первой версии программы только курсы доллара
// в дальнейшем можно будет указывать необходимые курсы валют
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

const uri string = "http://www.xe.com/currency/usd-us-dollar";

func main() {
    fmt.Println(reader())
}

func reader() string {
    resp, err := http.Get(uri)
    fmt.Println(resp.Body)

    if err != nil {
        return ""
    } else {
        defer resp.Body.Close()
        contents, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return ""
        }
        return string(contents);
    }
}

func formatter() string {
    return "";
}