package hanja

import (
	"testing"
)

func TestDownloadHanjaHangul(t *testing.T) {
	DownloadHanjaHangul()
}

func TestTranslate(t *testing.T) {
	var translated string
	for i := 0; i < 100; i++ {
		translated = Translate("大韓民國은 民主共和國이다.")
	}
	wanted := `대한민국은 민주공화국이다.`

	if translated != wanted {
		t.Fatalf(`Translate("大韓民國은 民主共和國이다.") == %q, but wanted "%v",`, translated, wanted)
	}
}
