package ast

import (
	"bytes"
	"seqa/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
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
	LeftNode Node
	Line Node
	RightNode Node
	Comment Statement
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
	Header Statement
	Nodes []Node
	Statements []Statement
	Footer Statement	
}

func (c *Context) TokenLiteral() string {
	return c.Header.TokenLiteral()
}

func (c *Context) String() string {
	var out bytes.Buffer

	//ここでNodeや長さを理解し構築する
	out.WriteString(c.Header.String())
	out.WriteString("\n")
	l := 30
	for _, n := range c.Nodes {
		out.WriteString(n.String())
		spaces := l - len(n.String())
		out.WriteString(strings.Repeat(" ", spaces))
	}
	for _, n := range c.Statements {
		//ここにも人手間
		out.WriteString(n.String())
	}
	out.WriteString("\n")
	//	out.WriteString(c.Footer.String())

	return out.String()
}



	
