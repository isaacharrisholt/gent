package python

import (
	"fmt"
	"github.com/tree-sitter/go-tree-sitter"
	"slices"
)

type SyntaxKind = string

var (
	SyntaxKind_AliasedImport           SyntaxKind = "aliased_import"
	SyntaxKind_ArgumentList            SyntaxKind = "argument_list"
	SyntaxKind_AsPattern               SyntaxKind = "as_pattern"
	SyntaxKind_AssertStatement         SyntaxKind = "assert_statement"
	SyntaxKind_Assignment              SyntaxKind = "assignment"
	SyntaxKind_Attribute               SyntaxKind = "attribute"
	SyntaxKind_AugmentedAssignment     SyntaxKind = "augmented_assignment"
	SyntaxKind_Await                   SyntaxKind = "await"
	SyntaxKind_BinaryOperator          SyntaxKind = "binary_operator"
	SyntaxKind_Block                   SyntaxKind = "block"
	SyntaxKind_BooleanOperator         SyntaxKind = "boolean_operator"
	SyntaxKind_BreakStatement          SyntaxKind = "break_statement"
	SyntaxKind_Call                    SyntaxKind = "call"
	SyntaxKind_CaseClause              SyntaxKind = "case_clause"
	SyntaxKind_CasePattern             SyntaxKind = "case_pattern"
	SyntaxKind_Chevron                 SyntaxKind = "chevron"
	SyntaxKind_ClassDefinition         SyntaxKind = "class_definition"
	SyntaxKind_ClassPattern            SyntaxKind = "class_pattern"
	SyntaxKind_ComparisonOperator      SyntaxKind = "comparison_operator"
	SyntaxKind_ComplexPattern          SyntaxKind = "complex_pattern"
	SyntaxKind_ConcatenatedString      SyntaxKind = "concatenated_string"
	SyntaxKind_ConditionalExpression   SyntaxKind = "conditional_expression"
	SyntaxKind_ConstrainedType         SyntaxKind = "constrained_type"
	SyntaxKind_ContinueStatement       SyntaxKind = "continue_statement"
	SyntaxKind_DecoratedDefinition     SyntaxKind = "decorated_definition"
	SyntaxKind_Decorator               SyntaxKind = "decorator"
	SyntaxKind_DefaultParameter        SyntaxKind = "default_parameter"
	SyntaxKind_DeleteStatement         SyntaxKind = "delete_statement"
	SyntaxKind_DictPattern             SyntaxKind = "dict_pattern"
	SyntaxKind_Dictionary              SyntaxKind = "dictionary"
	SyntaxKind_DictionaryComprehension SyntaxKind = "dictionary_comprehension"
	SyntaxKind_DictionarySplat         SyntaxKind = "dictionary_splat"
	SyntaxKind_DictionarySplatPattern  SyntaxKind = "dictionary_splat_pattern"
	SyntaxKind_DottedName              SyntaxKind = "dotted_name"
	SyntaxKind_ElifClause              SyntaxKind = "elif_clause"
	SyntaxKind_ElseClause              SyntaxKind = "else_clause"
	SyntaxKind_ExceptClause            SyntaxKind = "except_clause"
	SyntaxKind_ExceptGroupClause       SyntaxKind = "except_group_clause"
	SyntaxKind_ExecStatement           SyntaxKind = "exec_statement"
	SyntaxKind_ExpressionList          SyntaxKind = "expression_list"
	SyntaxKind_ExpressionStatement     SyntaxKind = "expression_statement"
	SyntaxKind_FinallyClause           SyntaxKind = "finally_clause"
	SyntaxKind_ForInClause             SyntaxKind = "for_in_clause"
	SyntaxKind_ForStatement            SyntaxKind = "for_statement"
	SyntaxKind_FormatExpression        SyntaxKind = "format_expression"
	SyntaxKind_FormatSpecifier         SyntaxKind = "format_specifier"
	SyntaxKind_FunctionDefinition      SyntaxKind = "function_definition"
	SyntaxKind_FutureImportStatement   SyntaxKind = "future_import_statement"
	SyntaxKind_GeneratorExpression     SyntaxKind = "generator_expression"
	SyntaxKind_GenericType             SyntaxKind = "generic_type"
	SyntaxKind_GlobalStatement         SyntaxKind = "global_statement"
	SyntaxKind_IfClause                SyntaxKind = "if_clause"
	SyntaxKind_IfStatement             SyntaxKind = "if_statement"
	SyntaxKind_ImportFromStatement     SyntaxKind = "import_from_statement"
	SyntaxKind_ImportPrefix            SyntaxKind = "import_prefix"
	SyntaxKind_ImportStatement         SyntaxKind = "import_statement"
	SyntaxKind_Interpolation           SyntaxKind = "interpolation"
	SyntaxKind_KeywordArgument         SyntaxKind = "keyword_argument"
	SyntaxKind_KeywordPattern          SyntaxKind = "keyword_pattern"
	SyntaxKind_KeywordSeparator        SyntaxKind = "keyword_separator"
	SyntaxKind_Lambda                  SyntaxKind = "lambda"
	SyntaxKind_LambdaParameters        SyntaxKind = "lambda_parameters"
	SyntaxKind_List                    SyntaxKind = "list"
	SyntaxKind_ListComprehension       SyntaxKind = "list_comprehension"
	SyntaxKind_ListPattern             SyntaxKind = "list_pattern"
	SyntaxKind_ListSplat               SyntaxKind = "list_splat"
	SyntaxKind_ListSplatPattern        SyntaxKind = "list_splat_pattern"
	SyntaxKind_MatchStatement          SyntaxKind = "match_statement"
	SyntaxKind_MemberType              SyntaxKind = "member_type"
	SyntaxKind_Module                  SyntaxKind = "module"
	SyntaxKind_NamedExpression         SyntaxKind = "named_expression"
	SyntaxKind_NonlocalStatement       SyntaxKind = "nonlocal_statement"
	SyntaxKind_NotOperator             SyntaxKind = "not_operator"
	SyntaxKind_Pair                    SyntaxKind = "pair"
	SyntaxKind_Parameters              SyntaxKind = "parameters"
	SyntaxKind_ParenthesizedExpression SyntaxKind = "parenthesized_expression"
	SyntaxKind_ParenthesizedListSplat  SyntaxKind = "parenthesized_list_splat"
	SyntaxKind_PassStatement           SyntaxKind = "pass_statement"
	SyntaxKind_PatternList             SyntaxKind = "pattern_list"
	SyntaxKind_PositionalSeparator     SyntaxKind = "positional_separator"
	SyntaxKind_PrintStatement          SyntaxKind = "print_statement"
	SyntaxKind_RaiseStatement          SyntaxKind = "raise_statement"
	SyntaxKind_RelativeImport          SyntaxKind = "relative_import"
	SyntaxKind_ReturnStatement         SyntaxKind = "return_statement"
	SyntaxKind_Set                     SyntaxKind = "set"
	SyntaxKind_SetComprehension        SyntaxKind = "set_comprehension"
	SyntaxKind_Slice                   SyntaxKind = "slice"
	SyntaxKind_SplatPattern            SyntaxKind = "splat_pattern"
	SyntaxKind_SplatType               SyntaxKind = "splat_type"
	SyntaxKind_String                  SyntaxKind = "string"
	SyntaxKind_StringContent           SyntaxKind = "string_content"
	SyntaxKind_Subscript               SyntaxKind = "subscript"
	SyntaxKind_TryStatement            SyntaxKind = "try_statement"
	SyntaxKind_Tuple                   SyntaxKind = "tuple"
	SyntaxKind_TuplePattern            SyntaxKind = "tuple_pattern"
	SyntaxKind_Type                    SyntaxKind = "type"
	SyntaxKind_TypeAliasStatement      SyntaxKind = "type_alias_statement"
	SyntaxKind_TypeParameter           SyntaxKind = "type_parameter"
	SyntaxKind_TypedDefaultParameter   SyntaxKind = "typed_default_parameter"
	SyntaxKind_TypedParameter          SyntaxKind = "typed_parameter"
	SyntaxKind_UnaryOperator           SyntaxKind = "unary_operator"
	SyntaxKind_UnionPattern            SyntaxKind = "union_pattern"
	SyntaxKind_UnionType               SyntaxKind = "union_type"
	SyntaxKind_WhileStatement          SyntaxKind = "while_statement"
	SyntaxKind_WildcardImport          SyntaxKind = "wildcard_import"
	SyntaxKind_WithClause              SyntaxKind = "with_clause"
	SyntaxKind_WithItem                SyntaxKind = "with_item"
	SyntaxKind_WithStatement           SyntaxKind = "with_statement"
	SyntaxKind_Yield                   SyntaxKind = "yield"
	SyntaxKind_Comment                 SyntaxKind = "comment"
	SyntaxKind_Ellipsis                SyntaxKind = "ellipsis"
	SyntaxKind_EscapeInterpolation     SyntaxKind = "escape_interpolation"
	SyntaxKind_EscapeSequence          SyntaxKind = "escape_sequence"
	SyntaxKind_False                   SyntaxKind = "false"
	SyntaxKind_Float                   SyntaxKind = "float"
	SyntaxKind_Identifier              SyntaxKind = "identifier"
	SyntaxKind_Integer                 SyntaxKind = "integer"
	SyntaxKind_LineContinuation        SyntaxKind = "line_continuation"
	SyntaxKind_None                    SyntaxKind = "none"
	SyntaxKind_StringEnd               SyntaxKind = "string_end"
	SyntaxKind_StringStart             SyntaxKind = "string_start"
	SyntaxKind_True                    SyntaxKind = "true"
	SyntaxKind_TypeConversion          SyntaxKind = "type_conversion"
	SyntaxKind_Unnamed_IsSpaceNot      SyntaxKind = "is not"
	SyntaxKind_Unnamed_NotSpaceIn      SyntaxKind = "not in"
	SyntaxKind_Unnamed_NotEq           SyntaxKind = "!="
	SyntaxKind_Unnamed_Mod             SyntaxKind = "%"
	SyntaxKind_Unnamed_ModEq           SyntaxKind = "%="
	SyntaxKind_Unnamed_Ampersand       SyntaxKind = "&"
	SyntaxKind_Unnamed_AmpersandEq     SyntaxKind = "&="
	SyntaxKind_Unnamed_LParen          SyntaxKind = "("
	SyntaxKind_Unnamed_RParen          SyntaxKind = ")"
	SyntaxKind_Unnamed_Mul             SyntaxKind = "*"
	SyntaxKind_Unnamed_MulMul          SyntaxKind = "**"
	SyntaxKind_Unnamed_MulMulEq        SyntaxKind = "**="
	SyntaxKind_Unnamed_MulEq           SyntaxKind = "*="
	SyntaxKind_Unnamed_Add             SyntaxKind = "+"
	SyntaxKind_Unnamed_AddEq           SyntaxKind = "+="
	SyntaxKind_Unnamed_Comma           SyntaxKind = ","
	SyntaxKind_Unnamed_Sub             SyntaxKind = "-"
	SyntaxKind_Unnamed_SubEq           SyntaxKind = "-="
	SyntaxKind_Unnamed_SubGt           SyntaxKind = "->"
	SyntaxKind_Unnamed_Dot             SyntaxKind = "."
	SyntaxKind_Unnamed_Div             SyntaxKind = "/"
	SyntaxKind_Unnamed_DivDiv          SyntaxKind = "//"
	SyntaxKind_Unnamed_DivDivEq        SyntaxKind = "//="
	SyntaxKind_Unnamed_DivEq           SyntaxKind = "/="
	SyntaxKind_Unnamed_Colon           SyntaxKind = ":"
	SyntaxKind_Unnamed_ColonEq         SyntaxKind = ":="
	SyntaxKind_Unnamed_Semicolon       SyntaxKind = ";"
	SyntaxKind_Unnamed_Lt              SyntaxKind = "<"
	SyntaxKind_Unnamed_LtLt            SyntaxKind = "<<"
	SyntaxKind_Unnamed_LtLtEq          SyntaxKind = "<<="
	SyntaxKind_Unnamed_LtEq            SyntaxKind = "<="
	SyntaxKind_Unnamed_LtGt            SyntaxKind = "<>"
	SyntaxKind_Unnamed_Eq              SyntaxKind = "="
	SyntaxKind_Unnamed_EqEq            SyntaxKind = "=="
	SyntaxKind_Unnamed_Gt              SyntaxKind = ">"
	SyntaxKind_Unnamed_GtEq            SyntaxKind = ">="
	SyntaxKind_Unnamed_GtGt            SyntaxKind = ">>"
	SyntaxKind_Unnamed_GtGtEq          SyntaxKind = ">>="
	SyntaxKind_Unnamed_At              SyntaxKind = "@"
	SyntaxKind_Unnamed_AtEq            SyntaxKind = "@="
	SyntaxKind_Unnamed_LBracket        SyntaxKind = "["
	SyntaxKind_Unnamed_Backslash       SyntaxKind = "\\"
	SyntaxKind_Unnamed_RBracket        SyntaxKind = "]"
	SyntaxKind_Unnamed_BitXor          SyntaxKind = "^"
	SyntaxKind_Unnamed_BitXorEq        SyntaxKind = "^="
	SyntaxKind_Unnamed_Underscore      SyntaxKind = "_"
	SyntaxKind_Unnamed_Future          SyntaxKind = "__future__"
	SyntaxKind_Unnamed_And             SyntaxKind = "and"
	SyntaxKind_Unnamed_As              SyntaxKind = "as"
	SyntaxKind_Unnamed_Assert          SyntaxKind = "assert"
	SyntaxKind_Unnamed_Async           SyntaxKind = "async"
	SyntaxKind_Unnamed_Await           SyntaxKind = "await"
	SyntaxKind_Unnamed_Break           SyntaxKind = "break"
	SyntaxKind_Unnamed_Case            SyntaxKind = "case"
	SyntaxKind_Unnamed_Class           SyntaxKind = "class"
	SyntaxKind_Unnamed_Continue        SyntaxKind = "continue"
	SyntaxKind_Unnamed_Def             SyntaxKind = "def"
	SyntaxKind_Unnamed_Del             SyntaxKind = "del"
	SyntaxKind_Unnamed_Elif            SyntaxKind = "elif"
	SyntaxKind_Unnamed_Else            SyntaxKind = "else"
	SyntaxKind_Unnamed_Except          SyntaxKind = "except"
	SyntaxKind_Unnamed_ExceptMul       SyntaxKind = "except*"
	SyntaxKind_Unnamed_Exec            SyntaxKind = "exec"
	SyntaxKind_Unnamed_Finally         SyntaxKind = "finally"
	SyntaxKind_Unnamed_For             SyntaxKind = "for"
	SyntaxKind_Unnamed_From            SyntaxKind = "from"
	SyntaxKind_Unnamed_Global          SyntaxKind = "global"
	SyntaxKind_Unnamed_If              SyntaxKind = "if"
	SyntaxKind_Unnamed_Import          SyntaxKind = "import"
	SyntaxKind_Unnamed_In              SyntaxKind = "in"
	SyntaxKind_Unnamed_Is              SyntaxKind = "is"
	SyntaxKind_Unnamed_Lambda          SyntaxKind = "lambda"
	SyntaxKind_Unnamed_Match           SyntaxKind = "match"
	SyntaxKind_Unnamed_Nonlocal        SyntaxKind = "nonlocal"
	SyntaxKind_Unnamed_Not             SyntaxKind = "not"
	SyntaxKind_Unnamed_Or              SyntaxKind = "or"
	SyntaxKind_Unnamed_Pass            SyntaxKind = "pass"
	SyntaxKind_Unnamed_Print           SyntaxKind = "print"
	SyntaxKind_Unnamed_Raise           SyntaxKind = "raise"
	SyntaxKind_Unnamed_Return          SyntaxKind = "return"
	SyntaxKind_Unnamed_Try             SyntaxKind = "try"
	SyntaxKind_Unnamed_Type            SyntaxKind = "type"
	SyntaxKind_Unnamed_While           SyntaxKind = "while"
	SyntaxKind_Unnamed_With            SyntaxKind = "with"
	SyntaxKind_Unnamed_Yield           SyntaxKind = "yield"
	SyntaxKind_Unnamed_LBrace          SyntaxKind = "{"
	SyntaxKind_Unnamed_Bar             SyntaxKind = "|"
	SyntaxKind_Unnamed_BarEq           SyntaxKind = "|="
	SyntaxKind_Unnamed_RBrace          SyntaxKind = "}"
	SyntaxKind_Unnamed_BitNot          SyntaxKind = "~"
	SyntaxKind_CompoundStatement       SyntaxKind = "_compound_statement"
	SyntaxKind_SimpleStatement         SyntaxKind = "_simple_statement"
	SyntaxKind_Expression              SyntaxKind = "expression"
	SyntaxKind_Parameter               SyntaxKind = "parameter"
	SyntaxKind_Pattern                 SyntaxKind = "pattern"
	SyntaxKind_PrimaryExpression       SyntaxKind = "primary_expression"
)

type AliasedImport struct {
	tree_sitter.Node
}

func (a *AliasedImport) Alias() (*Identifier, error) {
	child := a.Node.ChildByFieldName("alias")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "aliased_import", "alias")
	}
	return &Identifier{Node: *child}, nil
}
func (a *AliasedImport) Name() (*DottedName, error) {
	child := a.Node.ChildByFieldName("name")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "aliased_import", "name")
	}
	return &DottedName{Node: *child}, nil
}
func NewAliasedImport(node *tree_sitter.Node) (*AliasedImport, error) {
	if node.Kind() != "aliased_import" {
		return nil, fmt.Errorf("Node is not a %s", "aliased_import")
	}
	return &AliasedImport{Node: *node}, nil
}

type ArgumentList struct {
	tree_sitter.Node
}

func (a *ArgumentList) TypedChildren(cursor *tree_sitter.TreeCursor) []dictionarySplat_expression_keywordArgument_listSplat_parenthesizedExpression {
	children := a.Node.Children(cursor)
	output := []dictionarySplat_expression_keywordArgument_listSplat_parenthesizedExpression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, dictionarySplat_expression_keywordArgument_listSplat_parenthesizedExpression{Node: child})
		}
	}
	return output
}
func NewArgumentList(node *tree_sitter.Node) (*ArgumentList, error) {
	if node.Kind() != "argument_list" {
		return nil, fmt.Errorf("Node is not a %s", "argument_list")
	}
	return &ArgumentList{Node: *node}, nil
}

type AsPattern struct {
	tree_sitter.Node
}

func (a *AsPattern) Alias() (*Unknown__asPatternTarget, error) {
	child := a.Node.ChildByFieldName("alias")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "as_pattern", "alias")
	}
	return &Unknown__asPatternTarget{Node: *child}, nil
}
func (a *AsPattern) TypedChildren(cursor *tree_sitter.TreeCursor) []casePattern_expression_identifier {
	children := a.Node.Children(cursor)
	output := []casePattern_expression_identifier{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, casePattern_expression_identifier{Node: child})
		}
	}
	return output
}
func NewAsPattern(node *tree_sitter.Node) (*AsPattern, error) {
	if node.Kind() != "as_pattern" {
		return nil, fmt.Errorf("Node is not a %s", "as_pattern")
	}
	return &AsPattern{Node: *node}, nil
}

type AssertStatement struct {
	tree_sitter.Node
}

func (a *AssertStatement) TypedChildren(cursor *tree_sitter.TreeCursor) []Expression {
	children := a.Node.Children(cursor)
	output := []Expression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Expression{Node: child})
		}
	}
	return output
}
func NewAssertStatement(node *tree_sitter.Node) (*AssertStatement, error) {
	if node.Kind() != "assert_statement" {
		return nil, fmt.Errorf("Node is not a %s", "assert_statement")
	}
	return &AssertStatement{Node: *node}, nil
}

type Assignment struct {
	tree_sitter.Node
}

func (a *Assignment) Left() (*pattern_patternList, error) {
	child := a.Node.ChildByFieldName("left")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "assignment", "left")
	}
	return &pattern_patternList{Node: *child}, nil
}
func (a *Assignment) Right() (*assignment_augmentedAssignment_expression_expressionList_patternList_yield, error) {
	child := a.Node.ChildByFieldName("right")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "assignment", "right")
	}
	return &assignment_augmentedAssignment_expression_expressionList_patternList_yield{Node: *child}, nil
}
func (a *Assignment) Type_() (*Type, error) {
	child := a.Node.ChildByFieldName("type_")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "assignment", "type_")
	}
	return &Type{Node: *child}, nil
}
func NewAssignment(node *tree_sitter.Node) (*Assignment, error) {
	if node.Kind() != "assignment" {
		return nil, fmt.Errorf("Node is not a %s", "assignment")
	}
	return &Assignment{Node: *node}, nil
}

type Attribute struct {
	tree_sitter.Node
}

func (a *Attribute) Attribute() (*Identifier, error) {
	child := a.Node.ChildByFieldName("attribute")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "attribute", "attribute")
	}
	return &Identifier{Node: *child}, nil
}
func (a *Attribute) Object() (*PrimaryExpression, error) {
	child := a.Node.ChildByFieldName("object")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "attribute", "object")
	}
	return &PrimaryExpression{Node: *child}, nil
}
func NewAttribute(node *tree_sitter.Node) (*Attribute, error) {
	if node.Kind() != "attribute" {
		return nil, fmt.Errorf("Node is not a %s", "attribute")
	}
	return &Attribute{Node: *node}, nil
}

type AugmentedAssignment struct {
	tree_sitter.Node
}

func (a *AugmentedAssignment) Left() (*pattern_patternList, error) {
	child := a.Node.ChildByFieldName("left")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "augmented_assignment", "left")
	}
	return &pattern_patternList{Node: *child}, nil
}
func (a *AugmentedAssignment) Operator() (*modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq, error) {
	child := a.Node.ChildByFieldName("operator")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "augmented_assignment", "operator")
	}
	return &modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq{Node: *child}, nil
}
func (a *AugmentedAssignment) Right() (*assignment_augmentedAssignment_expression_expressionList_patternList_yield, error) {
	child := a.Node.ChildByFieldName("right")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "augmented_assignment", "right")
	}
	return &assignment_augmentedAssignment_expression_expressionList_patternList_yield{Node: *child}, nil
}
func NewAugmentedAssignment(node *tree_sitter.Node) (*AugmentedAssignment, error) {
	if node.Kind() != "augmented_assignment" {
		return nil, fmt.Errorf("Node is not a %s", "augmented_assignment")
	}
	return &AugmentedAssignment{Node: *node}, nil
}

type Await struct {
	tree_sitter.Node
}

func (a *Await) TypedChild(cursor *tree_sitter.TreeCursor) (PrimaryExpression, error) {
	children := a.Node.Children(cursor)
	output := []PrimaryExpression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, PrimaryExpression{Node: child})
		}
	}
	if len(output) == 0 {
		return PrimaryExpression{}, fmt.Errorf("No children found on node of kind %s", "await")
	}
	return output[0], nil
}
func NewAwait(node *tree_sitter.Node) (*Await, error) {
	if node.Kind() != "await" {
		return nil, fmt.Errorf("Node is not a %s", "await")
	}
	return &Await{Node: *node}, nil
}

type BinaryOperator struct {
	tree_sitter.Node
}

