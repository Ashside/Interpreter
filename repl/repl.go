package repl

import (
	"Interp/lexer"
	"Interp/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> " //提示符

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in) //创建一个scanner，用于读取输入
	//NOTE: 这里的scanner是一个指针，所以在调用Scan()方法时，会改变scanner的值

	//死循环，不断的读取输入
	for {
		//打印提示符
		_, err := fmt.Fprintf(out, PROMPT)
		if err != nil {
			return
		}
		//读取输入
		scanned := scanner.Scan()
		//如果读取失败，就退出循环
		if !scanned {
			return
		}

		line := scanner.Text()    //获取输入的内容
		l := lexer.NewLexer(line) //创建一个lexer，用于解析输入

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			_, err := fmt.Fprintf(out, "%+v\n", tok) //%+v表示打印结构体时会添加字段名
			if err != nil {
				return
			}
		}
	}
}
