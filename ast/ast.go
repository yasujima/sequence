package ast

import (
	"bytes"
	"seqa/token"
	"strings"
)

type Nodes struct {
	store []string
	keys map[string]int16
}

func NewNodes() *Nodes {
	ks := make(map[string]int16)
	return &Nodes{keys: ks}
}

func (ns *Nodes) String(l int) string {
	var out bytes.Buffer
	
	for _, n := range ns.store {
		if l < len(n) {
			out.WriteString(n[:l])
		} else {
			out.WriteString(n)
			out.WriteString(strings.Repeat(" ", l - len(n)))
		}

	}

	return out.String()
}

func (ns *Nodes) Size() int16 {
	return int16(len(ns.store))
}

func (ns *Nodes) Set(name string, comment string) (Node, int16) {
	index, ok := ns.keys[name]
	if ok {
		return &ActorNode{
			Comment: CommentStatement{
				Token: token.Token{Type:token.COLON},
				Description: comment,
			},
		}, index
	} else {
		ns.store = append(ns.store, name)
		ns.keys[name] = int16(len(ns.store))-1
		return &ActorNode{
			Comment: CommentStatement{
				Token: token.Token{Type:token.COLON},
				Description: comment,
			},},
			ns.keys[name]
	}
}

type Node interface {
	TokenLiteral() string
	String() string
}

type ActorNode struct {
	Token token.Token
	Comment CommentStatement
}

func (an *ActorNode) TokenLiteral() string {return an.Token.Literal}
func (an *ActorNode) String() string {
	var out bytes.Buffer

	out.WriteString(an.Token.Literal)
	out.WriteString(" ")
	out.WriteString(an.Comment.String())

	return out.String()
}

type ArrowNode struct {
	Token token.Token
}

func (an *ArrowNode) TokenLiteral() string {return an.Token.Literal}
func (an *ArrowNode) String() string {return an.Token.Literal}

type LineNode struct {
	Token token.Token
	length uint64
}

func (ln *LineNode) TokenLiteral() string {return ln.Token.Literal}
func (ln *LineNode) String() string {
	var out bytes.Buffer
	out.WriteString(strings.Repeat(ln.TokenLiteral(), int(ln.length/2))) //コンテンツ分を差し引くこと

	return out.String()
}

type Statement interface {
	Node
}

type CommentStatement struct {
	Token token.Token
	Description string
}

func (cs *CommentStatement) TokenLiteral() string {return cs.Token.Literal}
func (cs *CommentStatement) String() string {return cs.Description}

type ArrowStatement struct {
	Token token.Token
	LeftNode ArrowNode
	Line LineNode
	RightNode ArrowNode
	Comment CommentStatement
}

func (as *ArrowStatement) TokenLiteral() string {return as.Token.Literal}
func (as *ArrowStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.LeftNode.String())
	out.WriteString(as.Line.String())
	out.WriteString(as.Comment.String())
	out.WriteString(as.Line.String())	
	out.WriteString(as.RightNode.String())

	return out.String()
}

type Context struct {
	Header CommentStatement
	Nodes *Nodes
	Statements []Statement
	Footer CommentStatement	
}

func (c *Context) TokenLiteral() string {
	return c.Header.TokenLiteral()
}

func (c *Context) String() string {
	var out bytes.Buffer

	//ここでNodeや長さを理解し構築する
	out.WriteString(c.Header.String())
	out.WriteString("\n")
	l := 10
	out.WriteString(c.Nodes.String(l))
	for _, n := range c.Statements {
		//ここにも人手間
		out.WriteString(n.String())
	}
	out.WriteString("\n")
	out.WriteString(c.Footer.String())

	return out.String()
}



	
