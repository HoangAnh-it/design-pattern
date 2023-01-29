package prototype

type INode interface {
	Clone() INode
	PrintTree(space string)
}
