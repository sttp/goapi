// Code generated from FilterExpressionSyntax.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // FilterExpressionSyntax

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 100, 245,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 5, 2, 55, 10, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 7, 4, 63, 10, 4, 12, 4, 14, 4, 66,
	11, 4, 3, 4, 3, 4, 6, 4, 70, 10, 4, 13, 4, 14, 4, 71, 3, 4, 7, 4, 75, 10,
	4, 12, 4, 14, 4, 78, 11, 4, 3, 4, 7, 4, 81, 10, 4, 12, 4, 14, 4, 84, 11,
	4, 3, 5, 3, 5, 3, 5, 5, 5, 89, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 5,
	7, 96, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 106,
	10, 7, 12, 7, 14, 7, 109, 11, 7, 5, 7, 111, 10, 7, 3, 8, 5, 8, 114, 10,
	8, 3, 8, 3, 8, 3, 9, 5, 9, 119, 10, 9, 3, 9, 3, 9, 5, 9, 123, 10, 9, 3,
	10, 3, 10, 3, 10, 7, 10, 128, 10, 10, 12, 10, 14, 10, 131, 11, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 138, 10, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 7, 11, 144, 10, 11, 12, 11, 14, 11, 147, 11, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 158, 10, 12, 3, 12, 3,
	12, 5, 12, 162, 10, 12, 3, 12, 3, 12, 3, 12, 5, 12, 167, 10, 12, 3, 12,
	3, 12, 5, 12, 171, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 5, 12, 180, 10, 12, 3, 12, 7, 12, 183, 10, 12, 12, 12, 14, 12, 186,
	11, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 5, 13, 199, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 7, 13, 209, 10, 13, 12, 13, 14, 13, 212, 11, 13, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 5, 22, 233, 10, 22,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	26, 2, 5, 20, 22, 24, 27, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 2, 14, 3, 2, 92, 94, 3,
	2, 5, 6, 4, 2, 33, 33, 43, 43, 4, 2, 9, 9, 62, 62, 5, 2, 5, 6, 9, 10, 62,
	62, 4, 2, 11, 11, 34, 34, 3, 2, 11, 20, 5, 2, 21, 22, 32, 32, 66, 66, 4,
	2, 23, 27, 87, 87, 4, 2, 5, 6, 28, 30, 12, 2, 31, 31, 36, 42, 44, 44, 46,
	47, 49, 49, 51, 57, 59, 61, 63, 64, 68, 79, 81, 85, 6, 2, 65, 65, 88, 88,
	90, 92, 95, 96, 2, 251, 2, 54, 3, 2, 2, 2, 4, 58, 3, 2, 2, 2, 6, 64, 3,
	2, 2, 2, 8, 88, 3, 2, 2, 2, 10, 90, 3, 2, 2, 2, 12, 92, 3, 2, 2, 2, 14,
	113, 3, 2, 2, 2, 16, 118, 3, 2, 2, 2, 18, 124, 3, 2, 2, 2, 20, 137, 3,
	2, 2, 2, 22, 148, 3, 2, 2, 2, 24, 198, 3, 2, 2, 2, 26, 213, 3, 2, 2, 2,
	28, 215, 3, 2, 2, 2, 30, 217, 3, 2, 2, 2, 32, 219, 3, 2, 2, 2, 34, 221,
	3, 2, 2, 2, 36, 223, 3, 2, 2, 2, 38, 225, 3, 2, 2, 2, 40, 227, 3, 2, 2,
	2, 42, 229, 3, 2, 2, 2, 44, 236, 3, 2, 2, 2, 46, 238, 3, 2, 2, 2, 48, 240,
	3, 2, 2, 2, 50, 242, 3, 2, 2, 2, 52, 55, 5, 6, 4, 2, 53, 55, 5, 4, 3, 2,
	54, 52, 3, 2, 2, 2, 54, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 57, 7,
	2, 2, 3, 57, 3, 3, 2, 2, 2, 58, 59, 7, 100, 2, 2, 59, 60, 8, 3, 1, 2, 60,
	5, 3, 2, 2, 2, 61, 63, 7, 3, 2, 2, 62, 61, 3, 2, 2, 2, 63, 66, 3, 2, 2,
	2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 67, 3, 2, 2, 2, 66, 64,
	3, 2, 2, 2, 67, 76, 5, 8, 5, 2, 68, 70, 7, 3, 2, 2, 69, 68, 3, 2, 2, 2,
	70, 71, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 3,
	2, 2, 2, 73, 75, 5, 8, 5, 2, 74, 69, 3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76,
	74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 82, 3, 2, 2, 2, 78, 76, 3, 2, 2,
	2, 79, 81, 7, 3, 2, 2, 80, 79, 3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80,
	3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 7, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2,
	85, 89, 5, 10, 6, 2, 86, 89, 5, 12, 7, 2, 87, 89, 5, 20, 11, 2, 88, 85,
	3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 87, 3, 2, 2, 2, 89, 9, 3, 2, 2, 2,
	90, 91, 9, 2, 2, 2, 91, 11, 3, 2, 2, 2, 92, 95, 7, 45, 2, 2, 93, 94, 7,
	80, 2, 2, 94, 96, 5, 14, 8, 2, 95, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2,
	96, 97, 3, 2, 2, 2, 97, 98, 5, 46, 24, 2, 98, 99, 7, 86, 2, 2, 99, 110,
	5, 20, 11, 2, 100, 101, 7, 67, 2, 2, 101, 102, 7, 35, 2, 2, 102, 107, 5,
	16, 9, 2, 103, 104, 7, 4, 2, 2, 104, 106, 5, 16, 9, 2, 105, 103, 3, 2,
	2, 2, 106, 109, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2,
	108, 111, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 110, 100, 3, 2, 2, 2, 110,
	111, 3, 2, 2, 2, 111, 13, 3, 2, 2, 2, 112, 114, 9, 3, 2, 2, 113, 112, 3,
	2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 116, 7, 90, 2,
	2, 116, 15, 3, 2, 2, 2, 117, 119, 5, 30, 16, 2, 118, 117, 3, 2, 2, 2, 118,
	119, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 122, 5, 50, 26, 2, 121, 123,
	9, 4, 2, 2, 122, 121, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 17, 3, 2,
	2, 2, 124, 129, 5, 20, 11, 2, 125, 126, 7, 4, 2, 2, 126, 128, 5, 20, 11,
	2, 127, 125, 3, 2, 2, 2, 128, 131, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129,
	130, 3, 2, 2, 2, 130, 19, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 132, 133, 8,
	11, 1, 2, 133, 134, 5, 26, 14, 2, 134, 135, 5, 20, 11, 5, 135, 138, 3,
	2, 2, 2, 136, 138, 5, 22, 12, 2, 137, 132, 3, 2, 2, 2, 137, 136, 3, 2,
	2, 2, 138, 145, 3, 2, 2, 2, 139, 140, 12, 4, 2, 2, 140, 141, 5, 34, 18,
	2, 141, 142, 5, 20, 11, 5, 142, 144, 3, 2, 2, 2, 143, 139, 3, 2, 2, 2,
	144, 147, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146,
	21, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 148, 149, 8, 12, 1, 2, 149, 150,
	5, 24, 13, 2, 150, 184, 3, 2, 2, 2, 151, 152, 12, 5, 2, 2, 152, 153, 5,
	32, 17, 2, 153, 154, 5, 22, 12, 6, 154, 183, 3, 2, 2, 2, 155, 157, 12,
	4, 2, 2, 156, 158, 5, 26, 14, 2, 157, 156, 3, 2, 2, 2, 157, 158, 3, 2,
	2, 2, 158, 159, 3, 2, 2, 2, 159, 161, 7, 58, 2, 2, 160, 162, 5, 30, 16,
	2, 161, 160, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163,
	183, 5, 22, 12, 5, 164, 166, 12, 7, 2, 2, 165, 167, 5, 26, 14, 2, 166,
	165, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 170,
	7, 48, 2, 2, 169, 171, 5, 30, 16, 2, 170, 169, 3, 2, 2, 2, 170, 171, 3,
	2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 173, 7, 7, 2, 2, 173, 174, 5, 18, 10,
	2, 174, 175, 7, 8, 2, 2, 175, 183, 3, 2, 2, 2, 176, 177, 12, 6, 2, 2, 177,
	179, 7, 50, 2, 2, 178, 180, 5, 26, 14, 2, 179, 178, 3, 2, 2, 2, 179, 180,
	3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 183, 7, 65, 2, 2, 182, 151, 3, 2,
	2, 2, 182, 155, 3, 2, 2, 2, 182, 164, 3, 2, 2, 2, 182, 176, 3, 2, 2, 2,
	183, 186, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185,
	23, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 187, 188, 8, 13, 1, 2, 188, 199,
	5, 44, 23, 2, 189, 199, 5, 48, 25, 2, 190, 199, 5, 42, 22, 2, 191, 192,
	5, 28, 15, 2, 192, 193, 5, 24, 13, 6, 193, 199, 3, 2, 2, 2, 194, 195, 7,
	7, 2, 2, 195, 196, 5, 20, 11, 2, 196, 197, 7, 8, 2, 2, 197, 199, 3, 2,
	2, 2, 198, 187, 3, 2, 2, 2, 198, 189, 3, 2, 2, 2, 198, 190, 3, 2, 2, 2,
	198, 191, 3, 2, 2, 2, 198, 194, 3, 2, 2, 2, 199, 210, 3, 2, 2, 2, 200,
	201, 12, 4, 2, 2, 201, 202, 5, 38, 20, 2, 202, 203, 5, 24, 13, 5, 203,
	209, 3, 2, 2, 2, 204, 205, 12, 3, 2, 2, 205, 206, 5, 36, 19, 2, 206, 207,
	5, 24, 13, 4, 207, 209, 3, 2, 2, 2, 208, 200, 3, 2, 2, 2, 208, 204, 3,
	2, 2, 2, 209, 212, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2, 2,
	2, 211, 25, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 213, 214, 9, 5, 2, 2, 214,
	27, 3, 2, 2, 2, 215, 216, 9, 6, 2, 2, 216, 29, 3, 2, 2, 2, 217, 218, 9,
	7, 2, 2, 218, 31, 3, 2, 2, 2, 219, 220, 9, 8, 2, 2, 220, 33, 3, 2, 2, 2,
	221, 222, 9, 9, 2, 2, 222, 35, 3, 2, 2, 2, 223, 224, 9, 10, 2, 2, 224,
	37, 3, 2, 2, 2, 225, 226, 9, 11, 2, 2, 226, 39, 3, 2, 2, 2, 227, 228, 9,
	12, 2, 2, 228, 41, 3, 2, 2, 2, 229, 230, 5, 40, 21, 2, 230, 232, 7, 7,
	2, 2, 231, 233, 5, 18, 10, 2, 232, 231, 3, 2, 2, 2, 232, 233, 3, 2, 2,
	2, 233, 234, 3, 2, 2, 2, 234, 235, 7, 8, 2, 2, 235, 43, 3, 2, 2, 2, 236,
	237, 9, 13, 2, 2, 237, 45, 3, 2, 2, 2, 238, 239, 7, 89, 2, 2, 239, 47,
	3, 2, 2, 2, 240, 241, 7, 89, 2, 2, 241, 49, 3, 2, 2, 2, 242, 243, 7, 89,
	2, 2, 243, 51, 3, 2, 2, 2, 28, 54, 64, 71, 76, 82, 88, 95, 107, 110, 113,
	118, 122, 129, 137, 145, 157, 161, 166, 170, 179, 182, 184, 198, 208, 210,
	232,
}
var literalNames = []string{
	"", "';'", "','", "'-'", "'+'", "'('", "')'", "'!'", "'~'", "'==='", "'<'",
	"'<='", "'>'", "'>='", "'='", "'=='", "'!='", "'!=='", "'<>'", "'&&'",
	"'||'", "'<<'", "'>>'", "'&'", "'|'", "'^'", "'*'", "'/'", "'%'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "K_ABS", "K_AND", "K_ASC",
	"K_BINARY", "K_BY", "K_CEILING", "K_COALESCE", "K_CONVERT", "K_CONTAINS",
	"K_DATEADD", "K_DATEDIFF", "K_DATEPART", "K_DESC", "K_ENDSWITH", "K_FILTER",
	"K_FLOOR", "K_IIF", "K_IN", "K_INDEXOF", "K_IS", "K_ISDATE", "K_ISINTEGER",
	"K_ISGUID", "K_ISNULL", "K_ISNUMERIC", "K_LASTINDEXOF", "K_LEN", "K_LIKE",
	"K_LOWER", "K_MAXOF", "K_MINOF", "K_NOT", "K_NOW", "K_NTHINDEXOF", "K_NULL",
	"K_OR", "K_ORDER", "K_POWER", "K_REGEXMATCH", "K_REGEXVAL", "K_REPLACE",
	"K_REVERSE", "K_ROUND", "K_SQRT", "K_SPLIT", "K_STARTSWITH", "K_STRCOUNT",
	"K_STRCMP", "K_SUBSTR", "K_TOP", "K_TRIM", "K_TRIMLEFT", "K_TRIMRIGHT",
	"K_UPPER", "K_UTCNOW", "K_WHERE", "K_XOR", "BOOLEAN_LITERAL", "IDENTIFIER",
	"INTEGER_LITERAL", "NUMERIC_LITERAL", "GUID_LITERAL", "MEASUREMENT_KEY_LITERAL",
	"POINT_TAG_LITERAL", "STRING_LITERAL", "DATETIME_LITERAL", "SINGLE_LINE_COMMENT",
	"MULTILINE_COMMENT", "SPACES", "UNEXPECTED_CHAR",
}

