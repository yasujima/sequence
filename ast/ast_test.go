package ast

import (
	"seqa/token"
	"testing"
)

func TestString (t *testing.T) {
	context := &Context{
		Header: &CommentStatement{
			Token: token.Token{Type: token.STRING, Literal: "#"},
			Description: "this is header description",
		},
		// Nodes: []Node{
		// 	&ArrowStatement{
		// 		Token: token.Token{Type: token.BULLET, Literal: "-"},
		// 		LeftNode: &ArrowNode{
		// 			Token: token.Token{Type: token.GT, Literal: ">"},
		// 		},
		// 		Line: &LineNode{
		// 			Token: token.Token{Type: token.LINE, Literal: "-"},
		// 		},
		// 		RightNode: &ArrowNode{
		// 			Token: token.Token{Type: token.GT},
		// 		},
		// 		Comment: &CommentStatement{
		// 			Token: token.Token{Type: token.STRING, Literal: ":"},
		// 			Description: "comment description",
		// 		},
		// 	},
		// },
		// Statements : []Statement{},
	}

	t.Logf("context test .... %q", context.String())
	
}
		
			
		