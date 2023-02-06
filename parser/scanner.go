package parser

import "unicode"

/*
	对输入进行扫描，分割出各个符号，以换行符 \n 作为输入的结尾

*/

type Scanner struct {
	sentence	[]rune		// 输入的句子
	position	int			// 当前解析的位置
	length		int			// 长度
	ch			rune		// 当前符号源
}

const zeroInput = -1
const eofRune = -1

func NewScanner (sen string) *Scanner {
	runes := []rune(sen)
	n := len(runes)
	if n == 0 {
		runes = append(runes, rune(zeroInput))
	}
	return &Scanner{
		sentence: runes,
		position: 0,
		length: n,
		ch: runes[0],
	}
}

/*
	1. 是否读完
	2. 提取字符并
	3.
*/
func (scan *Scanner) read() {
	if scan.isEnd() == true {
		return
	}
}

func (scan *Scanner) isEnd() bool {
	return scan.position >= scan.length
}

// 碰到空格就跳过，反之则解析
func (scan *Scanner) skipBlankSpace() {
	for !scan.isEnd() && unicode.IsSpace(scan.)
}