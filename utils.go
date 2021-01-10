package unodeutils

import (
    "fmt"
    "net"
    "os"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "math/rand"
    "time"
    "github.com/google/uuid"
    "github.com/jackc/pgx"
)


func remove_last_char(s string) string {
    r := []rune(s)
    return string(r[:len(r)-1])
}


func genUUID() string {
    id := uuid.New()
    // fmt.Printf("github.com/google/uuid:         %s\n", id.String())
    // fmt.Printf("github.com/google/uuid:         %s\n", id.String())
    return id.String()
}


func IsValidUUID(u string) bool {
    _, err := uuid.Parse(u)
    return err == nil
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
    intranet_ipv4 := GetIntranetIpaddrs()
    fmt.Printf("intranet ipv4s: %s\n", intranet_ipv4)
    // internet_ipv4 := get_internet_ip()
    var internet_ipv4 = GetInternetIpv4()
    fmt.Printf("internet ipv4: %s\n", internet_ipv4)
}


func UtilHello() string {
    return "Hello, world."
}


func RandStringBytes(n int) string {
    const letterBytes = "abcdefghijklmnopqrstuvwxyz"
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}


func TestDb() {
    conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
        os.Exit(1)
    }
    defer conn.Close(context.Background())

    // var uuid string
    uuid := genUUID()
    // var alias = 'jtest'
    var alias = RandStringBytes(10) 
    err = conn.QueryRow(context.Background(), "insert into agents (uuid, alias) VALUES ($1, $2) returning uuid", &alias, &uuid).Scan(&resp)
    if err != nil {
        fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
        os.Exit(1)
    }

    fmt.Println(resp)
}
