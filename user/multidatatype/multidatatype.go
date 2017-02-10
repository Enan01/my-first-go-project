package main

import (
    "fmt"
    "bufio"
    "os"
    "encoding/json"
    "log"
)

func main() {
    array()
    array1()
    slice()

    a := [...]int{1, 2, 3, 4, 5}
    reverse(a[:])
    fmt.Println(a)

    appendtest()

    s := []int{5, 6, 7, 8, 9}
    fmt.Println(remove(s, 2))

    maptest()
    //dedup()

    sortArray := []int{3, 5, 7, 1, 9}
    Sort(sortArray)
    fmt.Println(sortArray)

    jsontest()
}

func array() {
    var a [3]int
    fmt.Println(a[0])
    fmt.Println(a[len(a) - 1])

    for i, v := range a {
        fmt.Printf("%d %d\n", i, v)
    }

    for _, v := range a {
        fmt.Printf("%d\n", v)
    }

    b := [3]int{1, 2, 3}
    b = [...]int{1, 2, 4}
    fmt.Println(b[2])
}

func array1() {
    type Currency int

    const (
        USD Currency = iota
        EUR
        GBP
        RMB
    )

    symbol := [...]string{USD : "$", EUR : "@", GBP : "*", RMB : "+"}
    fmt.Println(RMB, symbol[RMB])

}

func slice() {
    months := [...]string{1 : "jan", 2 : "feb", 3 : "mar", 4 : "apr", 5 : "may", 6 : "jun", 7 : "jul", 8 : "aug", 9 : "sep", 10 : "oct", 11 : "nov", 12 : "dec"}

    Q2 := months[4 : 7]
    summer := months[6 : 9]
    fmt.Println(Q2)
    fmt.Println(summer)

    endlessSummeer := summer[:5]
    fmt.Println(endlessSummeer)
}

func reverse(s []int) {
    for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
        s[i], s[j] = s[j], s[i]
    }
}

func equal(x, y []string) bool {
    if len(x) != len(y) {
        return false
    }
    for i := 0; i < len(x); i++ {
        if x[i] != y[i] {
            return false
        }
    }
    return true
}

func appendtest() {
    var runes []rune
    for _, r := range "Hello, 世界" {
        runes = append(runes, r)
    }
    fmt.Printf("%q\n", runes)
    runes = []rune("Hello, 世界")
    fmt.Printf("%q\n", runes)
}

func remove(slice []int, i int) []int {
    copy(slice[i:], slice[i + 1:])
    return slice[:len(slice) - 1]
}

func maptest() {
    ages := make(map[string]int)
    ages["alice"] = 33
    ages["charlie"] = 42

    fmt.Println(ages["alice"])

    delete(ages, "alice")

    fmt.Println(ages["alice"])
}

func dedup() {
    seen := make(map[string]bool)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        line := input.Text()
        if !seen[line] {
            seen[line] = true
            fmt.Println(line)
        }
    }

    if err := input.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
        os.Exit(1)
    }
}

type tree struct {
    value       int
    left, right *tree
}

func Sort(values []int) {
    var root *tree
    for _, v := range values {
        root = add(root, v)
    }
    appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
    if t != nil {
        values = appendValues(values, t.left)
        values = append(values, t.value)
        values = appendValues(values, t.right)
    }
    return values
}

func add(t *tree, value int) *tree {
    if t == nil {
        t = new(tree)
        t.value = value
        return t
    }
    if value < t.value {
        t.left = add(t.left, value)
    } else {
        t.right = add(t.right, value)
    }
    return t
}

type Movie struct {
    Title  string
    Year   int `json:"released"`
    Color  bool `json:"color,omitempty"`
    Actors []string
}

func jsontest() {
    var movies = []Movie{
        {Title: "Casablanca", Year: 1942, Color: false,
            Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
        {Title: "Cool Hand Luke", Year: 1967, Color: true,
            Actors: []string{"Paul Newman"}},
        {Title: "Bullitt", Year: 1968, Color: true,
            Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
        // ...
    }
    data, err := json.Marshal(movies)
    if err != nil {
        log.Fatalf("JSON marshaling failed: %s", err)
    }
    fmt.Printf("%s\n", data)

    data, err = json.MarshalIndent(movies, "", "    ")
    if err != nil {
        log.Fatalf("JSON marshaling failed: %s", err)
    }
    fmt.Printf("%s\n", data)

    var titles []struct{Title string}
    if err := json.Unmarshal(data, &titles); err != nil {
        log.Fatalf("JSON marshaling failed: %s", err)
    }
    fmt.Println(titles)
}

