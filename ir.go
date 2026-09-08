package gosygr

type Struct struct {
	stringFields map[string]string
	structFields map[string]*Struct
	arrayFields map[string]*Array
	order []string
}

type Array struct {
	strings []string
	structs []*Struct
	arrays []*Array
}

