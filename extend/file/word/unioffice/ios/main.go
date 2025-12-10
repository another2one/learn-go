package main

import (
	"bufio"
	"fmt"
	"github.com/gomutex/godocx"
	"github.com/gomutex/godocx/wml/stypes"
	"learn-go/common/tool"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	basePath  = tool.ProjectPath + "extend/file/word/unioffice/ios/"
	maxPage   int // 最大页数
	maxLine   int // 最大行数
	breakLine = stypes.BreakTypeTextWrapping
)

// 使用示例 C:\Users\lizhi\go\pkg\mod\baliance.com\gooxml@v1.0.1\_examples
// 主要是编辑word 读取非文字的时候会有点问题
func main() {
	start()
}

func start() {
	maxPage = 65
	maxLine = 50
	// /Users/lizhi/Desktop/app/mac/merchant_ios/LWMMerchantApp
	codePaths := map[string][]string{
		"xy_merchant": {
			"/Users/lizhi/Desktop/app/mac/merchant_ios/LWMMerchantApp/AppDelegate",
			"/Users/lizhi/Desktop/app/mac/merchant_ios/LWMMerchantApp/Category",
			"/Users/lizhi/Desktop/app/mac/merchant_ios/LWMMerchantApp/Modules/Home",
			"/Users/lizhi/Desktop/app/mac/merchant_ios/LWMMerchantApp/Modules/Login",
			"/Users/lizhi/Desktop/app/mac/merchant_ios/LWMMerchantApp/Modules/Mine",
		},
		"tc_merchant": {
			"/Users/lizhi/Desktop/app/mac/merchant_ios/LWMMerchantApp/Modules/More",
			"/Users/lizhi/Desktop/app/mac/merchant_ios/LWMMerchantApp/Modules/Shop",
			"/Users/lizhi/Desktop/app/mac/merchant_ios/LWMMerchantApp/QCY",
			"/Users/lizhi/Desktop/app/mac/merchant_ios/LWMMerchantApp/SLY",
		},
		"xy_customer": {
			"/Users/lizhi/Desktop/app/mac/customer_ios/乐外卖消费者/Classes/Tools",
			"/Users/lizhi/Desktop/app/mac/customer_ios/乐外卖消费者/Classes/Main",
			"/Users/lizhi/Desktop/app/mac/customer_ios/乐外卖消费者/Classes/Sections/Choose",
			"/Users/lizhi/Desktop/app/mac/customer_ios/乐外卖消费者/Classes/Sections/City",
			"/Users/lizhi/Desktop/app/mac/customer_ios/乐外卖消费者/Classes/Sections/ShoppingCar",
		},
		"tc_customer": {
			"/Users/lizhi/Desktop/app/mac/customer_ios/乐外卖消费者/Classes/Sections/Errand",
			"/Users/lizhi/Desktop/app/mac/customer_ios/乐外卖消费者/Classes/Sections/Home",
			"/Users/lizhi/Desktop/app/mac/customer_ios/乐外卖消费者/Classes/Sections/Login",
			"/Users/lizhi/Desktop/app/mac/customer_ios/乐外卖消费者/Classes/Sections/Mine",
			"/Users/lizhi/Desktop/app/mac/customer_ios/乐外卖消费者/Classes/Sections/Orders",
		},
		"xy_deliveryman": {
			"/Users/lizhi/Desktop/app/mac/deliveryman_ios/LWMDeliverymanApp/AppDelegate",
			"/Users/lizhi/Desktop/app/mac/deliveryman_ios/LWMDeliverymanApp/Category",
			"/Users/lizhi/Desktop/app/mac/deliveryman_ios/LWMDeliverymanApp/ThridParty",
			"/Users/lizhi/Desktop/app/mac/deliveryman_ios/LWMDeliverymanApp/Modules/Delivery",
		},
		"tc_deliveryman": {
			"/Users/lizhi/Desktop/app/mac/deliveryman_ios/LWMDeliverymanApp/Utils",
			"/Users/lizhi/Desktop/app/mac/deliveryman_ios/LWMDeliverymanApp/Modules/Errand",
			"/Users/lizhi/Desktop/app/mac/deliveryman_ios/LWMDeliverymanApp/Modules/Main",
			"/Users/lizhi/Desktop/app/mac/deliveryman_ios/LWMDeliverymanApp/Modules/Mine",
		},
		"xy_fast_distribution": {
			"/Users/lizhi/Desktop/app/mac/FastDistributionAPP/FastDistributionAPP/AppDelegate",
			"/Users/lizhi/Desktop/app/mac/FastDistributionAPP/FastDistributionAPP/Category",
			"/Users/lizhi/Desktop/app/mac/FastDistributionAPP/FastDistributionAPP/ThridParty",
			"/Users/lizhi/Desktop/app/mac/FastDistributionAPP/FastDistributionAPP/Modules/AddOrder",
			"/Users/lizhi/Desktop/app/mac/FastDistributionAPP/FastDistributionAPP/Modules/Home",
		},
		"tc_fast_distribution": {
			// 代码量不够，重复了部分
			"/Users/lizhi/Desktop/app/mac/FastDistributionAPP/FastDistributionAPP/Utils",
			"/Users/lizhi/Desktop/app/mac/FastDistributionAPP/FastDistributionAPP/Modules/Login",
			"/Users/lizhi/Desktop/app/mac/FastDistributionAPP/FastDistributionAPP/Modules/Main",
			"/Users/lizhi/Desktop/app/mac/FastDistributionAPP/FastDistributionAPP/Modules/Mine",
			"/Users/lizhi/Desktop/app/mac/FastDistributionAPP/FastDistributionAPP/Modules/Home",
			"/Users/lizhi/Desktop/app/mac/FastDistributionAPP/pods",
		},
	}
	for fileName, paths := range codePaths {
		fmt.Printf("\n软件名称: %s \n", fileName)
		ruanzhu(basePath+fileName+".docx", paths)
	}

}

