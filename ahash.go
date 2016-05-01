package ahash

import (
	"image"
	"image/color"

	"github.com/nfnt/resize"
)

// 8ビットグレースケール
func Generate(img image.Image) uint64 {
	reszImg := resize.Resize(8, 8, img, resize.NearestNeighbor)
	grayImg := toGray(reszImg)

	sum := 0
	for _, v := range grayImg.Pix {
		sum += int(v)
	}
	avg := byte(sum / len(grayImg.Pix))

	var hash uint64
	for i := 0; i < 64; i++ {
		hash <<= 1
		if avg <= grayImg.Pix[i] {
			hash |= 1
		}
	}

	return hash
}

func toGray(src image.Image) *image.Gray {
	rct := src.Bounds()
	dst := image.NewGray(rct)
	for x := rct.Min.X; x < rct.Max.X; x++ {
		for y := rct.Min.Y; y < rct.Max.Y; y++ {
			dst.Set(x, y, color.GrayModel.Convert(src.At(x, y)))
		}
	}
	return dst
}

// 16ビットグレースケール
func Generate16(img image.Image) uint64 {
	reszImg := resize.Resize(8, 8, img, resize.NearestNeighbor)
	grayImg := toGray16(reszImg)

	sum := 0
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			sum += int(grayImg.Gray16At(x, y).Y)
		}
	}
	avg := uint16(sum / 64)

	var hash uint64
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			hash <<= 1
			if avg <= grayImg.Gray16At(x, y).Y {
				hash |= 1
			}
		}
	}

	return hash
}

func toGray16(src image.Image) *image.Gray16 {
	rct := src.Bounds()
	dst := image.NewGray16(rct)
	for x := rct.Min.X; x < rct.Max.X; x++ {
		for y := rct.Min.Y; y < rct.Max.Y; y++ {
			dst.Set(x, y, color.Gray16Model.Convert(src.At(x, y)))
		}
	}
	return dst
}
