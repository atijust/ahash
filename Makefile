.PHONY: test bench

test:
	go vet `glide novendor`
	go test -race -v `glide novendor`

bench:
	go test -bench=. -benchmem `glide novendor`
