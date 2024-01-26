build: BetterDeadThanRed BetterDeadThanRed.exe BetterDeadThanRed.wasm

BetterDeadThanRed:
	GOOS=linux garble -tiny build

BetterDeadThanRed.exe:
	GOOS=windows garble -tiny build

BetterDeadThanRed.wasm:
	GOOS=js GOARCH=wasm garble -tiny build -o BetterDeadThanRed.wasm

clean:
	rm -f BetterDeadThanRed*

.PHONY: build clean
