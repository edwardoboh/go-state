package main

func main() {
	/*
		We user a wait group here, from package sync to wait for the completion
		of tasks been ran by several go routines
	*/
	wg()

	/*
		Here, we show how to send and receive data between several worker pools using channels in go
	*/
	wp()

	/*
		This function makes us of mutexes this time instead of an atomic counter when multiple go routines need to update a single value
	*/
	minMut()
}
