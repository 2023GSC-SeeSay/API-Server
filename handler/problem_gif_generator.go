package handler

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"log"
	"os"
)

type GifInfo struct {
	path         string
	durationTime int
}

func GenerateGIF(pathes []GifInfo, title string) string {
	// get the image path list and generate gif
	// TODO : get image from firebase
	// fixme : get image from firebase
	// print the pwd
	os_info := os.Getenv("OS")
	fmt.Printf("%s", os_info)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// add the image path to the pwd
	a := make([]GifInfo, len(pathes))
	for i, path := range pathes {
		a[i] = GifInfo{
			path:         dir + path.path,
			durationTime: path.durationTime,
		}
	}
	// fmt.Printf("%v", a)
	// create a new  GIF with a delay of 100ms between frames
	var images []*image.Paletted
	var delays []int

	// Extract the dominant colors from the image.
	palette := color.Palette{
		color.RGBA{0, 0, 0, 0},
		color.RGBA{255, 255, 255, 255},
		color.RGBA{200, 146, 112, 255},
		color.RGBA{105, 56, 52, 255},
		color.RGBA{191, 67, 63, 255},
		color.RGBA{152, 87, 83, 255},
		color.RGBA{11, 5, 2, 255},
		color.RGBA{103, 92, 86, 255},
		color.RGBA{65, 51, 42, 255},
		color.RGBA{127, 64, 62, 255},
		color.RGBA{74, 40, 29, 255},
		color.RGBA{183, 78, 75, 255},
		color.RGBA{189, 142, 112, 255},
		color.RGBA{192, 143, 111, 255},
		color.RGBA{181, 64, 59, 255},
		color.RGBA{181, 90, 86, 255},
		color.RGBA{95, 55, 55, 255},
		color.RGBA{179, 90, 89, 255},
		color.RGBA{174, 178, 177, 255},
		color.RGBA{168, 145, 136, 255},
		color.RGBA{185, 177, 167, 255},
		color.RGBA{159, 85, 82, 255},
		color.RGBA{154, 85, 80, 255},
		color.RGBA{87, 61, 58, 255},
		color.RGBA{197, 75, 71, 255},
		color.RGBA{163, 81, 78, 255},
		color.RGBA{107, 58, 53, 255},
		color.RGBA{99, 51, 49, 255},
		color.RGBA{166, 117, 101, 255},
		color.RGBA{52, 42, 33, 255},
		color.RGBA{38, 20, 19, 255},
		color.RGBA{132, 121, 121, 255},
		color.RGBA{191, 141, 108, 255},
		color.RGBA{181, 140, 115, 255},
		color.RGBA{136, 77, 66, 255},
		color.RGBA{173, 107, 102, 255},
		color.RGBA{195, 149, 116, 255},
		color.RGBA{142, 78, 77, 255},
		color.RGBA{176, 70, 64, 255},
		color.RGBA{173, 125, 103, 255},
		color.RGBA{189, 143, 110, 255},
		color.RGBA{160, 125, 103, 255},
		color.RGBA{144, 116, 95, 255},
		color.RGBA{83, 65, 53, 255},
		color.RGBA{7, 1, 1, 255},
		color.RGBA{154, 126, 105, 255},
		color.RGBA{159, 118, 98, 255},
		color.RGBA{144, 116, 95, 255},
		color.RGBA{137, 108, 85, 255},
		color.RGBA{189, 143, 116, 255},
		color.RGBA{156, 120, 96, 255},
		color.RGBA{154, 126, 105, 255},
		color.RGBA{154, 118, 94, 255},
		color.RGBA{62, 34, 33, 255},
	}

	for _, infos := range a {
		// read the image file

		f, err := os.Open(infos.path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// decode the image based on its file extension
		img, err := png.Decode(f)
		if err != nil {
			log.Printf("error decoding %s: %v", infos.path, err)
			continue
		}

		// convert the image to a paletted image
		palettedImg := image.NewPaletted(img.Bounds(), palette)
		draw.Draw(palettedImg, palettedImg.Bounds(), img, image.ZP, draw.Src)

		images = append(images, palettedImg)
		delays = append(delays, infos.durationTime)
	}

	// save the GIF to a file
	gif_path := fmt.Sprintf("%s/assets/video/%s.gif", dir, title)
	f, err := os.Create(gif_path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	if err != nil {
		log.Fatal(err)
	}
	// TODO : generate gif
	// TODO : save gif to firebase
	return gif_path
}

func TranslateTextToMouthPath(text string) []GifInfo {
	paths := []GifInfo{}

	for _, char := range text {
		pchar := mouth_seperate[char]
		// fmt.Printf("%v", pchar)
		if len(pchar) == 3 {
			paths = append(paths, GifInfo{path: mouth_image_windows[string(pchar[0:3])], durationTime: 100})
		} else {
			// fmt.Printf("%v", pchar[0:3])
			paths = append(paths, GifInfo{path: mouth_image_windows[string(pchar[0:3])], durationTime: 34})
			paths = append(paths, GifInfo{path: mouth_image_windows[string(pchar[3:6])], durationTime: 66})
		}
	}
	// fmt.Printf("%v, %v", paths[0].durationTime, paths[0].path)
	return paths
}

func TranslateTextToTonguePath(text string) []GifInfo {
	paths := []GifInfo{}
	fmt.Printf("%v", paths)
	for _, char := range text {
		paths = append(paths, GifInfo{path: tongue_image_windows[string(char)], durationTime: 100})
	}
	// fmt.Printf("%v", paths)
	return paths
}
