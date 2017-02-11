package main

import (
    "fmt"
    "sync"
)

func main() {
    //c := make(chan bool)
    //go func() {
    //    fmt.Println("Go Go Go!!!")
    //    c <- false
    //    close(c)
    //}()
    //
    //for v := range c {
    //    fmt.Println(v)
    //}

    //runtime.GOMAXPROCS(runtime.NumCPU())
    //wg := sync.WaitGroup{} //同步包
    //wg.Add(10)
    //for i := 0; i < 10; i++ {
    //    go Go(&wg, i)
    //}
    //wg.Wait()

    c1, c2 := make(chan int), make(chan string)
    o := make(chan bool)
    go func() {
        for {
            select {
            case v, ok := <- c1:
                if !ok {
                    o <- true
                    break
                }
                fmt.Println("c1", v)
            case v, ok := <- c2:
                if !ok {
                    o <- true
                    break
                }
                fmt.Println("c2", v)
            }
        }
    }()

    c1 <- 1
    c2 <- "hi"
    c1 <- 2
    c2 <- "hello"

    close(c1)
    close(c2)

    <- o
}

func Go(wg *sync.WaitGroup, index int) {
    a := 1
    for i := 0; i < 100000000; i++ {
        a += i
    }
    fmt.Println(index, a)

    wg.Done()
}

func GoRoutineTest(c chan int) {




}
