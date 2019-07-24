package main

import "fmt"

func main() {
	fmt.Println(`State machines usually create structs with methods IsValid(), Action() and Next(), so you can run them in a loop like this:
while state.Valid() == true {
    state.Action()
    state = state.Next()
}

Instead of creating a interactive flow like this, I created a lot of tests in states module and a flow test in this folder, so we can see the progres.

I suggest to run make in root folder and look for the logs :) `)
}
