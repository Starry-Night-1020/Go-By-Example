package buffering

import "fmt"

func PrintValue() {
	messages := make(chan string, 2)

	messages <- ("hello")
	messages <- ("channel_buffering")

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
