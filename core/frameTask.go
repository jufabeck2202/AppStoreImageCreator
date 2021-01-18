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
	outputImagePath      string
	hexColor1, hexColor2 string
	heading              string //if not center image
	resizeToOriginal     bool
	hasFrame             bool
	mergeIntoSingleImage bool
	image                image.Image
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
func CreateNewFrameTask(inputImagePath, outputImagePath, hexColor1, hexColor2, heading string, resizeToOriginal, hasFrame, mergeIntoSingleImage bool) *addFrameTask {
	t := new(addFrameTask)
	t.inputImagePath = inputImagePath
	t.outputImagePath = outputImagePath
	t.hexColor1 = hexColor1
	t.hexColor2 = hexColor2
	t.heading = heading
	t.resizeToOriginal = resizeToOriginal
	t.hasFrame = hasFrame
	t.mergeIntoSingleImage = mergeIntoSingleImage
	return t
}
