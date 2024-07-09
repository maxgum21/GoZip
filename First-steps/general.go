package main

import (
    "fmt";
    "time";
)

func maximum(args...int) (int, int) {
    var max_id, max_val int = 0, args[0];
    for id, val := range args {
        if val > max_val {
            max_val = val;
            max_id = id;
        }
    }
    return max_id, max_val;
}

func worker_routine(done chan bool) {
    fmt.Println("Working");
    time.Sleep(time.Second);
    fmt.Println("Stopped working");

    done <- true;
}

func main() {
    fmt.Println("Hello world!");

    // Variabels and operations
    var b int = 7;
    fmt.Println(b << 1);

    // Loops
    for i := 0; i < 7; i++ {
        if i % 2 == 1 {
            fmt.Println("got i =", i)
        }
    }

    // Maps
    var m = map[string]int{"monday": 1, "tuesday": 2};
    _, present := m["wednesday"];
    fmt.Println(present);

    // Variadic functions
    fmt.Println(maximum(1, 2, 7, 4, 5, 6));
    var nums = []int{2346, 1234, 23546, 122};
    fmt.Println(maximum(nums...));

    // Goroutines
    done := make(chan bool, 1);
    go worker_routine(done);

    fmt.Println("Main thread waiting...");

    <- done;

    fmt.Println("Sync complete!");

    /* YOU CAN'T DO THAT:
    var g = []int{1, 2, 3} ;

    var i1, i2, i3 int = g;*/

}
