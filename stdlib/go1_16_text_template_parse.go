// Code generated by 'yaegi extract text/template/parse'. DO NOT EDIT.

// +build go1.16,!go1.17

package stdlib

import (
	"reflect"
	"text/template/parse"
)

func init() {
	Symbols["text/template/parse/parse"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"IsEmptyTree":    reflect.ValueOf(parse.IsEmptyTree),
		"New":            reflect.ValueOf(parse.New),
		"NewIdentifier":  reflect.ValueOf(parse.NewIdentifier),
		"NodeAction":     reflect.ValueOf(parse.NodeAction),
		"NodeBool":       reflect.ValueOf(parse.NodeBool),
		"NodeChain":      reflect.ValueOf(parse.NodeChain),
		"NodeCommand":    reflect.ValueOf(parse.NodeCommand),
		"NodeComment":    reflect.ValueOf(parse.NodeComment),
		"NodeDot":        reflect.ValueOf(parse.NodeDot),
		"NodeField":      reflect.ValueOf(parse.NodeField),
		"NodeIdentifier": reflect.ValueOf(parse.NodeIdentifier),
		"NodeIf":         reflect.ValueOf(parse.NodeIf),
		"NodeList":       reflect.ValueOf(parse.NodeList),
		"NodeNil":        reflect.ValueOf(parse.NodeNil),
		"NodeNumber":     reflect.ValueOf(parse.NodeNumber),
		"NodePipe":       reflect.ValueOf(parse.NodePipe),
		"NodeRange":      reflect.ValueOf(parse.NodeRange),
		"NodeString":     reflect.ValueOf(parse.NodeString),
		"NodeTemplate":   reflect.ValueOf(parse.NodeTemplate),
		"NodeText":       reflect.ValueOf(parse.NodeText),
		"NodeVariable":   reflect.ValueOf(parse.NodeVariable),
		"NodeWith":       reflect.ValueOf(parse.NodeWith),
		"Parse":          reflect.ValueOf(parse.Parse),
		"ParseComments":  reflect.ValueOf(parse.ParseComments),

		// type definitions
		"ActionNode":     reflect.ValueOf((*parse.ActionNode)(nil)),
		"BoolNode":       reflect.ValueOf((*parse.BoolNode)(nil)),
		"BranchNode":     reflect.ValueOf((*parse.BranchNode)(nil)),
		"ChainNode":      reflect.ValueOf((*parse.ChainNode)(nil)),
		"CommandNode":    reflect.ValueOf((*parse.CommandNode)(nil)),
		"CommentNode":    reflect.ValueOf((*parse.CommentNode)(nil)),
		"DotNode":        reflect.ValueOf((*parse.DotNode)(nil)),
		"FieldNode":      reflect.ValueOf((*parse.FieldNode)(nil)),
		"IdentifierNode": reflect.ValueOf((*parse.IdentifierNode)(nil)),
		"IfNode":         reflect.ValueOf((*parse.IfNode)(nil)),
		"ListNode":       reflect.ValueOf((*parse.ListNode)(nil)),
		"Mode":           reflect.ValueOf((*parse.Mode)(nil)),
		"NilNode":        reflect.ValueOf((*parse.NilNode)(nil)),
		"Node":           reflect.ValueOf((*parse.Node)(nil)),
		"NodeType":       reflect.ValueOf((*parse.NodeType)(nil)),
		"NumberNode":     reflect.ValueOf((*parse.NumberNode)(nil)),
		"PipeNode":       reflect.ValueOf((*parse.PipeNode)(nil)),
		"Pos":            reflect.ValueOf((*parse.Pos)(nil)),
		"RangeNode":      reflect.ValueOf((*parse.RangeNode)(nil)),
		"StringNode":     reflect.ValueOf((*parse.StringNode)(nil)),
		"TemplateNode":   reflect.ValueOf((*parse.TemplateNode)(nil)),
		"TextNode":       reflect.ValueOf((*parse.TextNode)(nil)),
		"Tree":           reflect.ValueOf((*parse.Tree)(nil)),
		"VariableNode":   reflect.ValueOf((*parse.VariableNode)(nil)),
		"WithNode":       reflect.ValueOf((*parse.WithNode)(nil)),

		// interface wrapper definitions
		"_Node": reflect.ValueOf((*_text_template_parse_Node)(nil)),
	}
}

// _text_template_parse_Node is an interface wrapper for Node type
type _text_template_parse_Node struct {
	IValue    interface{}
	WCopy     func() parse.Node
	WPosition func() parse.Pos
	WString   func() string
	WType     func() parse.NodeType
}

func (W _text_template_parse_Node) Copy() parse.Node {
	return W.WCopy()
}
func (W _text_template_parse_Node) Position() parse.Pos {
	return W.WPosition()
}
func (W _text_template_parse_Node) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}
func (W _text_template_parse_Node) Type() parse.NodeType {
	return W.WType()
}