var ruleNames = []string{
	"parse", "err", "filterExpressionStatementList", "filterExpressionStatement",
	"identifierStatement", "filterStatement", "topLimit", "orderingTerm", "expressionList",
	"expression", "predicateExpression", "valueExpression", "notOperator",
	"unaryOperator", "exactMatchModifier", "comparisonOperator", "logicalOperator",
	"bitwiseOperator", "mathOperator", "functionName", "functionExpression",
	"literalValue", "tableName", "columnName", "orderByColumnName",
}

type FilterExpressionSyntaxParser struct {
	*antlr.BaseParser
}

// NewFilterExpressionSyntaxParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *FilterExpressionSyntaxParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewFilterExpressionSyntaxParser(input antlr.TokenStream) *FilterExpressionSyntaxParser {
	this := new(FilterExpressionSyntaxParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "FilterExpressionSyntax.g4"

	return this
}

// FilterExpressionSyntaxParser tokens.
const (
	FilterExpressionSyntaxParserEOF                     = antlr.TokenEOF
	FilterExpressionSyntaxParserT__0                    = 1
	FilterExpressionSyntaxParserT__1                    = 2
	FilterExpressionSyntaxParserT__2                    = 3
	FilterExpressionSyntaxParserT__3                    = 4
	FilterExpressionSyntaxParserT__4                    = 5
	FilterExpressionSyntaxParserT__5                    = 6
	FilterExpressionSyntaxParserT__6                    = 7
	FilterExpressionSyntaxParserT__7                    = 8
	FilterExpressionSyntaxParserT__8                    = 9
	FilterExpressionSyntaxParserT__9                    = 10
	FilterExpressionSyntaxParserT__10                   = 11
	FilterExpressionSyntaxParserT__11                   = 12
	FilterExpressionSyntaxParserT__12                   = 13
	FilterExpressionSyntaxParserT__13                   = 14
	FilterExpressionSyntaxParserT__14                   = 15
	FilterExpressionSyntaxParserT__15                   = 16
	FilterExpressionSyntaxParserT__16                   = 17
	FilterExpressionSyntaxParserT__17                   = 18
	FilterExpressionSyntaxParserT__18                   = 19
	FilterExpressionSyntaxParserT__19                   = 20
	FilterExpressionSyntaxParserT__20                   = 21
	FilterExpressionSyntaxParserT__21                   = 22
	FilterExpressionSyntaxParserT__22                   = 23
	FilterExpressionSyntaxParserT__23                   = 24
	FilterExpressionSyntaxParserT__24                   = 25
	FilterExpressionSyntaxParserT__25                   = 26
	FilterExpressionSyntaxParserT__26                   = 27
	FilterExpressionSyntaxParserT__27                   = 28
	FilterExpressionSyntaxParserK_ABS                   = 29
	FilterExpressionSyntaxParserK_AND                   = 30
	FilterExpressionSyntaxParserK_ASC                   = 31
	FilterExpressionSyntaxParserK_BINARY                = 32
	FilterExpressionSyntaxParserK_BY                    = 33
	FilterExpressionSyntaxParserK_CEILING               = 34
	FilterExpressionSyntaxParserK_COALESCE              = 35
	FilterExpressionSyntaxParserK_CONVERT               = 36
	FilterExpressionSyntaxParserK_CONTAINS              = 37
	FilterExpressionSyntaxParserK_DATEADD               = 38
	FilterExpressionSyntaxParserK_DATEDIFF              = 39
	FilterExpressionSyntaxParserK_DATEPART              = 40
	FilterExpressionSyntaxParserK_DESC                  = 41
	FilterExpressionSyntaxParserK_ENDSWITH              = 42
	FilterExpressionSyntaxParserK_FILTER                = 43
	FilterExpressionSyntaxParserK_FLOOR                 = 44
	FilterExpressionSyntaxParserK_IIF                   = 45
	FilterExpressionSyntaxParserK_IN                    = 46
	FilterExpressionSyntaxParserK_INDEXOF               = 47
	FilterExpressionSyntaxParserK_IS                    = 48
	FilterExpressionSyntaxParserK_ISDATE                = 49
	FilterExpressionSyntaxParserK_ISINTEGER             = 50
	FilterExpressionSyntaxParserK_ISGUID                = 51
	FilterExpressionSyntaxParserK_ISNULL                = 52
	FilterExpressionSyntaxParserK_ISNUMERIC             = 53
	FilterExpressionSyntaxParserK_LASTINDEXOF           = 54
	FilterExpressionSyntaxParserK_LEN                   = 55
	FilterExpressionSyntaxParserK_LIKE                  = 56
	FilterExpressionSyntaxParserK_LOWER                 = 57
	FilterExpressionSyntaxParserK_MAXOF                 = 58
	FilterExpressionSyntaxParserK_MINOF                 = 59
	FilterExpressionSyntaxParserK_NOT                   = 60
	FilterExpressionSyntaxParserK_NOW                   = 61
	FilterExpressionSyntaxParserK_NTHINDEXOF            = 62
	FilterExpressionSyntaxParserK_NULL                  = 63
	FilterExpressionSyntaxParserK_OR                    = 64
	FilterExpressionSyntaxParserK_ORDER                 = 65
	FilterExpressionSyntaxParserK_POWER                 = 66
	FilterExpressionSyntaxParserK_REGEXMATCH            = 67
	FilterExpressionSyntaxParserK_REGEXVAL              = 68
	FilterExpressionSyntaxParserK_REPLACE               = 69
	FilterExpressionSyntaxParserK_REVERSE               = 70
	FilterExpressionSyntaxParserK_ROUND                 = 71
	FilterExpressionSyntaxParserK_SQRT                  = 72
	FilterExpressionSyntaxParserK_SPLIT                 = 73
	FilterExpressionSyntaxParserK_STARTSWITH            = 74
	FilterExpressionSyntaxParserK_STRCOUNT              = 75
	FilterExpressionSyntaxParserK_STRCMP                = 76
	FilterExpressionSyntaxParserK_SUBSTR                = 77
	FilterExpressionSyntaxParserK_TOP                   = 78
	FilterExpressionSyntaxParserK_TRIM                  = 79
	FilterExpressionSyntaxParserK_TRIMLEFT              = 80
	FilterExpressionSyntaxParserK_TRIMRIGHT             = 81
	FilterExpressionSyntaxParserK_UPPER                 = 82
	FilterExpressionSyntaxParserK_UTCNOW                = 83
	FilterExpressionSyntaxParserK_WHERE                 = 84
	FilterExpressionSyntaxParserK_XOR                   = 85
	FilterExpressionSyntaxParserBOOLEAN_LITERAL         = 86
	FilterExpressionSyntaxParserIDENTIFIER              = 87
	FilterExpressionSyntaxParserINTEGER_LITERAL         = 88
	FilterExpressionSyntaxParserNUMERIC_LITERAL         = 89
	FilterExpressionSyntaxParserGUID_LITERAL            = 90
	FilterExpressionSyntaxParserMEASUREMENT_KEY_LITERAL = 91
	FilterExpressionSyntaxParserPOINT_TAG_LITERAL       = 92
	FilterExpressionSyntaxParserSTRING_LITERAL          = 93
	FilterExpressionSyntaxParserDATETIME_LITERAL        = 94
	FilterExpressionSyntaxParserSINGLE_LINE_COMMENT     = 95
	FilterExpressionSyntaxParserMULTILINE_COMMENT       = 96
	FilterExpressionSyntaxParserSPACES                  = 97
	FilterExpressionSyntaxParserUNEXPECTED_CHAR         = 98
)

