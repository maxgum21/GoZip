package main

func enqueue[T comparable](val T, q []T) []T {
    return append(q, val);
}

func dequeue[T comparable](q []T) (T, []T) {
    return q[0], q[1:];
}
