package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
/ 词法定义，即对不同的数据类型规定其可使用的标识符
 */

type LexerState struct {
	isEOF	bool	//标识该类型能否作为结尾
	TypeDefine	[]TokenType		//罗列其后面可用的标识符
}

var vaildLexerStates = map[TokenType]LexerState {
	/*
		在表达式开头可使用的
	*/
	Illegal: {
		isEOF: false,
		TypeDefine: []TokenType{
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		//12.3
			StringKind,		//"123"
			BoolKind,		// true, false

			LeftBracket,	// (

			Addition,		// +
			Subtraction,	// -
			Not,			// !
		},
	},

	Identifier: {
		isEOF: true,
		TypeDefine: []TokenType{
			RightBracket,	// )

			Or,				// ||
			And,			// &&

			Addition,		// +
			Subtraction,	// -
			Multiply,		// *
			Divide,			// /
			Model,			// %

			GreaterThan,	// >
			LessThan,		// <
			GreaterEqual,	// >=
			LessEqual,		// <=
			Equal,			// ==
			NotEqual,		// !=
		},
	},

	IntKind: {
		isEOF: true,
		TypeDefine: []TokenType{
			RightBracket,	// )

			Or,				// ||
			And,			// &&

			Addition,		// +
			Subtraction,	// -
			Multiply,		// *
			Divide,			// /
			Model,			// %

			GreaterThan,	// >
			LessThan,		// <
			GreaterEqual,	// >=
			LessEqual,		// <=
			Equal,			// ==
			NotEqual,		// !=
		},
	},

	FloatKind: {
		isEOF: true,
		TypeDefine: []TokenType{
			RightBracket,	// )

			Or,				// ||
			And,			// &&

			Addition,		// +
			Subtraction,	// -
			Multiply,		// *
			Divide,			// /
			Model,			// %

			GreaterThan,	// >
			LessThan,		// <
			GreaterEqual,	// >=
			LessEqual,		// <=
			Equal,			// ==
			NotEqual,		// !=
		},
	},

	StringKind: {
		isEOF: true,
		TypeDefine: []TokenType{
			RightBracket,	// )

			Or,				// ||
			And,			// &&

			Addition,		// +
			/*
			string类型的 > < >= <= 均无意义
			*/
			Equal,			// ==
			NotEqual,		// !=
		},
	},

	BoolKind: {
		isEOF: true,
		TypeDefine: []TokenType{
			RightBracket,	// )

			Or,				// ||
			And,			// &&

			Equal,			// ==
			NotEqual,		// !=
		},
	},

	LeftBracket: {
		isEOF: false,
		TypeDefine: []TokenType{
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		// 12.3
			StringKind,		// "123"
			BoolKind,		// true, false

			Not,			// !

			Addition,		// +
			Subtraction,	// -

		},
	},

	RightBracket: {
		isEOF: true,
		TypeDefine: []TokenType{
			Or,				// ||
			And,			// &&

			Addition,		// +
			Subtraction,	// -
			Multiply,		// *
			Divide,			// /
			Model,			// %

			GreaterThan,	// >
			LessThan,		// <
			GreaterEqual,	// >=
			LessEqual,		// <=
			Equal,			// ==
			NotEqual,		// !=
		},
	},

	Or: {
		isEOF: false,
		TypeDefine: []TokenType{
			LeftBracket,	// (
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		// 12.3
			StringKind,		// "123"
			BoolKind,		// true, false

			Not,			// !

			Addition,		// +
			Subtraction,	// -
		},
	},

	And: {
		isEOF: false,
		TypeDefine: []TokenType{
			LeftBracket,	// (
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		// 12.3
			StringKind,		// "123"
			BoolKind,		// true, false

			Not,			// !

			Addition,		// +
			Subtraction,	// -
		},
	},

	Not:{
		isEOF: false,
		TypeDefine: []TokenType{
			LeftBracket,	// (
			Identifier,		// a
			BoolKind,		// true, false
		},
	},

	Addition: {
		isEOF: false,
		TypeDefine: []TokenType{
			LeftBracket,	// (
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		// 12.3
			StringKind,		// "123"

			Addition,		// +
			Subtraction,	// -
		},
	},

	Subtraction: {
		isEOF: false,
		TypeDefine: []TokenType{
			LeftBracket,	// (
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		// 12.3

			Addition,		// +
			Subtraction,	// -
		},
	},

	Multiply: {
		isEOF: false,
		TypeDefine: []TokenType{
			LeftBracket,	// (
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		// 12.3

			Addition,		// +
			Subtraction,	// -
		},
	},

	Divide: {
		isEOF: false,
		TypeDefine: []TokenType{
			LeftBracket,	// (
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		// 12.3

			Addition,		// +
			Subtraction,	// -
		},
	},

	Model: {
		isEOF: false,
		TypeDefine: []TokenType{
			LeftBracket,	// (
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		// 12.3
		},
	},

	GreaterThan: {
		isEOF: false,
		TypeDefine: []TokenType{
			LeftBracket,	// (
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		// 12.3

			Not,			// !

			Addition,		// +
			Subtraction,	// -
		},
	},

	LessThan: {
		isEOF: false,
		TypeDefine: []TokenType{
			LeftBracket,	// (
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		// 12.3

			Not,			// !

			Addition,		// +
			Subtraction,	// -
		},
	},

	GreaterEqual: {
		isEOF: false,
		TypeDefine: []TokenType{
			LeftBracket,	// (
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		// 12.3

			Not,			// !

			Addition,		// +
			Subtraction,	// -
		},
	},

	LessEqual: {
		isEOF: false,
		TypeDefine: []TokenType{
			LeftBracket,	// (
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		// 12.3

			Not,			// !

			Addition,		// +
			Subtraction,	// -
		},
	},

	Equal: {
		isEOF: false,
		TypeDefine: []TokenType{
			LeftBracket,	// (
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		// 12.3
			StringKind,		// "123"
			BoolKind,		// true, false

			Not,			// !

			Addition,		// +
			Subtraction,	// -
		},
	},

	NotEqual: {
		isEOF: false,
		TypeDefine: []TokenType{
			LeftBracket,	// (
			Identifier,		// a
			IntKind,		// 123
			FloatKind,		// 12.3
			StringKind,		// "123"
			BoolKind,		// true, false

			Not,			// !

			Addition,		// +
			Subtraction,	// -
		},
	},

}


