package main

import (
	"fmt"
	"learn-go/common/funcs"
	"log"

	"baliance.com/gooxml/document"
)

// 使用示例 C:\Users\lizhi\go\pkg\mod\baliance.com\gooxml@v1.0.1\_examples
// 主要是编辑word 读取非文字的时候会有点问题
func main() {
	test()
}

func test() {
	// /Users/lizhi/Desktop/app/go/learn-go/extend/file/word/goxml/test.docx
	doc, err := document.Open(funcs.ProjectPath + "extend/file/word/goxml/test.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	//doc.Paragraphs()得到包含文档所有的段落的切片
	for i, para := range doc.Paragraphs() {
		//run为每个段落相同格式的文字组成的片段
		fmt.Println("-----------第", i, "段-------------")
		for j, run := range para.Runs() {
			fmt.Print("\t-----------第", j, "格式片段-------------")
			fmt.Printf("%v \n", run.Text())
		}
	}
}