func (b *BinaryOperator) Left() (*PrimaryExpression, error) {
	child := b.Node.ChildByFieldName("left")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "binary_operator", "left")
	}
	return &PrimaryExpression{Node: *child}, nil
}
func (b *BinaryOperator) Operator() (*mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar, error) {
	child := b.Node.ChildByFieldName("operator")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "binary_operator", "operator")
	}
	return &mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar{Node: *child}, nil
}
func (b *BinaryOperator) Right() (*PrimaryExpression, error) {
	child := b.Node.ChildByFieldName("right")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "binary_operator", "right")
	}
	return &PrimaryExpression{Node: *child}, nil
}
func NewBinaryOperator(node *tree_sitter.Node) (*BinaryOperator, error) {
	if node.Kind() != "binary_operator" {
		return nil, fmt.Errorf("Node is not a %s", "binary_operator")
	}
	return &BinaryOperator{Node: *node}, nil
}

type Block struct {
	tree_sitter.Node
}

func (b *Block) Alternative(cursor *tree_sitter.TreeCursor) []*CaseClause {
	children := b.Node.ChildrenByFieldName("alternative", cursor)
	output := []*CaseClause{}
	for _, child := range children {
		output = append(output, &CaseClause{Node: child})
	}
	return output
}
func (b *Block) TypedChildren(cursor *tree_sitter.TreeCursor) []compoundStatement_simpleStatement {
	children := b.Node.Children(cursor)
	output := []compoundStatement_simpleStatement{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, compoundStatement_simpleStatement{Node: child})
		}
	}
	return output
}
func NewBlock(node *tree_sitter.Node) (*Block, error) {
	if node.Kind() != "block" {
		return nil, fmt.Errorf("Node is not a %s", "block")
	}
	return &Block{Node: *node}, nil
}

type BooleanOperator struct {
	tree_sitter.Node
}

func (b *BooleanOperator) Left() (*Expression, error) {
	child := b.Node.ChildByFieldName("left")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "boolean_operator", "left")
	}
	return &Expression{Node: *child}, nil
}
func (b *BooleanOperator) Operator() (*and_or, error) {
	child := b.Node.ChildByFieldName("operator")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "boolean_operator", "operator")
	}
	return &and_or{Node: *child}, nil
}
func (b *BooleanOperator) Right() (*Expression, error) {
	child := b.Node.ChildByFieldName("right")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "boolean_operator", "right")
	}
	return &Expression{Node: *child}, nil
}
func NewBooleanOperator(node *tree_sitter.Node) (*BooleanOperator, error) {
	if node.Kind() != "boolean_operator" {
		return nil, fmt.Errorf("Node is not a %s", "boolean_operator")
	}
	return &BooleanOperator{Node: *node}, nil
}

type BreakStatement struct {
	tree_sitter.Node
}

func NewBreakStatement(node *tree_sitter.Node) (*BreakStatement, error) {
	if node.Kind() != "break_statement" {
		return nil, fmt.Errorf("Node is not a %s", "break_statement")
	}
	return &BreakStatement{Node: *node}, nil
}

type Call struct {
	tree_sitter.Node
}

func (c *Call) Arguments() (*argumentList_generatorExpression, error) {
	child := c.Node.ChildByFieldName("arguments")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "call", "arguments")
	}
	return &argumentList_generatorExpression{Node: *child}, nil
}
func (c *Call) Function() (*PrimaryExpression, error) {
	child := c.Node.ChildByFieldName("function")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "call", "function")
	}
	return &PrimaryExpression{Node: *child}, nil
}
func NewCall(node *tree_sitter.Node) (*Call, error) {
	if node.Kind() != "call" {
		return nil, fmt.Errorf("Node is not a %s", "call")
	}
	return &Call{Node: *node}, nil
}

type CaseClause struct {
	tree_sitter.Node
}

func (c *CaseClause) Consequence() (*Block, error) {
	child := c.Node.ChildByFieldName("consequence")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "case_clause", "consequence")
	}
	return &Block{Node: *child}, nil
}
func (c *CaseClause) Guard() (*IfClause, error) {
	child := c.Node.ChildByFieldName("guard")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "case_clause", "guard")
	}
	return &IfClause{Node: *child}, nil
}
func (c *CaseClause) TypedChildren(cursor *tree_sitter.TreeCursor) []CasePattern {
	children := c.Node.Children(cursor)
	output := []CasePattern{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, CasePattern{Node: child})
		}
	}
	return output
}
func NewCaseClause(node *tree_sitter.Node) (*CaseClause, error) {
	if node.Kind() != "case_clause" {
		return nil, fmt.Errorf("Node is not a %s", "case_clause")
	}
	return &CaseClause{Node: *node}, nil
}

type CasePattern struct {
	tree_sitter.Node
}

func (c *CasePattern) TypedChild(cursor *tree_sitter.TreeCursor) (asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern, error) {
	children := c.Node.Children(cursor)
	output := []asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern{Node: child})
		}
	}
	if len(output) == 0 {
		return asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern{}, fmt.Errorf("No children found on node of kind %s", "case_pattern")
	}
	return output[0], nil
}
func NewCasePattern(node *tree_sitter.Node) (*CasePattern, error) {
	if node.Kind() != "case_pattern" {
		return nil, fmt.Errorf("Node is not a %s", "case_pattern")
	}
	return &CasePattern{Node: *node}, nil
}

type Chevron struct {
	tree_sitter.Node
}

func (c *Chevron) TypedChild(cursor *tree_sitter.TreeCursor) (Expression, error) {
	children := c.Node.Children(cursor)
	output := []Expression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Expression{Node: child})
		}
	}
	if len(output) == 0 {
		return Expression{}, fmt.Errorf("No children found on node of kind %s", "chevron")
	}
	return output[0], nil
}
func NewChevron(node *tree_sitter.Node) (*Chevron, error) {
	if node.Kind() != "chevron" {
		return nil, fmt.Errorf("Node is not a %s", "chevron")
	}
	return &Chevron{Node: *node}, nil
}

type ClassDefinition struct {
	tree_sitter.Node
}

func (c *ClassDefinition) Body() (*Block, error) {
	child := c.Node.ChildByFieldName("body")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "class_definition", "body")
	}
	return &Block{Node: *child}, nil
}
func (c *ClassDefinition) Name() (*Identifier, error) {
	child := c.Node.ChildByFieldName("name")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "class_definition", "name")
	}
	return &Identifier{Node: *child}, nil
}
func (c *ClassDefinition) Superclasses() (*ArgumentList, error) {
	child := c.Node.ChildByFieldName("superclasses")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "class_definition", "superclasses")
	}
	return &ArgumentList{Node: *child}, nil
}
func (c *ClassDefinition) TypeParameters() (*TypeParameter, error) {
	child := c.Node.ChildByFieldName("typeParameters")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "class_definition", "typeParameters")
	}
	return &TypeParameter{Node: *child}, nil
}
func NewClassDefinition(node *tree_sitter.Node) (*ClassDefinition, error) {
	if node.Kind() != "class_definition" {
		return nil, fmt.Errorf("Node is not a %s", "class_definition")
	}
	return &ClassDefinition{Node: *node}, nil
}

type ClassPattern struct {
	tree_sitter.Node
}

func (c *ClassPattern) TypedChildren(cursor *tree_sitter.TreeCursor) []casePattern_dottedName {
	children := c.Node.Children(cursor)
	output := []casePattern_dottedName{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, casePattern_dottedName{Node: child})
		}
	}
	return output
}
func NewClassPattern(node *tree_sitter.Node) (*ClassPattern, error) {
	if node.Kind() != "class_pattern" {
		return nil, fmt.Errorf("Node is not a %s", "class_pattern")
	}
	return &ClassPattern{Node: *node}, nil
}

type ComparisonOperator struct {
	tree_sitter.Node
}

func (c *ComparisonOperator) Operators(cursor *tree_sitter.TreeCursor) []*notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn {
	children := c.Node.ChildrenByFieldName("operators", cursor)
	output := []*notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn{}
	for _, child := range children {
		output = append(output, &notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn{Node: child})
	}
	return output
}
func (c *ComparisonOperator) TypedChildren(cursor *tree_sitter.TreeCursor) []PrimaryExpression {
	children := c.Node.Children(cursor)
	output := []PrimaryExpression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, PrimaryExpression{Node: child})
		}
	}
	return output
}
func NewComparisonOperator(node *tree_sitter.Node) (*ComparisonOperator, error) {
	if node.Kind() != "comparison_operator" {
		return nil, fmt.Errorf("Node is not a %s", "comparison_operator")
	}
	return &ComparisonOperator{Node: *node}, nil
}

type ComplexPattern struct {
	tree_sitter.Node
}

func (c *ComplexPattern) TypedChildren(cursor *tree_sitter.TreeCursor) []float_integer {
	children := c.Node.Children(cursor)
	output := []float_integer{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, float_integer{Node: child})
		}
	}
	return output
}
func NewComplexPattern(node *tree_sitter.Node) (*ComplexPattern, error) {
	if node.Kind() != "complex_pattern" {
		return nil, fmt.Errorf("Node is not a %s", "complex_pattern")
	}
	return &ComplexPattern{Node: *node}, nil
}

type ConcatenatedString struct {
	tree_sitter.Node
}

func (c *ConcatenatedString) TypedChildren(cursor *tree_sitter.TreeCursor) []String {
	children := c.Node.Children(cursor)
	output := []String{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, String{Node: child})
		}
	}
	return output
}
func NewConcatenatedString(node *tree_sitter.Node) (*ConcatenatedString, error) {
	if node.Kind() != "concatenated_string" {
		return nil, fmt.Errorf("Node is not a %s", "concatenated_string")
	}
	return &ConcatenatedString{Node: *node}, nil
}

type ConditionalExpression struct {
	tree_sitter.Node
}

func (c *ConditionalExpression) TypedChildren(cursor *tree_sitter.TreeCursor) []Expression {
	children := c.Node.Children(cursor)
	output := []Expression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Expression{Node: child})
		}
	}
	return output
}
func NewConditionalExpression(node *tree_sitter.Node) (*ConditionalExpression, error) {
	if node.Kind() != "conditional_expression" {
		return nil, fmt.Errorf("Node is not a %s", "conditional_expression")
	}
	return &ConditionalExpression{Node: *node}, nil
}

type ConstrainedType struct {
	tree_sitter.Node
}

func (c *ConstrainedType) TypedChildren(cursor *tree_sitter.TreeCursor) []Type {
	children := c.Node.Children(cursor)
	output := []Type{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Type{Node: child})
		}
	}
	return output
}
func NewConstrainedType(node *tree_sitter.Node) (*ConstrainedType, error) {
	if node.Kind() != "constrained_type" {
		return nil, fmt.Errorf("Node is not a %s", "constrained_type")
	}
	return &ConstrainedType{Node: *node}, nil
}

type ContinueStatement struct {
	tree_sitter.Node
}

func NewContinueStatement(node *tree_sitter.Node) (*ContinueStatement, error) {
	if node.Kind() != "continue_statement" {
		return nil, fmt.Errorf("Node is not a %s", "continue_statement")
	}
	return &ContinueStatement{Node: *node}, nil
}

type DecoratedDefinition struct {
	tree_sitter.Node
}

func (d *DecoratedDefinition) Definition() (*classDefinition_functionDefinition, error) {
	child := d.Node.ChildByFieldName("definition")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "decorated_definition", "definition")
	}
	return &classDefinition_functionDefinition{Node: *child}, nil
}
func (d *DecoratedDefinition) TypedChildren(cursor *tree_sitter.TreeCursor) []Decorator {
	children := d.Node.Children(cursor)
	output := []Decorator{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Decorator{Node: child})
		}
	}
	return output
}
func NewDecoratedDefinition(node *tree_sitter.Node) (*DecoratedDefinition, error) {
	if node.Kind() != "decorated_definition" {
		return nil, fmt.Errorf("Node is not a %s", "decorated_definition")
	}
	return &DecoratedDefinition{Node: *node}, nil
}

type Decorator struct {
	tree_sitter.Node
}

func (d *Decorator) TypedChild(cursor *tree_sitter.TreeCursor) (Expression, error) {
	children := d.Node.Children(cursor)
	output := []Expression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Expression{Node: child})
		}
	}
	if len(output) == 0 {
		return Expression{}, fmt.Errorf("No children found on node of kind %s", "decorator")
	}
	return output[0], nil
}
func NewDecorator(node *tree_sitter.Node) (*Decorator, error) {
	if node.Kind() != "decorator" {
		return nil, fmt.Errorf("Node is not a %s", "decorator")
	}
	return &Decorator{Node: *node}, nil
}

type DefaultParameter struct {
	tree_sitter.Node
}

func (d *DefaultParameter) Name() (*identifier_tuplePattern, error) {
	child := d.Node.ChildByFieldName("name")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "default_parameter", "name")
	}
	return &identifier_tuplePattern{Node: *child}, nil
}
func (d *DefaultParameter) Value() (*Expression, error) {
	child := d.Node.ChildByFieldName("value")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "default_parameter", "value")
	}
	return &Expression{Node: *child}, nil
}
func NewDefaultParameter(node *tree_sitter.Node) (*DefaultParameter, error) {
	if node.Kind() != "default_parameter" {
		return nil, fmt.Errorf("Node is not a %s", "default_parameter")
	}
	return &DefaultParameter{Node: *node}, nil
}

type DeleteStatement struct {
	tree_sitter.Node
}

func (d *DeleteStatement) TypedChild(cursor *tree_sitter.TreeCursor) (expression_expressionList, error) {
	children := d.Node.Children(cursor)
	output := []expression_expressionList{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, expression_expressionList{Node: child})
		}
	}
	if len(output) == 0 {
		return expression_expressionList{}, fmt.Errorf("No children found on node of kind %s", "delete_statement")
	}
	return output[0], nil
}
func NewDeleteStatement(node *tree_sitter.Node) (*DeleteStatement, error) {
	if node.Kind() != "delete_statement" {
		return nil, fmt.Errorf("Node is not a %s", "delete_statement")
	}
	return &DeleteStatement{Node: *node}, nil
}

type DictPattern struct {
	tree_sitter.Node
}

func (d *DictPattern) Key(cursor *tree_sitter.TreeCursor) []*sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern {
	children := d.Node.ChildrenByFieldName("key", cursor)
	output := []*sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern{}
	for _, child := range children {
		output = append(output, &sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern{Node: child})
	}
	return output
}
func (d *DictPattern) Value(cursor *tree_sitter.TreeCursor) []*CasePattern {
	children := d.Node.ChildrenByFieldName("value", cursor)
	output := []*CasePattern{}
	for _, child := range children {
		output = append(output, &CasePattern{Node: child})
	}
	return output
}
func (d *DictPattern) TypedChildren(cursor *tree_sitter.TreeCursor) []SplatPattern {
	children := d.Node.Children(cursor)
	output := []SplatPattern{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, SplatPattern{Node: child})
		}
	}
	return output
}
func NewDictPattern(node *tree_sitter.Node) (*DictPattern, error) {
	if node.Kind() != "dict_pattern" {
		return nil, fmt.Errorf("Node is not a %s", "dict_pattern")
	}
	return &DictPattern{Node: *node}, nil
}

type Dictionary struct {
	tree_sitter.Node
}

func (d *Dictionary) TypedChildren(cursor *tree_sitter.TreeCursor) []dictionarySplat_pair {
	children := d.Node.Children(cursor)
	output := []dictionarySplat_pair{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, dictionarySplat_pair{Node: child})
		}
	}
	return output
}
func NewDictionary(node *tree_sitter.Node) (*Dictionary, error) {
	if node.Kind() != "dictionary" {
		return nil, fmt.Errorf("Node is not a %s", "dictionary")
	}
	return &Dictionary{Node: *node}, nil
}

type DictionaryComprehension struct {
	tree_sitter.Node
}

func (d *DictionaryComprehension) Body() (*Pair, error) {
	child := d.Node.ChildByFieldName("body")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "dictionary_comprehension", "body")
	}
	return &Pair{Node: *child}, nil
}
func (d *DictionaryComprehension) TypedChildren(cursor *tree_sitter.TreeCursor) []forInClause_ifClause {
	children := d.Node.Children(cursor)
	output := []forInClause_ifClause{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, forInClause_ifClause{Node: child})
		}
	}
	return output
}
func NewDictionaryComprehension(node *tree_sitter.Node) (*DictionaryComprehension, error) {
	if node.Kind() != "dictionary_comprehension" {
		return nil, fmt.Errorf("Node is not a %s", "dictionary_comprehension")
	}
	return &DictionaryComprehension{Node: *node}, nil
}

type DictionarySplat struct {
	tree_sitter.Node
}

func (d *DictionarySplat) TypedChild(cursor *tree_sitter.TreeCursor) (Expression, error) {
	children := d.Node.Children(cursor)
	output := []Expression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Expression{Node: child})
		}
	}
	if len(output) == 0 {
		return Expression{}, fmt.Errorf("No children found on node of kind %s", "dictionary_splat")
	}
	return output[0], nil
}
func NewDictionarySplat(node *tree_sitter.Node) (*DictionarySplat, error) {
	if node.Kind() != "dictionary_splat" {
		return nil, fmt.Errorf("Node is not a %s", "dictionary_splat")
	}
	return &DictionarySplat{Node: *node}, nil
}

type DictionarySplatPattern struct {
	tree_sitter.Node
}

func (d *DictionarySplatPattern) TypedChild(cursor *tree_sitter.TreeCursor) (attribute_identifier_subscript, error) {
	children := d.Node.Children(cursor)
	output := []attribute_identifier_subscript{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, attribute_identifier_subscript{Node: child})
		}
	}
	if len(output) == 0 {
		return attribute_identifier_subscript{}, fmt.Errorf("No children found on node of kind %s", "dictionary_splat_pattern")
	}
	return output[0], nil
}
func NewDictionarySplatPattern(node *tree_sitter.Node) (*DictionarySplatPattern, error) {
	if node.Kind() != "dictionary_splat_pattern" {
		return nil, fmt.Errorf("Node is not a %s", "dictionary_splat_pattern")
	}
	return &DictionarySplatPattern{Node: *node}, nil
}

type DottedName struct {
	tree_sitter.Node
}

func (d *DottedName) TypedChildren(cursor *tree_sitter.TreeCursor) []Identifier {
	children := d.Node.Children(cursor)
	output := []Identifier{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Identifier{Node: child})
		}
	}
	return output
}
func NewDottedName(node *tree_sitter.Node) (*DottedName, error) {
	if node.Kind() != "dotted_name" {
		return nil, fmt.Errorf("Node is not a %s", "dotted_name")
	}
	return &DottedName{Node: *node}, nil
}

type ElifClause struct {
	tree_sitter.Node
}

func (e *ElifClause) Condition() (*Expression, error) {
	child := e.Node.ChildByFieldName("condition")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "elif_clause", "condition")
	}
	return &Expression{Node: *child}, nil
}
func (e *ElifClause) Consequence() (*Block, error) {
	child := e.Node.ChildByFieldName("consequence")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "elif_clause", "consequence")
	}
	return &Block{Node: *child}, nil
}
func NewElifClause(node *tree_sitter.Node) (*ElifClause, error) {
	if node.Kind() != "elif_clause" {
		return nil, fmt.Errorf("Node is not a %s", "elif_clause")
	}
	return &ElifClause{Node: *node}, nil
}

type ElseClause struct {
	tree_sitter.Node
}

func (e *ElseClause) Body() (*Block, error) {
	child := e.Node.ChildByFieldName("body")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "else_clause", "body")
	}
	return &Block{Node: *child}, nil
}
func NewElseClause(node *tree_sitter.Node) (*ElseClause, error) {
	if node.Kind() != "else_clause" {
		return nil, fmt.Errorf("Node is not a %s", "else_clause")
	}
	return &ElseClause{Node: *node}, nil
}

type ExceptClause struct {
	tree_sitter.Node
}

func (e *ExceptClause) Alias() (*Expression, error) {
	child := e.Node.ChildByFieldName("alias")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "except_clause", "alias")
	}
	return &Expression{Node: *child}, nil
}
func (e *ExceptClause) Value() (*Expression, error) {
	child := e.Node.ChildByFieldName("value")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "except_clause", "value")
	}
	return &Expression{Node: *child}, nil
}
func (e *ExceptClause) TypedChild(cursor *tree_sitter.TreeCursor) (Block, error) {
	children := e.Node.Children(cursor)
	output := []Block{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Block{Node: child})
		}
	}
	if len(output) == 0 {
		return Block{}, fmt.Errorf("No children found on node of kind %s", "except_clause")
	}
	return output[0], nil
}
func NewExceptClause(node *tree_sitter.Node) (*ExceptClause, error) {
	if node.Kind() != "except_clause" {
		return nil, fmt.Errorf("Node is not a %s", "except_clause")
	}
	return &ExceptClause{Node: *node}, nil
}

type ExceptGroupClause struct {
	tree_sitter.Node
}

func (e *ExceptGroupClause) TypedChildren(cursor *tree_sitter.TreeCursor) []block_expression {
	children := e.Node.Children(cursor)
	output := []block_expression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, block_expression{Node: child})
		}
	}
	return output
}
func NewExceptGroupClause(node *tree_sitter.Node) (*ExceptGroupClause, error) {
	if node.Kind() != "except_group_clause" {
		return nil, fmt.Errorf("Node is not a %s", "except_group_clause")
	}
	return &ExceptGroupClause{Node: *node}, nil
}

type ExecStatement struct {
	tree_sitter.Node
}

func (e *ExecStatement) Code() (*identifier_string, error) {
	child := e.Node.ChildByFieldName("code")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "exec_statement", "code")
	}
	return &identifier_string{Node: *child}, nil
}
func (e *ExecStatement) TypedChildren(cursor *tree_sitter.TreeCursor) []Expression {
	children := e.Node.Children(cursor)
	output := []Expression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Expression{Node: child})
		}
	}
	return output
}
func NewExecStatement(node *tree_sitter.Node) (*ExecStatement, error) {
	if node.Kind() != "exec_statement" {
		return nil, fmt.Errorf("Node is not a %s", "exec_statement")
	}
	return &ExecStatement{Node: *node}, nil
}

type ExpressionList struct {
	tree_sitter.Node
}

func (e *ExpressionList) TypedChildren(cursor *tree_sitter.TreeCursor) []Expression {
	children := e.Node.Children(cursor)
	output := []Expression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Expression{Node: child})
		}
	}
	return output
}
func NewExpressionList(node *tree_sitter.Node) (*ExpressionList, error) {
	if node.Kind() != "expression_list" {
		return nil, fmt.Errorf("Node is not a %s", "expression_list")
	}
	return &ExpressionList{Node: *node}, nil
}

type ExpressionStatement struct {
	tree_sitter.Node
}

func (e *ExpressionStatement) TypedChildren(cursor *tree_sitter.TreeCursor) []assignment_augmentedAssignment_expression_yield {
	children := e.Node.Children(cursor)
	output := []assignment_augmentedAssignment_expression_yield{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, assignment_augmentedAssignment_expression_yield{Node: child})
		}
	}
	return output
}
func NewExpressionStatement(node *tree_sitter.Node) (*ExpressionStatement, error) {
	if node.Kind() != "expression_statement" {
		return nil, fmt.Errorf("Node is not a %s", "expression_statement")
	}
	return &ExpressionStatement{Node: *node}, nil
}

type FinallyClause struct {
	tree_sitter.Node
}

func (f *FinallyClause) TypedChild(cursor *tree_sitter.TreeCursor) (Block, error) {
	children := f.Node.Children(cursor)
	output := []Block{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Block{Node: child})
		}
	}
	if len(output) == 0 {
		return Block{}, fmt.Errorf("No children found on node of kind %s", "finally_clause")
	}
	return output[0], nil
}
func NewFinallyClause(node *tree_sitter.Node) (*FinallyClause, error) {
	if node.Kind() != "finally_clause" {
		return nil, fmt.Errorf("Node is not a %s", "finally_clause")
	}
	return &FinallyClause{Node: *node}, nil
}

