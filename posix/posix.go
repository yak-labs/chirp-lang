package posix

import (
	"bufio"
	"bytes"
	. "fmt"
	. "github.com/yak-labs/chirp-lang"
	"io"
	"os"
	"os/exec"
	R "reflect"
	"runtime/pprof"
	"strings"
	"sync"
)

type terpFile struct {
	f *os.File
	r *bufio.Reader
	w *bufio.Writer
}

func MkFile(f *os.File) *terpFile {
	return &terpFile{f: f}
}

// *terpFile implements T

func (t terpFile) ChirpKind() string { return "File" }
func (t *terpFile) Raw() interface{} {
	return t.f
}
func (t *terpFile) String() string {
	return Sprintf("file%d", t.f.Fd())
}
func (t *terpFile) Float() float64 {
	panic("not implemented on terpFile (Float)")
}
func (t *terpFile) Int() int64 {
	panic("not implemented on terpFile (Int)")
}
func (t *terpFile) Uint() uint64 {
	panic("not implemented on terpFile (Uint)")
}
func (t *terpFile) ListElementString() string {
	return t.String()
}
func (t *terpFile) QuickString() string {
	return t.String()
}
func (t *terpFile) Bool() bool {
	panic("terpFile cannot be used as Bool")
}
func (t *terpFile) IsEmpty() bool {
	return false
}
func (t *terpFile) IsPreservedByList() bool { return true }
func (t *terpFile) IsQuickInt() bool     { return false }
func (t *terpFile) IsQuickNumber() bool     { return false }
func (t *terpFile) List() []T {
	return []T{t}
}
func (t *terpFile) HeadTail() (hd, tl T) {
	return MkList(t.List()).HeadTail()
}
func (t *terpFile) Hash() (Hash, *sync.Mutex) {
	panic("a terpFile is not a Hash")
}
func (t *terpFile) GetAt(key T) T {
	panic("a terpFile cannot GetAt")
}
func (t *terpFile) PutAt(value T, key T) {
	panic("a terpFile cannot PutAt")
}
func (t *terpFile) QuickReflectValue() R.Value {
	panic("a terpFile cannot QuickReflectValue")
}
func (t terpFile) EvalSeq(fr *Frame) T         { return Parse2EvalSeqStr(fr, t.String()) }
func (t terpFile) EvalExpr(fr *Frame) T        { return Parse2EvalExprStr(fr, t.String()) }
func (t terpFile) Apply(fr *Frame, args []T) T { panic("Cannot apply terpFile as command") }

func cmdOpen(fr *Frame, argv []T) T {
	nameT, args := Arg1v(argv)
	name := nameT.String()

	// access defaults to "r" if no extra arg.
	access := "r"
	if len(args) > 0 {
		access = args[0].String()
	}

	return Open(name, access)
}

func Open(name string, access string) T {
	var f *os.File
	var err error
	switch access {
	case "r":
		f, err = os.OpenFile(name, os.O_RDONLY, 0666)
	case "r+":
		f, err = os.OpenFile(name, os.O_RDWR, 0666)
	case "w":
		f, err = os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	case "w+":
		f, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
	case "a":
		f, err = os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	case "a+":
		f, err = os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	default:
		panic(Sprintf(`Unknown access mode in "open" command: %q`, access))
	}

	if err != nil {
		panic(Sprintf(`Cannot "open" file %q because %q`, name, err.Error()))
	}

	return &terpFile{f: f}
}

func cmdFlush(fr *Frame, argv []T) T {
	fileT := Arg1(argv)
	tf := fileT.(*terpFile)
	Flush(tf)
	return Empty
}

func Flush(tf *terpFile) {
	if tf.w != nil {
		tf.w.Flush()
	}
}

func cmdClose(fr *Frame, argv []T) T {
	fileT := Arg1(argv)
	tf := fileT.(*terpFile)
	Close(tf)
	return Empty
}

