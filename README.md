# A home for usual tools

Sometimes you need to create a simple or complex tool for a specific purpose that usually comes as a corollary of some other, often more specific, task. This is a home for such tools.

Testing may be deemed necessary or unnecessary depending on the context and precision required by the tool.

## Tool Descriptions
### inflate.go
This is a very small tool written in Go, which I wrote as an accompaniment to James Coglan's excellent book [Building Git](https://shop.jcoglan.com/building-git/). James uses Ruby in Building Git, and so this is a Go port of a tool he uses in the book. This tool essentially decompresses data which has been compressed into the zlib compression file format. It's useful because it allows us to see how Git stores objects stored under `.git/objects/xx/yyy...`.
#### Usage
```
$ go build inflate.go
$ cat my-project/.git/objects/xx/yyy... | ./inflate
```
