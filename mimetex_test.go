package mimetex

import "testing"

var expr string = `\mathcal{M}^{ax}_{in}\llbracket\gamma, \delta\rrbracket`

// I'm not sure if mimeTeX renders images hardware-independently,
// so assume everythining is OK, otherwise test manually with example/main.go
func TestRender(t *testing.T) {
	data, err := RenderGifBytes(expr, 5)
	if err != nil {
		t.Fatalf("mimtex error: %s", err.Error())
	}
	if len(data) < 1 {
		t.Fatalf("empty results")
	}
}

func BenchmarkRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RenderGifBytes(expr, 5)
	}
}
