package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {

	// 只是写一个本地文件，用于证明确实能够执行可执行文件，在实际进行钓鱼攻击的时候可以替换为木马文件直接执行啥的
	message := fmt.Sprintf("[%s] 您好，当您看到此行文本的时候，您已经被成功钓鱼，请及时打开杀软扫描清理！", time.Now().Format("2006-01-02 15:04:05"))
	content := []byte(message)
	err := ioutil.WriteFile("prove.txt", content, 0644)
	if err != nil {
		panic(err)
	}
}