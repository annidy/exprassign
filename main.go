package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/annidy/exprassign/pkg/expr"
)

var exprList = []expr.AssignExpr{}

const COMPILE_TIME = "2024年 4月10日 星期三 20时21分11秒 CST"

func main() {
	var (
		file string
		buf  bytes.Buffer
	)
	// 定义-f file参数
	flag.StringVar(&file, "f", "", "optional file path")
	flag.BoolFunc("v", "print version", func(_ string) error { fmt.Println(COMPILE_TIME); os.Exit(0); return nil })
	flag.Parse()

	args := flag.Args()
	for _, arg := range args {
		exprs := strings.Split(arg, "=")
		if len(exprs) != 2 {
			panic("invalid argument: " + arg)
		}
		exprList = append(exprList, expr.AssignExpr{Key: exprs[0], Value: exprs[1]})
	}

	if len(exprList) == 0 {
		log.Fatal("no expression found")
	}

	w := bufio.NewWriter(&buf)
	scanner := bufio.NewScanner(os.Stdin)
	if file != "" {
		fd, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(fd)
		defer fd.Close()
	}

	for scanner.Scan() {
		t := scanner.Text()
		ok := false
		for _, expr := range exprList {
			var r string
			r, ok = expr.Assign(t)
			if ok {
				fmt.Fprintln(w, r)
				break
			}
		}
		if !ok {
			fmt.Fprintln(w, t)
		}
	}
	w.Flush()
	if file != "" {
		os.WriteFile(file, buf.Bytes(), 0644)
	} else {
		fmt.Print(buf.String())
	}
}