func (l *LexerState) IsEOF() bool {
	return l.isEOF
}


type Lexer struct {
	sentence	[]rune		// 输入的句子
	position	int			// 当前解析的位置
	length		int			// 长度
	ch			rune		// 当前符号源
}

const eofRune = -1

func NewLexer (sen string) *Lexer {
	runes := []rune(sen)
	n := len(runes)
	if n == 0 {
		runes = append(runes, rune(eofRune))
	}
	return &Lexer{
		sentence: runes,
		position: 0,
		length: n,
		ch: runes[0],
	}
}

/*
	1. 是否结尾，是否空格
	2. read是处理空格的
	3. 返回l.ch后一位的值
*/
// l.ch里面存的是下一位字符
func (l *Lexer) read() {
	l.ch = l.peek()
	l.position += 1
}


func (l *Lexer) isNotEnd() bool {
	return l.position < l.length
}

func (l *Lexer) peek() rune {
	if l.position < l.length - 1{
		return l.sentence[l.position + 1]
	}
	return eofRune
}

//碰到空格就跳过，反之则解析
func (l *Lexer) skipBlankSpace() {
	for l.isNotEnd() && unicode.IsSpace(l.ch) {
		l.read()
	}
}

//解析变量
func (l *Lexer) scanIdentifier() string {
	var ans []rune
	ans = append(ans, l.ch)
	for isLetter(l.peek()) {
		l.read()
		ans = append(ans, l.ch)
	}
	return string(ans)
}

