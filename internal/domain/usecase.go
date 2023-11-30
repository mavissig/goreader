package domain

type Reader interface {
	Read(string) (string, MainStruct)
}

type Printer interface {
	Print(string, MainStruct)
}

type Comparator interface {
	Eq(MainStruct, MainStruct)
}

type UseCase struct {
	read    Reader
	print   Printer
	compare Comparator
}

func (uc *UseCase) ConvertAndPrintFile(path string) {
	sfx, data := uc.read.Read(path)
	if sfx == ".json" {
		uc.print.Print(".xml", data)
	} else if sfx == ".xml" {
		uc.print.Print(".json", data)
	}
}

func (uc *UseCase) CompareFiles(path1, path2 string) {
	_, data1 := uc.read.Read(path1)
	_, data2 := uc.read.Read(path2)
	uc.compare.Eq(data1, data2)
}

func New(reader Reader, printer Printer, comparator Comparator) *UseCase {
	return &UseCase{
		read:    reader,
		print:   printer,
		compare: comparator,
	}
}