func Close(tf *terpFile) {
	if tf.w != nil {
		tf.w.Flush()
	}
	if tf.f != nil {
		tf.f.Close()
	}

	tf.f = nil
	tf.r = nil
	tf.w = nil
}

func cmdGets(fr *Frame, argv []T) T {
	fileT, args := Arg1v(argv)
	var varName string
	if len(args) > 1 {
		panic(`Too many args to "gets"`)
	}
	if len(args) > 0 {
		varName = args[0].String()
	}
	f := fileT.(*terpFile)

	if f.r == nil {
		f.r = bufio.NewReader(f.f)
	}

	data, err := f.r.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(Sprintf(`Error duing "gets": %s`, err.Error()))
	}
	if len(data) > 0 {
		if data[len(data)-1] == '\n' {
			data = data[:len(data)-1]
		}
	}
	dataT := MkString(data)

	if len(varName) > 0 {
		fr.SetVar(varName, dataT)
		return MkInt(int64(len(data)))
	}
	// else:
	return dataT
}

func cmdPuts(fr *Frame, argv []T) T {
	i := 1
	noNewLine := false
	if len(argv) > i {
		if strings.HasPrefix(argv[i].String(), "-n") && strings.HasPrefix("-nonewline", argv[i].String()) {
			noNewLine = true
			i++
		}
	}

	var t *terpFile
	var data string
	switch len(argv) {
	case i + 1:
		data = argv[i].String()
	case i + 2:
		var ok bool
		t, ok = argv[i].(*terpFile)
		if !ok {
			panic(Sprintf(`Bad args to "puts". Expected file as arg %d.`, i))
		}
		data = argv[i+1].String()
	default:
		panic(`Bad args to "puts"`)
	}

	Puts(noNewLine, t, data)
	return Empty
}

func Puts(noNewLine bool, t *terpFile, data string) {
	var err error
	if t == nil {
		if noNewLine {
			_, err = Print(data)
		} else {
			_, err = Println(data)
		}
	} else {
		if t.w == nil {
			t.w = bufio.NewWriter(t.f)
		}
		if noNewLine {
			_, err = Fprint(t.w, data)
		} else {
			_, err = Fprintln(t.w, data)
		}
	}
	if err != nil {
		panic(Sprintf(`Error during "puts": %s`, err.Error()))
	}
}

var fileEnsemble = []EnsembleItem{
	EnsembleItem{Name: "separator", Cmd: cmdFileSeparator},
	EnsembleItem{Name: "tempdir", Cmd: cmdFileTempdir},
	EnsembleItem{Name: "join", Cmd: cmdFileJoin},
}

func cmdFileSeparator(fr *Frame, argv []T) T {
	Arg0(argv)
	return MkString(string(os.PathSeparator))
}

func cmdFileTempdir(fr *Frame, argv []T) T {
	Arg0(argv)
	return MkString(os.TempDir())
}

func cmdFileJoin(fr *Frame, argv []T) T {
	panic("TODO")
}

type pipeNotification struct {
	which   int
	writing bool
	err     error
}

type pipeBuffer struct {
	buf      []byte
	producer io.ReadCloser
	consumer io.WriteCloser
	which    int // which pipe is this?
	notifier chan<- *pipeNotification
}

// run() could be faster if more than 1 buffer & more than 1 goroutine?
func (p *pipeBuffer) run() {
	var rc int
	var err error
	for {
		rc, err = p.producer.Read(p.buf) // rc will be used by Iterated write & final write.
		if err == io.EOF {
			break
		}
		if err != nil {
			p.notifier <- &pipeNotification{which: p.which, writing: false, err: err}
			break
		}
		if rc == 0 {
			continue
		}

		_, err = p.consumer.Write(p.buf[:rc]) // Iterated write.
		rc = 0
		if err != nil {
			p.notifier <- &pipeNotification{which: p.which, writing: true, err: err}
			break
		}
	}

	err = p.producer.Close()
	if err != nil {
		p.notifier <- &pipeNotification{which: p.which, writing: false, err: err}
	}

	if rc > 0 {
		_, err = p.consumer.Write(p.buf[:rc]) // Final write.
		if err != nil {
			p.notifier <- &pipeNotification{which: p.which, writing: true, err: err}
		}
	}

	err = p.consumer.Close()
	if err != nil {
		p.notifier <- &pipeNotification{which: p.which, writing: true, err: err}
	}

	p.notifier <- nil // The end.
}

