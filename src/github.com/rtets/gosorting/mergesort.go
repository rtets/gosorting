package main

import (
    "fmt"
    "os"
)

func mergesort(input []int) {
    ilen := len(input)
    if (ilen<2) {
        return input
    }
    left,right := split(input,ilen)
    return merge(mergesort(left),mergesort(right))
}

func merge(left []int,right []int) (merged []int){
    llen,rlen := len(left), len(right)
    mlen = llen+rlen
    merged := make([]int, mlen)
    l,r = 0,0
    for (i := 0; i<mlen; i++) {
        if (left[l] < right[r]) {

        }
    }
}

func split(input []int, ilen int) (left []int, right []int){
    mid := ilen/2
    left, right := make([]int, mid), make([]int, ilen-mid)
    copy(left, input[0:mid])
    copy(right, input[mid:ilen])
    return
}

func main() {
    fmt.Println("Hello world!")
    
}