type ForInClause struct {
	tree_sitter.Node
}

func (f *ForInClause) Left() (*pattern_patternList, error) {
	child := f.Node.ChildByFieldName("left")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "for_in_clause", "left")
	}
	return &pattern_patternList{Node: *child}, nil
}
func (f *ForInClause) Right(cursor *tree_sitter.TreeCursor) []*comma_expression {
	children := f.Node.ChildrenByFieldName("right", cursor)
	output := []*comma_expression{}
	for _, child := range children {
		output = append(output, &comma_expression{Node: child})
	}
	return output
}
func NewForInClause(node *tree_sitter.Node) (*ForInClause, error) {
	if node.Kind() != "for_in_clause" {
		return nil, fmt.Errorf("Node is not a %s", "for_in_clause")
	}
	return &ForInClause{Node: *node}, nil
}

type ForStatement struct {
	tree_sitter.Node
}

func (f *ForStatement) Alternative() (*ElseClause, error) {
	child := f.Node.ChildByFieldName("alternative")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "for_statement", "alternative")
	}
	return &ElseClause{Node: *child}, nil
}
func (f *ForStatement) Body() (*Block, error) {
	child := f.Node.ChildByFieldName("body")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "for_statement", "body")
	}
	return &Block{Node: *child}, nil
}
func (f *ForStatement) Left() (*pattern_patternList, error) {
	child := f.Node.ChildByFieldName("left")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "for_statement", "left")
	}
	return &pattern_patternList{Node: *child}, nil
}
func (f *ForStatement) Right() (*expression_expressionList, error) {
	child := f.Node.ChildByFieldName("right")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "for_statement", "right")
	}
	return &expression_expressionList{Node: *child}, nil
}
func NewForStatement(node *tree_sitter.Node) (*ForStatement, error) {
	if node.Kind() != "for_statement" {
		return nil, fmt.Errorf("Node is not a %s", "for_statement")
	}
	return &ForStatement{Node: *node}, nil
}

type FormatExpression struct {
	tree_sitter.Node
}

func (f *FormatExpression) Expression() (*expression_expressionList_patternList_yield, error) {
	child := f.Node.ChildByFieldName("expression")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "format_expression", "expression")
	}
	return &expression_expressionList_patternList_yield{Node: *child}, nil
}
func (f *FormatExpression) FormatSpecifier() (*FormatSpecifier, error) {
	child := f.Node.ChildByFieldName("formatSpecifier")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "format_expression", "formatSpecifier")
	}
	return &FormatSpecifier{Node: *child}, nil
}
func (f *FormatExpression) TypeConversion() (*TypeConversion, error) {
	child := f.Node.ChildByFieldName("typeConversion")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "format_expression", "typeConversion")
	}
	return &TypeConversion{Node: *child}, nil
}
func NewFormatExpression(node *tree_sitter.Node) (*FormatExpression, error) {
	if node.Kind() != "format_expression" {
		return nil, fmt.Errorf("Node is not a %s", "format_expression")
	}
	return &FormatExpression{Node: *node}, nil
}

type FormatSpecifier struct {
	tree_sitter.Node
}

func (f *FormatSpecifier) TypedChildren(cursor *tree_sitter.TreeCursor) []FormatExpression {
	children := f.Node.Children(cursor)
	output := []FormatExpression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, FormatExpression{Node: child})
		}
	}
	return output
}
func NewFormatSpecifier(node *tree_sitter.Node) (*FormatSpecifier, error) {
	if node.Kind() != "format_specifier" {
		return nil, fmt.Errorf("Node is not a %s", "format_specifier")
	}
	return &FormatSpecifier{Node: *node}, nil
}

type FunctionDefinition struct {
	tree_sitter.Node
}

func (f *FunctionDefinition) Body() (*Block, error) {
	child := f.Node.ChildByFieldName("body")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "function_definition", "body")
	}
	return &Block{Node: *child}, nil
}
func (f *FunctionDefinition) Name() (*Identifier, error) {
	child := f.Node.ChildByFieldName("name")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "function_definition", "name")
	}
	return &Identifier{Node: *child}, nil
}
func (f *FunctionDefinition) Parameters() (*Parameters, error) {
	child := f.Node.ChildByFieldName("parameters")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "function_definition", "parameters")
	}
	return &Parameters{Node: *child}, nil
}
func (f *FunctionDefinition) ReturnType() (*Type, error) {
	child := f.Node.ChildByFieldName("returnType")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "function_definition", "returnType")
	}
	return &Type{Node: *child}, nil
}
func (f *FunctionDefinition) TypeParameters() (*TypeParameter, error) {
	child := f.Node.ChildByFieldName("typeParameters")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "function_definition", "typeParameters")
	}
	return &TypeParameter{Node: *child}, nil
}
func NewFunctionDefinition(node *tree_sitter.Node) (*FunctionDefinition, error) {
	if node.Kind() != "function_definition" {
		return nil, fmt.Errorf("Node is not a %s", "function_definition")
	}
	return &FunctionDefinition{Node: *node}, nil
}

type FutureImportStatement struct {
	tree_sitter.Node
}

func (f *FutureImportStatement) Name(cursor *tree_sitter.TreeCursor) []*aliasedImport_dottedName {
	children := f.Node.ChildrenByFieldName("name", cursor)
	output := []*aliasedImport_dottedName{}
	for _, child := range children {
		output = append(output, &aliasedImport_dottedName{Node: child})
	}
	return output
}
func NewFutureImportStatement(node *tree_sitter.Node) (*FutureImportStatement, error) {
	if node.Kind() != "future_import_statement" {
		return nil, fmt.Errorf("Node is not a %s", "future_import_statement")
	}
	return &FutureImportStatement{Node: *node}, nil
}

type GeneratorExpression struct {
	tree_sitter.Node
}

func (g *GeneratorExpression) Body() (*Expression, error) {
	child := g.Node.ChildByFieldName("body")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "generator_expression", "body")
	}
	return &Expression{Node: *child}, nil
}
func (g *GeneratorExpression) TypedChildren(cursor *tree_sitter.TreeCursor) []forInClause_ifClause {
	children := g.Node.Children(cursor)
	output := []forInClause_ifClause{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, forInClause_ifClause{Node: child})
		}
	}
	return output
}
func NewGeneratorExpression(node *tree_sitter.Node) (*GeneratorExpression, error) {
	if node.Kind() != "generator_expression" {
		return nil, fmt.Errorf("Node is not a %s", "generator_expression")
	}
	return &GeneratorExpression{Node: *node}, nil
}

type GenericType struct {
	tree_sitter.Node
}

func (g *GenericType) TypedChildren(cursor *tree_sitter.TreeCursor) []identifier_typeParameter {
	children := g.Node.Children(cursor)
	output := []identifier_typeParameter{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, identifier_typeParameter{Node: child})
		}
	}
	return output
}
func NewGenericType(node *tree_sitter.Node) (*GenericType, error) {
	if node.Kind() != "generic_type" {
		return nil, fmt.Errorf("Node is not a %s", "generic_type")
	}
	return &GenericType{Node: *node}, nil
}

type GlobalStatement struct {
	tree_sitter.Node
}

func (g *GlobalStatement) TypedChildren(cursor *tree_sitter.TreeCursor) []Identifier {
	children := g.Node.Children(cursor)
	output := []Identifier{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Identifier{Node: child})
		}
	}
	return output
}
func NewGlobalStatement(node *tree_sitter.Node) (*GlobalStatement, error) {
	if node.Kind() != "global_statement" {
		return nil, fmt.Errorf("Node is not a %s", "global_statement")
	}
	return &GlobalStatement{Node: *node}, nil
}

type IfClause struct {
	tree_sitter.Node
}

func (i *IfClause) TypedChild(cursor *tree_sitter.TreeCursor) (Expression, error) {
	children := i.Node.Children(cursor)
	output := []Expression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Expression{Node: child})
		}
	}
	if len(output) == 0 {
		return Expression{}, fmt.Errorf("No children found on node of kind %s", "if_clause")
	}
	return output[0], nil
}
func NewIfClause(node *tree_sitter.Node) (*IfClause, error) {
	if node.Kind() != "if_clause" {
		return nil, fmt.Errorf("Node is not a %s", "if_clause")
	}
	return &IfClause{Node: *node}, nil
}

type IfStatement struct {
	tree_sitter.Node
}

func (i *IfStatement) Alternative(cursor *tree_sitter.TreeCursor) []*elifClause_elseClause {
	children := i.Node.ChildrenByFieldName("alternative", cursor)
	output := []*elifClause_elseClause{}
	for _, child := range children {
		output = append(output, &elifClause_elseClause{Node: child})
	}
	return output
}
func (i *IfStatement) Condition() (*Expression, error) {
	child := i.Node.ChildByFieldName("condition")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "if_statement", "condition")
	}
	return &Expression{Node: *child}, nil
}
func (i *IfStatement) Consequence() (*Block, error) {
	child := i.Node.ChildByFieldName("consequence")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "if_statement", "consequence")
	}
	return &Block{Node: *child}, nil
}
func NewIfStatement(node *tree_sitter.Node) (*IfStatement, error) {
	if node.Kind() != "if_statement" {
		return nil, fmt.Errorf("Node is not a %s", "if_statement")
	}
	return &IfStatement{Node: *node}, nil
}

type ImportFromStatement struct {
	tree_sitter.Node
}

func (i *ImportFromStatement) ModuleName() (*dottedName_relativeImport, error) {
	child := i.Node.ChildByFieldName("moduleName")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "import_from_statement", "moduleName")
	}
	return &dottedName_relativeImport{Node: *child}, nil
}
func (i *ImportFromStatement) Name(cursor *tree_sitter.TreeCursor) []*aliasedImport_dottedName {
	children := i.Node.ChildrenByFieldName("name", cursor)
	output := []*aliasedImport_dottedName{}
	for _, child := range children {
		output = append(output, &aliasedImport_dottedName{Node: child})
	}
	return output
}
func (i *ImportFromStatement) TypedChild(cursor *tree_sitter.TreeCursor) (WildcardImport, error) {
	children := i.Node.Children(cursor)
	output := []WildcardImport{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, WildcardImport{Node: child})
		}
	}
	if len(output) == 0 {
		return WildcardImport{}, fmt.Errorf("No children found on node of kind %s", "import_from_statement")
	}
	return output[0], nil
}
func NewImportFromStatement(node *tree_sitter.Node) (*ImportFromStatement, error) {
	if node.Kind() != "import_from_statement" {
		return nil, fmt.Errorf("Node is not a %s", "import_from_statement")
	}
	return &ImportFromStatement{Node: *node}, nil
}

type ImportPrefix struct {
	tree_sitter.Node
}

func NewImportPrefix(node *tree_sitter.Node) (*ImportPrefix, error) {
	if node.Kind() != "import_prefix" {
		return nil, fmt.Errorf("Node is not a %s", "import_prefix")
	}
	return &ImportPrefix{Node: *node}, nil
}

type ImportStatement struct {
	tree_sitter.Node
}

func (i *ImportStatement) Name(cursor *tree_sitter.TreeCursor) []*aliasedImport_dottedName {
	children := i.Node.ChildrenByFieldName("name", cursor)
	output := []*aliasedImport_dottedName{}
	for _, child := range children {
		output = append(output, &aliasedImport_dottedName{Node: child})
	}
	return output
}
func NewImportStatement(node *tree_sitter.Node) (*ImportStatement, error) {
	if node.Kind() != "import_statement" {
		return nil, fmt.Errorf("Node is not a %s", "import_statement")
	}
	return &ImportStatement{Node: *node}, nil
}

type Interpolation struct {
	tree_sitter.Node
}

func (i *Interpolation) Expression() (*expression_expressionList_patternList_yield, error) {
	child := i.Node.ChildByFieldName("expression")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "interpolation", "expression")
	}
	return &expression_expressionList_patternList_yield{Node: *child}, nil
}
func (i *Interpolation) FormatSpecifier() (*FormatSpecifier, error) {
	child := i.Node.ChildByFieldName("formatSpecifier")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "interpolation", "formatSpecifier")
	}
	return &FormatSpecifier{Node: *child}, nil
}
func (i *Interpolation) TypeConversion() (*TypeConversion, error) {
	child := i.Node.ChildByFieldName("typeConversion")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "interpolation", "typeConversion")
	}
	return &TypeConversion{Node: *child}, nil
}
func NewInterpolation(node *tree_sitter.Node) (*Interpolation, error) {
	if node.Kind() != "interpolation" {
		return nil, fmt.Errorf("Node is not a %s", "interpolation")
	}
	return &Interpolation{Node: *node}, nil
}

type Unnamed_IsSpaceNot struct {
	tree_sitter.Node
}

func NewUnnamed_IsSpaceNot(node *tree_sitter.Node) (*Unnamed_IsSpaceNot, error) {
	if node.Kind() != "is not" {
		return nil, fmt.Errorf("Node is not a %s", "is not")
	}
	return &Unnamed_IsSpaceNot{Node: *node}, nil
}

type KeywordArgument struct {
	tree_sitter.Node
}

func (k *KeywordArgument) Name() (*Identifier, error) {
	child := k.Node.ChildByFieldName("name")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "keyword_argument", "name")
	}
	return &Identifier{Node: *child}, nil
}
func (k *KeywordArgument) Value() (*Expression, error) {
	child := k.Node.ChildByFieldName("value")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "keyword_argument", "value")
	}
	return &Expression{Node: *child}, nil
}
func NewKeywordArgument(node *tree_sitter.Node) (*KeywordArgument, error) {
	if node.Kind() != "keyword_argument" {
		return nil, fmt.Errorf("Node is not a %s", "keyword_argument")
	}
	return &KeywordArgument{Node: *node}, nil
}

type KeywordPattern struct {
	tree_sitter.Node
}

func (k *KeywordPattern) TypedChildren(cursor *tree_sitter.TreeCursor) []classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern {
	children := k.Node.Children(cursor)
	output := []classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern{Node: child})
		}
	}
	return output
}
func NewKeywordPattern(node *tree_sitter.Node) (*KeywordPattern, error) {
	if node.Kind() != "keyword_pattern" {
		return nil, fmt.Errorf("Node is not a %s", "keyword_pattern")
	}
	return &KeywordPattern{Node: *node}, nil
}

type KeywordSeparator struct {
	tree_sitter.Node
}

func NewKeywordSeparator(node *tree_sitter.Node) (*KeywordSeparator, error) {
	if node.Kind() != "keyword_separator" {
		return nil, fmt.Errorf("Node is not a %s", "keyword_separator")
	}
	return &KeywordSeparator{Node: *node}, nil
}

type Lambda struct {
	tree_sitter.Node
}

func (l *Lambda) Body() (*Expression, error) {
	child := l.Node.ChildByFieldName("body")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "lambda", "body")
	}
	return &Expression{Node: *child}, nil
}
func (l *Lambda) Parameters() (*LambdaParameters, error) {
	child := l.Node.ChildByFieldName("parameters")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "lambda", "parameters")
	}
	return &LambdaParameters{Node: *child}, nil
}
func NewLambda(node *tree_sitter.Node) (*Lambda, error) {
	if node.Kind() != "lambda" {
		return nil, fmt.Errorf("Node is not a %s", "lambda")
	}
	return &Lambda{Node: *node}, nil
}

type LambdaParameters struct {
	tree_sitter.Node
}

func (l *LambdaParameters) TypedChildren(cursor *tree_sitter.TreeCursor) []Parameter {
	children := l.Node.Children(cursor)
	output := []Parameter{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Parameter{Node: child})
		}
	}
	return output
}
func NewLambdaParameters(node *tree_sitter.Node) (*LambdaParameters, error) {
	if node.Kind() != "lambda_parameters" {
		return nil, fmt.Errorf("Node is not a %s", "lambda_parameters")
	}
	return &LambdaParameters{Node: *node}, nil
}

type List struct {
	tree_sitter.Node
}

func (l *List) TypedChildren(cursor *tree_sitter.TreeCursor) []expression_listSplat_parenthesizedListSplat_yield {
	children := l.Node.Children(cursor)
	output := []expression_listSplat_parenthesizedListSplat_yield{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, expression_listSplat_parenthesizedListSplat_yield{Node: child})
		}
	}
	return output
}
func NewList(node *tree_sitter.Node) (*List, error) {
	if node.Kind() != "list" {
		return nil, fmt.Errorf("Node is not a %s", "list")
	}
	return &List{Node: *node}, nil
}

type ListComprehension struct {
	tree_sitter.Node
}

func (l *ListComprehension) Body() (*Expression, error) {
	child := l.Node.ChildByFieldName("body")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "list_comprehension", "body")
	}
	return &Expression{Node: *child}, nil
}
func (l *ListComprehension) TypedChildren(cursor *tree_sitter.TreeCursor) []forInClause_ifClause {
	children := l.Node.Children(cursor)
	output := []forInClause_ifClause{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, forInClause_ifClause{Node: child})
		}
	}
	return output
}
func NewListComprehension(node *tree_sitter.Node) (*ListComprehension, error) {
	if node.Kind() != "list_comprehension" {
		return nil, fmt.Errorf("Node is not a %s", "list_comprehension")
	}
	return &ListComprehension{Node: *node}, nil
}

type ListPattern struct {
	tree_sitter.Node
}

func (l *ListPattern) TypedChildren(cursor *tree_sitter.TreeCursor) []casePattern_pattern {
	children := l.Node.Children(cursor)
	output := []casePattern_pattern{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, casePattern_pattern{Node: child})
		}
	}
	return output
}
func NewListPattern(node *tree_sitter.Node) (*ListPattern, error) {
	if node.Kind() != "list_pattern" {
		return nil, fmt.Errorf("Node is not a %s", "list_pattern")
	}
	return &ListPattern{Node: *node}, nil
}

type ListSplat struct {
	tree_sitter.Node
}

func (l *ListSplat) TypedChild(cursor *tree_sitter.TreeCursor) (attribute_expression_identifier_subscript, error) {
	children := l.Node.Children(cursor)
	output := []attribute_expression_identifier_subscript{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, attribute_expression_identifier_subscript{Node: child})
		}
	}
	if len(output) == 0 {
		return attribute_expression_identifier_subscript{}, fmt.Errorf("No children found on node of kind %s", "list_splat")
	}
	return output[0], nil
}
func NewListSplat(node *tree_sitter.Node) (*ListSplat, error) {
	if node.Kind() != "list_splat" {
		return nil, fmt.Errorf("Node is not a %s", "list_splat")
	}
	return &ListSplat{Node: *node}, nil
}

type ListSplatPattern struct {
	tree_sitter.Node
}

func (l *ListSplatPattern) TypedChild(cursor *tree_sitter.TreeCursor) (attribute_identifier_subscript, error) {
	children := l.Node.Children(cursor)
	output := []attribute_identifier_subscript{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, attribute_identifier_subscript{Node: child})
		}
	}
	if len(output) == 0 {
		return attribute_identifier_subscript{}, fmt.Errorf("No children found on node of kind %s", "list_splat_pattern")
	}
	return output[0], nil
}
func NewListSplatPattern(node *tree_sitter.Node) (*ListSplatPattern, error) {
	if node.Kind() != "list_splat_pattern" {
		return nil, fmt.Errorf("Node is not a %s", "list_splat_pattern")
	}
	return &ListSplatPattern{Node: *node}, nil
}

type MatchStatement struct {
	tree_sitter.Node
}

func (m *MatchStatement) Body() (*Block, error) {
	child := m.Node.ChildByFieldName("body")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "match_statement", "body")
	}
	return &Block{Node: *child}, nil
}
func (m *MatchStatement) Subject(cursor *tree_sitter.TreeCursor) []*Expression {
	children := m.Node.ChildrenByFieldName("subject", cursor)
	output := []*Expression{}
	for _, child := range children {
		output = append(output, &Expression{Node: child})
	}
	return output
}
func NewMatchStatement(node *tree_sitter.Node) (*MatchStatement, error) {
	if node.Kind() != "match_statement" {
		return nil, fmt.Errorf("Node is not a %s", "match_statement")
	}
	return &MatchStatement{Node: *node}, nil
}

type MemberType struct {
	tree_sitter.Node
}

func (m *MemberType) TypedChildren(cursor *tree_sitter.TreeCursor) []identifier_type_ {
	children := m.Node.Children(cursor)
	output := []identifier_type_{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, identifier_type_{Node: child})
		}
	}
	return output
}
func NewMemberType(node *tree_sitter.Node) (*MemberType, error) {
	if node.Kind() != "member_type" {
		return nil, fmt.Errorf("Node is not a %s", "member_type")
	}
	return &MemberType{Node: *node}, nil
}

type Module struct {
	tree_sitter.Node
}

func (m *Module) TypedChildren(cursor *tree_sitter.TreeCursor) []compoundStatement_simpleStatement {
	children := m.Node.Children(cursor)
	output := []compoundStatement_simpleStatement{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, compoundStatement_simpleStatement{Node: child})
		}
	}
	return output
}
func NewModule(node *tree_sitter.Node) (*Module, error) {
	if node.Kind() != "module" {
		return nil, fmt.Errorf("Node is not a %s", "module")
	}
	return &Module{Node: *node}, nil
}

type NamedExpression struct {
	tree_sitter.Node
}

func (n *NamedExpression) Name() (*Identifier, error) {
	child := n.Node.ChildByFieldName("name")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "named_expression", "name")
	}
	return &Identifier{Node: *child}, nil
}
func (n *NamedExpression) Value() (*Expression, error) {
	child := n.Node.ChildByFieldName("value")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "named_expression", "value")
	}
	return &Expression{Node: *child}, nil
}
func NewNamedExpression(node *tree_sitter.Node) (*NamedExpression, error) {
	if node.Kind() != "named_expression" {
		return nil, fmt.Errorf("Node is not a %s", "named_expression")
	}
	return &NamedExpression{Node: *node}, nil
}

type NonlocalStatement struct {
	tree_sitter.Node
}

func (n *NonlocalStatement) TypedChildren(cursor *tree_sitter.TreeCursor) []Identifier {
	children := n.Node.Children(cursor)
	output := []Identifier{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Identifier{Node: child})
		}
	}
	return output
}
func NewNonlocalStatement(node *tree_sitter.Node) (*NonlocalStatement, error) {
	if node.Kind() != "nonlocal_statement" {
		return nil, fmt.Errorf("Node is not a %s", "nonlocal_statement")
	}
	return &NonlocalStatement{Node: *node}, nil
}

type Unnamed_NotSpaceIn struct {
	tree_sitter.Node
}

func NewUnnamed_NotSpaceIn(node *tree_sitter.Node) (*Unnamed_NotSpaceIn, error) {
	if node.Kind() != "not in" {
		return nil, fmt.Errorf("Node is not a %s", "not in")
	}
	return &Unnamed_NotSpaceIn{Node: *node}, nil
}

type NotOperator struct {
	tree_sitter.Node
}

func (n *NotOperator) Argument() (*Expression, error) {
	child := n.Node.ChildByFieldName("argument")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "not_operator", "argument")
	}
	return &Expression{Node: *child}, nil
}
func NewNotOperator(node *tree_sitter.Node) (*NotOperator, error) {
	if node.Kind() != "not_operator" {
		return nil, fmt.Errorf("Node is not a %s", "not_operator")
	}
	return &NotOperator{Node: *node}, nil
}

type Pair struct {
	tree_sitter.Node
}