// FilterExpressionSyntaxParser rules.
const (
	FilterExpressionSyntaxParserRULE_parse                         = 0
	FilterExpressionSyntaxParserRULE_err                           = 1
	FilterExpressionSyntaxParserRULE_filterExpressionStatementList = 2
	FilterExpressionSyntaxParserRULE_filterExpressionStatement     = 3
	FilterExpressionSyntaxParserRULE_identifierStatement           = 4
	FilterExpressionSyntaxParserRULE_filterStatement               = 5
	FilterExpressionSyntaxParserRULE_topLimit                      = 6
	FilterExpressionSyntaxParserRULE_orderingTerm                  = 7
	FilterExpressionSyntaxParserRULE_expressionList                = 8
	FilterExpressionSyntaxParserRULE_expression                    = 9
	FilterExpressionSyntaxParserRULE_predicateExpression           = 10
	FilterExpressionSyntaxParserRULE_valueExpression               = 11
	FilterExpressionSyntaxParserRULE_notOperator                   = 12
	FilterExpressionSyntaxParserRULE_unaryOperator                 = 13
	FilterExpressionSyntaxParserRULE_exactMatchModifier            = 14
	FilterExpressionSyntaxParserRULE_comparisonOperator            = 15
	FilterExpressionSyntaxParserRULE_logicalOperator               = 16
	FilterExpressionSyntaxParserRULE_bitwiseOperator               = 17
	FilterExpressionSyntaxParserRULE_mathOperator                  = 18
	FilterExpressionSyntaxParserRULE_functionName                  = 19
	FilterExpressionSyntaxParserRULE_functionExpression            = 20
	FilterExpressionSyntaxParserRULE_literalValue                  = 21
	FilterExpressionSyntaxParserRULE_tableName                     = 22
	FilterExpressionSyntaxParserRULE_columnName                    = 23
	FilterExpressionSyntaxParserRULE_orderByColumnName             = 24
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserEOF, 0)
}

func (s *ParseContext) FilterExpressionStatementList() IFilterExpressionStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterExpressionStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterExpressionStatementListContext)
}

func (s *ParseContext) Err() IErrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IErrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IErrContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitParse(s)
	}
}

func (p *FilterExpressionSyntaxParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FilterExpressionSyntaxParserRULE_parse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FilterExpressionSyntaxParserT__0, FilterExpressionSyntaxParserT__2, FilterExpressionSyntaxParserT__3, FilterExpressionSyntaxParserT__4, FilterExpressionSyntaxParserT__6, FilterExpressionSyntaxParserT__7, FilterExpressionSyntaxParserK_ABS, FilterExpressionSyntaxParserK_CEILING, FilterExpressionSyntaxParserK_COALESCE, FilterExpressionSyntaxParserK_CONVERT, FilterExpressionSyntaxParserK_CONTAINS, FilterExpressionSyntaxParserK_DATEADD, FilterExpressionSyntaxParserK_DATEDIFF, FilterExpressionSyntaxParserK_DATEPART, FilterExpressionSyntaxParserK_ENDSWITH, FilterExpressionSyntaxParserK_FILTER, FilterExpressionSyntaxParserK_FLOOR, FilterExpressionSyntaxParserK_IIF, FilterExpressionSyntaxParserK_INDEXOF, FilterExpressionSyntaxParserK_ISDATE, FilterExpressionSyntaxParserK_ISINTEGER, FilterExpressionSyntaxParserK_ISGUID, FilterExpressionSyntaxParserK_ISNULL, FilterExpressionSyntaxParserK_ISNUMERIC, FilterExpressionSyntaxParserK_LASTINDEXOF, FilterExpressionSyntaxParserK_LEN, FilterExpressionSyntaxParserK_LOWER, FilterExpressionSyntaxParserK_MAXOF, FilterExpressionSyntaxParserK_MINOF, FilterExpressionSyntaxParserK_NOT, FilterExpressionSyntaxParserK_NOW, FilterExpressionSyntaxParserK_NTHINDEXOF, FilterExpressionSyntaxParserK_NULL, FilterExpressionSyntaxParserK_POWER, FilterExpressionSyntaxParserK_REGEXMATCH, FilterExpressionSyntaxParserK_REGEXVAL, FilterExpressionSyntaxParserK_REPLACE, FilterExpressionSyntaxParserK_REVERSE, FilterExpressionSyntaxParserK_ROUND, FilterExpressionSyntaxParserK_SQRT, FilterExpressionSyntaxParserK_SPLIT, FilterExpressionSyntaxParserK_STARTSWITH, FilterExpressionSyntaxParserK_STRCOUNT, FilterExpressionSyntaxParserK_STRCMP, FilterExpressionSyntaxParserK_SUBSTR, FilterExpressionSyntaxParserK_TRIM, FilterExpressionSyntaxParserK_TRIMLEFT, FilterExpressionSyntaxParserK_TRIMRIGHT, FilterExpressionSyntaxParserK_UPPER, FilterExpressionSyntaxParserK_UTCNOW, FilterExpressionSyntaxParserBOOLEAN_LITERAL, FilterExpressionSyntaxParserIDENTIFIER, FilterExpressionSyntaxParserINTEGER_LITERAL, FilterExpressionSyntaxParserNUMERIC_LITERAL, FilterExpressionSyntaxParserGUID_LITERAL, FilterExpressionSyntaxParserMEASUREMENT_KEY_LITERAL, FilterExpressionSyntaxParserPOINT_TAG_LITERAL, FilterExpressionSyntaxParserSTRING_LITERAL, FilterExpressionSyntaxParserDATETIME_LITERAL:
		{
			p.SetState(50)
			p.FilterExpressionStatementList()
		}

	case FilterExpressionSyntaxParserUNEXPECTED_CHAR:
		{
			p.SetState(51)
			p.Err()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(54)
		p.Match(FilterExpressionSyntaxParserEOF)
	}

	return localctx
}

// IErrContext is an interface to support dynamic dispatch.
type IErrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_UNEXPECTED_CHAR returns the _UNEXPECTED_CHAR token.
	Get_UNEXPECTED_CHAR() antlr.Token

	// Set_UNEXPECTED_CHAR sets the _UNEXPECTED_CHAR token.
	Set_UNEXPECTED_CHAR(antlr.Token)

	// IsErrContext differentiates from other interfaces.
	IsErrContext()
}

type ErrContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	_UNEXPECTED_CHAR antlr.Token
}

func NewEmptyErrContext() *ErrContext {
	var p = new(ErrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_err
	return p
}

func (*ErrContext) IsErrContext() {}

func NewErrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrContext {
	var p = new(ErrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_err

	return p
}

func (s *ErrContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrContext) Get_UNEXPECTED_CHAR() antlr.Token { return s._UNEXPECTED_CHAR }

func (s *ErrContext) Set_UNEXPECTED_CHAR(v antlr.Token) { s._UNEXPECTED_CHAR = v }

func (s *ErrContext) UNEXPECTED_CHAR() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserUNEXPECTED_CHAR, 0)
}

func (s *ErrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ErrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterErr(s)
	}
}

func (s *ErrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitErr(s)
	}
}

