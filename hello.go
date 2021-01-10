package unodeutils 

import "rsc.io/quote/v3"

func Hello() string {
    return quote.HelloV3()
}

func Proverb() string {
    return quote.Concurrency()
}

func Hi() string {
    // return (println("Hello, world"))
    return ("Hello, world")
    // println("Hello, world")
}
