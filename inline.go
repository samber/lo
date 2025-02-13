package lo

import "io"

type Printer struct {
	fn func(...interface{})
}

func (x Printer) Print(v ...interface{}) { x.fn(v...) }

var _ interface{ Print(...interface{}) } = Printer{}

func BePrinter(fn func(...interface{})) Printer { return Printer{fn} }

type Printfer struct {
	fn func(string, ...interface{})
}

func (x Printfer) Printf(format string, v ...interface{}) { x.fn(format, v...) }

var _ interface{ Printf(string, ...interface{}) } = Printfer{}

func BePrintfer(fn func(string, ...interface{})) Printfer { return Printfer{fn} }

type Writer struct {
	fn func([]byte) (int, error)
}

func (x Writer) Write(p []byte) (int, error) { return x.fn(p) }

var _ io.Writer = Writer{}

func BeWriter(fn func([]byte) (int, error)) io.Writer { return Writer{fn} }

type Reader struct {
	fn func([]byte) (int, error)
}

func (x Reader) Read(p []byte) (int, error) { return x.fn(p) }

var _ io.Reader = Reader{}

func BeReader(fn func([]byte) (int, error)) io.Reader { return Reader{fn} }

type Closer struct {
	fn func() error
}

func (x Closer) Close() error { return x.fn() }

var _ io.Closer = Closer{}

func BeCloser(fn func() error) io.Closer { return Closer{fn} }

type Seeker struct {
	fn func(int64, int) (int64, error)
}

func (x Seeker) Seek(offset int64, whence int) (int64, error) { return x.fn(offset, whence) }

var _ io.Seeker = Seeker{}

func BeSeeker(fn func(int64, int) (int64, error)) io.Seeker { return Seeker{fn} }

type ReadWriter struct {
	Reader
	Writer
}

var _ io.ReadWriter = ReadWriter{}

func BeReadWriter(read func([]byte) (int, error), write func([]byte) (int, error)) io.ReadWriter {
	return ReadWriter{Reader{read}, Writer{write}}
}
