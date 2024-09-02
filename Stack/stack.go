package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(data string) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}

	s.items = s.items[:len(s.items)-1]
}
func (s *Stack) Top() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("Stack is empty")
	}

	return s.items[len(s.items)-1], nil
}
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// filepath.Clean(path) уже будет работать
func simplifyPath(path string) string {
	res := ""
	stackForTask := Stack{}
	pathWithoutSlash := strings.Split(path, "/")
	for _, elem := range pathWithoutSlash {
		if elem == "." {
			continue
		} else if elem == ".." {
			stackForTask.Pop()
			continue
		}
		stackForTask.Push(elem)
	}
	for !stackForTask.IsEmpty() {
		elem, err := stackForTask.Top()
		stackForTask.Pop()
		if err != nil {
			return ""
		}
		if elem != "" {
			res = "/" + elem + res
		}

	}
	if res == "" {
		res += "/"
	}
	return res
}

func main() {
	paths := []string{
		"/foo/bar/baz/../gar/././",
		"/bar/../../../../../../../",
		"/..",
		"/.././///..",
	}
	for _, path := range paths {
		fmt.Printf("Original: %s, Simplified: %s\n", path, simplifyPath(path))
	}
}
