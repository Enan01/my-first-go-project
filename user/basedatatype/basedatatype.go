package main

import (
    "fmt"
    "strings"
    "bytes"
    "strconv"
)

func main() {
    floattest()
    inttest()
    weiyunsuantest()
    intuinttest()
    typetransfer()
    jinzhitest()
    asciitest()
    stringtest()
    basename("basename.go")
    basename2("basename2.go")
    fmt.Println(comma("12321321312"))
    fmt.Println(intsToString([]int{1, 2, 3}))
    fmt.Println(stringToInt("923874"))
}

func floattest() {
    var f float32 = 16777216
    fmt.Println(f == f + 1)
}

func inttest() {
    var u uint8 = 255
    fmt.Println(u, u+1, u*u)
    var i int8 = 127
    fmt.Println(i, i+1, i*i)
}

func weiyunsuantest() {
    var x uint8 = 1 << 1 | 1 << 5
    var y uint8 = 1 << 1 | 1 << 2
    fmt.Printf("%08b\n", x)
    fmt.Printf("%08b\n", y)
    fmt.Printf("%08b\n", x&y)
    fmt.Printf("%08b\n", x|y)
    fmt.Printf("%08b\n", x^y)
    fmt.Printf("%08b\n", x&^y)

    for i := uint(0); i < 8; i++ {
        if x & (1 << i) != 0 {
            fmt.Println(i)
        }
    }

    fmt.Printf("%08b\n", x << 1)
    fmt.Printf("%08b\n", x >> 1)
}

func intuinttest() {
    medals := []string{"gold", "silver", "bronze"}
    for i := len(medals) - 1; i >= 0; i-- {
        fmt.Println(medals[i])
    }
}

func typetransfer() {
    f := 3.141
    i := int(f)
    fmt.Println(f, i)
    f = 1.99
    fmt.Println(int(f))
}

func jinzhitest() {
    o := 0666
    fmt.Printf("%d %[1]o %#[1]o\n", o)
    x := int64(0xdeadbeef)
    fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}

func asciitest() {
    ascii := 'a'
    unicode := '国'
    newline := '\n'
    fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
    fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
    fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
}

func stringtest() {
    s := "hello world"
    fmt.Println(len(s))
    fmt.Println(s[0], s[7])
    fmt.Println(s[:])
    fmt.Println(s[2:5])
}

func basename(s string) {
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '/' {
            s = s[i + 1:]
            break
        }
    }
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '.' {
            s = s[:i]
            break
        }
    }
    fmt.Println(s)
}

func basename2(s string) {
    slash := strings.LastIndex(s, "/")
    s = s[slash + 1:]
    if dot := strings.LastIndex(s, "."); dot >= 0 {
        s = s[:dot]
    }
    fmt.Println(s)
}

func comma(s string) string {
    n := len(s)
    if n <= 3 {
        return s
    }
    return comma(s[:n - 3]) + "," + s[n - 3:]
}

func intsToString(values []int) string {
    var buf bytes.Buffer
    buf.WriteByte('[')
    for i, v := range values {
        if i > 0 {
            buf.WriteString(", ")
        }
        fmt.Fprintf(&buf, "%d", v)
    }
    buf.WriteByte(']')
    return buf.String()
}

func stringToInt(value string) (int, int) {
    x, _ := strconv.ParseInt(value, 10, 0)
    y, _ := strconv.Atoi(value)
    return int(x), y;
}



