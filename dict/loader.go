package dict

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/huichen/sego"
)

//go:generate go-bindata -prefix "assets/" -pkg dict -o assets.go assets/
func Load(segmenter *sego.Segmenter) {
	var files = make([]string, len(AssetNames()))

	for i, v := range AssetNames() {
		data, err := Asset(v)
		if err != nil {
			log.Fatal(err)
		}

		tmpfile, err := ioutil.TempFile("", v)
		if err != nil {
			log.Fatal(err)
		}

		defer os.Remove(tmpfile.Name())
		files[i] = tmpfile.Name()

		if _, err := tmpfile.Write(data); err != nil {
			log.Fatal(err)
		}
		if err := tmpfile.Close(); err != nil {
			log.Fatal(err)
		}
	}

	segmenter.LoadDictionary(strings.Join(files, ","))
}
