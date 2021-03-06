package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func saveFile(src io.Reader) {
	f, err := os.Create("save.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(f, src)
}

func loadFileToConfig(path string) {
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	//設定をデコード
	config, format, err := image.DecodeConfig(f)
	if err != nil {
		log.Fatal(err)
	}

	//フォーマット名表示
	fmt.Println("画像フォーマット：" + format)
	//サイズ表示
	fmt.Println("横幅=" + strconv.Itoa(config.Width) + ", 縦幅=" + strconv.Itoa(config.Height))
}

func loadFileToImage(path string) {
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	i, s, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	f, err = os.Create("img.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, i, nil)
}

func main() {
	var url string = "http://placekitten.com/g/640/340"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	saveFile(response.Body)

	loadFileToConfig("save.jpg")

	loadFileToImage("save.jpg")
}
