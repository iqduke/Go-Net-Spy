package main

import (
    "io/ioutil"
    "net/http"
    "fmt"
    "os"
)

func main()  {
    client := &http.Client{}
    request,err := http.NewRequest("GET","https://www.nichijou.com/p/n3jtw/",nil)

    if err != nil {
        fmt.Println("Fatal error",err.Error())
        os.Exit(0)
    }

    response, err := client.Do(request)

    defer response.Body.Close()

    body,err1:= ioutil.ReadAll(response.Body)

    if err1 != nil {
        fmt.Println("body err")
    }

    fmt.Println(string(body))
    
}