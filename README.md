# MandelbrotSet
Go package to create image of colorful Mandelbrot set

## Install
> go get github.com/Phaseant/MandelbrotSet@latest

## Output
![Set_1920_1920.png](https://user-images.githubusercontent.com/100575059/216792132-b9f1bf9c-c2ff-4ce6-8737-2ba7a57144f4.png)

## Signature
// Creates PNG image with mandelbrod set. Image filename looks like "filename_width x width.png" <br>
`func CreateMandelbrotImage(width int, filename string)`

## Usage 
```
import "github.com/Phaseant/MandelbrotSet"

func main() {
    mandelbrotSet.CreateMandelbrotImage(1920, "set")
}
```
