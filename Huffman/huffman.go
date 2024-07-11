package main

import (
	"fmt"
	"sort"
)

type hnode[T comparable] struct {
    val T
    freq int
    is_leaf bool
    l, r *hnode[T]
}

func huffEncode[T comparable](m map[T]int) hnode[T] {
    
    var q1, q2 []hnode[T];   

    for v, f := range m {
        q1 = enqueue(hnode[T]{v, f, true, nil, nil}, q1)
    }

    sort.Slice(q1, func(i, j int) bool {
        return q1[i].freq < q1[j].freq;
    })

    for len(q1) != 0 || len(q2) > 1 {
        var n1 hnode[T];
        if len(q1) != 0 && (len(q2) == 0 || q1[0].freq < q2[0].freq) {
            n1, q1 = dequeue(q1);
        } else {
            n1, q2 = dequeue(q2);
        }
        
        var n2 hnode[T];
        if len(q1) != 0 && (len(q2) == 0 || q1[0].freq < q2[0].freq) {
            n2, q1 = dequeue(q1);
        } else {
            n2, q2 = dequeue(q2);
        }
        q2 = enqueue(hnode[T]{*new(T), n1.freq + n2.freq, false, &n1, &n2}, q2);
    }

    root, q2 := dequeue(q2);
    return root;
}

func huffMap[T comparable](root hnode[T]) map[T]string {
    res := make(map[T]string);

    var trav func(n hnode[T], path string);

    trav = func(n hnode[T], path string) {
        if n.is_leaf {
            res[n.val] = path;
            return;
        }
        if n.l != nil {
            trav(*n.l, path + "0");
        }
        if n.r != nil {
            trav(*n.r, path + "1");
        }
    }

    trav(root, "");
    return res;
}

func main() {
    var val string;
    var freq int;
    m := make(map[string]int);

    for n, _ := fmt.Scanf("%s %d", &val, &freq); n != 0; n, _ = fmt.Scanf("%s %d", &val, &freq) {
        m[val] = freq;
    }

    root := huffEncode(m);

    fmt.Println(huffMap(root));

    var str string;

    fmt.Scanf("%s", &str);

    start := root;
    for i := 0; i < len(str); i++ {
        if str[i] == '1' {
            start = *start.r;
        } else if str[i] == '0' {
            start = *start.l;
        }

        if start.is_leaf {
            fmt.Print(start.val);
            start = root;
        }
    }
}