func (p *FilterExpressionSyntaxParser) Err() (localctx IErrContext) {
	localctx = NewErrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FilterExpressionSyntaxParserRULE_err)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)

		var _m = p.Match(FilterExpressionSyntaxParserUNEXPECTED_CHAR)

		localctx.(*ErrContext)._UNEXPECTED_CHAR = _m
	}

	panic("Unexpected character: " + (func() string {
		if localctx.(*ErrContext).Get_UNEXPECTED_CHAR() == nil {
			return ""
		} else {
			return localctx.(*ErrContext).Get_UNEXPECTED_CHAR().GetText()
		}
	}()))

	//return localctx
}

// IFilterExpressionStatementListContext is an interface to support dynamic dispatch.
type IFilterExpressionStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterExpressionStatementListContext differentiates from other interfaces.
	IsFilterExpressionStatementListContext()
}

type FilterExpressionStatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterExpressionStatementListContext() *FilterExpressionStatementListContext {
	var p = new(FilterExpressionStatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_filterExpressionStatementList
	return p
}

func (*FilterExpressionStatementListContext) IsFilterExpressionStatementListContext() {}

func NewFilterExpressionStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterExpressionStatementListContext {
	var p = new(FilterExpressionStatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_filterExpressionStatementList

	return p
}

func (s *FilterExpressionStatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterExpressionStatementListContext) AllFilterExpressionStatement() []IFilterExpressionStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilterExpressionStatementContext)(nil)).Elem())
	var tst = make([]IFilterExpressionStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilterExpressionStatementContext)
		}
	}

	return tst
}

func (s *FilterExpressionStatementListContext) FilterExpressionStatement(i int) IFilterExpressionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterExpressionStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilterExpressionStatementContext)
}

func (s *FilterExpressionStatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterExpressionStatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterExpressionStatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterFilterExpressionStatementList(s)
	}
}

func (s *FilterExpressionStatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitFilterExpressionStatementList(s)
	}
}

func (p *FilterExpressionSyntaxParser) FilterExpressionStatementList() (localctx IFilterExpressionStatementListContext) {
	localctx = NewFilterExpressionStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FilterExpressionSyntaxParserRULE_filterExpressionStatementList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FilterExpressionSyntaxParserT__0 {
		{
			p.SetState(59)
			p.Match(FilterExpressionSyntaxParserT__0)
		}

		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(65)
		p.FilterExpressionStatement()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == FilterExpressionSyntaxParserT__0 {
				{
					p.SetState(66)
					p.Match(FilterExpressionSyntaxParserT__0)
				}

				p.SetState(69)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(71)
				p.FilterExpressionStatement()
			}

		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FilterExpressionSyntaxParserT__0 {
		{
			p.SetState(77)
			p.Match(FilterExpressionSyntaxParserT__0)
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFilterExpressionStatementContext is an interface to support dynamic dispatch.
type IFilterExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterExpressionStatementContext differentiates from other interfaces.
	IsFilterExpressionStatementContext()
}

type FilterExpressionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterExpressionStatementContext() *FilterExpressionStatementContext {
	var p = new(FilterExpressionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_filterExpressionStatement
	return p
}

func (*FilterExpressionStatementContext) IsFilterExpressionStatementContext() {}

func NewFilterExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterExpressionStatementContext {
	var p = new(FilterExpressionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_filterExpressionStatement

	return p
}

func (s *FilterExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterExpressionStatementContext) IdentifierStatement() IIdentifierStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierStatementContext)
}

func (s *FilterExpressionStatementContext) FilterStatement() IFilterStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterStatementContext)
}

func (s *FilterExpressionStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FilterExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterFilterExpressionStatement(s)
	}
}

func (s *FilterExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitFilterExpressionStatement(s)
	}
}

func (p *FilterExpressionSyntaxParser) FilterExpressionStatement() (localctx IFilterExpressionStatementContext) {
	localctx = NewFilterExpressionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FilterExpressionSyntaxParserRULE_filterExpressionStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.IdentifierStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.FilterStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(85)
			p.expression(0)
		}

	}

	return localctx
}

// IIdentifierStatementContext is an interface to support dynamic dispatch.
type IIdentifierStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierStatementContext differentiates from other interfaces.
	IsIdentifierStatementContext()
}

type IdentifierStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierStatementContext() *IdentifierStatementContext {
	var p = new(IdentifierStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_identifierStatement
	return p
}

func (*IdentifierStatementContext) IsIdentifierStatementContext() {}

func NewIdentifierStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierStatementContext {
	var p = new(IdentifierStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_identifierStatement

	return p
}

func (s *IdentifierStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierStatementContext) GUID_LITERAL() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserGUID_LITERAL, 0)
}

func (s *IdentifierStatementContext) MEASUREMENT_KEY_LITERAL() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserMEASUREMENT_KEY_LITERAL, 0)
}

func (s *IdentifierStatementContext) POINT_TAG_LITERAL() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserPOINT_TAG_LITERAL, 0)
}

func (s *IdentifierStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterIdentifierStatement(s)
	}
}

func (s *IdentifierStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitIdentifierStatement(s)
	}
}

func (p *FilterExpressionSyntaxParser) IdentifierStatement() (localctx IIdentifierStatementContext) {
	localctx = NewIdentifierStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FilterExpressionSyntaxParserRULE_identifierStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-90)&-(0x1f+1)) == 0 && ((1<<uint((_la-90)))&((1<<(FilterExpressionSyntaxParserGUID_LITERAL-90))|(1<<(FilterExpressionSyntaxParserMEASUREMENT_KEY_LITERAL-90))|(1<<(FilterExpressionSyntaxParserPOINT_TAG_LITERAL-90)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFilterStatementContext is an interface to support dynamic dispatch.
type IFilterStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterStatementContext differentiates from other interfaces.
	IsFilterStatementContext()
}

type FilterStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterStatementContext() *FilterStatementContext {
	var p = new(FilterStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_filterStatement
	return p
}

func (*FilterStatementContext) IsFilterStatementContext() {}

func NewFilterStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterStatementContext {
	var p = new(FilterStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_filterStatement

	return p
}

func (s *FilterStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterStatementContext) K_FILTER() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_FILTER, 0)
}

func (s *FilterStatementContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *FilterStatementContext) K_WHERE() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_WHERE, 0)
}

func (s *FilterStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FilterStatementContext) K_TOP() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_TOP, 0)
}

func (s *FilterStatementContext) TopLimit() ITopLimitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopLimitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITopLimitContext)
}

func (s *FilterStatementContext) K_ORDER() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_ORDER, 0)
}

func (s *FilterStatementContext) K_BY() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_BY, 0)
}

func (s *FilterStatementContext) AllOrderingTerm() []IOrderingTermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOrderingTermContext)(nil)).Elem())
	var tst = make([]IOrderingTermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOrderingTermContext)
		}
	}

	return tst
}

func (s *FilterStatementContext) OrderingTerm(i int) IOrderingTermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderingTermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOrderingTermContext)
}

func (s *FilterStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterFilterStatement(s)
	}
}

func (s *FilterStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitFilterStatement(s)
	}
}

func (p *FilterExpressionSyntaxParser) FilterStatement() (localctx IFilterStatementContext) {
	localctx = NewFilterStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FilterExpressionSyntaxParserRULE_filterStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(FilterExpressionSyntaxParserK_FILTER)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FilterExpressionSyntaxParserK_TOP {
		{
			p.SetState(91)
			p.Match(FilterExpressionSyntaxParserK_TOP)
		}
		{
			p.SetState(92)
			p.TopLimit()
		}

	}
	{
		p.SetState(95)
		p.TableName()
	}
	{
		p.SetState(96)
		p.Match(FilterExpressionSyntaxParserK_WHERE)
	}
	{
		p.SetState(97)
		p.expression(0)
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FilterExpressionSyntaxParserK_ORDER {
		{
			p.SetState(98)
			p.Match(FilterExpressionSyntaxParserK_ORDER)
		}
		{
			p.SetState(99)
			p.Match(FilterExpressionSyntaxParserK_BY)
		}
		{
			p.SetState(100)
			p.OrderingTerm()
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FilterExpressionSyntaxParserT__1 {
			{
				p.SetState(101)
				p.Match(FilterExpressionSyntaxParserT__1)
			}
			{
				p.SetState(102)
				p.OrderingTerm()
			}

			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// ITopLimitContext is an interface to support dynamic dispatch.
type ITopLimitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopLimitContext differentiates from other interfaces.
	IsTopLimitContext()
}

type TopLimitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLimitContext() *TopLimitContext {
	var p = new(TopLimitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_topLimit
	return p
}

func (*TopLimitContext) IsTopLimitContext() {}

func NewTopLimitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLimitContext {
	var p = new(TopLimitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_topLimit

	return p
}

func (s *TopLimitContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLimitContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserINTEGER_LITERAL, 0)
}

func (s *TopLimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLimitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLimitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterTopLimit(s)
	}
}

func (s *TopLimitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitTopLimit(s)
	}
}

func (p *FilterExpressionSyntaxParser) TopLimit() (localctx ITopLimitContext) {
	localctx = NewTopLimitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FilterExpressionSyntaxParserRULE_topLimit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FilterExpressionSyntaxParserT__2 || _la == FilterExpressionSyntaxParserT__3 {
		{
			p.SetState(110)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FilterExpressionSyntaxParserT__2 || _la == FilterExpressionSyntaxParserT__3) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(113)
		p.Match(FilterExpressionSyntaxParserINTEGER_LITERAL)
	}

	return localctx
}

// IOrderingTermContext is an interface to support dynamic dispatch.
type IOrderingTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderingTermContext differentiates from other interfaces.
	IsOrderingTermContext()
}

type OrderingTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderingTermContext() *OrderingTermContext {
	var p = new(OrderingTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_orderingTerm
	return p
}

func (*OrderingTermContext) IsOrderingTermContext() {}

func NewOrderingTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderingTermContext {
	var p = new(OrderingTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_orderingTerm

	return p
}

func (s *OrderingTermContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderingTermContext) OrderByColumnName() IOrderByColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderByColumnNameContext)
}

func (s *OrderingTermContext) ExactMatchModifier() IExactMatchModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExactMatchModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExactMatchModifierContext)
}

