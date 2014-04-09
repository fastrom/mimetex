// Package mimeTeX provides functions to render LaTeX documents
// into images instantly and without any external dependencies.
package mimetex

import (
	"bytes"
	"image"
	"image/gif"

	"gopkg.in/mimetex.v1/bridge"
)

// RenderImage produces an image.Image of rendered LaTeX expression expr,
// the scale of rendered expression is defined by size argument.
func RenderImage(expr string, size int) (image.Image, error) {
	data, err := bridge.RenderExprToGif(expr, size)
	if err != nil {
		return nil, err
	}

	img, err := gif.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return img, nil
}

// RenderGifBytes produces a byte slice containing encoded GIF image
// of LaTeX expression expr, the scale of rendered expression
// is defined by size argument.
func RenderGifBytes(expr string, size int) ([]byte, error) {
	data, err := bridge.RenderExprToGif(expr, size)
	if err != nil {
		return nil, err
	}
	return data, nil
}