func (p *Pair) Key() (*Expression, error) {
	child := p.Node.ChildByFieldName("key")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "pair", "key")
	}
	return &Expression{Node: *child}, nil
}
func (p *Pair) Value() (*Expression, error) {
	child := p.Node.ChildByFieldName("value")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "pair", "value")
	}
	return &Expression{Node: *child}, nil
}
func NewPair(node *tree_sitter.Node) (*Pair, error) {
	if node.Kind() != "pair" {
		return nil, fmt.Errorf("Node is not a %s", "pair")
	}
	return &Pair{Node: *node}, nil
}

type Parameters struct {
	tree_sitter.Node
}

func (p *Parameters) TypedChildren(cursor *tree_sitter.TreeCursor) []Parameter {
	children := p.Node.Children(cursor)
	output := []Parameter{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Parameter{Node: child})
		}
	}
	return output
}
func NewParameters(node *tree_sitter.Node) (*Parameters, error) {
	if node.Kind() != "parameters" {
		return nil, fmt.Errorf("Node is not a %s", "parameters")
	}
	return &Parameters{Node: *node}, nil
}

type ParenthesizedExpression struct {
	tree_sitter.Node
}

func (p *ParenthesizedExpression) TypedChild(cursor *tree_sitter.TreeCursor) (expression_listSplat_parenthesizedExpression_yield, error) {
	children := p.Node.Children(cursor)
	output := []expression_listSplat_parenthesizedExpression_yield{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, expression_listSplat_parenthesizedExpression_yield{Node: child})
		}
	}
	if len(output) == 0 {
		return expression_listSplat_parenthesizedExpression_yield{}, fmt.Errorf("No children found on node of kind %s", "parenthesized_expression")
	}
	return output[0], nil
}
func NewParenthesizedExpression(node *tree_sitter.Node) (*ParenthesizedExpression, error) {
	if node.Kind() != "parenthesized_expression" {
		return nil, fmt.Errorf("Node is not a %s", "parenthesized_expression")
	}
	return &ParenthesizedExpression{Node: *node}, nil
}

type ParenthesizedListSplat struct {
	tree_sitter.Node
}

func (p *ParenthesizedListSplat) TypedChild(cursor *tree_sitter.TreeCursor) (listSplat_parenthesizedExpression, error) {
	children := p.Node.Children(cursor)
	output := []listSplat_parenthesizedExpression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, listSplat_parenthesizedExpression{Node: child})
		}
	}
	if len(output) == 0 {
		return listSplat_parenthesizedExpression{}, fmt.Errorf("No children found on node of kind %s", "parenthesized_list_splat")
	}
	return output[0], nil
}
func NewParenthesizedListSplat(node *tree_sitter.Node) (*ParenthesizedListSplat, error) {
	if node.Kind() != "parenthesized_list_splat" {
		return nil, fmt.Errorf("Node is not a %s", "parenthesized_list_splat")
	}
	return &ParenthesizedListSplat{Node: *node}, nil
}

type PassStatement struct {
	tree_sitter.Node
}

func NewPassStatement(node *tree_sitter.Node) (*PassStatement, error) {
	if node.Kind() != "pass_statement" {
		return nil, fmt.Errorf("Node is not a %s", "pass_statement")
	}
	return &PassStatement{Node: *node}, nil
}

type PatternList struct {
	tree_sitter.Node
}

func (p *PatternList) TypedChildren(cursor *tree_sitter.TreeCursor) []Pattern {
	children := p.Node.Children(cursor)
	output := []Pattern{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Pattern{Node: child})
		}
	}
	return output
}
func NewPatternList(node *tree_sitter.Node) (*PatternList, error) {
	if node.Kind() != "pattern_list" {
		return nil, fmt.Errorf("Node is not a %s", "pattern_list")
	}
	return &PatternList{Node: *node}, nil
}

type PositionalSeparator struct {
	tree_sitter.Node
}

func NewPositionalSeparator(node *tree_sitter.Node) (*PositionalSeparator, error) {
	if node.Kind() != "positional_separator" {
		return nil, fmt.Errorf("Node is not a %s", "positional_separator")
	}
	return &PositionalSeparator{Node: *node}, nil
}

type PrintStatement struct {
	tree_sitter.Node
}

func (p *PrintStatement) Argument(cursor *tree_sitter.TreeCursor) []*Expression {
	children := p.Node.ChildrenByFieldName("argument", cursor)
	output := []*Expression{}
	for _, child := range children {
		output = append(output, &Expression{Node: child})
	}
	return output
}
func (p *PrintStatement) TypedChild(cursor *tree_sitter.TreeCursor) (Chevron, error) {
	children := p.Node.Children(cursor)
	output := []Chevron{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Chevron{Node: child})
		}
	}
	if len(output) == 0 {
		return Chevron{}, fmt.Errorf("No children found on node of kind %s", "print_statement")
	}
	return output[0], nil
}
func NewPrintStatement(node *tree_sitter.Node) (*PrintStatement, error) {
	if node.Kind() != "print_statement" {
		return nil, fmt.Errorf("Node is not a %s", "print_statement")
	}
	return &PrintStatement{Node: *node}, nil
}

type RaiseStatement struct {
	tree_sitter.Node
}

func (r *RaiseStatement) Cause() (*Expression, error) {
	child := r.Node.ChildByFieldName("cause")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "raise_statement", "cause")
	}
	return &Expression{Node: *child}, nil
}
func (r *RaiseStatement) TypedChild(cursor *tree_sitter.TreeCursor) (expression_expressionList, error) {
	children := r.Node.Children(cursor)
	output := []expression_expressionList{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, expression_expressionList{Node: child})
		}
	}
	if len(output) == 0 {
		return expression_expressionList{}, fmt.Errorf("No children found on node of kind %s", "raise_statement")
	}
	return output[0], nil
}
func NewRaiseStatement(node *tree_sitter.Node) (*RaiseStatement, error) {
	if node.Kind() != "raise_statement" {
		return nil, fmt.Errorf("Node is not a %s", "raise_statement")
	}
	return &RaiseStatement{Node: *node}, nil
}

type RelativeImport struct {
	tree_sitter.Node
}

func (r *RelativeImport) TypedChildren(cursor *tree_sitter.TreeCursor) []dottedName_importPrefix {
	children := r.Node.Children(cursor)
	output := []dottedName_importPrefix{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, dottedName_importPrefix{Node: child})
		}
	}
	return output
}
func NewRelativeImport(node *tree_sitter.Node) (*RelativeImport, error) {
	if node.Kind() != "relative_import" {
		return nil, fmt.Errorf("Node is not a %s", "relative_import")
	}
	return &RelativeImport{Node: *node}, nil
}

type ReturnStatement struct {
	tree_sitter.Node
}

func (r *ReturnStatement) TypedChild(cursor *tree_sitter.TreeCursor) (expression_expressionList, error) {
	children := r.Node.Children(cursor)
	output := []expression_expressionList{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, expression_expressionList{Node: child})
		}
	}
	if len(output) == 0 {
		return expression_expressionList{}, fmt.Errorf("No children found on node of kind %s", "return_statement")
	}
	return output[0], nil
}
func NewReturnStatement(node *tree_sitter.Node) (*ReturnStatement, error) {
	if node.Kind() != "return_statement" {
		return nil, fmt.Errorf("Node is not a %s", "return_statement")
	}
	return &ReturnStatement{Node: *node}, nil
}

type Set struct {
	tree_sitter.Node
}

func (s *Set) TypedChildren(cursor *tree_sitter.TreeCursor) []expression_listSplat_parenthesizedListSplat_yield {
	children := s.Node.Children(cursor)
	output := []expression_listSplat_parenthesizedListSplat_yield{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, expression_listSplat_parenthesizedListSplat_yield{Node: child})
		}
	}
	return output
}
func NewSet(node *tree_sitter.Node) (*Set, error) {
	if node.Kind() != "set" {
		return nil, fmt.Errorf("Node is not a %s", "set")
	}
	return &Set{Node: *node}, nil
}

type SetComprehension struct {
	tree_sitter.Node
}

func (s *SetComprehension) Body() (*Expression, error) {
	child := s.Node.ChildByFieldName("body")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "set_comprehension", "body")
	}
	return &Expression{Node: *child}, nil
}
func (s *SetComprehension) TypedChildren(cursor *tree_sitter.TreeCursor) []forInClause_ifClause {
	children := s.Node.Children(cursor)
	output := []forInClause_ifClause{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, forInClause_ifClause{Node: child})
		}
	}
	return output
}
func NewSetComprehension(node *tree_sitter.Node) (*SetComprehension, error) {
	if node.Kind() != "set_comprehension" {
		return nil, fmt.Errorf("Node is not a %s", "set_comprehension")
	}
	return &SetComprehension{Node: *node}, nil
}

type Slice struct {
	tree_sitter.Node
}

func (s *Slice) TypedChildren(cursor *tree_sitter.TreeCursor) []Expression {
	children := s.Node.Children(cursor)
	output := []Expression{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Expression{Node: child})
		}
	}
	return output
}
func NewSlice(node *tree_sitter.Node) (*Slice, error) {
	if node.Kind() != "slice" {
		return nil, fmt.Errorf("Node is not a %s", "slice")
	}
	return &Slice{Node: *node}, nil
}

type SplatPattern struct {
	tree_sitter.Node
}

func (s *SplatPattern) TypedChild(cursor *tree_sitter.TreeCursor) (Identifier, error) {
	children := s.Node.Children(cursor)
	output := []Identifier{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Identifier{Node: child})
		}
	}
	if len(output) == 0 {
		return Identifier{}, fmt.Errorf("No children found on node of kind %s", "splat_pattern")
	}
	return output[0], nil
}
func NewSplatPattern(node *tree_sitter.Node) (*SplatPattern, error) {
	if node.Kind() != "splat_pattern" {
		return nil, fmt.Errorf("Node is not a %s", "splat_pattern")
	}
	return &SplatPattern{Node: *node}, nil
}

type SplatType struct {
	tree_sitter.Node
}

func (s *SplatType) TypedChild(cursor *tree_sitter.TreeCursor) (Identifier, error) {
	children := s.Node.Children(cursor)
	output := []Identifier{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Identifier{Node: child})
		}
	}
	if len(output) == 0 {
		return Identifier{}, fmt.Errorf("No children found on node of kind %s", "splat_type")
	}
	return output[0], nil
}
func NewSplatType(node *tree_sitter.Node) (*SplatType, error) {
	if node.Kind() != "splat_type" {
		return nil, fmt.Errorf("Node is not a %s", "splat_type")
	}
	return &SplatType{Node: *node}, nil
}

type String struct {
	tree_sitter.Node
}

func (s *String) TypedChildren(cursor *tree_sitter.TreeCursor) []interpolation_stringContent_stringEnd_stringStart {
	children := s.Node.Children(cursor)
	output := []interpolation_stringContent_stringEnd_stringStart{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, interpolation_stringContent_stringEnd_stringStart{Node: child})
		}
	}
	return output
}
func NewString(node *tree_sitter.Node) (*String, error) {
	if node.Kind() != "string" {
		return nil, fmt.Errorf("Node is not a %s", "string")
	}
	return &String{Node: *node}, nil
}

type StringContent struct {
	tree_sitter.Node
}

func (s *StringContent) TypedChildren(cursor *tree_sitter.TreeCursor) []escapeInterpolation_escapeSequence {
	children := s.Node.Children(cursor)
	output := []escapeInterpolation_escapeSequence{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, escapeInterpolation_escapeSequence{Node: child})
		}
	}
	return output
}
func NewStringContent(node *tree_sitter.Node) (*StringContent, error) {
	if node.Kind() != "string_content" {
		return nil, fmt.Errorf("Node is not a %s", "string_content")
	}
	return &StringContent{Node: *node}, nil
}

type Subscript struct {
	tree_sitter.Node
}

func (s *Subscript) Subscript(cursor *tree_sitter.TreeCursor) []*expression_slice {
	children := s.Node.ChildrenByFieldName("subscript", cursor)
	output := []*expression_slice{}
	for _, child := range children {
		output = append(output, &expression_slice{Node: child})
	}
	return output
}
func (s *Subscript) Value() (*PrimaryExpression, error) {
	child := s.Node.ChildByFieldName("value")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "subscript", "value")
	}
	return &PrimaryExpression{Node: *child}, nil
}
func NewSubscript(node *tree_sitter.Node) (*Subscript, error) {
	if node.Kind() != "subscript" {
		return nil, fmt.Errorf("Node is not a %s", "subscript")
	}
	return &Subscript{Node: *node}, nil
}

type TryStatement struct {
	tree_sitter.Node
}

func (t *TryStatement) Body() (*Block, error) {
	child := t.Node.ChildByFieldName("body")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "try_statement", "body")
	}
	return &Block{Node: *child}, nil
}
func (t *TryStatement) TypedChildren(cursor *tree_sitter.TreeCursor) []elseClause_exceptClause_exceptGroupClause_finallyClause {
	children := t.Node.Children(cursor)
	output := []elseClause_exceptClause_exceptGroupClause_finallyClause{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, elseClause_exceptClause_exceptGroupClause_finallyClause{Node: child})
		}
	}
	return output
}
func NewTryStatement(node *tree_sitter.Node) (*TryStatement, error) {
	if node.Kind() != "try_statement" {
		return nil, fmt.Errorf("Node is not a %s", "try_statement")
	}
	return &TryStatement{Node: *node}, nil
}

type Tuple struct {
	tree_sitter.Node
}

func (t *Tuple) TypedChildren(cursor *tree_sitter.TreeCursor) []expression_listSplat_parenthesizedListSplat_yield {
	children := t.Node.Children(cursor)
	output := []expression_listSplat_parenthesizedListSplat_yield{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, expression_listSplat_parenthesizedListSplat_yield{Node: child})
		}
	}
	return output
}
func NewTuple(node *tree_sitter.Node) (*Tuple, error) {
	if node.Kind() != "tuple" {
		return nil, fmt.Errorf("Node is not a %s", "tuple")
	}
	return &Tuple{Node: *node}, nil
}

type TuplePattern struct {
	tree_sitter.Node
}

func (t *TuplePattern) TypedChildren(cursor *tree_sitter.TreeCursor) []casePattern_pattern {
	children := t.Node.Children(cursor)
	output := []casePattern_pattern{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, casePattern_pattern{Node: child})
		}
	}
	return output
}
func NewTuplePattern(node *tree_sitter.Node) (*TuplePattern, error) {
	if node.Kind() != "tuple_pattern" {
		return nil, fmt.Errorf("Node is not a %s", "tuple_pattern")
	}
	return &TuplePattern{Node: *node}, nil
}

type Type struct {
	tree_sitter.Node
}

func (t *Type) TypedChild(cursor *tree_sitter.TreeCursor) (constrainedType_expression_genericType_memberType_splatType_unionType, error) {
	children := t.Node.Children(cursor)
	output := []constrainedType_expression_genericType_memberType_splatType_unionType{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, constrainedType_expression_genericType_memberType_splatType_unionType{Node: child})
		}
	}
	if len(output) == 0 {
		return constrainedType_expression_genericType_memberType_splatType_unionType{}, fmt.Errorf("No children found on node of kind %s", "type")
	}
	return output[0], nil
}
func NewType(node *tree_sitter.Node) (*Type, error) {
	if node.Kind() != "type" {
		return nil, fmt.Errorf("Node is not a %s", "type")
	}
	return &Type{Node: *node}, nil
}

type TypeAliasStatement struct {
	tree_sitter.Node
}

func (t *TypeAliasStatement) Left() (*Type, error) {
	child := t.Node.ChildByFieldName("left")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "type_alias_statement", "left")
	}
	return &Type{Node: *child}, nil
}
func (t *TypeAliasStatement) Right() (*Type, error) {
	child := t.Node.ChildByFieldName("right")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "type_alias_statement", "right")
	}
	return &Type{Node: *child}, nil
}
func NewTypeAliasStatement(node *tree_sitter.Node) (*TypeAliasStatement, error) {
	if node.Kind() != "type_alias_statement" {
		return nil, fmt.Errorf("Node is not a %s", "type_alias_statement")
	}
	return &TypeAliasStatement{Node: *node}, nil
}

type TypeParameter struct {
	tree_sitter.Node
}

func (t *TypeParameter) TypedChildren(cursor *tree_sitter.TreeCursor) []Type {
	children := t.Node.Children(cursor)
	output := []Type{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Type{Node: child})
		}
	}
	return output
}
func NewTypeParameter(node *tree_sitter.Node) (*TypeParameter, error) {
	if node.Kind() != "type_parameter" {
		return nil, fmt.Errorf("Node is not a %s", "type_parameter")
	}
	return &TypeParameter{Node: *node}, nil
}

type TypedDefaultParameter struct {
	tree_sitter.Node
}

func (t *TypedDefaultParameter) Name() (*Identifier, error) {
	child := t.Node.ChildByFieldName("name")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "typed_default_parameter", "name")
	}
	return &Identifier{Node: *child}, nil
}
func (t *TypedDefaultParameter) Type_() (*Type, error) {
	child := t.Node.ChildByFieldName("type_")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "typed_default_parameter", "type_")
	}
	return &Type{Node: *child}, nil
}
func (t *TypedDefaultParameter) Value() (*Expression, error) {
	child := t.Node.ChildByFieldName("value")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "typed_default_parameter", "value")
	}
	return &Expression{Node: *child}, nil
}
func NewTypedDefaultParameter(node *tree_sitter.Node) (*TypedDefaultParameter, error) {
	if node.Kind() != "typed_default_parameter" {
		return nil, fmt.Errorf("Node is not a %s", "typed_default_parameter")
	}
	return &TypedDefaultParameter{Node: *node}, nil
}

type TypedParameter struct {
	tree_sitter.Node
}

func (t *TypedParameter) Type_() (*Type, error) {
	child := t.Node.ChildByFieldName("type_")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "typed_parameter", "type_")
	}
	return &Type{Node: *child}, nil
}
func (t *TypedParameter) TypedChild(cursor *tree_sitter.TreeCursor) (dictionarySplatPattern_identifier_listSplatPattern, error) {
	children := t.Node.Children(cursor)
	output := []dictionarySplatPattern_identifier_listSplatPattern{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, dictionarySplatPattern_identifier_listSplatPattern{Node: child})
		}
	}
	if len(output) == 0 {
		return dictionarySplatPattern_identifier_listSplatPattern{}, fmt.Errorf("No children found on node of kind %s", "typed_parameter")
	}
	return output[0], nil
}
func NewTypedParameter(node *tree_sitter.Node) (*TypedParameter, error) {
	if node.Kind() != "typed_parameter" {
		return nil, fmt.Errorf("Node is not a %s", "typed_parameter")
	}
	return &TypedParameter{Node: *node}, nil
}

type UnaryOperator struct {
	tree_sitter.Node
}

func (u *UnaryOperator) Argument() (*PrimaryExpression, error) {
	child := u.Node.ChildByFieldName("argument")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "unary_operator", "argument")
	}
	return &PrimaryExpression{Node: *child}, nil
}
func (u *UnaryOperator) Operator() (*add_sub_bitNot, error) {
	child := u.Node.ChildByFieldName("operator")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "unary_operator", "operator")
	}
	return &add_sub_bitNot{Node: *child}, nil
}
func NewUnaryOperator(node *tree_sitter.Node) (*UnaryOperator, error) {
	if node.Kind() != "unary_operator" {
		return nil, fmt.Errorf("Node is not a %s", "unary_operator")
	}
	return &UnaryOperator{Node: *node}, nil
}

type UnionPattern struct {
	tree_sitter.Node
}

func (u *UnionPattern) TypedChildren(cursor *tree_sitter.TreeCursor) []classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern {
	children := u.Node.Children(cursor)
	output := []classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern{Node: child})
		}
	}
	return output
}
func NewUnionPattern(node *tree_sitter.Node) (*UnionPattern, error) {
	if node.Kind() != "union_pattern" {
		return nil, fmt.Errorf("Node is not a %s", "union_pattern")
	}
	return &UnionPattern{Node: *node}, nil
}

type UnionType struct {
	tree_sitter.Node
}

func (u *UnionType) TypedChildren(cursor *tree_sitter.TreeCursor) []Type {
	children := u.Node.Children(cursor)
	output := []Type{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, Type{Node: child})
		}
	}
	return output
}
func NewUnionType(node *tree_sitter.Node) (*UnionType, error) {
	if node.Kind() != "union_type" {
		return nil, fmt.Errorf("Node is not a %s", "union_type")
	}
	return &UnionType{Node: *node}, nil
}

type WhileStatement struct {
	tree_sitter.Node
}

func (w *WhileStatement) Alternative() (*ElseClause, error) {
	child := w.Node.ChildByFieldName("alternative")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "while_statement", "alternative")
	}
	return &ElseClause{Node: *child}, nil
}
func (w *WhileStatement) Body() (*Block, error) {
	child := w.Node.ChildByFieldName("body")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "while_statement", "body")
	}
	return &Block{Node: *child}, nil
}
func (w *WhileStatement) Condition() (*Expression, error) {
	child := w.Node.ChildByFieldName("condition")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "while_statement", "condition")
	}
	return &Expression{Node: *child}, nil
}
func NewWhileStatement(node *tree_sitter.Node) (*WhileStatement, error) {
	if node.Kind() != "while_statement" {
		return nil, fmt.Errorf("Node is not a %s", "while_statement")
	}
	return &WhileStatement{Node: *node}, nil
}

type WildcardImport struct {
	tree_sitter.Node
}

func NewWildcardImport(node *tree_sitter.Node) (*WildcardImport, error) {
	if node.Kind() != "wildcard_import" {
		return nil, fmt.Errorf("Node is not a %s", "wildcard_import")
	}
	return &WildcardImport{Node: *node}, nil
}

type WithClause struct {
	tree_sitter.Node
}

func (w *WithClause) TypedChildren(cursor *tree_sitter.TreeCursor) []WithItem {
	children := w.Node.Children(cursor)
	output := []WithItem{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, WithItem{Node: child})
		}
	}
	return output
}
func NewWithClause(node *tree_sitter.Node) (*WithClause, error) {
	if node.Kind() != "with_clause" {
		return nil, fmt.Errorf("Node is not a %s", "with_clause")
	}
	return &WithClause{Node: *node}, nil
}

type WithItem struct {
	tree_sitter.Node
}

func (w *WithItem) Value() (*Expression, error) {
	child := w.Node.ChildByFieldName("value")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "with_item", "value")
	}
	return &Expression{Node: *child}, nil
}
func NewWithItem(node *tree_sitter.Node) (*WithItem, error) {
	if node.Kind() != "with_item" {
		return nil, fmt.Errorf("Node is not a %s", "with_item")
	}
	return &WithItem{Node: *node}, nil
}

type WithStatement struct {
	tree_sitter.Node
}

func (w *WithStatement) Body() (*Block, error) {
	child := w.Node.ChildByFieldName("body")
	if child == nil {
		return nil, fmt.Errorf("Node of kind %s has no child of name %s", "with_statement", "body")
	}
	return &Block{Node: *child}, nil
}
func (w *WithStatement) TypedChild(cursor *tree_sitter.TreeCursor) (WithClause, error) {
	children := w.Node.Children(cursor)
	output := []WithClause{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, WithClause{Node: child})
		}
	}
	if len(output) == 0 {
		return WithClause{}, fmt.Errorf("No children found on node of kind %s", "with_statement")
	}
	return output[0], nil
}
func NewWithStatement(node *tree_sitter.Node) (*WithStatement, error) {
	if node.Kind() != "with_statement" {
		return nil, fmt.Errorf("Node is not a %s", "with_statement")
	}
	return &WithStatement{Node: *node}, nil
}

