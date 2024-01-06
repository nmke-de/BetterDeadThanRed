package main

func unwrap(err error) {
	if err != nil {
		panic(err)
	}
}
