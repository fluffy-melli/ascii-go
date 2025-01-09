package ascii

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"golang.org/x/image/draw"
)

var ASCII_CHARS = []string{
	"$", "@", "B", "%", "8", "&", "W", "M", "#", "*", "o", "a", "h", "k", "b", "d",
	"p", "q", "w", "m", "Z", "O", "0", "Q", "L", "C", "J", "U", "Y", "X", "z", "c",
	"v", "u", "n", "x", "r", "j", "f", "t", "/", "\\", "|", "(", ")", "1", "{", "}",
	"[", "]", "?", "-", "_", "+", "~", "<", ">", "i", "!", "l", "I", ";", ":", ",",
	"\"", "^", "`", "'", ".", " ",
}

func ReadImage(filepath string) (image.Image, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	image, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return image, nil
}

func Brightness(c color.Color) int {
	r, g, b, a := c.RGBA()
	r = r >> 8
	g = g >> 8
	b = b >> 8
	a = a >> 8
	if a == 0 {
		return 0
	}
	brightness := (r + g + b) / 3
	brightness = brightness * uint32(a) / 255
	return int(brightness)
}

type Ascii [][]string

func (a *Ascii) ToStr() string {
	respond := ""
	for _, y := range *a {
		for _, x := range y {
			respond += x
		}
		respond += "\n"
	}
	return respond
}

func (a *Ascii) ToArray() [][]string {
	return *a
}

func Render(img image.Image, w int) Ascii {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	ratio := float32(height / width)
	height = int(float32(w) * ratio * 0.65)
	resized_image := image.NewRGBA(image.Rect(0, 0, w, height))
	draw.NearestNeighbor.Scale(resized_image, image.Rectangle{Min: bounds.Min, Max: image.Point{X: w, Y: height}}, img, img.Bounds(), draw.Over, nil)
	respond := make(Ascii, 0)
	for y := 0; y < height; y++ {
		ascii_char := make([]string, 0)
		for x := 0; x < w; x++ {
			pixel := resized_image.At(x, y)
			brightness := Brightness(pixel)
			if brightness == 0 {
				ascii_char = append(ascii_char, " ")
			} else {
				index := brightness * (len(ASCII_CHARS) - 1) / 255
				ascii_char = append(ascii_char, ASCII_CHARS[index])
			}
		}
		respond = append(respond, ascii_char)
	}
	return respond
}