//解析字符串.在读完后需要多操作一次，把尾部的"顶掉
func (l *Lexer) scanString() string {
	var ans []rune
	for !isString(l.peek()) {
		l.read()
		ans = append(ans, l.ch)
	}
	l.read()
	return string(ans)
}

func (l *Lexer) scanNumber() string {
	var ans []rune
	ans = append(ans, l.ch)
	for isDigit(l.peek()) || isDot(l.peek()) {
		l.read()
		ans = append(ans, l.ch)
	}
	return string(ans)
}

func (l *Lexer) scanTwoOperator(front TokenType, need rune, follow TokenType)(TokenType, string){
	s := string(l.ch)
	tok := l.peek()
	if tok == need {
		l.read()
		s += string(tok)
		return follow, s
	}
	return front, s
}



func isEOF(ch rune) bool {
	return ch == eofRune
}

// 字母、下划线
func isLetter(ch rune) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9' || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}

func isDot(ch rune) bool {
	return ch == '.'
}

func isString(ch rune) bool{
	return ch == '"'
}

func toLower(ch rune) rune {
	return ('a' - 'A') | ch
}

func isBool(str string) bool {
	return str == "true" || str == "false"
}

/*
	单步解析，数字、点、字母、下划线、符号、字符串
*/
func (l *Lexer) lexerParse() (Token, error){
	var tokens Token
	var err error
	l.skipBlankSpace()
	tokens.index = l.position

	switch ch := l.ch; {
	case isEOF(ch):
		tokens.Type = EOF
	case isLetter(ch):	//判断变量 or 关键字 or Bool Literal
		literal := l.scanIdentifier()	// a or break
		tokens.Value = literal
		tokens.Type = Lookup(literal)
		if isBool(literal) {	//bool
			tokens.Type = BoolKind
		}
	case isString(ch):
		tokens.Value = l.scanString()
		tokens.Type = StringKind
	case isDigit(ch) || isDot(ch):	// 123    12.3    .32  0.32  3.2.1
		literal := l.scanNumber()
		if strings.Contains(literal, ".") {	//float
			tokens.Value, err = strconv.ParseFloat(literal, 64)
			tokens.Type = FloatKind
		} else {	//int
			tokens.Value, err = strconv.ParseInt(literal, 10, 64)
			tokens.Type = IntKind
		}

		if err != nil{
			errorMsg := fmt.Sprintf("Unable to compiler numeric value '%v'", literal)
			return tokens, errors.New(errorMsg)
		}
	default:	//基础类型完成了，剩下符号
		switch ch {
		case '+', '-', '*', '/', '%':	//基础的运算符
			tokens.Type = LookupOperator(string(ch))
			tokens.Value = ch
			l.read()
		case '<':
			tokens.Type, tokens.Value = l.scanTwoOperator(LessThan, '=', LessEqual)
		case '>':
			tokens.Type, tokens.Value = l.scanTwoOperator(GreaterThan, '=', GreaterEqual)
		case '!':
			tokens.Type, tokens.Value = l.scanTwoOperator(Not, '=', NotEqual)
		case '=':
			tokens.Type, tokens.Value = l.scanTwoOperator(Illegal, '=', Equal)
			if tokens.Type == Illegal {
				return tokens, errors.New("expected to get '==', but only found '='")
			}
		case '&':
			tokens.Type, tokens.Value = l.scanTwoOperator(Illegal, '&', And)
			if tokens.Type == Illegal {
				return tokens, errors.New("expected to get '&&', but only found '&'")
			}
		case '|':
			tokens.Type, tokens.Value = l.scanTwoOperator(Illegal, '|', Or)
			if tokens.Type == Illegal {
				return tokens, errors.New("expected to get '||', but only found '|'")
			}
		}

	}


}