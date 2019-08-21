package panic

import "os"

func PrintValue() {
	// panic("a problem")

	_, err := os.Create("/temp/file")
	if err != nil {
		// panic(err)
	}
}
