// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		draw(w)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func draw(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height-1; py += 2 {
		y := []float64{
			float64(py)/height*(ymax-ymin) + ymin,
			float64(py+1)/height*(ymax-ymin) + ymin,
		}

		for px := 0; px < width-1; px += 2 {
			x := []float64{
				float64(px)/width*(xmax-xmin) + xmin,
				float64(px+1)/width*(xmax-xmin) + xmin,
			}

			m := []color.Color{
				mandelbrot(complex(x[0], y[0])),
				mandelbrot(complex(x[0], y[1])),
				mandelbrot(complex(x[1], y[0])),
				mandelbrot(complex(x[1], y[1])),
			}

			R, G, B, A := make([]uint32, 4), make([]uint32, 4), make([]uint32, 4), make([]uint32, 4)
			for i := 0; i < 4; i++ {
				R[i], G[i], B[i], A[i] = m[i].RGBA()
			}

			avg := []uint32{
				(R[0] + R[1] + R[2] + R[3]) / 4,
				(G[0] + G[1] + G[2] + G[3]) / 4,
				(B[0] + B[1] + B[2] + B[3]) / 4,
				(A[0] + A[1] + A[2] + A[3]) / 4,
			}

			avgColor := color.RGBA{uint8(avg[0]), uint8(avg[1]), uint8(avg[2]), uint8(avg[3])}

			// Image point (px, py) represents complex value z.
			img.Set(px, py, avgColor)
			img.Set(px, py+1, avgColor)
			img.Set(px+1, py, avgColor)
			img.Set(px+1, py+1, avgColor)
		}
	}

	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z

		if cmplx.Abs(v) > 2 {
			return color.RGBA{0 + contrast*n, 255 - contrast*n, 255 - contrast*n, 255}
		}
	}

	return color.Black
}
