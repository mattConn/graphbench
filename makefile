BIN=graphbench

$(BIN): interactive/*.go graph/*.go
	go build -o $@

check:
	./$(BIN)
clean:
	rm $(BIN)
