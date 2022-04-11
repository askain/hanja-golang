# hanja-golang: 한자-한글 변환 라이브러리


> 주의해주세요.
>
> 이 라이브러리는 한자 데이터베이스로 [suminb님의 hanja table.yml](https://raw.githubusercontent.com/suminb/hanja/develop/hanja/table.yml)파일을 다운로드 해서 사용합니다.
>
> 사용 허락을 안 받았으니 라이센스 취득하시려면 suminb님의 허락을 받으세요.
>
> 그리고 사용 하시다가 빠진 한자 또는 틀린 독음을 발견하시면 [suminb님의 링크](https://docs.google.com/forms/d/e/1FAIpQLScAtw6ylAhy1t0hMn5K25ZbN1vSNPlRdUtebS9PVtKeLQRfvw/viewform)를 통해 제보해주세요.


## 설치

윈도우(파워쉘)
```
$Env:GOPRIVATE="github.com/askain"		<- private repository
go get github.com/askain/hanja-golang
```

리눅스
```
GOPRIVATE="github.com/askain" go get github.com/askain/hanja-golang
```

## 사용법
이 모듈은 다른 기능은 없고 한자를 한글 독음으로 변환하는 기능만 있습니다.

`go test` 를 실행하면 table.yml을 [suminb님의 hanja table.yml](https://raw.githubusercontent.com/suminb/hanja/develop/hanja/table.yml)에서 새로 다운로드 받습니다.

```
translated := hanja.Translate("大韓民國은 民主共和國이다.")	// translated == `대한민국은 민주공화국이다.`
```


<details>
  <summary>재미로 보는 Go와 Python의 속도비교</summary>

Python

```
import hanja
import datetime

if __name__ == '__main__':
    a = datetime.datetime.now()
    for i in range(99):
        translated = hanja.translate('大韓民國은 民主共和國이다.', 'substitution')
    wanted = r'대한민국은 민주공화국이다.'
    
    if translated != wanted:
        print("hanja.translate('大韓民國은 民主共和國이다.') == '{0}', but wanted '{1}'".format(translated, wanted))

    b = datetime.datetime.now()
    print(b-a)
```

Go
```
package hanja

import (
	"testing"
)

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
```


| 언어 | 1 | 2 | 3 | 4 |
|------|---|---|---|---|
|Python3.9|2.050s|2.089s|2.000s|2.024s|
|Go1.17  |0.889s|0.865s|0.434s|0.451s|


> 특이사항
> - Go는 yaml을 파싱하기위해 gopkg.in/yaml.v3 을 사용한 경우 Python을 사용할때보다 오히려 더 느렸음
> - Go는 gopkg.in/yaml.v3 라이브러리를 사용하지않고 text 파일처럼 읽어서 속도 향상: 소요시간 차이는 대충 6.0초 -> 0.1초
> - Go는 왜 3, 4회 실행시 속도가 늘어나는지는 의문

</details>

<br>

## 라이센스

MIT (table.yml 제외)