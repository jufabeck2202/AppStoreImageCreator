package core

type DeviceFrame struct {
	Name                              string
	path                              string
	screenshotWidth, screenshotHeight int
	frameWidth, frameHeight           int
	xOffset, YOffset                  int
}
type Frames struct {

}

func (f Frames) get() []DeviceFrame {
	frames:= []DeviceFrame{
		{
			Name:             "iPhone 12 Pro Max",
			path:             "12_Pro_MAX.png",
			screenshotWidth:  1284,
			screenshotHeight: 2778,
			frameWidth:       1484,
			frameHeight:      2972,
			xOffset:          179,
			YOffset:          155,
		},
		{
			Name:             "iPhone 11 Pro Max",
			path:             "11_PRO_MAX.png",
			screenshotWidth:  1242,
			screenshotHeight: 2688,
			frameWidth:       1600,
			frameHeight:      3000,
			xOffset:          179,
			YOffset:          155,
		},
		{
			Name:             "iPhone 11/XR",
			path:             "11.png",
			screenshotWidth:  828,
			screenshotHeight: 1792,
			frameWidth:       1000,
			frameHeight:      2000,
			xOffset:          179,
			YOffset:          155,
		},
		{
			Name:             "iPhone 6/7/8Plus",
			path:             "6_7_8_PLUS.png",
			screenshotWidth:  1080,
			screenshotHeight: 1920,
			frameWidth:       1280,
			frameHeight:      2540,
			xOffset:          100,
			YOffset:          310,
		},
		{
			Name:             "iPhone 11 Pro/X",
			path:             "11_PRO.png",
			screenshotWidth:  1125,
			screenshotHeight: 2436,
			frameWidth:       1600,
			frameHeight:      2800,
			xOffset:          238,
			YOffset:          182,
		},
		{
			Name:             "iPhone 8/SE",
			path:             "8_SE.png",
			screenshotWidth:  750,
			screenshotHeight: 1334,
			frameWidth:       1000,
			frameHeight:      2000,
			xOffset:          125,
			YOffset:          334,
		},
		{
			Name:             "iPhone 12, iPhone 12 Pro",
			path:             "12_12_PRO.png",
			screenshotWidth:  1170,
			screenshotHeight: 2532,
			frameWidth:       1800,
			frameHeight:      3400,
			xOffset:          0,
			YOffset:          0,
		},
		{
			Name:             "iPhone 12 mini",
			path:             "12_MINI.png",
			screenshotWidth:  1080,
			screenshotHeight: 2340,
			frameWidth:       1325,
			frameHeight:      2616,
			xOffset:          0,
			YOffset:          0,
		},
	}
	return frames
}

func (f Frames) GetForSize(width, height int) DeviceFrame {
	for _, v := range f.get() {
		if v.screenshotWidth == width && v.screenshotHeight == height {
			return v
		}
	}
	return DeviceFrame{Name: "unknown"}
}