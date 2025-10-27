package generator

import (
	"bytes"
	"image/color"
	"image/png"

	"github.com/zxc7563598/avatar-service/internal/utils"

	"github.com/fogleman/gg"
)

// GenerateIdenticon 根据 seed 生成一致性 identicon 风格头像
func GenerateIdenticon(seed string, size int) ([]byte, error) {
	hash := utils.HashToBytes(seed)
	margin := 25
	grid := 7
	usableSize := size - 2*margin
	cellSize := usableSize / grid
	dc := gg.NewContext(size, size)
	bg := color.RGBA{255, 255, 255, 255}
	dc.SetColor(bg)
	dc.Clear()
	// 前景色基于 hash
	fg := color.RGBA{
		R: hash[0],
		G: hash[1],
		B: hash[2],
		A: 255,
	}
	for x := 0; x < grid/2+1; x++ {
		for y := 0; y < grid; y++ {
			idx := x*grid + y
			if hash[idx]%2 == 0 {
				dc.SetColor(fg)
				dc.DrawRectangle(float64(margin+x*cellSize), float64(margin+y*cellSize), float64(cellSize), float64(cellSize))
				dc.Fill()
				dc.DrawRectangle(float64(margin+(grid-1-x)*cellSize), float64(margin+y*cellSize), float64(cellSize), float64(cellSize))
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