func runBuffer(producer io.ReadCloser, consumer io.WriteCloser, which int, notifier chan<- *pipeNotification) {
	pb := &pipeBuffer{
		buf:      make([]byte, 4096),
		producer: producer,
		consumer: consumer,
		which:    which,
		notifier: notifier,
	}
	go pb.run()
}

type ExecStdoutError struct {
	which int
	data  string
}

func (p ExecStdoutError) Error() string { return p.data }

type Process struct {
	// Default stdin is process's normal stdin.
	stdin io.Reader

	// Default stdout is captured, and becomes result of exec command, unless background.
	outBuf bytes.Buffer
	stdout io.Writer

	// Default stderr is captured, and becomes panic text, unless background.
	errBuf bytes.Buffer
	stderr io.Writer

	background bool
	piped      bool

	words []string

	cmd *exec.Cmd
}

func (p *Process) runAndReportStatus(reporter chan<- error) {
	// Sayf("START %v", p)
	report := p.cmd.Run()
	Sayf("GOT_FROM_RUN %v", report)
	reporter <- report
	// Sayf("FINISH %v", p)
}

func (p *Process) useStandardDefaults() {
	if p.stdin == nil {
		p.stdin = os.Stdin
	}
	if p.stdout == nil {
		if p.background {
			p.stdout = os.Stdout
		} else {
			p.stdout = &p.outBuf
		}
	}
	if p.stderr == nil {
		if p.background {
			p.stderr = os.Stderr
		} else {
			p.stderr = &p.errBuf
		}
	}
}

func consumeAndIgnoreNotifications(notifier <-chan *pipeNotification) {
	for {
		note := <-notifier
		if note == nil { // When we read nil, that's the end.
			return
		} else {
			Sayf("Background Pipe Status: %s", note)
		}
	}
}

func consumeAndIgnoreReports(n int, notifier <-chan error) {
	i := 0
	for {
		report := <-notifier
		if len(report.Error()) > 0 {
			Sayf("Background Exit Status: %s", report)
		}
		i++
		if i == n {
			return
		}
	}
}

