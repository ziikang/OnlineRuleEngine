package parser

/*
/ 词法定义，即对不同的数据类型规定其可使用的标识符
 */

type LexterState struct {
	isEOF	bool	//标识该类型能否作为结尾
	TypeDefine	[]TokenType		//罗列其后面可用的标识符
}

var vaildLexerStates = map[TokenType]LexterState {
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


func (l *LexterState) IsEOF() bool {
	return l.isEOF
}