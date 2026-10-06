// Gera animaçoes GIF de figuras de Lissjous aleatórias
// go build .\exer-1.5.go
// .\exer-1.5.exe >out.gif

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand/v2"
	"os"
)

var pallete = []color.Color{
	color.Black,                                // Índice 0: Fundo
	color.RGBA{R: 0, G: 255, B: 0, A: 255},     // Índice 1: Verde
	color.RGBA{R: 255, G: 0, B: 0, A: 255},     // Índice 2: Vermelho
	color.RGBA{R: 0, G: 0, B: 255, A: 255},     // Índice 3: Azul
	color.RGBA{R: 255, G: 255, B: 255, A: 255}, // Índice 4: Branco
	color.RGBA{R: 255, G: 255, B: 0, A: 255},   // Índice 5: Amarelo
}

const (
	whiteIndex = 0 // primeira cor da paleta
	blackIndex = 1 // próxima cor da paleta
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Número de revoluções completas do oscilaor x
		res     = 0.001 // resolução angular
		size    = 100   // canvas da imagem cobre de [-size..+size]
		nframes = 64    // número de quadros da animação
		delay   = 8     // tempo entre quadros em unidades de 10ms
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// LÓGICA INTERESSANTE:
			// Mudamos a cor baseada no tempo 't' ou no frame 'i'
			// O operador % (módulo) garante que fiquemos dentro do range da paleta (1 a 5)
			// colorIndex := uint8((int(t*10) % (len(pallete) - 1)) + 1)
			colorIndex := uint8((i/10)%5 + 1) // a cor mude gradualmente a cada 10 frames
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Nota: Ignorando erros de codificação
}
