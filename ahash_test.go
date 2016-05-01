package ahash

import (
	"image"
	_ "image/png"
	"os"
	"testing"
)

func load(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func TestGenerate(t *testing.T) {
	img, err := load("./testdata/gophercolor.png")
	if err != nil {
		t.Fatal(err)
	}

	if hash := Generate(img); hash != 0x1c7c7c3e3e3e1e18 {
		t.Errorf("ハッシュが正しくない; 0x%x", hash)
	}
}

func TestGenerate16(t *testing.T) {
	img, err := load("./testdata/gophercolor.png")
	if err != nil {
		t.Fatal(err)
	}

	if hash := Generate16(img); hash != 0x607cfffefe1e00 {
		t.Errorf("ハッシュが正しくない; 0x%x", hash)
	}
}

func BenchmarkGenerate(b *testing.B) {
	img, err := load("./testdata/gophercolor.png")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Generate(img)
	}
}

func BenchmarkGenerate16(b *testing.B) {
	img, err := load("./testdata/gophercolor.png")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Generate16(img)
	}
}
