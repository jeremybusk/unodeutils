package unodeutils 

import "testing"

func TestUtilHello(t *testing.T) {
    want := "Hello, world."
    if got := UtilHello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
