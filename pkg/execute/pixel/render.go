package pixel

import (
	"fmt"
	"image"
	"math"
	"sync"
	"time"

	fpixel "github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/ysanson/AdventOfCode/pkg/twod"
	"golang.org/x/image/colornames"
)

var (
	started   = false
	picUpdate = make(chan fpixel.Picture)
	fps       = (<-chan time.Time)(nil)

	lock = sync.Mutex{}
)

func SetFPS(rate int) {
	lock.Lock()
	if rate == 0 {
		fps = nil
	} else {
		fps = time.Tick(time.Second / time.Duration(rate))
	}
	lock.Unlock()
}

func render(m twod.Map) {
	if !started {
		go mainLoop()
		started = true
	}
	picUpdate <- loadPicture(m)
}

func mainLoop() {
	maxX, maxY := pixelgl.PrimaryMonitor().Size()
	cfg := pixelgl.WindowConfig{
		Title:  "Advent of code!",
		Bounds: fpixel.R(0, 0, maxX, maxY),
		//Monitor:   pixelgl.PrimaryMonitor(),
		//Undecorated: true,
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.White)

	camPos := fpixel.ZV
	camSpeed := 500.0
	camZoom := 10.0
	camZoomSpeed := 1.2
	frames := 0
	second := time.Tick(time.Second)
	lastTime := time.Now()
	lastPic := fpixel.Picture(nil)

	for !win.Closed() {
		dt := time.Since(lastTime).Seconds()
		lastTime = time.Now()
		if win.Pressed(pixelgl.KeyLeft) {
			camPos.X -= camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyRight) {
			camPos.X += camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyDown) {
			camPos.Y -= camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyUp) {
			camPos.Y += camSpeed * dt
		}
		camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)
		if win.MouseScroll().Y != 0 {
			drawPic(win, lastPic, camZoom, camPos)
		}

		frames++
		select {
		case pic := <-picUpdate:
			drawPic(win, pic, camZoom, camPos)
			lastPic = pic
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
		win.Update()
		lock.Lock()
		if fps != nil {
			<-fps
		}
		lock.Unlock()
	}
}

func drawPic(win *pixelgl.Window, pic fpixel.Picture, camZoom float64, camPos fpixel.Vec) {
	if pic == nil {
		return
	}
	win.Clear(colornames.White)
	sprite := fpixel.NewSprite(pic, pic.Bounds())

	mat := fpixel.IM
	mat = mat.Scaled(fpixel.ZV, camZoom)
	mat = mat.Moved(win.Bounds().Center().Sub(camPos))
	sprite.Draw(win, mat)
}

func rectangle(m twod.Map) image.Rectangle {
	botL := m.FindBottomLeft()
	topR := m.FindTopRight()

	upLeft := image.Point{botL.X(), botL.Y()}
	lowRight := image.Point{topR.X() + 1, topR.Y() + 1}
	return image.Rectangle{upLeft, lowRight}
}

func loadPicture(m twod.Map) fpixel.Picture {
	img := image.NewRGBA(rectangle(m))
	for k, v := range m {
		matchingColor, exist := twod.RenderingMap[v]
		if exist {
			img.Set(k.X(), k.Y(), matchingColor)
		}
	}
	return fpixel.PictureDataFromImage(img)
}