// ruanzhu 软件源代码（前30页和后30页，每页50行；不足60页则提交全部）
func ruanzhu(docFile string, codePaths []string) {
	if err := os.Remove(docFile); err != nil {
		fmt.Printf("删除文件失败: %s \n", err)
	}

	document, err := godocx.NewDocument()
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	page := 1

	// 遍历目录下的 php 和 vue 文件, 读取前 maxLine 行
	for _, path := range codePaths {

		if page > maxPage {
			break
		}

		err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				fmt.Printf("访问出错: %v \n", err)
				return nil
			}

			// 检查文件扩展名
			if !d.IsDir() && (strings.HasSuffix(strings.ToLower(path), ".m") ||
				strings.HasSuffix(strings.ToLower(path), ".h")) {

				// 读取文件前50行
				textArr, err := readLines(path)
				if err != nil {
					//fmt.Printf("读取文件(%s)失败: %v", d.Name(), err)
					return nil
				}

				if page > maxPage {
					return nil
				}
				p := document.AddParagraph("")
				for _, text := range textArr {
					p.AddText(text).Size(10).AddBreak(&breakLine)
				}
				if page != maxPage {
					document.AddPageBreak()
				}
				fmt.Printf("文件%d: %s \n", page, d.Name())
				page++
			}

			return nil
		})

		if err != nil {
			fmt.Printf("遍历失败: %v \n", err)
		}
	}

	err = document.SaveTo(docFile)
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
		// 去除空行和注释
		str := strings.TrimSpace(scanner.Text())
		if len(str) == 0 || tool.HasAnyPrefix(str, "/*", "*", "//") {
			continue
		}
		textArr[lineCount] = scanner.Text()
		lineCount++
	}

	if lineCount < 50 {
		return nil, fmt.Errorf("文件共 %d 行 \n", lineCount)
	}

	return textArr, scanner.Err()
}
