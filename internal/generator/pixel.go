package generator

import (
	"bytes"
	"fmt"
	"image/color"
	"image/png"
	"math/rand"

	"github.com/zxc7563598/avatar-service/internal/utils"

	"github.com/fogleman/gg"
)

// GeneratePixelAvatar 生成像素块风格头像
func GeneratePixelAvatar(seed string, size int) ([]byte, error) {
	hash := utils.HashToBytes(seed)
	r := rand.New(rand.NewSource(int64(utils.BinaryHash(hash))))
	fmt.Println(r)
	grid := 8
	cell := size / grid
	dc := gg.NewContext(size, size)
	bg := color.RGBA{255, 255, 255, 255}
	dc.SetColor(bg)
	dc.Clear()
	for x := 0; x < grid; x++ {
		for y := 0; y < grid; y++ {
			rv := r.Intn(100)
			if rv < 60 {
				col := color.RGBA{
					R: uint8(100 + r.Intn(155)),
					G: uint8(100 + r.Intn(155)),
					B: uint8(100 + r.Intn(155)),
					A: 255,
				}
				dc.SetColor(col)
				dc.DrawRectangle(float64(x*cell), float64(y*cell), float64(cell), float64(cell))
				dc.Fill()
			}
		}
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, dc.Image()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
