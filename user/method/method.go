package main

import (
    "math"
    "fmt"
)

func main() {
    p := Point{1, 2}
    q := Point{4, 6}
    fmt.Println(p.Distance(q))

    r := &Point{1, 2}
    r.ScaleBy(2)
    fmt.Println(*r)
    p.ScaleBy(3)
    fmt.Println(p)

    fmt.Println("------------------------")

    m := Values{"lang": {"en"}} // direct construction
    fmt.Printf("%p\n", m)
    m.Add("item", "1")
    m.Add("item", "2")
    fmt.Println(m.Get("lang")) // "en"
    fmt.Println(m.Get("q"))    // ""
    fmt.Println(m.Get("item")) // "1"      (first value)
    fmt.Println(m["item"])     // "[1 2]"  (direct map access)
    m = nil
    fmt.Println(m.Get("item")) // ""
    //m.Add("item", "3")         // panic: assignment to entry in nil map

    ma := map[string]string{}
    fmt.Printf("%p\n", ma)
    ma["name"] = "cool"
    fmt.Println(ma)
    change(ma)
    fmt.Println(ma)

    e := Employee{"enan"}
    fmt.Println(e.Name)
    e.ChangeName("enan01")
    fmt.Println(e.Name)
    e.ChangeName2("enan02")
    fmt.Println(e.Name)

    fmt.Println("------------------")
    e2 := &e
    fmt.Printf("%p\n", e2)
    fmt.Printf("%p\n", &e2)
}

type Point struct {
    X, Y float64
}

func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}

type Values map[string][]string
// Get returns the first value associated with the given key,
// or "" if there are none.
func (v Values) Get(key string) string {
    fmt.Printf("%p\n", v)
    if vs := (v)[key]; len(vs) > 0 {
        return vs[0]
    }
    return ""
}
// Add adds the value to key.
// It appends to any existing values associated with key.
func (v Values) Add(key, value string) {
    fmt.Printf("%p\n", v)
    (v)[key] = append((v)[key], value)
}



func change(m map[string]string) {
    fmt.Printf("%p\n", m)
    m["name"] = "kuring"
}

type Employee struct {
    Name string
}

func (e Employee) ChangeName(newName string) {
    e.Name = newName
}

func (e *Employee) ChangeName2(newName string) {
    e.Name = newName
}