func (s *OrderingTermContext) K_ASC() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_ASC, 0)
}

func (s *OrderingTermContext) K_DESC() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_DESC, 0)
}

func (s *OrderingTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderingTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderingTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterOrderingTerm(s)
	}
}

func (s *OrderingTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitOrderingTerm(s)
	}
}

func (p *FilterExpressionSyntaxParser) OrderingTerm() (localctx IOrderingTermContext) {
	localctx = NewOrderingTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FilterExpressionSyntaxParserRULE_orderingTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FilterExpressionSyntaxParserT__8 || _la == FilterExpressionSyntaxParserK_BINARY {
		{
			p.SetState(115)
			p.ExactMatchModifier()
		}

	}
	{
		p.SetState(118)
		p.OrderByColumnName()
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FilterExpressionSyntaxParserK_ASC || _la == FilterExpressionSyntaxParserK_DESC {
		{
			p.SetState(119)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FilterExpressionSyntaxParserK_ASC || _la == FilterExpressionSyntaxParserK_DESC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (p *FilterExpressionSyntaxParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FilterExpressionSyntaxParserRULE_expressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.expression(0)
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FilterExpressionSyntaxParserT__1 {
		{
			p.SetState(123)
			p.Match(FilterExpressionSyntaxParserT__1)
		}
		{
			p.SetState(124)
			p.expression(0)
		}

		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) NotOperator() INotOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotOperatorContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) PredicateExpression() IPredicateExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateExpressionContext)
}

func (s *ExpressionContext) LogicalOperator() ILogicalOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalOperatorContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FilterExpressionSyntaxParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *FilterExpressionSyntaxParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, FilterExpressionSyntaxParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(131)
			p.NotOperator()
		}
		{
			p.SetState(132)
			p.expression(3)
		}

	case 2:
		{
			p.SetState(134)
			p.predicateExpression(0)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, FilterExpressionSyntaxParserRULE_expression)
			p.SetState(137)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(138)
				p.LogicalOperator()
			}
			{
				p.SetState(139)
				p.expression(3)
			}

		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

// IPredicateExpressionContext is an interface to support dynamic dispatch.
type IPredicateExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateExpressionContext differentiates from other interfaces.
	IsPredicateExpressionContext()
}

type PredicateExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateExpressionContext() *PredicateExpressionContext {
	var p = new(PredicateExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_predicateExpression
	return p
}

func (*PredicateExpressionContext) IsPredicateExpressionContext() {}

func NewPredicateExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateExpressionContext {
	var p = new(PredicateExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_predicateExpression

	return p
}

func (s *PredicateExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateExpressionContext) ValueExpression() IValueExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueExpressionContext)
}

func (s *PredicateExpressionContext) AllPredicateExpression() []IPredicateExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPredicateExpressionContext)(nil)).Elem())
	var tst = make([]IPredicateExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPredicateExpressionContext)
		}
	}

	return tst
}

func (s *PredicateExpressionContext) PredicateExpression(i int) IPredicateExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPredicateExpressionContext)
}

func (s *PredicateExpressionContext) ComparisonOperator() IComparisonOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *PredicateExpressionContext) K_LIKE() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_LIKE, 0)
}

func (s *PredicateExpressionContext) NotOperator() INotOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotOperatorContext)
}

func (s *PredicateExpressionContext) ExactMatchModifier() IExactMatchModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExactMatchModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExactMatchModifierContext)
}

func (s *PredicateExpressionContext) K_IN() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_IN, 0)
}

func (s *PredicateExpressionContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *PredicateExpressionContext) K_IS() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_IS, 0)
}

func (s *PredicateExpressionContext) K_NULL() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_NULL, 0)
}

func (s *PredicateExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterPredicateExpression(s)
	}
}

func (s *PredicateExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitPredicateExpression(s)
	}
}

func (p *FilterExpressionSyntaxParser) PredicateExpression() (localctx IPredicateExpressionContext) {
	return p.predicateExpression(0)
}

//gocyclo:ignore
func (p *FilterExpressionSyntaxParser) predicateExpression(_p int) (localctx IPredicateExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPredicateExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPredicateExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, FilterExpressionSyntaxParserRULE_predicateExpression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.valueExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(180)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPredicateExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FilterExpressionSyntaxParserRULE_predicateExpression)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(150)
					p.ComparisonOperator()
				}
				{
					p.SetState(151)
					p.predicateExpression(4)
				}

			case 2:
				localctx = NewPredicateExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FilterExpressionSyntaxParserRULE_predicateExpression)
				p.SetState(153)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				p.SetState(155)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == FilterExpressionSyntaxParserT__6 || _la == FilterExpressionSyntaxParserK_NOT {
					{
						p.SetState(154)
						p.NotOperator()
					}

				}
				{
					p.SetState(157)
					p.Match(FilterExpressionSyntaxParserK_LIKE)
				}
				p.SetState(159)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == FilterExpressionSyntaxParserT__8 || _la == FilterExpressionSyntaxParserK_BINARY {
					{
						p.SetState(158)
						p.ExactMatchModifier()
					}

				}
				{
					p.SetState(161)
					p.predicateExpression(3)
				}

			case 3:
				localctx = NewPredicateExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FilterExpressionSyntaxParserRULE_predicateExpression)
				p.SetState(162)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(164)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == FilterExpressionSyntaxParserT__6 || _la == FilterExpressionSyntaxParserK_NOT {
					{
						p.SetState(163)
						p.NotOperator()
					}

				}
				{
					p.SetState(166)
					p.Match(FilterExpressionSyntaxParserK_IN)
				}
				p.SetState(168)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == FilterExpressionSyntaxParserT__8 || _la == FilterExpressionSyntaxParserK_BINARY {
					{
						p.SetState(167)
						p.ExactMatchModifier()
					}

				}
				{
					p.SetState(170)
					p.Match(FilterExpressionSyntaxParserT__4)
				}
				{
					p.SetState(171)
					p.ExpressionList()
				}
				{
					p.SetState(172)
					p.Match(FilterExpressionSyntaxParserT__5)
				}

			case 4:
				localctx = NewPredicateExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FilterExpressionSyntaxParserRULE_predicateExpression)
				p.SetState(174)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(175)
					p.Match(FilterExpressionSyntaxParserK_IS)
				}
				p.SetState(177)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == FilterExpressionSyntaxParserT__6 || _la == FilterExpressionSyntaxParserK_NOT {
					{
						p.SetState(176)
						p.NotOperator()
					}

				}
				{
					p.SetState(179)
					p.Match(FilterExpressionSyntaxParserK_NULL)
				}

			}

		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}

	return localctx
}

// IValueExpressionContext is an interface to support dynamic dispatch.
type IValueExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueExpressionContext differentiates from other interfaces.
	IsValueExpressionContext()
}

type ValueExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueExpressionContext() *ValueExpressionContext {
	var p = new(ValueExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_valueExpression
	return p
}

func (*ValueExpressionContext) IsValueExpressionContext() {}

func NewValueExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueExpressionContext {
	var p = new(ValueExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_valueExpression

	return p
}

func (s *ValueExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueExpressionContext) LiteralValue() ILiteralValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralValueContext)
}

func (s *ValueExpressionContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ValueExpressionContext) FunctionExpression() IFunctionExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionExpressionContext)
}

func (s *ValueExpressionContext) UnaryOperator() IUnaryOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOperatorContext)
}

func (s *ValueExpressionContext) AllValueExpression() []IValueExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueExpressionContext)(nil)).Elem())
	var tst = make([]IValueExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueExpressionContext)
		}
	}

	return tst
}

func (s *ValueExpressionContext) ValueExpression(i int) IValueExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueExpressionContext)
}

func (s *ValueExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ValueExpressionContext) MathOperator() IMathOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathOperatorContext)
}

func (s *ValueExpressionContext) BitwiseOperator() IBitwiseOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitwiseOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBitwiseOperatorContext)
}

func (s *ValueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterValueExpression(s)
	}
}

func (s *ValueExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitValueExpression(s)
	}
}

func (p *FilterExpressionSyntaxParser) ValueExpression() (localctx IValueExpressionContext) {
	return p.valueExpression(0)
}

