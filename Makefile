src=main.go
bin=crashdb

$(bin): $(src)
	go build -o $(bin)

.PHONY: sadness
sadness: $(bin)
	./sadness.sh

.PHONY: clean
clean:
	rm -f $(bin)
