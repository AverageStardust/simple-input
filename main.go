package main

func init() {
	for key := range RawTerminalKeys() {
		println(key)
	}
}
