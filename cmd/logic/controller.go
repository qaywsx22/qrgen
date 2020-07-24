package logic

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(data string) image.Image {
	var png []byte

	fmt.Println("generateQRCode called with data=" + data)

	png, err := qrcode.Encode(data, qrcode.Medium, 250)
	if err != nil {
		return nil
	}

	img := convertBytesToGoImage(png)
	return img
}

func convertBytesToGoImage(imgByte []byte) image.Image {
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return img
}
