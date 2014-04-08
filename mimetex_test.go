package mimetex

import "testing"

// I'm not sure if mimeTeX renders images hardware-independently,
// so assume everythining is OK, otherwise head to example/main.go

// But we'll do a benchmark.
func BenchmarkRender(b *testing.B) {
	expr := `\mathcal{M}^{ax}_{in}\llbracket\gamma, \delta\rrbracket`
	for i := 0; i < b.N; i++ {
		data, _ := RenderGifBytes(expr, 5)
	}
}
