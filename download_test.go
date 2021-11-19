package main

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	err := Download(`https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png`, `/tmp/notexist/test.jpg`)
	if err != nil {
		t.Errorf("download error: %+v", err)
		return
	}

	file, err := os.Open(`/tmp/notexist/test.jpg`)
	if err != nil {
		t.Errorf("open file error: %+v", err)
		return
	}

	bFile, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf("read file error: %+v", err)
		return
	}

	baMD5 := make([]byte, 16, 16)
	for index, byteItem := range md5.Sum(bFile) {
		baMD5[index] = byteItem
	}

	if "d9c8750bed0b3c7d089fa7d55720d6cf" != hex.EncodeToString(baMD5) {
		t.Errorf("file md5 not equal to 'd9c8750bed0b3c7d089fa7d55720d6cf'")
		return
	}
}
