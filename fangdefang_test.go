package gofangdefang_test

import (
	"os"
	"testing"

	"github.com/atakanaydinbas/gofangdefang"
)

func TestDefangText(t *testing.T) {
	defangedBytes, err := os.ReadFile("testfiles/sample_text_defanged.txt")
	if err != nil {
		t.Error(err)
	}

	fangedBytes, err := os.ReadFile("testfiles/sample_text_fanged.txt")
	if err != nil {
		t.Error(err)
	}

	if string(defangedBytes) != gofangdefang.DefangText(string(fangedBytes)) {
		t.Error("DefangText function is not working properly.")
	}

	if string(fangedBytes) != gofangdefang.FangText(string(defangedBytes)) {
		t.Error("FangText function is not working properly.")
	}

}
