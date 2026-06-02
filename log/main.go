package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	myLog := log.New(os.Stdout, "my:", log.LstdFlags)
	myLog.Println("from myLog")

	myLog.SetPrefix("ohmy:")
	myLog.Println("from mylog")

	var buf bytes.Buffer
	bufLog := log.New(&buf, "buf", log.LstdFlags)
	bufLog.Println("hello")

	fmt.Println("from buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("in there")

	myslog.Info("Hello agin", "key", "val", "age", 25)
}
