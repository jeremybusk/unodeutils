# Getting Started

```
go get github.com/jeremybusk/unodeutils
or
go get -u github.com/jeremybusk/unodeutils
```

It will put it somewhere like
```
~/go/pkg/mod/github.com/jeremybusk/unodeutils...
check it
ls ~/go/src/github.com/jeremybusk/unodeutils/
```

myexample.go
```
package main

import (
    "fmt"
    "github.com/jeremybusk/unodeutils"
)

func main() {
    msg := unodeutils.Hello()
    fmt.Printf("line1: %s\n", msg)
    ipv4 := unodeutils.GetInternetIpv4()
    fmt.Printf("line2: %s\n", ipv4)
    ipaddrs := unodeutils.GetIntranetIpaddrs()
    fmt.Printf("line2: %s\n", ipaddrs)
}
```

```
go build myexample.go
./myexample
```

