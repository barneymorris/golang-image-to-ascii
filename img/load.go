package img

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
)

func Load() (image.Image, error) {
	fmt.Printf("Please url to the image: ")

	var url string

	_, err := fmt.Fscan(os.Stdin, &url)

	if err != nil {
		panic(err)
	}

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	img, _, err := image.Decode(response.Body)

	if err != nil {
		panic(err)
	}

	if err != nil {
		return nil, err
	}

	return img, nil
}