type Yield struct {
	tree_sitter.Node
}

func (y *Yield) TypedChild(cursor *tree_sitter.TreeCursor) (expression_expressionList, error) {
	children := y.Node.Children(cursor)
	output := []expression_expressionList{}
	for _, child := range children {
		if child.IsNamed() {
			output = append(output, expression_expressionList{Node: child})
		}
	}
	if len(output) == 0 {
		return expression_expressionList{}, fmt.Errorf("No children found on node of kind %s", "yield")
	}
	return output[0], nil
}
func NewYield(node *tree_sitter.Node) (*Yield, error) {
	if node.Kind() != "yield" {
		return nil, fmt.Errorf("Node is not a %s", "yield")
	}
	return &Yield{Node: *node}, nil
}

type Unnamed_NotEq struct {
	tree_sitter.Node
}

func NewUnnamed_NotEq(node *tree_sitter.Node) (*Unnamed_NotEq, error) {
	if node.Kind() != "!=" {
		return nil, fmt.Errorf("Node is not a %s", "!=")
	}
	return &Unnamed_NotEq{Node: *node}, nil
}

type Unnamed_Mod struct {
	tree_sitter.Node
}

func NewUnnamed_Mod(node *tree_sitter.Node) (*Unnamed_Mod, error) {
	if node.Kind() != "%" {
		return nil, fmt.Errorf("Node is not a %s", "%")
	}
	return &Unnamed_Mod{Node: *node}, nil
}

type Unnamed_ModEq struct {
	tree_sitter.Node
}

func NewUnnamed_ModEq(node *tree_sitter.Node) (*Unnamed_ModEq, error) {
	if node.Kind() != "%=" {
		return nil, fmt.Errorf("Node is not a %s", "%=")
	}
	return &Unnamed_ModEq{Node: *node}, nil
}

type Unnamed_Ampersand struct {
	tree_sitter.Node
}

func NewUnnamed_Ampersand(node *tree_sitter.Node) (*Unnamed_Ampersand, error) {
	if node.Kind() != "&" {
		return nil, fmt.Errorf("Node is not a %s", "&")
	}
	return &Unnamed_Ampersand{Node: *node}, nil
}

type Unnamed_AmpersandEq struct {
	tree_sitter.Node
}

func NewUnnamed_AmpersandEq(node *tree_sitter.Node) (*Unnamed_AmpersandEq, error) {
	if node.Kind() != "&=" {
		return nil, fmt.Errorf("Node is not a %s", "&=")
	}
	return &Unnamed_AmpersandEq{Node: *node}, nil
}

type Unnamed_LParen struct {
	tree_sitter.Node
}

func NewUnnamed_LParen(node *tree_sitter.Node) (*Unnamed_LParen, error) {
	if node.Kind() != "(" {
		return nil, fmt.Errorf("Node is not a %s", "(")
	}
	return &Unnamed_LParen{Node: *node}, nil
}

type Unnamed_RParen struct {
	tree_sitter.Node
}

func NewUnnamed_RParen(node *tree_sitter.Node) (*Unnamed_RParen, error) {
	if node.Kind() != ")" {
		return nil, fmt.Errorf("Node is not a %s", ")")
	}
	return &Unnamed_RParen{Node: *node}, nil
}

type Unnamed_Mul struct {
	tree_sitter.Node
}

func NewUnnamed_Mul(node *tree_sitter.Node) (*Unnamed_Mul, error) {
	if node.Kind() != "*" {
		return nil, fmt.Errorf("Node is not a %s", "*")
	}
	return &Unnamed_Mul{Node: *node}, nil
}

type Unnamed_MulMul struct {
	tree_sitter.Node
}

func NewUnnamed_MulMul(node *tree_sitter.Node) (*Unnamed_MulMul, error) {
	if node.Kind() != "**" {
		return nil, fmt.Errorf("Node is not a %s", "**")
	}
	return &Unnamed_MulMul{Node: *node}, nil
}

type Unnamed_MulMulEq struct {
	tree_sitter.Node
}

func NewUnnamed_MulMulEq(node *tree_sitter.Node) (*Unnamed_MulMulEq, error) {
	if node.Kind() != "**=" {
		return nil, fmt.Errorf("Node is not a %s", "**=")
	}
	return &Unnamed_MulMulEq{Node: *node}, nil
}

type Unnamed_MulEq struct {
	tree_sitter.Node
}

func NewUnnamed_MulEq(node *tree_sitter.Node) (*Unnamed_MulEq, error) {
	if node.Kind() != "*=" {
		return nil, fmt.Errorf("Node is not a %s", "*=")
	}
	return &Unnamed_MulEq{Node: *node}, nil
}

type Unnamed_Add struct {
	tree_sitter.Node
}

func NewUnnamed_Add(node *tree_sitter.Node) (*Unnamed_Add, error) {
	if node.Kind() != "+" {
		return nil, fmt.Errorf("Node is not a %s", "+")
	}
	return &Unnamed_Add{Node: *node}, nil
}

type Unnamed_AddEq struct {
	tree_sitter.Node
}

func NewUnnamed_AddEq(node *tree_sitter.Node) (*Unnamed_AddEq, error) {
	if node.Kind() != "+=" {
		return nil, fmt.Errorf("Node is not a %s", "+=")
	}
	return &Unnamed_AddEq{Node: *node}, nil
}

type Unnamed_Comma struct {
	tree_sitter.Node
}

func NewUnnamed_Comma(node *tree_sitter.Node) (*Unnamed_Comma, error) {
	if node.Kind() != "," {
		return nil, fmt.Errorf("Node is not a %s", ",")
	}
	return &Unnamed_Comma{Node: *node}, nil
}

type Unnamed_Sub struct {
	tree_sitter.Node
}

func NewUnnamed_Sub(node *tree_sitter.Node) (*Unnamed_Sub, error) {
	if node.Kind() != "-" {
		return nil, fmt.Errorf("Node is not a %s", "-")
	}
	return &Unnamed_Sub{Node: *node}, nil
}

type Unnamed_SubEq struct {
	tree_sitter.Node
}

func NewUnnamed_SubEq(node *tree_sitter.Node) (*Unnamed_SubEq, error) {
	if node.Kind() != "-=" {
		return nil, fmt.Errorf("Node is not a %s", "-=")
	}
	return &Unnamed_SubEq{Node: *node}, nil
}

type Unnamed_SubGt struct {
	tree_sitter.Node
}

func NewUnnamed_SubGt(node *tree_sitter.Node) (*Unnamed_SubGt, error) {
	if node.Kind() != "->" {
		return nil, fmt.Errorf("Node is not a %s", "->")
	}
	return &Unnamed_SubGt{Node: *node}, nil
}

type Unnamed_Dot struct {
	tree_sitter.Node
}

func NewUnnamed_Dot(node *tree_sitter.Node) (*Unnamed_Dot, error) {
	if node.Kind() != "." {
		return nil, fmt.Errorf("Node is not a %s", ".")
	}
	return &Unnamed_Dot{Node: *node}, nil
}

type Unnamed_Div struct {
	tree_sitter.Node
}

func NewUnnamed_Div(node *tree_sitter.Node) (*Unnamed_Div, error) {
	if node.Kind() != "/" {
		return nil, fmt.Errorf("Node is not a %s", "/")
	}
	return &Unnamed_Div{Node: *node}, nil
}

type Unnamed_DivDiv struct {
	tree_sitter.Node
}

func NewUnnamed_DivDiv(node *tree_sitter.Node) (*Unnamed_DivDiv, error) {
	if node.Kind() != "//" {
		return nil, fmt.Errorf("Node is not a %s", "//")
	}
	return &Unnamed_DivDiv{Node: *node}, nil
}

type Unnamed_DivDivEq struct {
	tree_sitter.Node
}

func NewUnnamed_DivDivEq(node *tree_sitter.Node) (*Unnamed_DivDivEq, error) {
	if node.Kind() != "//=" {
		return nil, fmt.Errorf("Node is not a %s", "//=")
	}
	return &Unnamed_DivDivEq{Node: *node}, nil
}

type Unnamed_DivEq struct {
	tree_sitter.Node
}

func NewUnnamed_DivEq(node *tree_sitter.Node) (*Unnamed_DivEq, error) {
	if node.Kind() != "/=" {
		return nil, fmt.Errorf("Node is not a %s", "/=")
	}
	return &Unnamed_DivEq{Node: *node}, nil
}

type Unnamed_Colon struct {
	tree_sitter.Node
}

func NewUnnamed_Colon(node *tree_sitter.Node) (*Unnamed_Colon, error) {
	if node.Kind() != ":" {
		return nil, fmt.Errorf("Node is not a %s", ":")
	}
	return &Unnamed_Colon{Node: *node}, nil
}

type Unnamed_ColonEq struct {
	tree_sitter.Node
}

func NewUnnamed_ColonEq(node *tree_sitter.Node) (*Unnamed_ColonEq, error) {
	if node.Kind() != ":=" {
		return nil, fmt.Errorf("Node is not a %s", ":=")
	}
	return &Unnamed_ColonEq{Node: *node}, nil
}

type Unnamed_Semicolon struct {
	tree_sitter.Node
}

func NewUnnamed_Semicolon(node *tree_sitter.Node) (*Unnamed_Semicolon, error) {
	if node.Kind() != ";" {
		return nil, fmt.Errorf("Node is not a %s", ";")
	}
	return &Unnamed_Semicolon{Node: *node}, nil
}

type Unnamed_Lt struct {
	tree_sitter.Node
}

func NewUnnamed_Lt(node *tree_sitter.Node) (*Unnamed_Lt, error) {
	if node.Kind() != "<" {
		return nil, fmt.Errorf("Node is not a %s", "<")
	}
	return &Unnamed_Lt{Node: *node}, nil
}

type Unnamed_LtLt struct {
	tree_sitter.Node
}

func NewUnnamed_LtLt(node *tree_sitter.Node) (*Unnamed_LtLt, error) {
	if node.Kind() != "<<" {
		return nil, fmt.Errorf("Node is not a %s", "<<")
	}
	return &Unnamed_LtLt{Node: *node}, nil
}

type Unnamed_LtLtEq struct {
	tree_sitter.Node
}

func NewUnnamed_LtLtEq(node *tree_sitter.Node) (*Unnamed_LtLtEq, error) {
	if node.Kind() != "<<=" {
		return nil, fmt.Errorf("Node is not a %s", "<<=")
	}
	return &Unnamed_LtLtEq{Node: *node}, nil
}

type Unnamed_LtEq struct {
	tree_sitter.Node
}

func NewUnnamed_LtEq(node *tree_sitter.Node) (*Unnamed_LtEq, error) {
	if node.Kind() != "<=" {
		return nil, fmt.Errorf("Node is not a %s", "<=")
	}
	return &Unnamed_LtEq{Node: *node}, nil
}

type Unnamed_LtGt struct {
	tree_sitter.Node
}

func NewUnnamed_LtGt(node *tree_sitter.Node) (*Unnamed_LtGt, error) {
	if node.Kind() != "<>" {
		return nil, fmt.Errorf("Node is not a %s", "<>")
	}
	return &Unnamed_LtGt{Node: *node}, nil
}

type Unnamed_Eq struct {
	tree_sitter.Node
}

func NewUnnamed_Eq(node *tree_sitter.Node) (*Unnamed_Eq, error) {
	if node.Kind() != "=" {
		return nil, fmt.Errorf("Node is not a %s", "=")
	}
	return &Unnamed_Eq{Node: *node}, nil
}

type Unnamed_EqEq struct {
	tree_sitter.Node
}

func NewUnnamed_EqEq(node *tree_sitter.Node) (*Unnamed_EqEq, error) {
	if node.Kind() != "==" {
		return nil, fmt.Errorf("Node is not a %s", "==")
	}
	return &Unnamed_EqEq{Node: *node}, nil
}

type Unnamed_Gt struct {
	tree_sitter.Node
}

func NewUnnamed_Gt(node *tree_sitter.Node) (*Unnamed_Gt, error) {
	if node.Kind() != ">" {
		return nil, fmt.Errorf("Node is not a %s", ">")
	}
	return &Unnamed_Gt{Node: *node}, nil
}

type Unnamed_GtEq struct {
	tree_sitter.Node
}

func NewUnnamed_GtEq(node *tree_sitter.Node) (*Unnamed_GtEq, error) {
	if node.Kind() != ">=" {
		return nil, fmt.Errorf("Node is not a %s", ">=")
	}
	return &Unnamed_GtEq{Node: *node}, nil
}

type Unnamed_GtGt struct {
	tree_sitter.Node
}

func NewUnnamed_GtGt(node *tree_sitter.Node) (*Unnamed_GtGt, error) {
	if node.Kind() != ">>" {
		return nil, fmt.Errorf("Node is not a %s", ">>")
	}
	return &Unnamed_GtGt{Node: *node}, nil
}

type Unnamed_GtGtEq struct {
	tree_sitter.Node
}

func NewUnnamed_GtGtEq(node *tree_sitter.Node) (*Unnamed_GtGtEq, error) {
	if node.Kind() != ">>=" {
		return nil, fmt.Errorf("Node is not a %s", ">>=")
	}
	return &Unnamed_GtGtEq{Node: *node}, nil
}

type Unnamed_At struct {
	tree_sitter.Node
}

func NewUnnamed_At(node *tree_sitter.Node) (*Unnamed_At, error) {
	if node.Kind() != "@" {
		return nil, fmt.Errorf("Node is not a %s", "@")
	}
	return &Unnamed_At{Node: *node}, nil
}

type Unnamed_AtEq struct {
	tree_sitter.Node
}

func NewUnnamed_AtEq(node *tree_sitter.Node) (*Unnamed_AtEq, error) {
	if node.Kind() != "@=" {
		return nil, fmt.Errorf("Node is not a %s", "@=")
	}
	return &Unnamed_AtEq{Node: *node}, nil
}

type Unnamed_LBracket struct {
	tree_sitter.Node
}

func NewUnnamed_LBracket(node *tree_sitter.Node) (*Unnamed_LBracket, error) {
	if node.Kind() != "[" {
		return nil, fmt.Errorf("Node is not a %s", "[")
	}
	return &Unnamed_LBracket{Node: *node}, nil
}

type Unnamed_Backslash struct {
	tree_sitter.Node
}

func NewUnnamed_Backslash(node *tree_sitter.Node) (*Unnamed_Backslash, error) {
	if node.Kind() != "\\" {
		return nil, fmt.Errorf("Node is not a %s", "\\")
	}
	return &Unnamed_Backslash{Node: *node}, nil
}

type Unnamed_RBracket struct {
	tree_sitter.Node
}

func NewUnnamed_RBracket(node *tree_sitter.Node) (*Unnamed_RBracket, error) {
	if node.Kind() != "]" {
		return nil, fmt.Errorf("Node is not a %s", "]")
	}
	return &Unnamed_RBracket{Node: *node}, nil
}

type Unnamed_BitXor struct {
	tree_sitter.Node
}

func NewUnnamed_BitXor(node *tree_sitter.Node) (*Unnamed_BitXor, error) {
	if node.Kind() != "^" {
		return nil, fmt.Errorf("Node is not a %s", "^")
	}
	return &Unnamed_BitXor{Node: *node}, nil
}

type Unnamed_BitXorEq struct {
	tree_sitter.Node
}

func NewUnnamed_BitXorEq(node *tree_sitter.Node) (*Unnamed_BitXorEq, error) {
	if node.Kind() != "^=" {
		return nil, fmt.Errorf("Node is not a %s", "^=")
	}
	return &Unnamed_BitXorEq{Node: *node}, nil
}

type Unnamed_Underscore struct {
	tree_sitter.Node
}

func NewUnnamed_Underscore(node *tree_sitter.Node) (*Unnamed_Underscore, error) {
	if node.Kind() != "_" {
		return nil, fmt.Errorf("Node is not a %s", "_")
	}
	return &Unnamed_Underscore{Node: *node}, nil
}

type Unnamed_Future struct {
	tree_sitter.Node
}

func NewUnnamed_Future(node *tree_sitter.Node) (*Unnamed_Future, error) {
	if node.Kind() != "__future__" {
		return nil, fmt.Errorf("Node is not a %s", "__future__")
	}
	return &Unnamed_Future{Node: *node}, nil
}

type Unnamed_And struct {
	tree_sitter.Node
}

func NewUnnamed_And(node *tree_sitter.Node) (*Unnamed_And, error) {
	if node.Kind() != "and" {
		return nil, fmt.Errorf("Node is not a %s", "and")
	}
	return &Unnamed_And{Node: *node}, nil
}

type Unnamed_As struct {
	tree_sitter.Node
}

func NewUnnamed_As(node *tree_sitter.Node) (*Unnamed_As, error) {
	if node.Kind() != "as" {
		return nil, fmt.Errorf("Node is not a %s", "as")
	}
	return &Unnamed_As{Node: *node}, nil
}

type Unnamed_Assert struct {
	tree_sitter.Node
}

func NewUnnamed_Assert(node *tree_sitter.Node) (*Unnamed_Assert, error) {
	if node.Kind() != "assert" {
		return nil, fmt.Errorf("Node is not a %s", "assert")
	}
	return &Unnamed_Assert{Node: *node}, nil
}

type Unnamed_Async struct {
	tree_sitter.Node
}

func NewUnnamed_Async(node *tree_sitter.Node) (*Unnamed_Async, error) {
	if node.Kind() != "async" {
		return nil, fmt.Errorf("Node is not a %s", "async")
	}
	return &Unnamed_Async{Node: *node}, nil
}

type Unnamed_Await struct {
	tree_sitter.Node
}

func NewUnnamed_Await(node *tree_sitter.Node) (*Unnamed_Await, error) {
	if node.Kind() != "await" {
		return nil, fmt.Errorf("Node is not a %s", "await")
	}
	return &Unnamed_Await{Node: *node}, nil
}

type Unnamed_Break struct {
	tree_sitter.Node
}

func NewUnnamed_Break(node *tree_sitter.Node) (*Unnamed_Break, error) {
	if node.Kind() != "break" {
		return nil, fmt.Errorf("Node is not a %s", "break")
	}
	return &Unnamed_Break{Node: *node}, nil
}

type Unnamed_Case struct {
	tree_sitter.Node
}

func NewUnnamed_Case(node *tree_sitter.Node) (*Unnamed_Case, error) {
	if node.Kind() != "case" {
		return nil, fmt.Errorf("Node is not a %s", "case")
	}
	return &Unnamed_Case{Node: *node}, nil
}

type Unnamed_Class struct {
	tree_sitter.Node
}

func NewUnnamed_Class(node *tree_sitter.Node) (*Unnamed_Class, error) {
	if node.Kind() != "class" {
		return nil, fmt.Errorf("Node is not a %s", "class")
	}
	return &Unnamed_Class{Node: *node}, nil
}

type Comment struct {
	tree_sitter.Node
}

func NewComment(node *tree_sitter.Node) (*Comment, error) {
	if node.Kind() != "comment" {
		return nil, fmt.Errorf("Node is not a %s", "comment")
	}
	return &Comment{Node: *node}, nil
}

type Unnamed_Continue struct {
	tree_sitter.Node
}

func NewUnnamed_Continue(node *tree_sitter.Node) (*Unnamed_Continue, error) {
	if node.Kind() != "continue" {
		return nil, fmt.Errorf("Node is not a %s", "continue")
	}
	return &Unnamed_Continue{Node: *node}, nil
}

type Unnamed_Def struct {
	tree_sitter.Node
}

func NewUnnamed_Def(node *tree_sitter.Node) (*Unnamed_Def, error) {
	if node.Kind() != "def" {
		return nil, fmt.Errorf("Node is not a %s", "def")
	}
	return &Unnamed_Def{Node: *node}, nil
}

type Unnamed_Del struct {
	tree_sitter.Node
}

func NewUnnamed_Del(node *tree_sitter.Node) (*Unnamed_Del, error) {
	if node.Kind() != "del" {
		return nil, fmt.Errorf("Node is not a %s", "del")
	}
	return &Unnamed_Del{Node: *node}, nil
}

type Unnamed_Elif struct {
	tree_sitter.Node
}

func NewUnnamed_Elif(node *tree_sitter.Node) (*Unnamed_Elif, error) {
	if node.Kind() != "elif" {
		return nil, fmt.Errorf("Node is not a %s", "elif")
	}
	return &Unnamed_Elif{Node: *node}, nil
}

type Ellipsis struct {
	tree_sitter.Node
}

func NewEllipsis(node *tree_sitter.Node) (*Ellipsis, error) {
	if node.Kind() != "ellipsis" {
		return nil, fmt.Errorf("Node is not a %s", "ellipsis")
	}
	return &Ellipsis{Node: *node}, nil
}

type Unnamed_Else struct {
	tree_sitter.Node
}

func NewUnnamed_Else(node *tree_sitter.Node) (*Unnamed_Else, error) {
	if node.Kind() != "else" {
		return nil, fmt.Errorf("Node is not a %s", "else")
	}
	return &Unnamed_Else{Node: *node}, nil
}

type EscapeInterpolation struct {
	tree_sitter.Node
}

func NewEscapeInterpolation(node *tree_sitter.Node) (*EscapeInterpolation, error) {
	if node.Kind() != "escape_interpolation" {
		return nil, fmt.Errorf("Node is not a %s", "escape_interpolation")
	}
	return &EscapeInterpolation{Node: *node}, nil
}

type EscapeSequence struct {
	tree_sitter.Node
}

func NewEscapeSequence(node *tree_sitter.Node) (*EscapeSequence, error) {
	if node.Kind() != "escape_sequence" {
		return nil, fmt.Errorf("Node is not a %s", "escape_sequence")
	}
	return &EscapeSequence{Node: *node}, nil
}

type Unnamed_Except struct {
	tree_sitter.Node
}

func NewUnnamed_Except(node *tree_sitter.Node) (*Unnamed_Except, error) {
	if node.Kind() != "except" {
		return nil, fmt.Errorf("Node is not a %s", "except")
	}
	return &Unnamed_Except{Node: *node}, nil
}

type Unnamed_ExceptMul struct {
	tree_sitter.Node
}

func NewUnnamed_ExceptMul(node *tree_sitter.Node) (*Unnamed_ExceptMul, error) {
	if node.Kind() != "except*" {
		return nil, fmt.Errorf("Node is not a %s", "except*")
	}
	return &Unnamed_ExceptMul{Node: *node}, nil
}

type Unnamed_Exec struct {
	tree_sitter.Node
}

func NewUnnamed_Exec(node *tree_sitter.Node) (*Unnamed_Exec, error) {
	if node.Kind() != "exec" {
		return nil, fmt.Errorf("Node is not a %s", "exec")
	}
	return &Unnamed_Exec{Node: *node}, nil
}

type False struct {
	tree_sitter.Node
}

func NewFalse(node *tree_sitter.Node) (*False, error) {
	if node.Kind() != "false" {
		return nil, fmt.Errorf("Node is not a %s", "false")
	}
	return &False{Node: *node}, nil
}

type Unnamed_Finally struct {
	tree_sitter.Node
}

func NewUnnamed_Finally(node *tree_sitter.Node) (*Unnamed_Finally, error) {
	if node.Kind() != "finally" {
		return nil, fmt.Errorf("Node is not a %s", "finally")
	}
	return &Unnamed_Finally{Node: *node}, nil
}

type Float struct {
	tree_sitter.Node
}

func NewFloat(node *tree_sitter.Node) (*Float, error) {
	if node.Kind() != "float" {
		return nil, fmt.Errorf("Node is not a %s", "float")
	}
	return &Float{Node: *node}, nil
}

