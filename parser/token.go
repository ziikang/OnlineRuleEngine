package parser

type TokenType int
/*
	在常量定义中利用_beg 和 _end 来表示开头结尾
*/
const (
	Illegal TokenType = iota	//是否合法
	EOF
	/*
	/对于不同的数据类型有不同的语法对应，需要分别指明
	/123,12.3,"123",true,a
	*/

	literalEeg
	Identifier		// variables
	IntKind			// 123
	FloatKind		// 12.3
	StringKind		// "123"
	BoolKind		// true
	literalEnd

	/*
	/ bracket
	 */
	operatorBeg
	LeftBracket		// (
	RightBracket	// )

	/*
	/ calculate operator
	 */
	Addition		// +
	Subtraction		// -
	Multiply		// *
	Divide			// /
	Model			// %

	/*
	/ comparison operator
	 */
	GreaterThan		// >
	LessThan		// <
	GreaterEqual	// >=
	LessEqual		// <=
	Equal			// ==
	NotEqual		// !=

	/*
		/ logical operator
	*/
	And				// &&
	Or				// ||
	Not				// !
	operatorEnd
	/*
	keyword
	*/
	keywordBeg
	BREAK
	CASE
	CHAN
	CONST
	CONTINUE

	DEFAULT
	DEFER
	ELSE
	FALLTHROUGH
	FOR

	FUNC
	GO
	GOTO
	IF
	IMPORT

	INTERFACE
	MAP
	PACKAGE
	RANGE
	RETURN

	SELECT
	STRUCT
	SWITCH
	TYPE
	VAR
	keywordEnd
)

var tokens = [...]string{
	/*
	* special tokens
	* */
	Illegal: "Illegal",
	EOF:     "EOF",

	/*
	* literals of nil, bool, number, string
	* */
	Identifier:     "Identifier",
	IntKind: "IntKind",
	FloatKind:   "FloatKind",
	StringKind:  "StringKind",
	BoolKind:    "BoolKind",

	/*
	* single character operator
	* */
	LeftBracket:  "(",
	RightBracket: ")",

	/*
	* arithmetic operator
	* */
	Addition:    "+",
	Subtraction: "-",
	Multiply:    "*",
	Divide:      "/",
	Model:     "%",

	/*
	* cmp operator
	* */
	GreaterThan:  ">",
	LessThan:     "<",
	GreaterEqual: ">=",
	LessEqual:    "<=",
	Equal:        "==",
	NotEqual:     "!=",

	/*
	* logic operator
	* */
	And: "&&",
	Or:  "||",
	Not: "!",

	BREAK:    "break",
	CASE:     "case",
	CHAN:     "chan",
	CONST:    "const",
	CONTINUE: "continue",

	DEFAULT:     "default",
	DEFER:       "defer",
	ELSE:        "else",
	FALLTHROUGH: "fallthrough",
	FOR:         "for",

	FUNC:   "func",
	GO:     "go",
	GOTO:   "goto",
	IF:     "if",
	IMPORT: "import",

	INTERFACE: "interface",
	MAP:       "map",
	PACKAGE:   "package",
	RANGE:     "range",
	RETURN:    "return",

	SELECT: "select",
	STRUCT: "struct",
	SWITCH: "switch",
	TYPE:   "type",
	VAR:    "var",

}

var operatorToKind = map[string]TokenType{
	"(" : LeftBracket,
	")" : RightBracket,

	"||" : Or,
	"&&" : And,
	"!" : Not,

	"+" : Addition,
	"-" : Subtraction,
	"*" : Multiply,
	"/" : Divide,
	"%" : Model,

	">" : GreaterThan,
	"<" : LessThan,
	">=" : GreaterEqual,
	"<=" : LessEqual,
	"==" : Equal,
	"!=" : NotEqual,
}

type Token struct {
	Type	TokenType
	Value	interface{}
	index	int				// index表示在表达式中的位置
}

var keywords map[string]TokenType

func init() {
	keywords = make(map[string]TokenType)
	for i := keywordBeg + 1; i < keywordEnd;i++{
		keywords[tokens[i]] = i
	}
}

/*
	check whether it's keyword
*/
func Lookup(ident string) TokenType {
	if tok, isKeyword := keywords[ident]; isKeyword {
		return tok
	}
	return Identifier
}

func LookupOperator(str string) TokenType {

}
