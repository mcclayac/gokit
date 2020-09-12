package main

// StringService provides operations on strings.
import "context"

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}


