// https://go.dev/play/p/bPTXaPHm6jD
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type SimpleMessage struct {
	Content string `json:"content"`
}

func main() {
	fmt.Println("--- Testing Trailing Garbage Data with json.Decoder ---")
	// 一个有效的 JSON 对象，后面跟着 "恶意payload"
	jsonDataWithTrailing := `{"content":"legit data"} {"content":"legit data"}`
	reader := bytes.NewReader([]byte(jsonDataWithTrailing))
	decoder := json.NewDecoder(reader)

	var msg SimpleMessage
	// Decoder.Decode() 会尝试解码流中的下一个 JSON 值
	err := decoder.Decode(&msg)
	if err != nil {
		// 如果 JSON 本身格式错误，这里会报错
		fmt.Println("Initial Decode Error:", err)
	} else {
		// 第一个 JSON 对象被成功解码
		fmt.Printf("Successfully Decoded Message: %+v\n", msg)
	}

	// 关键：检查 Decode 之后流中是否还有剩余数据
	// Trail of Bits 指出这是 encoding/json 的一个开放 issue (golang/go#13407)，
	// 即 Decoder.Decode 后面跟非空白字符不报错。
	// 通常需要额外调用 decoder.Token() 并检查是否为 io.EOF 来确保流已耗尽。
	var buf [1]byte
	n, errPeek := reader.Read(buf[:]) // 尝试读取 Decode 之后的数据
	if n > 0 {
		fmt.Printf("!!! VULNERABILITY RISK: Trailing garbage data "+
			"found after valid JSON: '%s'\n", string(buf[:n]))
		// 在某些场景下，如果应用只调用 Decode() 一次且不检查流的末尾，
		// 攻击者可能通过附加数据来尝试进行其他类型的攻击。
	} else if errPeek == io.EOF {
		fmt.Println("Stream fully consumed as expected.")
	} else if errPeek != nil {
		fmt.Println("Error peeking after decode:", errPeek)
	} else {
		fmt.Println("No trailing data or EOF not reached clearly.")
	}

	// 更规范的检查方式是使用 decoder.More() 或尝试再解码一个Token
	fmt.Println("\n--- Proper check for trailing data ---")
	reader2 := bytes.NewReader([]byte(jsonDataWithTrailing))
	decoder2 := json.NewDecoder(reader2)
	var msg2 SimpleMessage
	decoder2.Decode(&msg2) // 解码第一个

	// 尝试解码下一个token，期望是EOF
	tok, errTok := decoder2.Token()
	if errTok == io.EOF {
		fmt.Println("Proper check: Stream fully consumed (EOF).")
	} else if errTok != nil {
		fmt.Printf("Proper check: Error after expected JSON object:"+
			" %v (Token: %v)\n", errTok, tok)
	} else if tok != nil {
		fmt.Printf("!!! VULNERABILITY RISK (Proper check):"+
			" Unexpected token after first JSON object: %v\n", tok)
	}
}
