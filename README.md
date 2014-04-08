## mimeTeX for Go

This is a Go package that allows rendering LaTeX documents to images directly, without any external dependencies and toolchains. That is implemented as native cgo binding to slightly modified [mimeTeX] codebase and uses its C API.
[mimeTeX]: http://www.forkosh.com/mimetex.html

#### I. Usage examples

```Go
package main

import "gopkg.in/mimetex.v1"

func main() {
	expr := `x*y`
	img, err := mimetex.RenderImage(expr, *size)
	if err != nil {
		log.Fatalln("mimetex error: " + err.Error())
	}
	// do whatever with image.Image
}
```

Also one might test with **example/main.go**:

	$ go build
	$ ./example -s=4 'P = \bigoplus_{k,t \in \bar{\mathbb{Z}}} s(k,t)\gamma^k\delta^t'
	$ ls -la out.gif
	
![rendered-expr](http://cl.ly/UtBr/out.gif)

**DISCLAIMER:** Do not call this rendering "crappy". Pay some respect to the author of the mimeTeX, it took 10 years to write 10k SLOC in C and 10k lines of comments in total for each of them. See the source [mimetex.c](http://cl.ly/UtGw). I personally hate his approach, but the work is purely titanic.

#### II. Benchmark

	$ go test -bench=.
	PASS
	BenchmarkRender	    5000	    776704 ns/op
	ok  	gopkg.in/mimetex.v1	3.997s
	
`0.77ms` to render `\mathcal{M}^{ax}_{in}\llbracket\gamma, \delta\rrbracket` on Intel i5. Also the package overhead is only `1.7MB`.

#### III. Further work

Implementing LaTeX rendering in pure language is always extremely hard. Take look at Java's pure implementation — [JLaTeXMath](http://forge.scilab.org/index.php/p/jlatexmath/source/tree/master/). For my needs this simple binding to mimeTeX is OK, at least until someone go crazy an this with Go.

#### IV. Licenses

* mimeTeX: GNU GPL
* mimetex-go: [MIT](http://xlab.mit-license.org)