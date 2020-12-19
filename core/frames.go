package core

type DeviceFrame struct {
	name string
	path string
	screenshotWidth, screenshotHeight int
	frameWidth, frameHeight int
	xOffset, YOffset int
}
type Frames struct {

}
func (v Frames) get() []DeviceFrame {
	frames:= []DeviceFrame{
		{
			name:             "iPhone 11 Pro Max",
			path:             "",
			screenshotWidth:  0,
			screenshotHeight: 0,
			frameWidth:       0,
			frameHeight:      0,
			xOffset:          0,
			YOffset:          0,
		},

	}
	return frames
}