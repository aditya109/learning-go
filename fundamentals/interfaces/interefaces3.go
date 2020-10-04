package main

func main() {

}

type Reader interface {
	Read(b [] byte) (int, error)
}


type Writer__ interface {
	Write(b [] byte) (int, error)
}

type ReaderWriterUnion1 interface {
	Read(b[] byte) (int, error)
	Write(b[] byte) (int, error)
}


type ReaderWriterUnion2 interface {
	Reader
	Writer__
}