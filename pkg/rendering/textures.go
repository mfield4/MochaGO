package rendering

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func NewTextureD(file string) (uint32, error) {
	return NewTexture(gl.REPEAT, gl.REPEAT, gl.LINEAR, gl.LINEAR, file)
}

func NewTexture(wrapS, wrapT, minF, magF int32, file string) (uint32, error) {
	rgba, _ := imageToPixelData(file)

	var texture uint32
	gl.GenTextures(1, &texture)
	// gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, wrapS)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, wrapT)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, minF)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, magF)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	gl.BindTexture(gl.TEXTURE_2D, 0)

	return texture, nil
}

func NewCubeMapD(faces [6]string) (uint32, error) {
	return NewCubeMap(gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE, gl.LINEAR, gl.LINEAR, faces)
}

func NewCubeMap(wrapS, wrapT, wrapR, minF, magF int32, faces [6]string) (cubeMapId uint32, err error) {
	gl.GenTextures(1, &cubeMapId)
	gl.BindTexture(gl.TEXTURE_CUBE_MAP, cubeMapId)

	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_S, wrapS)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_T, wrapT)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_R, wrapR)

	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MIN_FILTER, minF)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MAG_FILTER, magF)

	for idx, f := range faces {
		rgba, err := imageToPixelData(f)
		if err != nil {
			log.Fatalf("Failed with error: %v\n", err)
			return 0, err
		}

		gl.TexImage2D(
			uint32(gl.TEXTURE_CUBE_MAP_POSITIVE_X+idx),
			0,
			gl.RGB,
			int32(rgba.Rect.Size().X),
			int32(rgba.Rect.Size().Y),
			0,
			gl.RGBA,
			gl.UNSIGNED_BYTE,
			gl.Ptr(rgba.Pix))
	}

	gl.BindTexture(gl.TEXTURE_CUBE_MAP, 0)

	return cubeMapId, err
}

func imageToPixelData(file string) (*image.RGBA, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("texture %q not found on disk: %v", file, err)
	}
	defer func(imgFile *os.File) {
		err := imgFile.Close()
		if err != nil {

		}
	}(imgFile)

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride")
	}

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	return rgba, nil
}
