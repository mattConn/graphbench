BIN=graphbench

$(BIN): interactive/*.go graph/*.go
	go build -o $@

check:
	./$(BIN) -v
clean:
	rm $(BIN)
