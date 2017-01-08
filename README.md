# Jump Consistent Hash

Jump consistent hash, a Go library which implements the jump consistent hash algorithm. As described by John Lamping and Eric Veach in the paper [A Fast, Minimal Memory, Consistent Hash Algorithm](http://arxiv.org/pdf/1406.2294v1.pdf).

From the paper:

*jump consistent hash requires no storage, is faster, and does a better job of evenly dividing the key space among the buckets and of evenly dividing the workload when the number of buckets changes. Its main limitation is that the buckets must be numbered sequentially, which makes it more suitable for data storage applications than for distributed web caching.*

# Installation

Requires go 1.7+

> go get github.com/umahmood/jump-consistent-hash <br/>
> cd $GOPATH/github.com/umahmood/jump-consistent-hash <br/>
> go test 

# Usage

```
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
```

# Documentation

> http://godoc.org/github.com/umahmood/jump-consistent-hash

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
