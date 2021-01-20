/**
 * @Author: julianbeck
 * @Description:
 * @File:  frameTask
 * @Version: 1.0.0
 * @Date: 2021/01/18 09:53
 */

package core

import "image"

type addFrameTask struct {
	inputImagePath       string
	hexColor1, hexColor2 string
	heading              string //if not center image
	resizeToOriginal     bool
	image                image.Image
}

type ReturnFrame struct {
	Frame image.Image
	path  string
}


func (t *addFrameTask) hasGradient() bool {
	if t.hexColor1 == "" {
		return false
	} else if t.hexColor2 == "" {
		return false
	} else {
		return true
	}
}
func (t *addFrameTask) hasText() bool {
	if t.heading == "" {
		return false
	} else {
		return true
	}
}
func CreateNewFrameTask(inputImagePath, hexColor1, hexColor2, heading string, resizeToOriginal bool) *addFrameTask {
	t := new(addFrameTask)
	t.inputImagePath = inputImagePath
	t.hexColor1 = hexColor1
	t.hexColor2 = hexColor2
	t.heading = heading
	t.resizeToOriginal = resizeToOriginal
	return t
}
