package main

import (
	"bufio"
	"fmt"
	"github.com/gomutex/godocx"
	"github.com/gomutex/godocx/common/units"
	"github.com/gomutex/godocx/wml/stypes"
	"learn-go/common/funcs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	Path      = funcs.ProjectPath + "extend/file/word/unioffice/zqcrm.docx"
	basePath  = funcs.ProjectPath + "extend/file/word/unioffice/"
	maxPage   int      // 最大页数
	maxLine   int      // 最大行数
	codePaths []string // 代码路径集合（按顺序）
	breakLine = stypes.BreakTypeTextWrapping
)

// 使用示例 C:\Users\lizhi\go\pkg\mod\baliance.com\gooxml@v1.0.1\_examples
// 主要是编辑word 读取非文字的时候会有点问题
func main() {
	test()
}

// ruanzhu 软件源代码（前30页和后30页，每页50行；不足60页则提交全部）
func ruanzhu() {
	maxPage = 60
	maxLine = 50
	CodePaths := []string{
		"/Users/lizhi/Desktop/app/php/waimai/zhuqu-crm/protected/controllers",
		"/Users/lizhi/Desktop/app/php/waimai/zhuqu-crm/protected/include",
		"/Users/lizhi/Desktop/app/php/waimai/zhuqu-crm/web/src/views",
		"/Users/lizhi/Desktop/app/php/waimai/zhuqu-crm/protected/lib",
	}
	if err := os.Remove(Path); err != nil {
		fmt.Printf("删除文件失败: %s \n", err)
	}

	document, err := godocx.NewDocument()
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	page := 1

	// 遍历目录下的 php 和 vue 文件, 读取前 maxLine 行
	for _, path := range CodePaths {

		if page > maxPage {
			break
		}

		err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				fmt.Printf("访问出错: %v \n", err)
				return nil
			}

			// 检查文件扩展名
			if !d.IsDir() && (strings.HasSuffix(strings.ToLower(path), ".php") ||
				strings.HasSuffix(strings.ToLower(path), ".vue")) {

				// 读取文件前50行
				textArr, err := readLines(path)
				if err != nil {
					fmt.Printf("读取文件(%s)失败: %v", d.Name(), err)
					return nil
				}

				if page > maxPage {
					return nil
				}
				p := document.AddParagraph("")
				for _, text := range textArr {
					p.AddText(text).AddBreak(&breakLine)
				}
				document.AddPageBreak()
				fmt.Printf("文件 %d: %s \n", page, d.Name())
				page++
			}

			return nil
		})

		if err != nil {
			fmt.Printf("遍历失败: %v \n", err)
		}
	}

	err = document.SaveTo(Path)
	if err != nil {
		log.Fatalf("word保存失败: %s", err)
	}
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("关闭文件(%s)失败: %s \n\n", filePath, err)
		}
	}(file)

	textArr := make([]string, 50)

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() && lineCount < maxLine {
		// 去除空行
		if len(strings.TrimSpace(scanner.Text())) == 0 {
			continue
		}
		textArr[lineCount] = scanner.Text()
		lineCount++
	}

	if lineCount < 50 {
		return nil, fmt.Errorf("(文件共 %d 行) \n", lineCount)
	}

	return textArr, scanner.Err()
}

// test https://github.com/gomutex/godocx-examples/blob/main/table/main.go
func test() {
	document, err := godocx.NewDocument()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// 添加文档标题
	document.AddHeading("Document Title", 0)

	// 添加一个新段落到文档
	p := document.AddParagraph("A plain paragraph having some ")
	p.AddText("bold").Bold(true) // 设置加粗文本
	p.AddText(" and some ").Color("FF0000")
	p.AddText("italic.").Italic(true).AddBreak(&breakLine) // 设置斜体文本

	// 添加一级标题
	document.AddHeading("Heading, level 1", 1)
	// 添加引用样式段落
	document.AddParagraph("Intense quote").Style("Intense Quote")
	// 添加无序列表项
	document.AddParagraph("first item in unordered list").Style("List Bullet")
	// 添加有序列表项
	document.AddParagraph("first item in ordered list").Style("List Number")
	// 添加图片
	pic1Path := basePath + "test.jpg"
	w, h, err := funcs.GetImageDimensions(pic1Path)
	if err != nil {
		fmt.Printf("获取图片(%s)尺寸失败: %s \n", pic1Path, err)
	} else {
		inch := 2
		_, err := document.AddPicture(pic1Path, units.Inch(inch), units.Inch(float32(h)*float32(inch)/float32(w)))
		if err != nil {
			log.Fatalf("添加图片(%s)失败: %s", pic1Path, err)
		}
	}

	// 准备表格数据
	records := []struct{ Qty, ID, Desc string }{
		{"5", "A001", "Laptop"},
		{"10", "B202", "Smartphone"},
		{"2", "E505", "Smartwatch"},
	}

	// 添加表格
	table := document.AddTable()
	table.Style("LightList-Accent4") // 设置表格样式

	// 添加表头行
	hdrRow := table.AddRow()
	hdrRow.AddCell().AddParagraph("Qty")
	hdrRow.AddCell().AddParagraph("ID")
	hdrRow.AddCell().AddParagraph("Description")

	// 添加数据行
	for _, record := range records {
		row := table.AddRow()
		row.AddCell().AddParagraph(record.Qty)
		row.AddCell().AddParagraph(record.ID)
		row.AddCell().AddParagraph(record.Desc)
	}

	document.AddEmptyParagraph()

	err = document.SaveTo(basePath + "test.docx")
	if err != nil {
		log.Fatalf("word保存失败: %s", err)
	}
}