//gocyclo:ignore
func (p *FilterExpressionSyntaxParser) valueExpression(_p int) (localctx IValueExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewValueExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IValueExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, FilterExpressionSyntaxParserRULE_valueExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FilterExpressionSyntaxParserK_NULL, FilterExpressionSyntaxParserBOOLEAN_LITERAL, FilterExpressionSyntaxParserINTEGER_LITERAL, FilterExpressionSyntaxParserNUMERIC_LITERAL, FilterExpressionSyntaxParserGUID_LITERAL, FilterExpressionSyntaxParserSTRING_LITERAL, FilterExpressionSyntaxParserDATETIME_LITERAL:
		{
			p.SetState(186)
			p.LiteralValue()
		}

	case FilterExpressionSyntaxParserIDENTIFIER:
		{
			p.SetState(187)
			p.ColumnName()
		}

	case FilterExpressionSyntaxParserK_ABS, FilterExpressionSyntaxParserK_CEILING, FilterExpressionSyntaxParserK_COALESCE, FilterExpressionSyntaxParserK_CONVERT, FilterExpressionSyntaxParserK_CONTAINS, FilterExpressionSyntaxParserK_DATEADD, FilterExpressionSyntaxParserK_DATEDIFF, FilterExpressionSyntaxParserK_DATEPART, FilterExpressionSyntaxParserK_ENDSWITH, FilterExpressionSyntaxParserK_FLOOR, FilterExpressionSyntaxParserK_IIF, FilterExpressionSyntaxParserK_INDEXOF, FilterExpressionSyntaxParserK_ISDATE, FilterExpressionSyntaxParserK_ISINTEGER, FilterExpressionSyntaxParserK_ISGUID, FilterExpressionSyntaxParserK_ISNULL, FilterExpressionSyntaxParserK_ISNUMERIC, FilterExpressionSyntaxParserK_LASTINDEXOF, FilterExpressionSyntaxParserK_LEN, FilterExpressionSyntaxParserK_LOWER, FilterExpressionSyntaxParserK_MAXOF, FilterExpressionSyntaxParserK_MINOF, FilterExpressionSyntaxParserK_NOW, FilterExpressionSyntaxParserK_NTHINDEXOF, FilterExpressionSyntaxParserK_POWER, FilterExpressionSyntaxParserK_REGEXMATCH, FilterExpressionSyntaxParserK_REGEXVAL, FilterExpressionSyntaxParserK_REPLACE, FilterExpressionSyntaxParserK_REVERSE, FilterExpressionSyntaxParserK_ROUND, FilterExpressionSyntaxParserK_SQRT, FilterExpressionSyntaxParserK_SPLIT, FilterExpressionSyntaxParserK_STARTSWITH, FilterExpressionSyntaxParserK_STRCOUNT, FilterExpressionSyntaxParserK_STRCMP, FilterExpressionSyntaxParserK_SUBSTR, FilterExpressionSyntaxParserK_TRIM, FilterExpressionSyntaxParserK_TRIMLEFT, FilterExpressionSyntaxParserK_TRIMRIGHT, FilterExpressionSyntaxParserK_UPPER, FilterExpressionSyntaxParserK_UTCNOW:
		{
			p.SetState(188)
			p.FunctionExpression()
		}

	case FilterExpressionSyntaxParserT__2, FilterExpressionSyntaxParserT__3, FilterExpressionSyntaxParserT__6, FilterExpressionSyntaxParserT__7, FilterExpressionSyntaxParserK_NOT:
		{
			p.SetState(189)
			p.UnaryOperator()
		}
		{
			p.SetState(190)
			p.valueExpression(4)
		}

	case FilterExpressionSyntaxParserT__4:
		{
			p.SetState(192)
			p.Match(FilterExpressionSyntaxParserT__4)
		}
		{
			p.SetState(193)
			p.expression(0)
		}
		{
			p.SetState(194)
			p.Match(FilterExpressionSyntaxParserT__5)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(206)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				localctx = NewValueExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FilterExpressionSyntaxParserRULE_valueExpression)
				p.SetState(198)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(199)
					p.MathOperator()
				}
				{
					p.SetState(200)
					p.valueExpression(3)
				}

			case 2:
				localctx = NewValueExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FilterExpressionSyntaxParserRULE_valueExpression)
				p.SetState(202)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(203)
					p.BitwiseOperator()
				}
				{
					p.SetState(204)
					p.valueExpression(2)
				}

			}

		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
}

// INotOperatorContext is an interface to support dynamic dispatch.
type INotOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotOperatorContext differentiates from other interfaces.
	IsNotOperatorContext()
}

type NotOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotOperatorContext() *NotOperatorContext {
	var p = new(NotOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_notOperator
	return p
}

func (*NotOperatorContext) IsNotOperatorContext() {}

func NewNotOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotOperatorContext {
	var p = new(NotOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_notOperator

	return p
}

func (s *NotOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *NotOperatorContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_NOT, 0)
}

func (s *NotOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterNotOperator(s)
	}
}

func (s *NotOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitNotOperator(s)
	}
}

func (p *FilterExpressionSyntaxParser) NotOperator() (localctx INotOperatorContext) {
	localctx = NewNotOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FilterExpressionSyntaxParserRULE_notOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FilterExpressionSyntaxParserT__6 || _la == FilterExpressionSyntaxParserK_NOT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnaryOperatorContext is an interface to support dynamic dispatch.
type IUnaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOperatorContext differentiates from other interfaces.
	IsUnaryOperatorContext()
}

type UnaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorContext() *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_unaryOperator
	return p
}

func (*UnaryOperatorContext) IsUnaryOperatorContext() {}

func NewUnaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_unaryOperator

	return p
}

func (s *UnaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperatorContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_NOT, 0)
}

func (s *UnaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitUnaryOperator(s)
	}
}

func (p *FilterExpressionSyntaxParser) UnaryOperator() (localctx IUnaryOperatorContext) {
	localctx = NewUnaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FilterExpressionSyntaxParserRULE_unaryOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FilterExpressionSyntaxParserT__2)|(1<<FilterExpressionSyntaxParserT__3)|(1<<FilterExpressionSyntaxParserT__6)|(1<<FilterExpressionSyntaxParserT__7))) != 0) || _la == FilterExpressionSyntaxParserK_NOT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExactMatchModifierContext is an interface to support dynamic dispatch.
type IExactMatchModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExactMatchModifierContext differentiates from other interfaces.
	IsExactMatchModifierContext()
}

type ExactMatchModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExactMatchModifierContext() *ExactMatchModifierContext {
	var p = new(ExactMatchModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_exactMatchModifier
	return p
}

func (*ExactMatchModifierContext) IsExactMatchModifierContext() {}

func NewExactMatchModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExactMatchModifierContext {
	var p = new(ExactMatchModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_exactMatchModifier

	return p
}

func (s *ExactMatchModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ExactMatchModifierContext) K_BINARY() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_BINARY, 0)
}

func (s *ExactMatchModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExactMatchModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExactMatchModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterExactMatchModifier(s)
	}
}

func (s *ExactMatchModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitExactMatchModifier(s)
	}
}

func (p *FilterExpressionSyntaxParser) ExactMatchModifier() (localctx IExactMatchModifierContext) {
	localctx = NewExactMatchModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FilterExpressionSyntaxParserRULE_exactMatchModifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FilterExpressionSyntaxParserT__8 || _la == FilterExpressionSyntaxParserK_BINARY) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (p *FilterExpressionSyntaxParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FilterExpressionSyntaxParserRULE_comparisonOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FilterExpressionSyntaxParserT__8)|(1<<FilterExpressionSyntaxParserT__9)|(1<<FilterExpressionSyntaxParserT__10)|(1<<FilterExpressionSyntaxParserT__11)|(1<<FilterExpressionSyntaxParserT__12)|(1<<FilterExpressionSyntaxParserT__13)|(1<<FilterExpressionSyntaxParserT__14)|(1<<FilterExpressionSyntaxParserT__15)|(1<<FilterExpressionSyntaxParserT__16)|(1<<FilterExpressionSyntaxParserT__17))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILogicalOperatorContext is an interface to support dynamic dispatch.
type ILogicalOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalOperatorContext differentiates from other interfaces.
	IsLogicalOperatorContext()
}

type LogicalOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOperatorContext() *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_logicalOperator
	return p
}

func (*LogicalOperatorContext) IsLogicalOperatorContext() {}

func NewLogicalOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_logicalOperator

	return p
}

func (s *LogicalOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOperatorContext) K_AND() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_AND, 0)
}

func (s *LogicalOperatorContext) K_OR() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_OR, 0)
}

func (s *LogicalOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitLogicalOperator(s)
	}
}

func (p *FilterExpressionSyntaxParser) LogicalOperator() (localctx ILogicalOperatorContext) {
	localctx = NewLogicalOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FilterExpressionSyntaxParserRULE_logicalOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FilterExpressionSyntaxParserT__18)|(1<<FilterExpressionSyntaxParserT__19)|(1<<FilterExpressionSyntaxParserK_AND))) != 0) || _la == FilterExpressionSyntaxParserK_OR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBitwiseOperatorContext is an interface to support dynamic dispatch.
type IBitwiseOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBitwiseOperatorContext differentiates from other interfaces.
	IsBitwiseOperatorContext()
}

type BitwiseOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitwiseOperatorContext() *BitwiseOperatorContext {
	var p = new(BitwiseOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_bitwiseOperator
	return p
}

func (*BitwiseOperatorContext) IsBitwiseOperatorContext() {}

func NewBitwiseOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitwiseOperatorContext {
	var p = new(BitwiseOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_bitwiseOperator

	return p
}

func (s *BitwiseOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *BitwiseOperatorContext) K_XOR() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_XOR, 0)
}

func (s *BitwiseOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitwiseOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitwiseOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterBitwiseOperator(s)
	}
}

func (s *BitwiseOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitBitwiseOperator(s)
	}
}

