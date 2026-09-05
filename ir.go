package ir

type Node struct {
	ValueIsSet bool
	Value string
	Deduplicated bool
	Children []*Node 
}

func (N Node) SetValue(v string) error {
	if N.ValueIsSet == false {
		N.Value = v
		N.ValueIsSet = true
		return nil
	}
	return fmt.Errorf("Value has already been set.")
}

func (N Node) UnsetValue() {
	N.ValueIsSet = false
	N.Value = ""
}

func (N Node) AppendChild(np *Node) {
	N.Children = append(N.Children, np)
}

func (N Node) InsertChild(np *Node, position int) {
	if position > len(N.Children) {
		return fmt.Errorf("Nonexisting position")
	}
	var newChildren = make([]*Node, len(N.Children) + 1)
	writen := copy(newChildren, N.Children[:position])
	if writen != position {
		return fmt.Errorf("Have copied: %d elements, expected to copy %d", writen, position)
	}
	newChildren[position] = np
	N.Children = append(newChildren, N.Children[position:])
	return nil
}



func (N Node) Deduplicate() {}

func (N Node) Reduplicate() {}


//type newIr interface {
//	GetStrings() map[string]string
//	GetIr() map[string]newIr
//	GetOrder() []*string
//}
//
//type IrStruct interface {
//	GetStrings() map[string]string
//	GetIrStructs() map[string]Ir
//	GetIrArrs() map[string]IrArr
//	GetOrder() []*string 
//	SetStrings(m map[string]string)
//	SetIrStructs(m map[string]Ir)
//	SetIrArrs(m map[string]IrArr)
//	SetOrder(l []*string)
//}
//
//// only heterogenic arrays available in go, no lists
//type IrArr interface {
//	GetStrings() []string
//	GetIrStructs() []IrStruct
//	GetIrLists() []IrList
//	SetStrings(s []string)
//	SetIrStructs(iS []IrStruct)
//	SetIrLists(iL []IrList)
//}
//
//type Ir struct {
//	strings map[string]string
//	irs map[string]Ir
//	order []*string
//	listOrder []bool
//	listS []string
//	listI []Ir
//}
//
//type IrS struct {
//	map[string]string
//	map[string]IrS
//	map[string]IrL
//	order []*string
//}
//
//type IrL struct {
//	strings []string
//	irss []IrS
//	irls []IrL
//	listOrder []string
//}
//
////type IrStruct interface {
////	GetStrings() map[string]string
////	GetIrStruct() map[string]IrStruct
////	GetIrList() map[string]IrList
////	GetOrder() []*string
////}
////
////type IrList interface {
////	GetPrimitive() []*string
////	GetNested() []IrStruct
////}
////
////type JustIrIntrfc interface {
////	GetStrings()
////	GetIrs()
////	GetOrder()
////	GetIrA()
////}
////
////type JustIrSetter interface {
////	SetStrings()
////	SetIrs()
////	SetOrder()
////	SetIrA()
////}
