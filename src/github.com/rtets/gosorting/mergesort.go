package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func mergesort(input []int) ([]int) {
    ilen := len(input)
    if (ilen<2) {
        return input
    }
    left,right := split(input,ilen)
    return merge(mergesort(left),mergesort(right))
}

func merge(left []int,right []int) (merged []int){
    llen,rlen := len(left), len(right)
    mlen := llen+rlen
    merged = make([]int, mlen)
    l,r := 0,0
    for i := 0; i<mlen; i++ {
        if (l==llen) {
            copy(merged[i:mlen], right[r:rlen])
            return
        } else if (r==rlen) {
            copy(merged[i:mlen], left[l:llen])
            return
        } else if (left[l] < right[r]) {
            merged[i] = left[l]
            l++
        } else {
            merged[i] = right[r]
            r++
        }
    }
    return
}

func split(input []int, ilen int) (left []int, right []int){
    mid := ilen/2
    left, right = make([]int, mid), make([]int, ilen-mid)
    copy(left, input[0:mid])
    copy(right, input[mid:ilen])
    return
}

func main() {
    // scanner := bufio.NewScanner(os.Stdin)
    // unsorted := []int{}
    // for scanner.Scan() {
    //     i,err := strconv.Atoi(scanner.Text())
    //     if (err==nil) {
    //         unsorted = append(unsorted,i)
    //     }
    // }
    scanner := bufio.NewScanner(os.Stdin)
    index:=0
    unsorted := make([]int, 1000000)
    for scanner.Scan() {
        v,err := strconv.Atoi(scanner.Text())
        if (err==nil) {
            unsorted[index] = v
            index++
        }
    }
    unsorted = mergesort(unsorted)
    for _,v := range unsorted {
        fmt.Println(v)
    }

    // fmt.Println(os.Args[0])
    // unsorted := []int{}
    // fmt.Println(unsorted)
    // unsorted = mergesort(unsorted)
    // fmt.Println(unsorted)
    
    
}