type Unnamed_For struct {
	tree_sitter.Node
}

func NewUnnamed_For(node *tree_sitter.Node) (*Unnamed_For, error) {
	if node.Kind() != "for" {
		return nil, fmt.Errorf("Node is not a %s", "for")
	}
	return &Unnamed_For{Node: *node}, nil
}

type Unnamed_From struct {
	tree_sitter.Node
}

func NewUnnamed_From(node *tree_sitter.Node) (*Unnamed_From, error) {
	if node.Kind() != "from" {
		return nil, fmt.Errorf("Node is not a %s", "from")
	}
	return &Unnamed_From{Node: *node}, nil
}

type Unnamed_Global struct {
	tree_sitter.Node
}

func NewUnnamed_Global(node *tree_sitter.Node) (*Unnamed_Global, error) {
	if node.Kind() != "global" {
		return nil, fmt.Errorf("Node is not a %s", "global")
	}
	return &Unnamed_Global{Node: *node}, nil
}

type Identifier struct {
	tree_sitter.Node
}

func NewIdentifier(node *tree_sitter.Node) (*Identifier, error) {
	if node.Kind() != "identifier" {
		return nil, fmt.Errorf("Node is not a %s", "identifier")
	}
	return &Identifier{Node: *node}, nil
}

type Unnamed_If struct {
	tree_sitter.Node
}

func NewUnnamed_If(node *tree_sitter.Node) (*Unnamed_If, error) {
	if node.Kind() != "if" {
		return nil, fmt.Errorf("Node is not a %s", "if")
	}
	return &Unnamed_If{Node: *node}, nil
}

type Unnamed_Import struct {
	tree_sitter.Node
}

func NewUnnamed_Import(node *tree_sitter.Node) (*Unnamed_Import, error) {
	if node.Kind() != "import" {
		return nil, fmt.Errorf("Node is not a %s", "import")
	}
	return &Unnamed_Import{Node: *node}, nil
}

type Unnamed_In struct {
	tree_sitter.Node
}

func NewUnnamed_In(node *tree_sitter.Node) (*Unnamed_In, error) {
	if node.Kind() != "in" {
		return nil, fmt.Errorf("Node is not a %s", "in")
	}
	return &Unnamed_In{Node: *node}, nil
}

type Integer struct {
	tree_sitter.Node
}

func NewInteger(node *tree_sitter.Node) (*Integer, error) {
	if node.Kind() != "integer" {
		return nil, fmt.Errorf("Node is not a %s", "integer")
	}
	return &Integer{Node: *node}, nil
}

type Unnamed_Is struct {
	tree_sitter.Node
}

func NewUnnamed_Is(node *tree_sitter.Node) (*Unnamed_Is, error) {
	if node.Kind() != "is" {
		return nil, fmt.Errorf("Node is not a %s", "is")
	}
	return &Unnamed_Is{Node: *node}, nil
}

type Unnamed_Lambda struct {
	tree_sitter.Node
}

func NewUnnamed_Lambda(node *tree_sitter.Node) (*Unnamed_Lambda, error) {
	if node.Kind() != "lambda" {
		return nil, fmt.Errorf("Node is not a %s", "lambda")
	}
	return &Unnamed_Lambda{Node: *node}, nil
}

type LineContinuation struct {
	tree_sitter.Node
}

func NewLineContinuation(node *tree_sitter.Node) (*LineContinuation, error) {
	if node.Kind() != "line_continuation" {
		return nil, fmt.Errorf("Node is not a %s", "line_continuation")
	}
	return &LineContinuation{Node: *node}, nil
}

type Unnamed_Match struct {
	tree_sitter.Node
}

func NewUnnamed_Match(node *tree_sitter.Node) (*Unnamed_Match, error) {
	if node.Kind() != "match" {
		return nil, fmt.Errorf("Node is not a %s", "match")
	}
	return &Unnamed_Match{Node: *node}, nil
}

type None struct {
	tree_sitter.Node
}

func NewNone(node *tree_sitter.Node) (*None, error) {
	if node.Kind() != "none" {
		return nil, fmt.Errorf("Node is not a %s", "none")
	}
	return &None{Node: *node}, nil
}

type Unnamed_Nonlocal struct {
	tree_sitter.Node
}

func NewUnnamed_Nonlocal(node *tree_sitter.Node) (*Unnamed_Nonlocal, error) {
	if node.Kind() != "nonlocal" {
		return nil, fmt.Errorf("Node is not a %s", "nonlocal")
	}
	return &Unnamed_Nonlocal{Node: *node}, nil
}

type Unnamed_Not struct {
	tree_sitter.Node
}

func NewUnnamed_Not(node *tree_sitter.Node) (*Unnamed_Not, error) {
	if node.Kind() != "not" {
		return nil, fmt.Errorf("Node is not a %s", "not")
	}
	return &Unnamed_Not{Node: *node}, nil
}

type Unnamed_Or struct {
	tree_sitter.Node
}

func NewUnnamed_Or(node *tree_sitter.Node) (*Unnamed_Or, error) {
	if node.Kind() != "or" {
		return nil, fmt.Errorf("Node is not a %s", "or")
	}
	return &Unnamed_Or{Node: *node}, nil
}

type Unnamed_Pass struct {
	tree_sitter.Node
}

func NewUnnamed_Pass(node *tree_sitter.Node) (*Unnamed_Pass, error) {
	if node.Kind() != "pass" {
		return nil, fmt.Errorf("Node is not a %s", "pass")
	}
	return &Unnamed_Pass{Node: *node}, nil
}

type Unnamed_Print struct {
	tree_sitter.Node
}

func NewUnnamed_Print(node *tree_sitter.Node) (*Unnamed_Print, error) {
	if node.Kind() != "print" {
		return nil, fmt.Errorf("Node is not a %s", "print")
	}
	return &Unnamed_Print{Node: *node}, nil
}

type Unnamed_Raise struct {
	tree_sitter.Node
}

func NewUnnamed_Raise(node *tree_sitter.Node) (*Unnamed_Raise, error) {
	if node.Kind() != "raise" {
		return nil, fmt.Errorf("Node is not a %s", "raise")
	}
	return &Unnamed_Raise{Node: *node}, nil
}

type Unnamed_Return struct {
	tree_sitter.Node
}

func NewUnnamed_Return(node *tree_sitter.Node) (*Unnamed_Return, error) {
	if node.Kind() != "return" {
		return nil, fmt.Errorf("Node is not a %s", "return")
	}
	return &Unnamed_Return{Node: *node}, nil
}

type StringEnd struct {
	tree_sitter.Node
}

func NewStringEnd(node *tree_sitter.Node) (*StringEnd, error) {
	if node.Kind() != "string_end" {
		return nil, fmt.Errorf("Node is not a %s", "string_end")
	}
	return &StringEnd{Node: *node}, nil
}

type StringStart struct {
	tree_sitter.Node
}

func NewStringStart(node *tree_sitter.Node) (*StringStart, error) {
	if node.Kind() != "string_start" {
		return nil, fmt.Errorf("Node is not a %s", "string_start")
	}
	return &StringStart{Node: *node}, nil
}

type True struct {
	tree_sitter.Node
}

func NewTrue(node *tree_sitter.Node) (*True, error) {
	if node.Kind() != "true" {
		return nil, fmt.Errorf("Node is not a %s", "true")
	}
	return &True{Node: *node}, nil
}

type Unnamed_Try struct {
	tree_sitter.Node
}

func NewUnnamed_Try(node *tree_sitter.Node) (*Unnamed_Try, error) {
	if node.Kind() != "try" {
		return nil, fmt.Errorf("Node is not a %s", "try")
	}
	return &Unnamed_Try{Node: *node}, nil
}

type Unnamed_Type struct {
	tree_sitter.Node
}

func NewUnnamed_Type(node *tree_sitter.Node) (*Unnamed_Type, error) {
	if node.Kind() != "type" {
		return nil, fmt.Errorf("Node is not a %s", "type")
	}
	return &Unnamed_Type{Node: *node}, nil
}

type TypeConversion struct {
	tree_sitter.Node
}

func NewTypeConversion(node *tree_sitter.Node) (*TypeConversion, error) {
	if node.Kind() != "type_conversion" {
		return nil, fmt.Errorf("Node is not a %s", "type_conversion")
	}
	return &TypeConversion{Node: *node}, nil
}

type Unnamed_While struct {
	tree_sitter.Node
}

func NewUnnamed_While(node *tree_sitter.Node) (*Unnamed_While, error) {
	if node.Kind() != "while" {
		return nil, fmt.Errorf("Node is not a %s", "while")
	}
	return &Unnamed_While{Node: *node}, nil
}

type Unnamed_With struct {
	tree_sitter.Node
}

func NewUnnamed_With(node *tree_sitter.Node) (*Unnamed_With, error) {
	if node.Kind() != "with" {
		return nil, fmt.Errorf("Node is not a %s", "with")
	}
	return &Unnamed_With{Node: *node}, nil
}

type Unnamed_Yield struct {
	tree_sitter.Node
}

func NewUnnamed_Yield(node *tree_sitter.Node) (*Unnamed_Yield, error) {
	if node.Kind() != "yield" {
		return nil, fmt.Errorf("Node is not a %s", "yield")
	}
	return &Unnamed_Yield{Node: *node}, nil
}

type Unnamed_LBrace struct {
	tree_sitter.Node
}

func NewUnnamed_LBrace(node *tree_sitter.Node) (*Unnamed_LBrace, error) {
	if node.Kind() != "{" {
		return nil, fmt.Errorf("Node is not a %s", "{")
	}
	return &Unnamed_LBrace{Node: *node}, nil
}

type Unnamed_Bar struct {
	tree_sitter.Node
}

func NewUnnamed_Bar(node *tree_sitter.Node) (*Unnamed_Bar, error) {
	if node.Kind() != "|" {
		return nil, fmt.Errorf("Node is not a %s", "|")
	}
	return &Unnamed_Bar{Node: *node}, nil
}

type Unnamed_BarEq struct {
	tree_sitter.Node
}

func NewUnnamed_BarEq(node *tree_sitter.Node) (*Unnamed_BarEq, error) {
	if node.Kind() != "|=" {
		return nil, fmt.Errorf("Node is not a %s", "|=")
	}
	return &Unnamed_BarEq{Node: *node}, nil
}

type Unnamed_RBrace struct {
	tree_sitter.Node
}

func NewUnnamed_RBrace(node *tree_sitter.Node) (*Unnamed_RBrace, error) {
	if node.Kind() != "}" {
		return nil, fmt.Errorf("Node is not a %s", "}")
	}
	return &Unnamed_RBrace{Node: *node}, nil
}

type Unnamed_BitNot struct {
	tree_sitter.Node
}

func NewUnnamed_BitNot(node *tree_sitter.Node) (*Unnamed_BitNot, error) {
	if node.Kind() != "~" {
		return nil, fmt.Errorf("Node is not a %s", "~")
	}
	return &Unnamed_BitNot{Node: *node}, nil
}

type CompoundStatement struct {
	tree_sitter.Node
}

