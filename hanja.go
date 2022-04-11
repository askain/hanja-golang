package hanja

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var (
	_, b, _, _    = runtime.Caller(0)
	basepath      = filepath.Dir(b)
	hanja_hangul  = make(map[string]string)
	table_file_nm = filepath.Join(basepath, "table.yml")
)

func Translate(o string) string {
	if len(hanja_hangul) == 0 {
		loadTable()
	}

	new_text := make([]string, len(o))

	for i, r := range o {
		char := string(r)
		if hangul, isSuccess := hanja_hangul[char]; isSuccess == true {
			new_text[i] = hangul
		} else {
			new_text[i] = char
		}
	}

	return strings.Join(new_text, ``)
}

func loadTable() {
	file, err := os.Open(table_file_nm)
	if err != nil {
		panic("Please download table.yml first, by executing DownloadHanjaHangul() manually. ")
	}
	defer file.Close()

	lineRegexp := regexp.MustCompile(`(.+): (.+)`)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		submatch := lineRegexp.FindStringSubmatch(scanner.Text())
		hanja, _ := strconv.Unquote(submatch[1])
		hangul, _ := strconv.Unquote(submatch[2])
		hanja_hangul[hanja] = hangul
	}

	if err := scanner.Err(); err != nil {
		panic("canner error")
	}
}

func DownloadHanjaHangul() {
	out, _ := os.Create(table_file_nm)
	defer out.Close()

	resp, _ := http.Get(`https://raw.githubusercontent.com/suminb/hanja/develop/hanja/table.yml`)
	defer resp.Body.Close()

	io.Copy(out, resp.Body)
}
