package parser

type TokenType int

const (
	Illegal TokenType = iota	//是否合法


	/*
	/对于不同的数据类型有不同的语法对应，需要分别指明
	/123,12.3,"123",true,a
	*/
	Identifier		// variables
	IntKind			// 123
	FloatKind		// 12.3
	StringKind		// "123"
	BoolKind		// true

	/*
	/ bracket
	 */
	LeftBracket		// (
	RightBracket	// )

	/*
	/ logical operator
	 */
	Or				// ||
	And				// &&
	Not				// !

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
)

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