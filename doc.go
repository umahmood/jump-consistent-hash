/*
Package jump implements the jump consistent hash algorithm. A fast, minimal
memory, consistent hash algorithm (see http://arxiv.org/pdf/1406.2294v1.pdf).

Example usage:

    package main

    import (
        jump "github.com/umahmood/jump-consistent-hash"
    )

    func main() {
        buckets := make([]int32, 100)
        for k := uint64(0); k < 100; k++ {
            buckets[jump.Hash(k, int32(len(buckets)))]++
        }
    }
*/
package jump
