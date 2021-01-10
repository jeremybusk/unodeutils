package unodeutils

import (
    "fmt"
    "net"
    "os"
    "io/ioutil"
    "net/http"
    "encoding/json"
)


func remove_last_char(s string) string {
    r := []rune(s)
    return string(r[:len(r)-1])
}


func GetInternetIpv4() (string) {
    url := "https://api64.ipify.org?format=json"
    rsp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer rsp.Body.Close()
    rsp_body, err := ioutil.ReadAll(rsp.Body)
    if err != nil {
        panic(err)
    }
    var json_rsp map[string]interface{}
    json.Unmarshal([]byte(rsp_body),&json_rsp)
    ipv4 := json_rsp["ip"].(string)
    return ipv4
}


func GetIntranetIpaddrs() (string) {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
        os.Exit(1)
    }
    var ipv4s string
    for _, a := range addrs {
        if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            ipv4s = ipv4s + ipnet.IP.String() + "|"
        }
    }
    nipv4s := remove_last_char(ipv4s)
    return nipv4s
}


func display(){
    intranet_ipv4 := get_intranet_ip()
    fmt.Printf("intranet ipv4s: %s\n", intranet_ipv4)
    // internet_ipv4 := get_internet_ip()
    var internet_ipv4 = get_internet_ip()
    fmt.Printf("internet ipv4: %s\n", internet_ipv4)
}


func UtilHello() string {
    return "Hello, world."
}
