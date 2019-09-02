package restclient

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func prettyPrint(v interface{}) {
	_, file, line, _ := runtime.Caller(1)
	lineNo := strconv.Itoa(line)
	file = filepath.Base(file)
	b, _ := json.MarshalIndent(v, "", "\t")
	s := file + ":" + lineNo + ": \n" + string(b) + "\n"
	os.Stderr.WriteString(s)
}
