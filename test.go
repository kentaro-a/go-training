package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "log"
    "io/ioutil"
)

type JsonData struct {
    usrName string `json:"usr_name"`
    Age uint16
    Address string
}

func main() {
    fmt.Println("start")
    data, err := getJson()
    if err != nil {
        panic("error!")
    }

    fmt.Printf("%#v\n", data)
    // for _,v := range data {
    //     fmt.Printf("%#v\n", v)
    // }
}

func getJson() (data JsonData, err error){
    res, _ := http.Get("http://localhost/json.php")
    body, _ := ioutil.ReadAll(res.Body)
    if err := json.Unmarshal(body, &data); err != nil {
        log.Fatal(err)
    }
    return
}
