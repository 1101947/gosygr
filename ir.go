package gosygr

type irStruct struct {
	stringFields map[string]string
	structFields map[string]*irStruct
	arrayFields map[string]*irArray
	order []*string
}

type irArray struct {
	strings []string
	structs []*irStruct
	arrays []*irArray
}

