package main

import (
	"fmt"
)

type node struct {
    key, val int
    left, right *node
}

func addNode(key, value int, head **node) bool {
    if *head == nil {
        *head = &node{key, value, nil, nil};
        return true;
    }

    if key > (*head).key {
        return addNode(key, value, &(*head).right);
    }

    if key < (*head).key {
        return addNode(key, value, &(*head).left);
    }

    (*head).val = value;
    return false
}

func findNode(key int, head **node) (int, *node, bool) {
    if key == (*head).key {
        return (*head).val, *head, true;
    }

    if key > (*head).key {
        if (*head).right == nil {
            return 0, nil, false;
        }
        return findNode(key, &(*head).right);
    }
    
    if (*head).left == nil {
        return 0, nil, false;
    }
    return findNode(key, &(*head).left);
}

func deleteNode(key int, head **node) bool {
    if *head == nil {
        return false
    }

    if key == (*head).key {
        if (*head).left == nil && (*head).right == nil {
            *head = nil;
            return true;
        }
        if (*head).left != nil && (*head).right == nil {
            *head = (*head).left;
            return true;
        } else if (*head).left == nil && (*head).right != nil {
            *head = (*head).right;
            return true;
        }

        par, succ := *head, (*head).right;
        for succ.left != nil {
            par = succ;
            succ = succ.left;
        }
        par.left = succ.right;
        succ.left = (*head).left;
        succ.right = (*head).right;
        *head = succ;
        return true;
    }

    if key > (*head).key {
        return deleteNode(key, &(*head).right);
    }
    return deleteNode(key, &(*head).left);
}

func main() {
    var root *node;

    addNode(10, 1, &root);
    addNode(5, 2, &root);
    addNode(15, 3, &root);
    addNode(12, 4, &root);

    deleteNode(15, &root);
    fmt.Println(root.right);
}
