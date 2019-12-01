package ast

import (
	"seqa/token"
	"testing"
)

func TestString (t *testing.T) {
	context := &Context{
		Header: CommentStatement{
			Token: token.Token{Type: token.STRING, Literal: "#"},
			Description: "this is header description",
		},
		Footer: CommentStatement{
			Token: token.Token{Type: token.STRING, Literal: "#"},
			Description: "this is footer description",
		},
		
		Statements : []Statement{
			&ArrowStatement{
				Token: token.Token{Type: token.BULLET, Literal: "-"},
				LeftNode: ArrowNode{
					Token: token.Token{Type: token.LT, Literal: "<"},
				},
				Line: LineNode{
					Token: token.Token{Type: token.LINE, Literal: "-"},
					length: 10,
				},
				RightNode: ArrowNode{
					Token: token.Token{Type: token.GT},
				},
				Comment: CommentStatement{
					Token: token.Token{Type: token.STRING, Literal: ":"},
					Description: "comment arrow",
				},
			},
		},
	}

	ns := NewNodes()
	context.Nodes = ns
	context.Nodes.Set("ABCDE", "this is aaa")
	context.Nodes.Set("VW", "this is bbb")	

	t.Logf("context test ....\n%s", context.String())
	
}
		
			
		
