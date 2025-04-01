 func main(){
	var wc writerCloser = NewBufferdWriter()
	we.Write([]byte("Hello, GO!")) // calling the write method of the interface and passing the string as byte array
	wc.Close() // calling the close method of the interface and passing the string as byte array
}
type Writer interface{
	write([]byte) (int, error)
}
type Closer interface{
	close() error
}
type writerCloser interface{
	Writer
Closer // embedding the writer and closer interface
}
type bufferedWriterCloser interface {
	buffer *bytes.BUfferbuffer
	
}