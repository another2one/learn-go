package utils

import "fmt"

type MediaPlayerInter interface {
	Play(fileName string)
}

type MediaPlayer struct{}

func (mediaPlayer MediaPlayer) Play(ext, fileName string) {
	var media MediaPlayerInter
	switch ext {
	case "mp3":
		media = Mp3{}
		media.Play(fileName)
	case "mp4":
		media = Mp4{}
		media.Play(fileName)
	default:
		fmt.Println("暂不支持: ", ext)
	}
}

type Mp3 struct{}

func (mp3 Mp3) Play(fileName string) {
	fmt.Println("mp3 playing: ", fileName)
}

type Mp4 struct{}

func (mp4 Mp4) Play(fileName string) {
	fmt.Println("mp4 playing: ", fileName)
}
