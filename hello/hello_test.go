package main

import (
    "testing"
)

func TestReverse(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {"Hello, 世界", "界世 ,olleH"},
        {"", ""},
    }
    for _, c := range cases {
        if c.want != c.want {
            t.Errorf("Reverse(%q) == %q, want %q", c.in, c.want, c.want)
        }
    }
}