func (c *CompoundStatement) ClassDefinition() (*ClassDefinition, error) {
	tsKinds := []string{"class_definition"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &ClassDefinition{Node: c.Node}, nil
}
func (c *CompoundStatement) DecoratedDefinition() (*DecoratedDefinition, error) {
	tsKinds := []string{"decorated_definition"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &DecoratedDefinition{Node: c.Node}, nil
}
func (c *CompoundStatement) ForStatement() (*ForStatement, error) {
	tsKinds := []string{"for_statement"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &ForStatement{Node: c.Node}, nil
}
func (c *CompoundStatement) FunctionDefinition() (*FunctionDefinition, error) {
	tsKinds := []string{"function_definition"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &FunctionDefinition{Node: c.Node}, nil
}
func (c *CompoundStatement) IfStatement() (*IfStatement, error) {
	tsKinds := []string{"if_statement"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &IfStatement{Node: c.Node}, nil
}
func (c *CompoundStatement) MatchStatement() (*MatchStatement, error) {
	tsKinds := []string{"match_statement"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &MatchStatement{Node: c.Node}, nil
}
func (c *CompoundStatement) TryStatement() (*TryStatement, error) {
	tsKinds := []string{"try_statement"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &TryStatement{Node: c.Node}, nil
}
func (c *CompoundStatement) WhileStatement() (*WhileStatement, error) {
	tsKinds := []string{"while_statement"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &WhileStatement{Node: c.Node}, nil
}
func (c *CompoundStatement) WithStatement() (*WithStatement, error) {
	tsKinds := []string{"with_statement"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &WithStatement{Node: c.Node}, nil
}

type SimpleStatement struct {
	tree_sitter.Node
}

func (s *SimpleStatement) AssertStatement() (*AssertStatement, error) {
	tsKinds := []string{"assert_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &AssertStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) BreakStatement() (*BreakStatement, error) {
	tsKinds := []string{"break_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &BreakStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) ContinueStatement() (*ContinueStatement, error) {
	tsKinds := []string{"continue_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &ContinueStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) DeleteStatement() (*DeleteStatement, error) {
	tsKinds := []string{"delete_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &DeleteStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) ExecStatement() (*ExecStatement, error) {
	tsKinds := []string{"exec_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &ExecStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) ExpressionStatement() (*ExpressionStatement, error) {
	tsKinds := []string{"expression_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &ExpressionStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) FutureImportStatement() (*FutureImportStatement, error) {
	tsKinds := []string{"future_import_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &FutureImportStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) GlobalStatement() (*GlobalStatement, error) {
	tsKinds := []string{"global_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &GlobalStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) ImportFromStatement() (*ImportFromStatement, error) {
	tsKinds := []string{"import_from_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &ImportFromStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) ImportStatement() (*ImportStatement, error) {
	tsKinds := []string{"import_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &ImportStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) NonlocalStatement() (*NonlocalStatement, error) {
	tsKinds := []string{"nonlocal_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &NonlocalStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) PassStatement() (*PassStatement, error) {
	tsKinds := []string{"pass_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &PassStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) PrintStatement() (*PrintStatement, error) {
	tsKinds := []string{"print_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &PrintStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) RaiseStatement() (*RaiseStatement, error) {
	tsKinds := []string{"raise_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &RaiseStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) ReturnStatement() (*ReturnStatement, error) {
	tsKinds := []string{"return_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &ReturnStatement{Node: s.Node}, nil
}
func (s *SimpleStatement) TypeAliasStatement() (*TypeAliasStatement, error) {
	tsKinds := []string{"type_alias_statement"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &TypeAliasStatement{Node: s.Node}, nil
}

type Expression struct {
	tree_sitter.Node
}

func (e *Expression) AsPattern() (*AsPattern, error) {
	tsKinds := []string{"as_pattern"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &AsPattern{Node: e.Node}, nil
}
func (e *Expression) BooleanOperator() (*BooleanOperator, error) {
	tsKinds := []string{"boolean_operator"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &BooleanOperator{Node: e.Node}, nil
}
func (e *Expression) ComparisonOperator() (*ComparisonOperator, error) {
	tsKinds := []string{"comparison_operator"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &ComparisonOperator{Node: e.Node}, nil
}
func (e *Expression) ConditionalExpression() (*ConditionalExpression, error) {
	tsKinds := []string{"conditional_expression"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &ConditionalExpression{Node: e.Node}, nil
}
func (e *Expression) Lambda() (*Lambda, error) {
	tsKinds := []string{"lambda"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &Lambda{Node: e.Node}, nil
}
func (e *Expression) NamedExpression() (*NamedExpression, error) {
	tsKinds := []string{"named_expression"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &NamedExpression{Node: e.Node}, nil
}
func (e *Expression) NotOperator() (*NotOperator, error) {
	tsKinds := []string{"not_operator"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &NotOperator{Node: e.Node}, nil
}
func (e *Expression) PrimaryExpression() (*PrimaryExpression, error) {
	tsKinds := []string{"attribute", "await", "binary_operator", "call", "concatenated_string", "dictionary", "dictionary_comprehension", "ellipsis", "false", "float", "generator_expression", "identifier", "integer", "list", "list_comprehension", "list_splat", "none", "parenthesized_expression", "set", "set_comprehension", "string", "subscript", "true", "tuple", "unary_operator"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &PrimaryExpression{Node: e.Node}, nil
}

type Parameter struct {
	tree_sitter.Node
}

func (p *Parameter) DefaultParameter() (*DefaultParameter, error) {
	tsKinds := []string{"default_parameter"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &DefaultParameter{Node: p.Node}, nil
}
func (p *Parameter) DictionarySplatPattern() (*DictionarySplatPattern, error) {
	tsKinds := []string{"dictionary_splat_pattern"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &DictionarySplatPattern{Node: p.Node}, nil
}
func (p *Parameter) Identifier() (*Identifier, error) {
	tsKinds := []string{"identifier"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Identifier{Node: p.Node}, nil
}
func (p *Parameter) KeywordSeparator() (*KeywordSeparator, error) {
	tsKinds := []string{"keyword_separator"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &KeywordSeparator{Node: p.Node}, nil
}
func (p *Parameter) ListSplatPattern() (*ListSplatPattern, error) {
	tsKinds := []string{"list_splat_pattern"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &ListSplatPattern{Node: p.Node}, nil
}
func (p *Parameter) PositionalSeparator() (*PositionalSeparator, error) {
	tsKinds := []string{"positional_separator"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &PositionalSeparator{Node: p.Node}, nil
}
func (p *Parameter) TuplePattern() (*TuplePattern, error) {
	tsKinds := []string{"tuple_pattern"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &TuplePattern{Node: p.Node}, nil
}
func (p *Parameter) TypedDefaultParameter() (*TypedDefaultParameter, error) {
	tsKinds := []string{"typed_default_parameter"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &TypedDefaultParameter{Node: p.Node}, nil
}
func (p *Parameter) TypedParameter() (*TypedParameter, error) {
	tsKinds := []string{"typed_parameter"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &TypedParameter{Node: p.Node}, nil
}

type Pattern struct {
	tree_sitter.Node
}

func (p *Pattern) Attribute() (*Attribute, error) {
	tsKinds := []string{"attribute"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Attribute{Node: p.Node}, nil
}
func (p *Pattern) Identifier() (*Identifier, error) {
	tsKinds := []string{"identifier"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Identifier{Node: p.Node}, nil
}
func (p *Pattern) ListPattern() (*ListPattern, error) {
	tsKinds := []string{"list_pattern"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &ListPattern{Node: p.Node}, nil
}
func (p *Pattern) ListSplatPattern() (*ListSplatPattern, error) {
	tsKinds := []string{"list_splat_pattern"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &ListSplatPattern{Node: p.Node}, nil
}
func (p *Pattern) Subscript() (*Subscript, error) {
	tsKinds := []string{"subscript"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Subscript{Node: p.Node}, nil
}
func (p *Pattern) TuplePattern() (*TuplePattern, error) {
	tsKinds := []string{"tuple_pattern"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &TuplePattern{Node: p.Node}, nil
}

type PrimaryExpression struct {
	tree_sitter.Node
}

func (p *PrimaryExpression) Attribute() (*Attribute, error) {
	tsKinds := []string{"attribute"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Attribute{Node: p.Node}, nil
}
func (p *PrimaryExpression) Await() (*Await, error) {
	tsKinds := []string{"await"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Await{Node: p.Node}, nil
}
func (p *PrimaryExpression) BinaryOperator() (*BinaryOperator, error) {
	tsKinds := []string{"binary_operator"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &BinaryOperator{Node: p.Node}, nil
}
func (p *PrimaryExpression) Call() (*Call, error) {
	tsKinds := []string{"call"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Call{Node: p.Node}, nil
}
func (p *PrimaryExpression) ConcatenatedString() (*ConcatenatedString, error) {
	tsKinds := []string{"concatenated_string"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &ConcatenatedString{Node: p.Node}, nil
}
func (p *PrimaryExpression) Dictionary() (*Dictionary, error) {
	tsKinds := []string{"dictionary"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Dictionary{Node: p.Node}, nil
}
func (p *PrimaryExpression) DictionaryComprehension() (*DictionaryComprehension, error) {
	tsKinds := []string{"dictionary_comprehension"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &DictionaryComprehension{Node: p.Node}, nil
}
func (p *PrimaryExpression) Ellipsis() (*Ellipsis, error) {
	tsKinds := []string{"ellipsis"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Ellipsis{Node: p.Node}, nil
}
func (p *PrimaryExpression) False() (*False, error) {
	tsKinds := []string{"false"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &False{Node: p.Node}, nil
}
func (p *PrimaryExpression) Float() (*Float, error) {
	tsKinds := []string{"float"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Float{Node: p.Node}, nil
}
func (p *PrimaryExpression) GeneratorExpression() (*GeneratorExpression, error) {
	tsKinds := []string{"generator_expression"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &GeneratorExpression{Node: p.Node}, nil
}
func (p *PrimaryExpression) Identifier() (*Identifier, error) {
	tsKinds := []string{"identifier"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Identifier{Node: p.Node}, nil
}
func (p *PrimaryExpression) Integer() (*Integer, error) {
	tsKinds := []string{"integer"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Integer{Node: p.Node}, nil
}
func (p *PrimaryExpression) List() (*List, error) {
	tsKinds := []string{"list"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &List{Node: p.Node}, nil
}
func (p *PrimaryExpression) ListComprehension() (*ListComprehension, error) {
	tsKinds := []string{"list_comprehension"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &ListComprehension{Node: p.Node}, nil
}
func (p *PrimaryExpression) ListSplat() (*ListSplat, error) {
	tsKinds := []string{"list_splat"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &ListSplat{Node: p.Node}, nil
}
func (p *PrimaryExpression) None() (*None, error) {
	tsKinds := []string{"none"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &None{Node: p.Node}, nil
}
func (p *PrimaryExpression) ParenthesizedExpression() (*ParenthesizedExpression, error) {
	tsKinds := []string{"parenthesized_expression"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &ParenthesizedExpression{Node: p.Node}, nil
}
func (p *PrimaryExpression) Set() (*Set, error) {
	tsKinds := []string{"set"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Set{Node: p.Node}, nil
}
func (p *PrimaryExpression) SetComprehension() (*SetComprehension, error) {
	tsKinds := []string{"set_comprehension"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &SetComprehension{Node: p.Node}, nil
}
func (p *PrimaryExpression) String() (*String, error) {
	tsKinds := []string{"string"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &String{Node: p.Node}, nil
}
func (p *PrimaryExpression) Subscript() (*Subscript, error) {
	tsKinds := []string{"subscript"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Subscript{Node: p.Node}, nil
}
func (p *PrimaryExpression) True() (*True, error) {
	tsKinds := []string{"true"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &True{Node: p.Node}, nil
}
func (p *PrimaryExpression) Tuple() (*Tuple, error) {
	tsKinds := []string{"tuple"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Tuple{Node: p.Node}, nil
}
func (p *PrimaryExpression) UnaryOperator() (*UnaryOperator, error) {
	tsKinds := []string{"unary_operator"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &UnaryOperator{Node: p.Node}, nil
}

type dictionarySplat_expression_keywordArgument_listSplat_parenthesizedExpression struct {
	tree_sitter.Node
}

func (d *dictionarySplat_expression_keywordArgument_listSplat_parenthesizedExpression) DictionarySplat() (*DictionarySplat, error) {
	tsKinds := []string{"dictionary_splat"}
	if !slices.Contains(tsKinds, d.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", d.Node.Kind(), tsKinds)
	}
	return &DictionarySplat{Node: d.Node}, nil
}
func (d *dictionarySplat_expression_keywordArgument_listSplat_parenthesizedExpression) Expression() (*Expression, error) {
	tsKinds := []string{"as_pattern", "boolean_operator", "comparison_operator", "conditional_expression", "lambda", "named_expression", "not_operator", "attribute", "await", "binary_operator", "call", "concatenated_string", "dictionary", "dictionary_comprehension", "ellipsis", "false", "float", "generator_expression", "identifier", "integer", "list", "list_comprehension", "list_splat", "none", "parenthesized_expression", "set", "set_comprehension", "string", "subscript", "true", "tuple", "unary_operator"}
	if !slices.Contains(tsKinds, d.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", d.Node.Kind(), tsKinds)
	}
	return &Expression{Node: d.Node}, nil
}
func (d *dictionarySplat_expression_keywordArgument_listSplat_parenthesizedExpression) KeywordArgument() (*KeywordArgument, error) {
	tsKinds := []string{"keyword_argument"}
	if !slices.Contains(tsKinds, d.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", d.Node.Kind(), tsKinds)
	}
	return &KeywordArgument{Node: d.Node}, nil
}
func (d *dictionarySplat_expression_keywordArgument_listSplat_parenthesizedExpression) ListSplat() (*ListSplat, error) {
	tsKinds := []string{"list_splat"}
	if !slices.Contains(tsKinds, d.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", d.Node.Kind(), tsKinds)
	}
	return &ListSplat{Node: d.Node}, nil
}
func (d *dictionarySplat_expression_keywordArgument_listSplat_parenthesizedExpression) ParenthesizedExpression() (*ParenthesizedExpression, error) {
	tsKinds := []string{"parenthesized_expression"}
	if !slices.Contains(tsKinds, d.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", d.Node.Kind(), tsKinds)
	}
	return &ParenthesizedExpression{Node: d.Node}, nil
}

type casePattern_expression_identifier struct {
	tree_sitter.Node
}

func (c *casePattern_expression_identifier) CasePattern() (*CasePattern, error) {
	tsKinds := []string{"case_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &CasePattern{Node: c.Node}, nil
}
func (c *casePattern_expression_identifier) Expression() (*Expression, error) {
	tsKinds := []string{"as_pattern", "boolean_operator", "comparison_operator", "conditional_expression", "lambda", "named_expression", "not_operator", "attribute", "await", "binary_operator", "call", "concatenated_string", "dictionary", "dictionary_comprehension", "ellipsis", "false", "float", "generator_expression", "identifier", "integer", "list", "list_comprehension", "list_splat", "none", "parenthesized_expression", "set", "set_comprehension", "string", "subscript", "true", "tuple", "unary_operator"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &Expression{Node: c.Node}, nil
}
func (c *casePattern_expression_identifier) Identifier() (*Identifier, error) {
	tsKinds := []string{"identifier"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &Identifier{Node: c.Node}, nil
}

type pattern_patternList struct {
	tree_sitter.Node
}

func (p *pattern_patternList) Pattern() (*Pattern, error) {
	tsKinds := []string{"attribute", "identifier", "list_pattern", "list_splat_pattern", "subscript", "tuple_pattern"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &Pattern{Node: p.Node}, nil
}
func (p *pattern_patternList) PatternList() (*PatternList, error) {
	tsKinds := []string{"pattern_list"}
	if !slices.Contains(tsKinds, p.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", p.Node.Kind(), tsKinds)
	}
	return &PatternList{Node: p.Node}, nil
}

type assignment_augmentedAssignment_expression_expressionList_patternList_yield struct {
	tree_sitter.Node
}

func (a *assignment_augmentedAssignment_expression_expressionList_patternList_yield) Assignment() (*Assignment, error) {
	tsKinds := []string{"assignment"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Assignment{Node: a.Node}, nil
}
func (a *assignment_augmentedAssignment_expression_expressionList_patternList_yield) AugmentedAssignment() (*AugmentedAssignment, error) {
	tsKinds := []string{"augmented_assignment"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &AugmentedAssignment{Node: a.Node}, nil
}
func (a *assignment_augmentedAssignment_expression_expressionList_patternList_yield) Expression() (*Expression, error) {
	tsKinds := []string{"as_pattern", "boolean_operator", "comparison_operator", "conditional_expression", "lambda", "named_expression", "not_operator", "attribute", "await", "binary_operator", "call", "concatenated_string", "dictionary", "dictionary_comprehension", "ellipsis", "false", "float", "generator_expression", "identifier", "integer", "list", "list_comprehension", "list_splat", "none", "parenthesized_expression", "set", "set_comprehension", "string", "subscript", "true", "tuple", "unary_operator"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Expression{Node: a.Node}, nil
}
func (a *assignment_augmentedAssignment_expression_expressionList_patternList_yield) ExpressionList() (*ExpressionList, error) {
	tsKinds := []string{"expression_list"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &ExpressionList{Node: a.Node}, nil
}
func (a *assignment_augmentedAssignment_expression_expressionList_patternList_yield) PatternList() (*PatternList, error) {
	tsKinds := []string{"pattern_list"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &PatternList{Node: a.Node}, nil
}
func (a *assignment_augmentedAssignment_expression_expressionList_patternList_yield) Yield() (*Yield, error) {
	tsKinds := []string{"yield"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Yield{Node: a.Node}, nil
}

type modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq struct {
	tree_sitter.Node
}

func (m *modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq) ModEq() (*Unnamed_ModEq, error) {
	tsKinds := []string{"%="}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_ModEq{Node: m.Node}, nil
}
func (m *modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq) AmpersandEq() (*Unnamed_AmpersandEq, error) {
	tsKinds := []string{"&="}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_AmpersandEq{Node: m.Node}, nil
}
func (m *modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq) MulMulEq() (*Unnamed_MulMulEq, error) {
	tsKinds := []string{"**="}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_MulMulEq{Node: m.Node}, nil
}
func (m *modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq) MulEq() (*Unnamed_MulEq, error) {
	tsKinds := []string{"*="}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_MulEq{Node: m.Node}, nil
}
func (m *modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq) AddEq() (*Unnamed_AddEq, error) {
	tsKinds := []string{"+="}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_AddEq{Node: m.Node}, nil
}
func (m *modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq) SubEq() (*Unnamed_SubEq, error) {
	tsKinds := []string{"-="}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_SubEq{Node: m.Node}, nil
}
func (m *modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq) DivDivEq() (*Unnamed_DivDivEq, error) {
	tsKinds := []string{"//="}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_DivDivEq{Node: m.Node}, nil
}
func (m *modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq) DivEq() (*Unnamed_DivEq, error) {
	tsKinds := []string{"/="}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_DivEq{Node: m.Node}, nil
}
func (m *modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq) LtLtEq() (*Unnamed_LtLtEq, error) {
	tsKinds := []string{"<<="}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_LtLtEq{Node: m.Node}, nil
}
func (m *modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq) GtGtEq() (*Unnamed_GtGtEq, error) {
	tsKinds := []string{">>="}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_GtGtEq{Node: m.Node}, nil
}
func (m *modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq) AtEq() (*Unnamed_AtEq, error) {
	tsKinds := []string{"@="}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_AtEq{Node: m.Node}, nil
}
func (m *modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq) BitXorEq() (*Unnamed_BitXorEq, error) {
	tsKinds := []string{"^="}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_BitXorEq{Node: m.Node}, nil
}
func (m *modEq_ampersandEq_mulMulEq_mulEq_addEq_subEq_divDivEq_divEq_ltLtEq_gtGtEq_atEq_bitXorEq_barEq) BarEq() (*Unnamed_BarEq, error) {
	tsKinds := []string{"|="}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_BarEq{Node: m.Node}, nil
}

type mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar struct {
	tree_sitter.Node
}

func (m *mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar) Mod() (*Unnamed_Mod, error) {
	tsKinds := []string{"%"}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_Mod{Node: m.Node}, nil
}
func (m *mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar) Ampersand() (*Unnamed_Ampersand, error) {
	tsKinds := []string{"&"}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_Ampersand{Node: m.Node}, nil
}
func (m *mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar) Mul() (*Unnamed_Mul, error) {
	tsKinds := []string{"*"}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_Mul{Node: m.Node}, nil
}
func (m *mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar) MulMul() (*Unnamed_MulMul, error) {
	tsKinds := []string{"**"}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_MulMul{Node: m.Node}, nil
}
func (m *mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar) Add() (*Unnamed_Add, error) {
	tsKinds := []string{"+"}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_Add{Node: m.Node}, nil
}
func (m *mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar) Sub() (*Unnamed_Sub, error) {
	tsKinds := []string{"-"}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_Sub{Node: m.Node}, nil
}
func (m *mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar) Div() (*Unnamed_Div, error) {
	tsKinds := []string{"/"}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_Div{Node: m.Node}, nil
}
func (m *mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar) DivDiv() (*Unnamed_DivDiv, error) {
	tsKinds := []string{"//"}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_DivDiv{Node: m.Node}, nil
}
func (m *mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar) LtLt() (*Unnamed_LtLt, error) {
	tsKinds := []string{"<<"}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_LtLt{Node: m.Node}, nil
}
func (m *mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar) GtGt() (*Unnamed_GtGt, error) {
	tsKinds := []string{">>"}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_GtGt{Node: m.Node}, nil
}
func (m *mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar) At() (*Unnamed_At, error) {
	tsKinds := []string{"@"}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_At{Node: m.Node}, nil
}
func (m *mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar) BitXor() (*Unnamed_BitXor, error) {
	tsKinds := []string{"^"}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_BitXor{Node: m.Node}, nil
}
func (m *mod_ampersand_mul_mulMul_add_sub_div_divDiv_ltLt_gtGt_at_bitXor_bar) Bar() (*Unnamed_Bar, error) {
	tsKinds := []string{"|"}
	if !slices.Contains(tsKinds, m.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", m.Node.Kind(), tsKinds)
	}
	return &Unnamed_Bar{Node: m.Node}, nil
}

type compoundStatement_simpleStatement struct {
	tree_sitter.Node
}

func (c *compoundStatement_simpleStatement) CompoundStatement() (*CompoundStatement, error) {
	tsKinds := []string{"class_definition", "decorated_definition", "for_statement", "function_definition", "if_statement", "match_statement", "try_statement", "while_statement", "with_statement"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &CompoundStatement{Node: c.Node}, nil
}
func (c *compoundStatement_simpleStatement) SimpleStatement() (*SimpleStatement, error) {
	tsKinds := []string{"assert_statement", "break_statement", "continue_statement", "delete_statement", "exec_statement", "expression_statement", "future_import_statement", "global_statement", "import_from_statement", "import_statement", "nonlocal_statement", "pass_statement", "print_statement", "raise_statement", "return_statement", "type_alias_statement"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &SimpleStatement{Node: c.Node}, nil
}

type and_or struct {
	tree_sitter.Node
}

func (a *and_or) And() (*Unnamed_And, error) {
	tsKinds := []string{"and"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Unnamed_And{Node: a.Node}, nil
}
func (a *and_or) Or() (*Unnamed_Or, error) {
	tsKinds := []string{"or"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Unnamed_Or{Node: a.Node}, nil
}

type argumentList_generatorExpression struct {
	tree_sitter.Node
}

func (a *argumentList_generatorExpression) ArgumentList() (*ArgumentList, error) {
	tsKinds := []string{"argument_list"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &ArgumentList{Node: a.Node}, nil
}
func (a *argumentList_generatorExpression) GeneratorExpression() (*GeneratorExpression, error) {
	tsKinds := []string{"generator_expression"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &GeneratorExpression{Node: a.Node}, nil
}

type asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern struct {
	tree_sitter.Node
}

func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) AsPattern() (*AsPattern, error) {
	tsKinds := []string{"as_pattern"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &AsPattern{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ClassPattern() (*ClassPattern, error) {
	tsKinds := []string{"class_pattern"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &ClassPattern{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ComplexPattern() (*ComplexPattern, error) {
	tsKinds := []string{"complex_pattern"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &ComplexPattern{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ConcatenatedString() (*ConcatenatedString, error) {
	tsKinds := []string{"concatenated_string"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &ConcatenatedString{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) DictPattern() (*DictPattern, error) {
	tsKinds := []string{"dict_pattern"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &DictPattern{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) DottedName() (*DottedName, error) {
	tsKinds := []string{"dotted_name"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &DottedName{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) False() (*False, error) {
	tsKinds := []string{"false"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &False{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) Float() (*Float, error) {
	tsKinds := []string{"float"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Float{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) Integer() (*Integer, error) {
	tsKinds := []string{"integer"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Integer{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) KeywordPattern() (*KeywordPattern, error) {
	tsKinds := []string{"keyword_pattern"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &KeywordPattern{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ListPattern() (*ListPattern, error) {
	tsKinds := []string{"list_pattern"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &ListPattern{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) None() (*None, error) {
	tsKinds := []string{"none"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &None{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) SplatPattern() (*SplatPattern, error) {
	tsKinds := []string{"splat_pattern"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &SplatPattern{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) String() (*String, error) {
	tsKinds := []string{"string"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &String{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) True() (*True, error) {
	tsKinds := []string{"true"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &True{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) TuplePattern() (*TuplePattern, error) {
	tsKinds := []string{"tuple_pattern"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &TuplePattern{Node: a.Node}, nil
}
func (a *asPattern_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_keywordPattern_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) UnionPattern() (*UnionPattern, error) {
	tsKinds := []string{"union_pattern"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &UnionPattern{Node: a.Node}, nil
}

type casePattern_dottedName struct {
	tree_sitter.Node
}

func (c *casePattern_dottedName) CasePattern() (*CasePattern, error) {
	tsKinds := []string{"case_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &CasePattern{Node: c.Node}, nil
}
func (c *casePattern_dottedName) DottedName() (*DottedName, error) {
	tsKinds := []string{"dotted_name"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &DottedName{Node: c.Node}, nil
}

type notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn struct {
	tree_sitter.Node
}

func (n *notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn) NotEq() (*Unnamed_NotEq, error) {
	tsKinds := []string{"!="}
	if !slices.Contains(tsKinds, n.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", n.Node.Kind(), tsKinds)
	}
	return &Unnamed_NotEq{Node: n.Node}, nil
}
func (n *notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn) Lt() (*Unnamed_Lt, error) {
	tsKinds := []string{"<"}
	if !slices.Contains(tsKinds, n.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", n.Node.Kind(), tsKinds)
	}
	return &Unnamed_Lt{Node: n.Node}, nil
}
func (n *notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn) LtEq() (*Unnamed_LtEq, error) {
	tsKinds := []string{"<="}
	if !slices.Contains(tsKinds, n.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", n.Node.Kind(), tsKinds)
	}
	return &Unnamed_LtEq{Node: n.Node}, nil
}
func (n *notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn) LtGt() (*Unnamed_LtGt, error) {
	tsKinds := []string{"<>"}
	if !slices.Contains(tsKinds, n.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", n.Node.Kind(), tsKinds)
	}
	return &Unnamed_LtGt{Node: n.Node}, nil
}
func (n *notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn) EqEq() (*Unnamed_EqEq, error) {
	tsKinds := []string{"=="}
	if !slices.Contains(tsKinds, n.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", n.Node.Kind(), tsKinds)
	}
	return &Unnamed_EqEq{Node: n.Node}, nil
}
func (n *notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn) Gt() (*Unnamed_Gt, error) {
	tsKinds := []string{">"}
	if !slices.Contains(tsKinds, n.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", n.Node.Kind(), tsKinds)
	}
	return &Unnamed_Gt{Node: n.Node}, nil
}
func (n *notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn) GtEq() (*Unnamed_GtEq, error) {
	tsKinds := []string{">="}
	if !slices.Contains(tsKinds, n.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", n.Node.Kind(), tsKinds)
	}
	return &Unnamed_GtEq{Node: n.Node}, nil
}
func (n *notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn) In() (*Unnamed_In, error) {
	tsKinds := []string{"in"}
	if !slices.Contains(tsKinds, n.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", n.Node.Kind(), tsKinds)
	}
	return &Unnamed_In{Node: n.Node}, nil
}
func (n *notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn) Is() (*Unnamed_Is, error) {
	tsKinds := []string{"is"}
	if !slices.Contains(tsKinds, n.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", n.Node.Kind(), tsKinds)
	}
	return &Unnamed_Is{Node: n.Node}, nil
}
func (n *notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn) IsSpaceNot() (*Unnamed_IsSpaceNot, error) {
	tsKinds := []string{"is not"}
	if !slices.Contains(tsKinds, n.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", n.Node.Kind(), tsKinds)
	}
	return &Unnamed_IsSpaceNot{Node: n.Node}, nil
}
func (n *notEq_lt_ltEq_ltGt_eqEq_gt_gtEq_in_is_isSpaceNot_notSpaceIn) NotSpaceIn() (*Unnamed_NotSpaceIn, error) {
	tsKinds := []string{"not in"}
	if !slices.Contains(tsKinds, n.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", n.Node.Kind(), tsKinds)
	}
	return &Unnamed_NotSpaceIn{Node: n.Node}, nil
}

type float_integer struct {
	tree_sitter.Node
}

func (f *float_integer) Float() (*Float, error) {
	tsKinds := []string{"float"}
	if !slices.Contains(tsKinds, f.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", f.Node.Kind(), tsKinds)
	}
	return &Float{Node: f.Node}, nil
}
func (f *float_integer) Integer() (*Integer, error) {
	tsKinds := []string{"integer"}
	if !slices.Contains(tsKinds, f.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", f.Node.Kind(), tsKinds)
	}
	return &Integer{Node: f.Node}, nil
}

type classDefinition_functionDefinition struct {
	tree_sitter.Node
}

func (c *classDefinition_functionDefinition) ClassDefinition() (*ClassDefinition, error) {
	tsKinds := []string{"class_definition"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &ClassDefinition{Node: c.Node}, nil
}
func (c *classDefinition_functionDefinition) FunctionDefinition() (*FunctionDefinition, error) {
	tsKinds := []string{"function_definition"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &FunctionDefinition{Node: c.Node}, nil
}

type identifier_tuplePattern struct {
	tree_sitter.Node
}

func (i *identifier_tuplePattern) Identifier() (*Identifier, error) {
	tsKinds := []string{"identifier"}
	if !slices.Contains(tsKinds, i.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", i.Node.Kind(), tsKinds)
	}
	return &Identifier{Node: i.Node}, nil
}
func (i *identifier_tuplePattern) TuplePattern() (*TuplePattern, error) {
	tsKinds := []string{"tuple_pattern"}
	if !slices.Contains(tsKinds, i.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", i.Node.Kind(), tsKinds)
	}
	return &TuplePattern{Node: i.Node}, nil
}

type expression_expressionList struct {
	tree_sitter.Node
}

func (e *expression_expressionList) Expression() (*Expression, error) {
	tsKinds := []string{"as_pattern", "boolean_operator", "comparison_operator", "conditional_expression", "lambda", "named_expression", "not_operator", "attribute", "await", "binary_operator", "call", "concatenated_string", "dictionary", "dictionary_comprehension", "ellipsis", "false", "float", "generator_expression", "identifier", "integer", "list", "list_comprehension", "list_splat", "none", "parenthesized_expression", "set", "set_comprehension", "string", "subscript", "true", "tuple", "unary_operator"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &Expression{Node: e.Node}, nil
}
func (e *expression_expressionList) ExpressionList() (*ExpressionList, error) {
	tsKinds := []string{"expression_list"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &ExpressionList{Node: e.Node}, nil
}

type sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern struct {
	tree_sitter.Node
}

func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) Sub() (*Unnamed_Sub, error) {
	tsKinds := []string{"-"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &Unnamed_Sub{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) Underscore() (*Unnamed_Underscore, error) {
	tsKinds := []string{"_"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &Unnamed_Underscore{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ClassPattern() (*ClassPattern, error) {
	tsKinds := []string{"class_pattern"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &ClassPattern{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ComplexPattern() (*ComplexPattern, error) {
	tsKinds := []string{"complex_pattern"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &ComplexPattern{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ConcatenatedString() (*ConcatenatedString, error) {
	tsKinds := []string{"concatenated_string"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &ConcatenatedString{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) DictPattern() (*DictPattern, error) {
	tsKinds := []string{"dict_pattern"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &DictPattern{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) DottedName() (*DottedName, error) {
	tsKinds := []string{"dotted_name"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &DottedName{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) False() (*False, error) {
	tsKinds := []string{"false"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &False{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) Float() (*Float, error) {
	tsKinds := []string{"float"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &Float{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) Integer() (*Integer, error) {
	tsKinds := []string{"integer"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &Integer{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ListPattern() (*ListPattern, error) {
	tsKinds := []string{"list_pattern"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &ListPattern{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) None() (*None, error) {
	tsKinds := []string{"none"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &None{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) SplatPattern() (*SplatPattern, error) {
	tsKinds := []string{"splat_pattern"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &SplatPattern{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) String() (*String, error) {
	tsKinds := []string{"string"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &String{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) True() (*True, error) {
	tsKinds := []string{"true"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &True{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) TuplePattern() (*TuplePattern, error) {
	tsKinds := []string{"tuple_pattern"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &TuplePattern{Node: s.Node}, nil
}
func (s *sub_underscore_classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) UnionPattern() (*UnionPattern, error) {
	tsKinds := []string{"union_pattern"}
	if !slices.Contains(tsKinds, s.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", s.Node.Kind(), tsKinds)
	}
	return &UnionPattern{Node: s.Node}, nil
}

type dictionarySplat_pair struct {
	tree_sitter.Node
}

func (d *dictionarySplat_pair) DictionarySplat() (*DictionarySplat, error) {
	tsKinds := []string{"dictionary_splat"}
	if !slices.Contains(tsKinds, d.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", d.Node.Kind(), tsKinds)
	}
	return &DictionarySplat{Node: d.Node}, nil
}
func (d *dictionarySplat_pair) Pair() (*Pair, error) {
	tsKinds := []string{"pair"}
	if !slices.Contains(tsKinds, d.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", d.Node.Kind(), tsKinds)
	}
	return &Pair{Node: d.Node}, nil
}

type forInClause_ifClause struct {
	tree_sitter.Node
}

func (f *forInClause_ifClause) ForInClause() (*ForInClause, error) {
	tsKinds := []string{"for_in_clause"}
	if !slices.Contains(tsKinds, f.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", f.Node.Kind(), tsKinds)
	}
	return &ForInClause{Node: f.Node}, nil
}
func (f *forInClause_ifClause) IfClause() (*IfClause, error) {
	tsKinds := []string{"if_clause"}
	if !slices.Contains(tsKinds, f.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", f.Node.Kind(), tsKinds)
	}
	return &IfClause{Node: f.Node}, nil
}

type attribute_identifier_subscript struct {
	tree_sitter.Node
}

func (a *attribute_identifier_subscript) Attribute() (*Attribute, error) {
	tsKinds := []string{"attribute"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Attribute{Node: a.Node}, nil
}
func (a *attribute_identifier_subscript) Identifier() (*Identifier, error) {
	tsKinds := []string{"identifier"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Identifier{Node: a.Node}, nil
}
func (a *attribute_identifier_subscript) Subscript() (*Subscript, error) {
	tsKinds := []string{"subscript"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Subscript{Node: a.Node}, nil
}

type block_expression struct {
	tree_sitter.Node
}

func (b *block_expression) Block() (*Block, error) {
	tsKinds := []string{"block"}
	if !slices.Contains(tsKinds, b.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", b.Node.Kind(), tsKinds)
	}
	return &Block{Node: b.Node}, nil
}
func (b *block_expression) Expression() (*Expression, error) {
	tsKinds := []string{"as_pattern", "boolean_operator", "comparison_operator", "conditional_expression", "lambda", "named_expression", "not_operator", "attribute", "await", "binary_operator", "call", "concatenated_string", "dictionary", "dictionary_comprehension", "ellipsis", "false", "float", "generator_expression", "identifier", "integer", "list", "list_comprehension", "list_splat", "none", "parenthesized_expression", "set", "set_comprehension", "string", "subscript", "true", "tuple", "unary_operator"}
	if !slices.Contains(tsKinds, b.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", b.Node.Kind(), tsKinds)
	}
	return &Expression{Node: b.Node}, nil
}

type identifier_string struct {
	tree_sitter.Node
}

func (i *identifier_string) Identifier() (*Identifier, error) {
	tsKinds := []string{"identifier"}
	if !slices.Contains(tsKinds, i.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", i.Node.Kind(), tsKinds)
	}
	return &Identifier{Node: i.Node}, nil
}
func (i *identifier_string) String() (*String, error) {
	tsKinds := []string{"string"}
	if !slices.Contains(tsKinds, i.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", i.Node.Kind(), tsKinds)
	}
	return &String{Node: i.Node}, nil
}

type assignment_augmentedAssignment_expression_yield struct {
	tree_sitter.Node
}

func (a *assignment_augmentedAssignment_expression_yield) Assignment() (*Assignment, error) {
	tsKinds := []string{"assignment"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Assignment{Node: a.Node}, nil
}
func (a *assignment_augmentedAssignment_expression_yield) AugmentedAssignment() (*AugmentedAssignment, error) {
	tsKinds := []string{"augmented_assignment"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &AugmentedAssignment{Node: a.Node}, nil
}
func (a *assignment_augmentedAssignment_expression_yield) Expression() (*Expression, error) {
	tsKinds := []string{"as_pattern", "boolean_operator", "comparison_operator", "conditional_expression", "lambda", "named_expression", "not_operator", "attribute", "await", "binary_operator", "call", "concatenated_string", "dictionary", "dictionary_comprehension", "ellipsis", "false", "float", "generator_expression", "identifier", "integer", "list", "list_comprehension", "list_splat", "none", "parenthesized_expression", "set", "set_comprehension", "string", "subscript", "true", "tuple", "unary_operator"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Expression{Node: a.Node}, nil
}
func (a *assignment_augmentedAssignment_expression_yield) Yield() (*Yield, error) {
	tsKinds := []string{"yield"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Yield{Node: a.Node}, nil
}

type comma_expression struct {
	tree_sitter.Node
}

func (c *comma_expression) Comma() (*Unnamed_Comma, error) {
	tsKinds := []string{","}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &Unnamed_Comma{Node: c.Node}, nil
}
func (c *comma_expression) Expression() (*Expression, error) {
	tsKinds := []string{"as_pattern", "boolean_operator", "comparison_operator", "conditional_expression", "lambda", "named_expression", "not_operator", "attribute", "await", "binary_operator", "call", "concatenated_string", "dictionary", "dictionary_comprehension", "ellipsis", "false", "float", "generator_expression", "identifier", "integer", "list", "list_comprehension", "list_splat", "none", "parenthesized_expression", "set", "set_comprehension", "string", "subscript", "true", "tuple", "unary_operator"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &Expression{Node: c.Node}, nil
}

type expression_expressionList_patternList_yield struct {
	tree_sitter.Node
}

func (e *expression_expressionList_patternList_yield) Expression() (*Expression, error) {
	tsKinds := []string{"as_pattern", "boolean_operator", "comparison_operator", "conditional_expression", "lambda", "named_expression", "not_operator", "attribute", "await", "binary_operator", "call", "concatenated_string", "dictionary", "dictionary_comprehension", "ellipsis", "false", "float", "generator_expression", "identifier", "integer", "list", "list_comprehension", "list_splat", "none", "parenthesized_expression", "set", "set_comprehension", "string", "subscript", "true", "tuple", "unary_operator"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &Expression{Node: e.Node}, nil
}
func (e *expression_expressionList_patternList_yield) ExpressionList() (*ExpressionList, error) {
	tsKinds := []string{"expression_list"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &ExpressionList{Node: e.Node}, nil
}
func (e *expression_expressionList_patternList_yield) PatternList() (*PatternList, error) {
	tsKinds := []string{"pattern_list"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &PatternList{Node: e.Node}, nil
}
func (e *expression_expressionList_patternList_yield) Yield() (*Yield, error) {
	tsKinds := []string{"yield"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &Yield{Node: e.Node}, nil
}

type aliasedImport_dottedName struct {
	tree_sitter.Node
}

func (a *aliasedImport_dottedName) AliasedImport() (*AliasedImport, error) {
	tsKinds := []string{"aliased_import"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &AliasedImport{Node: a.Node}, nil
}
func (a *aliasedImport_dottedName) DottedName() (*DottedName, error) {
	tsKinds := []string{"dotted_name"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &DottedName{Node: a.Node}, nil
}

type identifier_typeParameter struct {
	tree_sitter.Node
}

func (i *identifier_typeParameter) Identifier() (*Identifier, error) {
	tsKinds := []string{"identifier"}
	if !slices.Contains(tsKinds, i.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", i.Node.Kind(), tsKinds)
	}
	return &Identifier{Node: i.Node}, nil
}
func (i *identifier_typeParameter) TypeParameter() (*TypeParameter, error) {
	tsKinds := []string{"type_parameter"}
	if !slices.Contains(tsKinds, i.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", i.Node.Kind(), tsKinds)
	}
	return &TypeParameter{Node: i.Node}, nil
}

type elifClause_elseClause struct {
	tree_sitter.Node
}

func (e *elifClause_elseClause) ElifClause() (*ElifClause, error) {
	tsKinds := []string{"elif_clause"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &ElifClause{Node: e.Node}, nil
}
func (e *elifClause_elseClause) ElseClause() (*ElseClause, error) {
	tsKinds := []string{"else_clause"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &ElseClause{Node: e.Node}, nil
}

type dottedName_relativeImport struct {
	tree_sitter.Node
}

func (d *dottedName_relativeImport) DottedName() (*DottedName, error) {
	tsKinds := []string{"dotted_name"}
	if !slices.Contains(tsKinds, d.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", d.Node.Kind(), tsKinds)
	}
	return &DottedName{Node: d.Node}, nil
}
func (d *dottedName_relativeImport) RelativeImport() (*RelativeImport, error) {
	tsKinds := []string{"relative_import"}
	if !slices.Contains(tsKinds, d.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", d.Node.Kind(), tsKinds)
	}
	return &RelativeImport{Node: d.Node}, nil
}

type classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern struct {
	tree_sitter.Node
}

func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ClassPattern() (*ClassPattern, error) {
	tsKinds := []string{"class_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &ClassPattern{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ComplexPattern() (*ComplexPattern, error) {
	tsKinds := []string{"complex_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &ComplexPattern{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ConcatenatedString() (*ConcatenatedString, error) {
	tsKinds := []string{"concatenated_string"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &ConcatenatedString{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) DictPattern() (*DictPattern, error) {
	tsKinds := []string{"dict_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &DictPattern{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) DottedName() (*DottedName, error) {
	tsKinds := []string{"dotted_name"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &DottedName{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) False() (*False, error) {
	tsKinds := []string{"false"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &False{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) Float() (*Float, error) {
	tsKinds := []string{"float"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &Float{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) Identifier() (*Identifier, error) {
	tsKinds := []string{"identifier"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &Identifier{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) Integer() (*Integer, error) {
	tsKinds := []string{"integer"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &Integer{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ListPattern() (*ListPattern, error) {
	tsKinds := []string{"list_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &ListPattern{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) None() (*None, error) {
	tsKinds := []string{"none"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &None{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) SplatPattern() (*SplatPattern, error) {
	tsKinds := []string{"splat_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &SplatPattern{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) String() (*String, error) {
	tsKinds := []string{"string"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &String{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) True() (*True, error) {
	tsKinds := []string{"true"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &True{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) TuplePattern() (*TuplePattern, error) {
	tsKinds := []string{"tuple_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &TuplePattern{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_identifier_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) UnionPattern() (*UnionPattern, error) {
	tsKinds := []string{"union_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &UnionPattern{Node: c.Node}, nil
}

type expression_listSplat_parenthesizedListSplat_yield struct {
	tree_sitter.Node
}

func (e *expression_listSplat_parenthesizedListSplat_yield) Expression() (*Expression, error) {
	tsKinds := []string{"as_pattern", "boolean_operator", "comparison_operator", "conditional_expression", "lambda", "named_expression", "not_operator", "attribute", "await", "binary_operator", "call", "concatenated_string", "dictionary", "dictionary_comprehension", "ellipsis", "false", "float", "generator_expression", "identifier", "integer", "list", "list_comprehension", "list_splat", "none", "parenthesized_expression", "set", "set_comprehension", "string", "subscript", "true", "tuple", "unary_operator"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &Expression{Node: e.Node}, nil
}
func (e *expression_listSplat_parenthesizedListSplat_yield) ListSplat() (*ListSplat, error) {
	tsKinds := []string{"list_splat"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &ListSplat{Node: e.Node}, nil
}
func (e *expression_listSplat_parenthesizedListSplat_yield) ParenthesizedListSplat() (*ParenthesizedListSplat, error) {
	tsKinds := []string{"parenthesized_list_splat"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &ParenthesizedListSplat{Node: e.Node}, nil
}
func (e *expression_listSplat_parenthesizedListSplat_yield) Yield() (*Yield, error) {
	tsKinds := []string{"yield"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &Yield{Node: e.Node}, nil
}

type casePattern_pattern struct {
	tree_sitter.Node
}

func (c *casePattern_pattern) CasePattern() (*CasePattern, error) {
	tsKinds := []string{"case_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &CasePattern{Node: c.Node}, nil
}
func (c *casePattern_pattern) Pattern() (*Pattern, error) {
	tsKinds := []string{"attribute", "identifier", "list_pattern", "list_splat_pattern", "subscript", "tuple_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &Pattern{Node: c.Node}, nil
}

type attribute_expression_identifier_subscript struct {
	tree_sitter.Node
}

func (a *attribute_expression_identifier_subscript) Attribute() (*Attribute, error) {
	tsKinds := []string{"attribute"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Attribute{Node: a.Node}, nil
}
func (a *attribute_expression_identifier_subscript) Expression() (*Expression, error) {
	tsKinds := []string{"as_pattern", "boolean_operator", "comparison_operator", "conditional_expression", "lambda", "named_expression", "not_operator", "attribute", "await", "binary_operator", "call", "concatenated_string", "dictionary", "dictionary_comprehension", "ellipsis", "false", "float", "generator_expression", "identifier", "integer", "list", "list_comprehension", "list_splat", "none", "parenthesized_expression", "set", "set_comprehension", "string", "subscript", "true", "tuple", "unary_operator"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Expression{Node: a.Node}, nil
}
func (a *attribute_expression_identifier_subscript) Identifier() (*Identifier, error) {
	tsKinds := []string{"identifier"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Identifier{Node: a.Node}, nil
}
func (a *attribute_expression_identifier_subscript) Subscript() (*Subscript, error) {
	tsKinds := []string{"subscript"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Subscript{Node: a.Node}, nil
}

type identifier_type_ struct {
	tree_sitter.Node
}

func (i *identifier_type_) Identifier() (*Identifier, error) {
	tsKinds := []string{"identifier"}
	if !slices.Contains(tsKinds, i.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", i.Node.Kind(), tsKinds)
	}
	return &Identifier{Node: i.Node}, nil
}
func (i *identifier_type_) Type_() (*Type, error) {
	tsKinds := []string{"type"}
	if !slices.Contains(tsKinds, i.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", i.Node.Kind(), tsKinds)
	}
	return &Type{Node: i.Node}, nil
}

type expression_listSplat_parenthesizedExpression_yield struct {
	tree_sitter.Node
}

func (e *expression_listSplat_parenthesizedExpression_yield) Expression() (*Expression, error) {
	tsKinds := []string{"as_pattern", "boolean_operator", "comparison_operator", "conditional_expression", "lambda", "named_expression", "not_operator", "attribute", "await", "binary_operator", "call", "concatenated_string", "dictionary", "dictionary_comprehension", "ellipsis", "false", "float", "generator_expression", "identifier", "integer", "list", "list_comprehension", "list_splat", "none", "parenthesized_expression", "set", "set_comprehension", "string", "subscript", "true", "tuple", "unary_operator"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &Expression{Node: e.Node}, nil
}
func (e *expression_listSplat_parenthesizedExpression_yield) ListSplat() (*ListSplat, error) {
	tsKinds := []string{"list_splat"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &ListSplat{Node: e.Node}, nil
}
func (e *expression_listSplat_parenthesizedExpression_yield) ParenthesizedExpression() (*ParenthesizedExpression, error) {
	tsKinds := []string{"parenthesized_expression"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &ParenthesizedExpression{Node: e.Node}, nil
}
func (e *expression_listSplat_parenthesizedExpression_yield) Yield() (*Yield, error) {
	tsKinds := []string{"yield"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &Yield{Node: e.Node}, nil
}

type listSplat_parenthesizedExpression struct {
	tree_sitter.Node
}

func (l *listSplat_parenthesizedExpression) ListSplat() (*ListSplat, error) {
	tsKinds := []string{"list_splat"}
	if !slices.Contains(tsKinds, l.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", l.Node.Kind(), tsKinds)
	}
	return &ListSplat{Node: l.Node}, nil
}
func (l *listSplat_parenthesizedExpression) ParenthesizedExpression() (*ParenthesizedExpression, error) {
	tsKinds := []string{"parenthesized_expression"}
	if !slices.Contains(tsKinds, l.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", l.Node.Kind(), tsKinds)
	}
	return &ParenthesizedExpression{Node: l.Node}, nil
}

type dottedName_importPrefix struct {
	tree_sitter.Node
}

func (d *dottedName_importPrefix) DottedName() (*DottedName, error) {
	tsKinds := []string{"dotted_name"}
	if !slices.Contains(tsKinds, d.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", d.Node.Kind(), tsKinds)
	}
	return &DottedName{Node: d.Node}, nil
}
func (d *dottedName_importPrefix) ImportPrefix() (*ImportPrefix, error) {
	tsKinds := []string{"import_prefix"}
	if !slices.Contains(tsKinds, d.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", d.Node.Kind(), tsKinds)
	}
	return &ImportPrefix{Node: d.Node}, nil
}

type interpolation_stringContent_stringEnd_stringStart struct {
	tree_sitter.Node
}

func (i *interpolation_stringContent_stringEnd_stringStart) Interpolation() (*Interpolation, error) {
	tsKinds := []string{"interpolation"}
	if !slices.Contains(tsKinds, i.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", i.Node.Kind(), tsKinds)
	}
	return &Interpolation{Node: i.Node}, nil
}
func (i *interpolation_stringContent_stringEnd_stringStart) StringContent() (*StringContent, error) {
	tsKinds := []string{"string_content"}
	if !slices.Contains(tsKinds, i.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", i.Node.Kind(), tsKinds)
	}
	return &StringContent{Node: i.Node}, nil
}
func (i *interpolation_stringContent_stringEnd_stringStart) StringEnd() (*StringEnd, error) {
	tsKinds := []string{"string_end"}
	if !slices.Contains(tsKinds, i.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", i.Node.Kind(), tsKinds)
	}
	return &StringEnd{Node: i.Node}, nil
}
func (i *interpolation_stringContent_stringEnd_stringStart) StringStart() (*StringStart, error) {
	tsKinds := []string{"string_start"}
	if !slices.Contains(tsKinds, i.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", i.Node.Kind(), tsKinds)
	}
	return &StringStart{Node: i.Node}, nil
}

type escapeInterpolation_escapeSequence struct {
	tree_sitter.Node
}

func (e *escapeInterpolation_escapeSequence) EscapeInterpolation() (*EscapeInterpolation, error) {
	tsKinds := []string{"escape_interpolation"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &EscapeInterpolation{Node: e.Node}, nil
}
func (e *escapeInterpolation_escapeSequence) EscapeSequence() (*EscapeSequence, error) {
	tsKinds := []string{"escape_sequence"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &EscapeSequence{Node: e.Node}, nil
}

type expression_slice struct {
	tree_sitter.Node
}

func (e *expression_slice) Expression() (*Expression, error) {
	tsKinds := []string{"as_pattern", "boolean_operator", "comparison_operator", "conditional_expression", "lambda", "named_expression", "not_operator", "attribute", "await", "binary_operator", "call", "concatenated_string", "dictionary", "dictionary_comprehension", "ellipsis", "false", "float", "generator_expression", "identifier", "integer", "list", "list_comprehension", "list_splat", "none", "parenthesized_expression", "set", "set_comprehension", "string", "subscript", "true", "tuple", "unary_operator"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &Expression{Node: e.Node}, nil
}
func (e *expression_slice) Slice() (*Slice, error) {
	tsKinds := []string{"slice"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &Slice{Node: e.Node}, nil
}

type elseClause_exceptClause_exceptGroupClause_finallyClause struct {
	tree_sitter.Node
}

func (e *elseClause_exceptClause_exceptGroupClause_finallyClause) ElseClause() (*ElseClause, error) {
	tsKinds := []string{"else_clause"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &ElseClause{Node: e.Node}, nil
}
func (e *elseClause_exceptClause_exceptGroupClause_finallyClause) ExceptClause() (*ExceptClause, error) {
	tsKinds := []string{"except_clause"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &ExceptClause{Node: e.Node}, nil
}
func (e *elseClause_exceptClause_exceptGroupClause_finallyClause) ExceptGroupClause() (*ExceptGroupClause, error) {
	tsKinds := []string{"except_group_clause"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &ExceptGroupClause{Node: e.Node}, nil
}
func (e *elseClause_exceptClause_exceptGroupClause_finallyClause) FinallyClause() (*FinallyClause, error) {
	tsKinds := []string{"finally_clause"}
	if !slices.Contains(tsKinds, e.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", e.Node.Kind(), tsKinds)
	}
	return &FinallyClause{Node: e.Node}, nil
}

type constrainedType_expression_genericType_memberType_splatType_unionType struct {
	tree_sitter.Node
}

func (c *constrainedType_expression_genericType_memberType_splatType_unionType) ConstrainedType() (*ConstrainedType, error) {
	tsKinds := []string{"constrained_type"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &ConstrainedType{Node: c.Node}, nil
}
func (c *constrainedType_expression_genericType_memberType_splatType_unionType) Expression() (*Expression, error) {
	tsKinds := []string{"as_pattern", "boolean_operator", "comparison_operator", "conditional_expression", "lambda", "named_expression", "not_operator", "attribute", "await", "binary_operator", "call", "concatenated_string", "dictionary", "dictionary_comprehension", "ellipsis", "false", "float", "generator_expression", "identifier", "integer", "list", "list_comprehension", "list_splat", "none", "parenthesized_expression", "set", "set_comprehension", "string", "subscript", "true", "tuple", "unary_operator"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &Expression{Node: c.Node}, nil
}
func (c *constrainedType_expression_genericType_memberType_splatType_unionType) GenericType() (*GenericType, error) {
	tsKinds := []string{"generic_type"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &GenericType{Node: c.Node}, nil
}
func (c *constrainedType_expression_genericType_memberType_splatType_unionType) MemberType() (*MemberType, error) {
	tsKinds := []string{"member_type"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &MemberType{Node: c.Node}, nil
}
func (c *constrainedType_expression_genericType_memberType_splatType_unionType) SplatType() (*SplatType, error) {
	tsKinds := []string{"splat_type"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &SplatType{Node: c.Node}, nil
}
func (c *constrainedType_expression_genericType_memberType_splatType_unionType) UnionType() (*UnionType, error) {
	tsKinds := []string{"union_type"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &UnionType{Node: c.Node}, nil
}

type dictionarySplatPattern_identifier_listSplatPattern struct {
	tree_sitter.Node
}

func (d *dictionarySplatPattern_identifier_listSplatPattern) DictionarySplatPattern() (*DictionarySplatPattern, error) {
	tsKinds := []string{"dictionary_splat_pattern"}
	if !slices.Contains(tsKinds, d.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", d.Node.Kind(), tsKinds)
	}
	return &DictionarySplatPattern{Node: d.Node}, nil
}
func (d *dictionarySplatPattern_identifier_listSplatPattern) Identifier() (*Identifier, error) {
	tsKinds := []string{"identifier"}
	if !slices.Contains(tsKinds, d.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", d.Node.Kind(), tsKinds)
	}
	return &Identifier{Node: d.Node}, nil
}
func (d *dictionarySplatPattern_identifier_listSplatPattern) ListSplatPattern() (*ListSplatPattern, error) {
	tsKinds := []string{"list_splat_pattern"}
	if !slices.Contains(tsKinds, d.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", d.Node.Kind(), tsKinds)
	}
	return &ListSplatPattern{Node: d.Node}, nil
}

type add_sub_bitNot struct {
	tree_sitter.Node
}

func (a *add_sub_bitNot) Add() (*Unnamed_Add, error) {
	tsKinds := []string{"+"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Unnamed_Add{Node: a.Node}, nil
}
func (a *add_sub_bitNot) Sub() (*Unnamed_Sub, error) {
	tsKinds := []string{"-"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Unnamed_Sub{Node: a.Node}, nil
}
func (a *add_sub_bitNot) BitNot() (*Unnamed_BitNot, error) {
	tsKinds := []string{"~"}
	if !slices.Contains(tsKinds, a.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", a.Node.Kind(), tsKinds)
	}
	return &Unnamed_BitNot{Node: a.Node}, nil
}

type classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern struct {
	tree_sitter.Node
}

func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ClassPattern() (*ClassPattern, error) {
	tsKinds := []string{"class_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &ClassPattern{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ComplexPattern() (*ComplexPattern, error) {
	tsKinds := []string{"complex_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &ComplexPattern{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ConcatenatedString() (*ConcatenatedString, error) {
	tsKinds := []string{"concatenated_string"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &ConcatenatedString{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) DictPattern() (*DictPattern, error) {
	tsKinds := []string{"dict_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &DictPattern{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) DottedName() (*DottedName, error) {
	tsKinds := []string{"dotted_name"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &DottedName{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) False() (*False, error) {
	tsKinds := []string{"false"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &False{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) Float() (*Float, error) {
	tsKinds := []string{"float"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &Float{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) Integer() (*Integer, error) {
	tsKinds := []string{"integer"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &Integer{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) ListPattern() (*ListPattern, error) {
	tsKinds := []string{"list_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &ListPattern{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) None() (*None, error) {
	tsKinds := []string{"none"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &None{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) SplatPattern() (*SplatPattern, error) {
	tsKinds := []string{"splat_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &SplatPattern{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) String() (*String, error) {
	tsKinds := []string{"string"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &String{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) True() (*True, error) {
	tsKinds := []string{"true"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &True{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) TuplePattern() (*TuplePattern, error) {
	tsKinds := []string{"tuple_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &TuplePattern{Node: c.Node}, nil
}
func (c *classPattern_complexPattern_concatenatedString_dictPattern_dottedName_false_float_integer_listPattern_none_splatPattern_string_true_tuplePattern_unionPattern) UnionPattern() (*UnionPattern, error) {
	tsKinds := []string{"union_pattern"}
	if !slices.Contains(tsKinds, c.Node.Kind()) {
		return nil, fmt.Errorf("Node is a %s, not in %v", c.Node.Kind(), tsKinds)
	}
	return &UnionPattern{Node: c.Node}, nil
}

type Unknown__asPatternTarget struct {
	tree_sitter.Node
}