func (p *FilterExpressionSyntaxParser) BitwiseOperator() (localctx IBitwiseOperatorContext) {
	localctx = NewBitwiseOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FilterExpressionSyntaxParserRULE_bitwiseOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FilterExpressionSyntaxParserT__20)|(1<<FilterExpressionSyntaxParserT__21)|(1<<FilterExpressionSyntaxParserT__22)|(1<<FilterExpressionSyntaxParserT__23)|(1<<FilterExpressionSyntaxParserT__24))) != 0) || _la == FilterExpressionSyntaxParserK_XOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMathOperatorContext is an interface to support dynamic dispatch.
type IMathOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathOperatorContext differentiates from other interfaces.
	IsMathOperatorContext()
}

type MathOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathOperatorContext() *MathOperatorContext {
	var p = new(MathOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_mathOperator
	return p
}

func (*MathOperatorContext) IsMathOperatorContext() {}

func NewMathOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathOperatorContext {
	var p = new(MathOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_mathOperator

	return p
}

func (s *MathOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *MathOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterMathOperator(s)
	}
}

func (s *MathOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitMathOperator(s)
	}
}

func (p *FilterExpressionSyntaxParser) MathOperator() (localctx IMathOperatorContext) {
	localctx = NewMathOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FilterExpressionSyntaxParserRULE_mathOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FilterExpressionSyntaxParserT__2)|(1<<FilterExpressionSyntaxParserT__3)|(1<<FilterExpressionSyntaxParserT__25)|(1<<FilterExpressionSyntaxParserT__26)|(1<<FilterExpressionSyntaxParserT__27))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunctionNameContext is an interface to support dynamic dispatch.
type IFunctionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionNameContext differentiates from other interfaces.
	IsFunctionNameContext()
}

type FunctionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionNameContext() *FunctionNameContext {
	var p = new(FunctionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_functionName
	return p
}

func (*FunctionNameContext) IsFunctionNameContext() {}

func NewFunctionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionNameContext {
	var p = new(FunctionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_functionName

	return p
}

func (s *FunctionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionNameContext) K_ABS() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_ABS, 0)
}

func (s *FunctionNameContext) K_CEILING() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_CEILING, 0)
}

func (s *FunctionNameContext) K_COALESCE() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_COALESCE, 0)
}

func (s *FunctionNameContext) K_CONVERT() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_CONVERT, 0)
}

func (s *FunctionNameContext) K_CONTAINS() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_CONTAINS, 0)
}

func (s *FunctionNameContext) K_DATEADD() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_DATEADD, 0)
}

func (s *FunctionNameContext) K_DATEDIFF() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_DATEDIFF, 0)
}

func (s *FunctionNameContext) K_DATEPART() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_DATEPART, 0)
}

func (s *FunctionNameContext) K_ENDSWITH() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_ENDSWITH, 0)
}

func (s *FunctionNameContext) K_FLOOR() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_FLOOR, 0)
}

func (s *FunctionNameContext) K_IIF() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_IIF, 0)
}

func (s *FunctionNameContext) K_INDEXOF() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_INDEXOF, 0)
}

func (s *FunctionNameContext) K_ISDATE() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_ISDATE, 0)
}

func (s *FunctionNameContext) K_ISINTEGER() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_ISINTEGER, 0)
}

func (s *FunctionNameContext) K_ISGUID() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_ISGUID, 0)
}

func (s *FunctionNameContext) K_ISNULL() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_ISNULL, 0)
}

func (s *FunctionNameContext) K_ISNUMERIC() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_ISNUMERIC, 0)
}

func (s *FunctionNameContext) K_LASTINDEXOF() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_LASTINDEXOF, 0)
}

func (s *FunctionNameContext) K_LEN() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_LEN, 0)
}

func (s *FunctionNameContext) K_LOWER() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_LOWER, 0)
}

func (s *FunctionNameContext) K_MAXOF() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_MAXOF, 0)
}

func (s *FunctionNameContext) K_MINOF() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_MINOF, 0)
}

func (s *FunctionNameContext) K_NOW() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_NOW, 0)
}

func (s *FunctionNameContext) K_NTHINDEXOF() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_NTHINDEXOF, 0)
}

func (s *FunctionNameContext) K_POWER() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_POWER, 0)
}

func (s *FunctionNameContext) K_REGEXMATCH() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_REGEXMATCH, 0)
}

func (s *FunctionNameContext) K_REGEXVAL() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_REGEXVAL, 0)
}

func (s *FunctionNameContext) K_REPLACE() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_REPLACE, 0)
}

func (s *FunctionNameContext) K_REVERSE() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_REVERSE, 0)
}

func (s *FunctionNameContext) K_ROUND() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_ROUND, 0)
}

func (s *FunctionNameContext) K_SPLIT() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_SPLIT, 0)
}

func (s *FunctionNameContext) K_SQRT() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_SQRT, 0)
}

func (s *FunctionNameContext) K_STARTSWITH() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_STARTSWITH, 0)
}

func (s *FunctionNameContext) K_STRCOUNT() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_STRCOUNT, 0)
}

func (s *FunctionNameContext) K_STRCMP() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_STRCMP, 0)
}

func (s *FunctionNameContext) K_SUBSTR() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_SUBSTR, 0)
}

func (s *FunctionNameContext) K_TRIM() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_TRIM, 0)
}

func (s *FunctionNameContext) K_TRIMLEFT() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_TRIMLEFT, 0)
}

func (s *FunctionNameContext) K_TRIMRIGHT() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_TRIMRIGHT, 0)
}

func (s *FunctionNameContext) K_UPPER() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_UPPER, 0)
}

func (s *FunctionNameContext) K_UTCNOW() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_UTCNOW, 0)
}

func (s *FunctionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterFunctionName(s)
	}
}

func (s *FunctionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitFunctionName(s)
	}
}

func (p *FilterExpressionSyntaxParser) FunctionName() (localctx IFunctionNameContext) {
	localctx = NewFunctionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FilterExpressionSyntaxParserRULE_functionName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-29)&-(0x1f+1)) == 0 && ((1<<uint((_la-29)))&((1<<(FilterExpressionSyntaxParserK_ABS-29))|(1<<(FilterExpressionSyntaxParserK_CEILING-29))|(1<<(FilterExpressionSyntaxParserK_COALESCE-29))|(1<<(FilterExpressionSyntaxParserK_CONVERT-29))|(1<<(FilterExpressionSyntaxParserK_CONTAINS-29))|(1<<(FilterExpressionSyntaxParserK_DATEADD-29))|(1<<(FilterExpressionSyntaxParserK_DATEDIFF-29))|(1<<(FilterExpressionSyntaxParserK_DATEPART-29))|(1<<(FilterExpressionSyntaxParserK_ENDSWITH-29))|(1<<(FilterExpressionSyntaxParserK_FLOOR-29))|(1<<(FilterExpressionSyntaxParserK_IIF-29))|(1<<(FilterExpressionSyntaxParserK_INDEXOF-29))|(1<<(FilterExpressionSyntaxParserK_ISDATE-29))|(1<<(FilterExpressionSyntaxParserK_ISINTEGER-29))|(1<<(FilterExpressionSyntaxParserK_ISGUID-29))|(1<<(FilterExpressionSyntaxParserK_ISNULL-29))|(1<<(FilterExpressionSyntaxParserK_ISNUMERIC-29))|(1<<(FilterExpressionSyntaxParserK_LASTINDEXOF-29))|(1<<(FilterExpressionSyntaxParserK_LEN-29))|(1<<(FilterExpressionSyntaxParserK_LOWER-29))|(1<<(FilterExpressionSyntaxParserK_MAXOF-29))|(1<<(FilterExpressionSyntaxParserK_MINOF-29)))) != 0) || (((_la-61)&-(0x1f+1)) == 0 && ((1<<uint((_la-61)))&((1<<(FilterExpressionSyntaxParserK_NOW-61))|(1<<(FilterExpressionSyntaxParserK_NTHINDEXOF-61))|(1<<(FilterExpressionSyntaxParserK_POWER-61))|(1<<(FilterExpressionSyntaxParserK_REGEXMATCH-61))|(1<<(FilterExpressionSyntaxParserK_REGEXVAL-61))|(1<<(FilterExpressionSyntaxParserK_REPLACE-61))|(1<<(FilterExpressionSyntaxParserK_REVERSE-61))|(1<<(FilterExpressionSyntaxParserK_ROUND-61))|(1<<(FilterExpressionSyntaxParserK_SQRT-61))|(1<<(FilterExpressionSyntaxParserK_SPLIT-61))|(1<<(FilterExpressionSyntaxParserK_STARTSWITH-61))|(1<<(FilterExpressionSyntaxParserK_STRCOUNT-61))|(1<<(FilterExpressionSyntaxParserK_STRCMP-61))|(1<<(FilterExpressionSyntaxParserK_SUBSTR-61))|(1<<(FilterExpressionSyntaxParserK_TRIM-61))|(1<<(FilterExpressionSyntaxParserK_TRIMLEFT-61))|(1<<(FilterExpressionSyntaxParserK_TRIMRIGHT-61))|(1<<(FilterExpressionSyntaxParserK_UPPER-61))|(1<<(FilterExpressionSyntaxParserK_UTCNOW-61)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunctionExpressionContext is an interface to support dynamic dispatch.
type IFunctionExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionExpressionContext differentiates from other interfaces.
	IsFunctionExpressionContext()
}

type FunctionExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionExpressionContext() *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_functionExpression
	return p
}

func (*FunctionExpressionContext) IsFunctionExpressionContext() {}

func NewFunctionExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_functionExpression

	return p
}

func (s *FunctionExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionExpressionContext) FunctionName() IFunctionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionNameContext)
}

