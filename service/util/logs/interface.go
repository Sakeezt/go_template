package logs

//go:generate mockery --name=Log
type Log interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
	Panic(v ...interface{})
	System(v ...interface{})
	Struct(v interface{})
	Line()
}
