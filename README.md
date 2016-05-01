# Average Hash

A simple image hashing library written in Go.
 
## Example use:

```go
package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/atijust/ahash"
)

func main() {
	img1, err := load("test01.png")
	if err != nil {
		log.Fatal(err)
	}

	img2, err := load("test02.png")
	if err != nil {
		log.Fatal(err)
	}

	hash1 := ahash.Generate(img1)
	hash2 := ahash.Generate(img2)
	dist := ahash.Distance(hash1, hash2)

	log.Printf("hash1: 0x%x, hash2: 0x%x, hamming distance: %d\n", hash1, hash2, dist)
}

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
```