func (s *FunctionExpressionContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

func (p *FilterExpressionSyntaxParser) FunctionExpression() (localctx IFunctionExpressionContext) {
	localctx = NewFunctionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FilterExpressionSyntaxParserRULE_functionExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.FunctionName()
	}
	{
		p.SetState(228)
		p.Match(FilterExpressionSyntaxParserT__4)
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FilterExpressionSyntaxParserT__2)|(1<<FilterExpressionSyntaxParserT__3)|(1<<FilterExpressionSyntaxParserT__4)|(1<<FilterExpressionSyntaxParserT__6)|(1<<FilterExpressionSyntaxParserT__7)|(1<<FilterExpressionSyntaxParserK_ABS))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(FilterExpressionSyntaxParserK_CEILING-34))|(1<<(FilterExpressionSyntaxParserK_COALESCE-34))|(1<<(FilterExpressionSyntaxParserK_CONVERT-34))|(1<<(FilterExpressionSyntaxParserK_CONTAINS-34))|(1<<(FilterExpressionSyntaxParserK_DATEADD-34))|(1<<(FilterExpressionSyntaxParserK_DATEDIFF-34))|(1<<(FilterExpressionSyntaxParserK_DATEPART-34))|(1<<(FilterExpressionSyntaxParserK_ENDSWITH-34))|(1<<(FilterExpressionSyntaxParserK_FLOOR-34))|(1<<(FilterExpressionSyntaxParserK_IIF-34))|(1<<(FilterExpressionSyntaxParserK_INDEXOF-34))|(1<<(FilterExpressionSyntaxParserK_ISDATE-34))|(1<<(FilterExpressionSyntaxParserK_ISINTEGER-34))|(1<<(FilterExpressionSyntaxParserK_ISGUID-34))|(1<<(FilterExpressionSyntaxParserK_ISNULL-34))|(1<<(FilterExpressionSyntaxParserK_ISNUMERIC-34))|(1<<(FilterExpressionSyntaxParserK_LASTINDEXOF-34))|(1<<(FilterExpressionSyntaxParserK_LEN-34))|(1<<(FilterExpressionSyntaxParserK_LOWER-34))|(1<<(FilterExpressionSyntaxParserK_MAXOF-34))|(1<<(FilterExpressionSyntaxParserK_MINOF-34))|(1<<(FilterExpressionSyntaxParserK_NOT-34))|(1<<(FilterExpressionSyntaxParserK_NOW-34))|(1<<(FilterExpressionSyntaxParserK_NTHINDEXOF-34))|(1<<(FilterExpressionSyntaxParserK_NULL-34)))) != 0) || (((_la-66)&-(0x1f+1)) == 0 && ((1<<uint((_la-66)))&((1<<(FilterExpressionSyntaxParserK_POWER-66))|(1<<(FilterExpressionSyntaxParserK_REGEXMATCH-66))|(1<<(FilterExpressionSyntaxParserK_REGEXVAL-66))|(1<<(FilterExpressionSyntaxParserK_REPLACE-66))|(1<<(FilterExpressionSyntaxParserK_REVERSE-66))|(1<<(FilterExpressionSyntaxParserK_ROUND-66))|(1<<(FilterExpressionSyntaxParserK_SQRT-66))|(1<<(FilterExpressionSyntaxParserK_SPLIT-66))|(1<<(FilterExpressionSyntaxParserK_STARTSWITH-66))|(1<<(FilterExpressionSyntaxParserK_STRCOUNT-66))|(1<<(FilterExpressionSyntaxParserK_STRCMP-66))|(1<<(FilterExpressionSyntaxParserK_SUBSTR-66))|(1<<(FilterExpressionSyntaxParserK_TRIM-66))|(1<<(FilterExpressionSyntaxParserK_TRIMLEFT-66))|(1<<(FilterExpressionSyntaxParserK_TRIMRIGHT-66))|(1<<(FilterExpressionSyntaxParserK_UPPER-66))|(1<<(FilterExpressionSyntaxParserK_UTCNOW-66))|(1<<(FilterExpressionSyntaxParserBOOLEAN_LITERAL-66))|(1<<(FilterExpressionSyntaxParserIDENTIFIER-66))|(1<<(FilterExpressionSyntaxParserINTEGER_LITERAL-66))|(1<<(FilterExpressionSyntaxParserNUMERIC_LITERAL-66))|(1<<(FilterExpressionSyntaxParserGUID_LITERAL-66))|(1<<(FilterExpressionSyntaxParserSTRING_LITERAL-66))|(1<<(FilterExpressionSyntaxParserDATETIME_LITERAL-66)))) != 0) {
		{
			p.SetState(229)
			p.ExpressionList()
		}

	}
	{
		p.SetState(232)
		p.Match(FilterExpressionSyntaxParserT__5)
	}

	return localctx
}

// ILiteralValueContext is an interface to support dynamic dispatch.
type ILiteralValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralValueContext differentiates from other interfaces.
	IsLiteralValueContext()
}

type LiteralValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralValueContext() *LiteralValueContext {
	var p = new(LiteralValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_literalValue
	return p
}

func (*LiteralValueContext) IsLiteralValueContext() {}

func NewLiteralValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralValueContext {
	var p = new(LiteralValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_literalValue

	return p
}

func (s *LiteralValueContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralValueContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserINTEGER_LITERAL, 0)
}

func (s *LiteralValueContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserNUMERIC_LITERAL, 0)
}

func (s *LiteralValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserSTRING_LITERAL, 0)
}

func (s *LiteralValueContext) DATETIME_LITERAL() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserDATETIME_LITERAL, 0)
}

func (s *LiteralValueContext) GUID_LITERAL() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserGUID_LITERAL, 0)
}

func (s *LiteralValueContext) BOOLEAN_LITERAL() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserBOOLEAN_LITERAL, 0)
}

func (s *LiteralValueContext) K_NULL() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserK_NULL, 0)
}

func (s *LiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterLiteralValue(s)
	}
}

func (s *LiteralValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitLiteralValue(s)
	}
}

func (p *FilterExpressionSyntaxParser) LiteralValue() (localctx ILiteralValueContext) {
	localctx = NewLiteralValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FilterExpressionSyntaxParserRULE_literalValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-63)&-(0x1f+1)) == 0 && ((1<<uint((_la-63)))&((1<<(FilterExpressionSyntaxParserK_NULL-63))|(1<<(FilterExpressionSyntaxParserBOOLEAN_LITERAL-63))|(1<<(FilterExpressionSyntaxParserINTEGER_LITERAL-63))|(1<<(FilterExpressionSyntaxParserNUMERIC_LITERAL-63))|(1<<(FilterExpressionSyntaxParserGUID_LITERAL-63))|(1<<(FilterExpressionSyntaxParserSTRING_LITERAL-63))|(1<<(FilterExpressionSyntaxParserDATETIME_LITERAL-63)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserIDENTIFIER, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (p *FilterExpressionSyntaxParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FilterExpressionSyntaxParserRULE_tableName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(FilterExpressionSyntaxParserIDENTIFIER)
	}

	return localctx
}

// IColumnNameContext is an interface to support dynamic dispatch.
type IColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnNameContext differentiates from other interfaces.
	IsColumnNameContext()
}

type ColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnNameContext() *ColumnNameContext {
	var p = new(ColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_columnName
	return p
}

func (*ColumnNameContext) IsColumnNameContext() {}

func NewColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNameContext {
	var p = new(ColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_columnName

	return p
}

func (s *ColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserIDENTIFIER, 0)
}

func (s *ColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterColumnName(s)
	}
}

func (s *ColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitColumnName(s)
	}
}

func (p *FilterExpressionSyntaxParser) ColumnName() (localctx IColumnNameContext) {
	localctx = NewColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FilterExpressionSyntaxParserRULE_columnName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(FilterExpressionSyntaxParserIDENTIFIER)
	}

	return localctx
}

// IOrderByColumnNameContext is an interface to support dynamic dispatch.
type IOrderByColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderByColumnNameContext differentiates from other interfaces.
	IsOrderByColumnNameContext()
}

type OrderByColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByColumnNameContext() *OrderByColumnNameContext {
	var p = new(OrderByColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionSyntaxParserRULE_orderByColumnName
	return p
}

func (*OrderByColumnNameContext) IsOrderByColumnNameContext() {}

func NewOrderByColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByColumnNameContext {
	var p = new(OrderByColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionSyntaxParserRULE_orderByColumnName

	return p
}

func (s *OrderByColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByColumnNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSyntaxParserIDENTIFIER, 0)
}

func (s *OrderByColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByColumnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.EnterOrderByColumnName(s)
	}
}

func (s *OrderByColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterExpressionSyntaxListener); ok {
		listenerT.ExitOrderByColumnName(s)
	}
}

func (p *FilterExpressionSyntaxParser) OrderByColumnName() (localctx IOrderByColumnNameContext) {
	localctx = NewOrderByColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, FilterExpressionSyntaxParserRULE_orderByColumnName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(FilterExpressionSyntaxParserIDENTIFIER)
	}

	return localctx
}

func (p *FilterExpressionSyntaxParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 10:
		var t *PredicateExpressionContext = nil
		if localctx != nil {
			t = localctx.(*PredicateExpressionContext)
		}
		return p.PredicateExpression_Sempred(t, predIndex)

	case 11:
		var t *ValueExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ValueExpressionContext)
		}
		return p.ValueExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FilterExpressionSyntaxParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FilterExpressionSyntaxParser) PredicateExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FilterExpressionSyntaxParser) ValueExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