// "exec" command.  Supports "< << > >> 2> 2>> &" when they are separte words.
// TODO:  Pipes.
func cmdExec(fr *Frame, argv []T) T {
	argsT := Arg0v(argv)
	notifier := make(chan *pipeNotification, 8) // Notifies us of errors in pipe copiers.

	reporter := make(chan error, 8) // Report exit status of processes.

	p := new(Process)
	p.words = make([]string, 0)
	// Add p to processes when it is finished.
	processes := make([]*Process, 0)
	_ = processes

	state := ""
	args := make([]string, len(argsT))
	for i, a := range argsT {
		args[i] = a.String()
	}

	for _, a := range args {
		var err error
		switch state {
		case "":
			switch a {
			case "<":
				state = a
			case "<<":
				state = a
			case ">":
				state = a
			case ">>":
				state = a
			case "2>":
				state = a
			case "2>>":
				state = a
			case "&":
				p.background = true
			case "|":
				processes = append(processes, p)
				p = new(Process)
				p.words = make([]string, 0)
			default:
				p.words = append(p.words, a)
			}
		case "<":
			p.stdin, err = os.Open(a)
			state = ""
		case "<<":
			p.stdin = strings.NewReader(a)
			state = ""
		case ">":
			p.stdout, err = os.OpenFile(a, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			state = ""
		case ">>":
			p.stdout, err = os.OpenFile(a, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			state = ""
		case "2>":
			p.stderr, err = os.OpenFile(a, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			state = ""
		case "2>>":
			p.stderr, err = os.OpenFile(a, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			state = ""
		}
		if err != nil {
			panic(Sprintf(`ERROR in redirection in "exec" command: %s`, err.Error()))
		}
	}

	processes = append(processes, p)

	n := len(processes)
	for i := 0; i < n; i++ {
		p = processes[i]
		p.useStandardDefaults()

		p.cmd = exec.Command(p.words[0], p.words[1:]...)
		p.cmd.Stdin = p.stdin
		p.cmd.Stdout = p.stdout
		p.cmd.Stderr = p.stderr
	}

	for i := 0; i < n; i++ {
		p = processes[i]
		isFinal := (i+1 == n)
		if !isFinal {
			// This must pipe into the next one.
			p.cmd.Stdout = nil // Necessary, to avoid error in StdoutPipe()
			rc, err1 := p.cmd.StdoutPipe()
			if err1 != nil {
				panic(Sprintf("p.cmd.StdoutPipe: %s", err1))
			}

			q := processes[i+1] // q is next in the pipeline after p.
			q.cmd.Stdin = nil   // Necessary, to avoid error in StdinPipe()
			wc, err2 := q.cmd.StdinPipe()
			if err2 != nil {
				panic(Sprintf("p.cmd.StdinPipe: %s", err2))
			}

			runBuffer(rc, wc, i, notifier)
		} else {
			if p.background {
				err := p.cmd.Start()
				if err != nil {
					panic(Sprintf("ERROR in \"exec\" of background %q: %s", p.words[0], err.Error()))
				}
				consumeAndIgnoreNotifications(notifier)
				consumeAndIgnoreReports(n, reporter)
				return Empty
			}
		}
		go p.runAndReportStatus(reporter)
	}

	// First read notes from the pipes.
	for i := 0; i < n-1; {
		note := <-notifier
		if note == nil {
			i++
		} else {
			Sayf("Pipe problem: %v", *note)
		}
	}

	// Then collect exit status
	var anError error
	for i := 0; i < n; i++ {
		report := <-reporter
		Sayf("Collecting Exit Status: got%d: %v", i, report)
		if report != nil {
			Sayf("Bad Exit Status: %v", report)
			anError = report
		}
	}

	for i := 0; i < n; i++ {
		errStr := p.errBuf.String()
		if len(errStr) > 0 {
			Sayf("Stderr of command #%d: %v", i, errStr)
		}
		if anError == nil {
			anError = ExecStdoutError{which: i, data: errStr}
		}
	}

	if anError != nil && len(anError.Error()) > 0 {
		panic(Sprintf("ERROR in \"exec\": %s", anError.Error()))
	}

	return MkString(processes[n-1].outBuf.String())
}

func cmdEvalWithProfile(fr *Frame, argv []T) T {
	filenameT, seqT := Arg2(argv)
	filename := filenameT.String()
	w, err := os.Create(filename)
	if err != nil {
		panic(Sprintf("eval_with_profile: Cannot Create: %s: %s", filename, err.Error()))
	}

	err = pprof.StartCPUProfile(w)
	z := fr.Eval(seqT)
	pprof.StopCPUProfile()
	return z
}

func init() {
	if Unsafes == nil {
		Unsafes = make(map[string]Command, 333)
	}

	Unsafes["open"] = cmdOpen
	Unsafes["close"] = cmdClose
	Unsafes["file"] = MkEnsemble(fileEnsemble)
	Unsafes["gets"] = cmdGets
	Unsafes["puts"] = cmdPuts
	Unsafes["flush"] = cmdFlush
	Unsafes["exec"] = cmdExec
	Unsafes["eval_with_profile"] = cmdEvalWithProfile
}
