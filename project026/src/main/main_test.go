package main

import (
	"fmt"
	"testing"
)



func TestPow(t *testing.T)  {
	res := pow(6)
	if res == 216{
		fmt.Println("yep")
	}else {
		fmt.Println("nope")
	}
}

func TestAdd(t *testing.T)  {
	res := add(6)
	if res == 9{
		fmt.Println("yep")
	}else {
		fmt.Println("nope")
	}
}
