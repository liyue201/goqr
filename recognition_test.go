package goqr

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"

	"testing"
)

func recognizeFile(path string) (string, error) {
	imgdata, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	img, _, err := image.Decode(bytes.NewReader(imgdata))
	if err != nil {
		return "", err
	}
	qrCodes, err := Recognize(img)
	if err != nil {
		return "", nil
	}
	return string(qrCodes[0].Payload), nil
}

func TestRecognize(t *testing.T) {
	testCases := []struct {
		filePath string
		text     string
	}{
		{
			"example/testdata/000.jpg",
			"http://www.amazon.co.jp/gp/aw/rd.html?uid=NULLGWDOCOMO&url=/gp/aw/h.html&at=aw_intl6-22",
		},
		{
			"example/testdata/001.jpg",
			"http://www.amazon.co.jp/gp/aw/rd.html?uid=NULLGWDOCOMO&url=/gp/aw/h.html&at=aw_intl6-22",
		},
		{
			"example/testdata/002.jpg",
			"http://www.amazon.co.jp/gp/aw/rd.html?uid=NULLGWDOCOMO&url=/gp/aw/h.html&at=aw_intl6-22",
		},
		{
			"example/testdata/003.jpg",
			"http://www.amazon.co.jp/gp/aw/rd.html?uid=NULLGWDOCOMO&url=/gp/aw/h.html&at=aw_intl6-22",
		},
		{
			"example/testdata/004.png",
			"http://swtch.com/pjw/#523892624657510299353520120480795433576563223876200460867159368633143833417166162086873959805500633263545263286786346759633071266643358952888263169163143415896058956186071276000133287120333224396223386892286076080898690020480143143263415796162046552639449120143662",
		},
		{
			"example/testdata/005.png",
			"https://github.com/",
		},
		{
			"example/testdata/006.png",
			"https://github.com/liyue201",
		},
		{
			"example/testdata/007.png",
			"https://github.com/",
		},
		{
			"example/testdata/008.png",
			"https://github.com/liyue201/goqr",
		},
	}
	for _, testCase := range testCases {
		text, err := recognizeFile(testCase.filePath)
		assert.Nil(t, err)
		assert.Equal(t, testCase.text, text)
	}
}
