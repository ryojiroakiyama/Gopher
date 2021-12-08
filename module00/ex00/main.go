/*
** package main
 */
package main

import (
	"bytes"
	"fmt"
	//"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "invalid argument")
		return
	}
	fmt.Println("argument: ", os.Args[1])
	files, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for i, file := range files {
		fmt.Println(file.Name())
		fmt.Println(i)
	}
	//fmt.Println(files[0].Name())
	////file, err := os.OpenFile(string(files[0].Name()), os.O_WRONLY, 0644)
	////file, err := os.OpenFile("./test", os.O_WRONLY, 0644)
	//file, err := os.Open(files[0].Name())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//b, err := ioutil.ReadAll(file)
	//// 出力
	//fmt.Println("inside:", string(b))
	//// ファイルオブジェクトを画像オブジェクトに変換
	//img, _, err := image.Decode(file)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// 出力ファイルを生成
	//out, err := os.Create("tmp")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer out.Close()

	//// 画像ファイル出力
	////    jpeg.Encode(out, img, nil)
	//png.Encode(out, img)
}

func ToPng(imageBytes []byte) ([]byte, error) {
	contentType := http.DetectContentType(imageBytes)

	switch contentType {
	case "image/png":
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(imageBytes))
		if err != nil {
			return nil, fmt.Errorf("%v: unable to decode jpeg", err)
		}

		buf := new(bytes.Buffer)
		if err := png.Encode(buf, img); err != nil {
			return nil, fmt.Errorf("%v: unable to encode png", err)
		}
		return buf.Bytes(), nil
	}
	return nil, fmt.Errorf("unable to convert %#v to png", contentType)
}

//func main() {
//	const width, height = 256, 256

//	// Create a colored image of the given width and height.
//	img := image.NewNRGBA(image.Rect(0, 0, width, height))

//	for y := 0; y < height; y++ {
//		for x := 0; x < width; x++ {
//			img.Set(x, y, color.NRGBA{
//				R: uint8((x + y) & 255),
//				G: uint8((x + y) << 1 & 255),
//				B: uint8((x + y) << 2 & 255),
//				A: 255,
//			})
//		}
//	}

//	f, err := os.Create("image.png")
//	if err != nil {
//		log.Fatal(err)
//	}

//	if err := png.Encode(f, img); err != nil {
//		f.Close()
//		log.Fatal(err)
//	}

//	if err := f.Close(); err != nil {
//		log.Fatal(err)
//	}
//}
