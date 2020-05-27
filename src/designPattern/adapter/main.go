package main

import "designPattern/adapter/utils"

// 规范化接口：给出一套统一标准规范，所有实现和使用按规范来用
// 接口
func main() {
	media := new(utils.MediaPlayer)
	media.Play("mp3", "最炫民族风.mp3")
	media.Play("mp4", "叶问大战海绵宝宝.mp4")
	media.Play("mp5", "我是你的谁.mp5")
}
