package src

import "fmt"

func GetHelloFrom(name string) string {
	return fmt.Sprintf("Hello %s! How are you?", name)
}
