package classpath

type Entry interface {
	readClass(className string)([]byte,Entry,error)
	String() string
}

func newEntry(path string)Entry {
}
