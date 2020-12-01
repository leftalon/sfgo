

# e
`import "gitlab.xsjcs.cn/tc-fwk/tcgo/core/e"`

* [总览](#pkg-overview)
* [索引](#pkg-index)
* [示例](#pkg-examples)

## <a name="pkg-overview">总览</a>
### 错误处理模块
* 生成错误：

在将错误值返回上级函数时，包含最原始的错误上下文信息，比如：错误链，错误发生时的stacktrace。

有2个方法：


	1. New：生成全新的错误值。
	2. Wrap：新错误值中封装了原错误值。在外层进行错误处理时，可以追溯到内部任意层的原始错误。

* 错误处理：

在外层函数进行错误处理时，有2种做法：


	1. Is：判断当前错误是否来自某个哨兵错误值。哨兵错误是预生成的错误值(比如：io.EOF)，通常用于错误信息简单，无需包含附加信息的情况。
	2. As: 判断当前错误是否来自某种错误类型。错误类型通常包含一些错误发生时的附加信息(比如：错误码)，可根据这些附加信息进行复杂的错误处理。

尽量不要根据错误值的Error()的内容进行错误处理。




## <a name="pkg-index">索引</a>
* [func As(err error, target interface{}) bool](#As)
* [func Cause(err error) error](#Cause)
* [func Is(err, target error) bool](#Is)
* [func New(format string, args ...interface{}) error](#New)
* [func Unwrap(err error) error](#Unwrap)
* [func Wrap(err error, format string, args ...interface{}) error](#Wrap)
* [type Frame](#Frame)
  * [func (f Frame) Format(s fmt.State, verb rune)](#Frame.Format)
  * [func (f Frame) MarshalText() ([]byte, error)](#Frame.MarshalText)
* [type StackTrace](#StackTrace)
  * [func (st StackTrace) Format(s fmt.State, verb rune)](#StackTrace.Format)

#### <a name="pkg-examples">示例</a>
* [New](#example_New)
* [Wrap](#example_Wrap)
* [Wrap (Args)](#example_Wrap_args)
* [Wrap (As)](#example_Wrap_as)
* [Wrap (Is)](#example_Wrap_is)
* [Wrap (Nil)](#example_Wrap_nil)

#### <a name="pkg-files">包内源文件</a>

* [e.go](/src/gitlab.xsjcs.cn/tc-fwk/tcgo/core/e/e.go)
* [errors.go](/src/gitlab.xsjcs.cn/tc-fwk/tcgo/core/e/errors.go)
* [go113.go](/src/gitlab.xsjcs.cn/tc-fwk/tcgo/core/e/go113.go)
* [stack.go](/src/gitlab.xsjcs.cn/tc-fwk/tcgo/core/e/stack.go)





## <a name="As">func</a> [As](/tc-fwk/tcgo/blob/master/core/e/go113.go?s=1141:1184#L31)
``` go
func As(err error, target interface{}) bool
```
As finds the first error in err's chain that matches target, and if so, sets
target to that error value and returns true.

The chain consists of err itself followed by the sequence of errors obtained by
repeatedly calling Unwrap.

An error matches target if the error's concrete value is assignable to the value
pointed to by target, or if the error has a method As(interface{}) bool such that
As(target) returns true. In the latter case, the As method is responsible for
setting target.

As will panic if target is not a non-nil pointer to either a type that implements
error, or to any interface type. As returns false if err is nil.



## <a name="Cause">func</a> [Cause](/tc-fwk/tcgo/blob/master/core/e/e.go?s=2072:2099#L72)
``` go
func Cause(err error) error
```
Cause returns the underlying cause of the error, if possible.
An error value has a cause if it implements the following
interface:


	type causer interface {
	       Cause() error
	}

If the error does not implement Cause, the original error will
be returned. If the error is nil, nil will be returned without further
investigation.



## <a name="Is">func</a> [Is](/tc-fwk/tcgo/blob/master/core/e/go113.go?s=399:430#L16)
``` go
func Is(err, target error) bool
```
Is reports whether any error in err's chain matches target.

The chain consists of err itself followed by the sequence of errors obtained by
repeatedly calling Unwrap.

An error is considered to match a target if it is equal to that target or if
it implements a method Is(error) bool such that Is(target) returns true.



## <a name="New">func</a> [New](/tc-fwk/tcgo/blob/master/core/e/e.go?s=1106:1156#L30)
``` go
func New(format string, args ...interface{}) error
```
New formats according to a format specifier and returns the string
as a value that satisfies error.
New also records the stack trace at the point it was called.



## <a name="Unwrap">func</a> [Unwrap](/tc-fwk/tcgo/blob/master/core/e/go113.go?s=1383:1411#L36)
``` go
func Unwrap(err error) error
```
Unwrap returns the result of calling the Unwrap method on err, if err's
type contains an Unwrap method returning error.
Otherwise, Unwrap returns nil.



## <a name="Wrap">func</a> [Wrap](/tc-fwk/tcgo/blob/master/core/e/e.go?s=1401:1463#L40)
``` go
func Wrap(err error, format string, args ...interface{}) error
```
Wrap returns an error annotating err with a stack trace
at the point Wrapf is called, and the format specifier.
If err is nil, Wrapf returns nil.




## <a name="Frame">type</a> [Frame](/tc-fwk/tcgo/blob/master/core/e/stack.go?s=249:267#L15)
``` go
type Frame uintptr
```
Frame represents a program counter inside a stack frame.
For historical reasons if Frame is interpreted as a uintptr
its value represents the program counter + 1.










### <a name="Frame.Format">func</a> (Frame) [Format](/tc-fwk/tcgo/blob/master/core/e/stack.go?s=1498:1543#L64)
``` go
func (f Frame) Format(s fmt.State, verb rune)
```
Format formats the frame according to the fmt.Formatter interface.


	%s    source file
	%d    source line
	%n    function name
	%v    equivalent to %s:%d

Format accepts flags that alter the printing of some verbs, as follows:


	%+s   function name and path of source file relative to the compile time
	      GOPATH separated by \n\t (<funcname>\n\t<path>)
	%+v   equivalent to %+s:%d




### <a name="Frame.MarshalText">func</a> (Frame) [MarshalText](/tc-fwk/tcgo/blob/master/core/e/stack.go?s=2087:2131#L88)
``` go
func (f Frame) MarshalText() ([]byte, error)
```
MarshalText formats a stacktrace Frame as a text string. The output is the
same as that of fmt.Sprintf("%+v", f), but without newlines or tabs.




## <a name="StackTrace">type</a> [StackTrace](/tc-fwk/tcgo/blob/master/core/e/stack.go?s=2360:2383#L97)
``` go
type StackTrace []Frame
```
StackTrace is stack of Frames from innermost (newest) to outermost (oldest).










### <a name="StackTrace.Format">func</a> (StackTrace) [Format](/tc-fwk/tcgo/blob/master/core/e/stack.go?s=2764:2815#L107)
``` go
func (st StackTrace) Format(s fmt.State, verb rune)
```
Format formats the stack of Frames according to the fmt.Formatter interface.


	%s	lists source files for each Frame in the stack
	%v	lists the source file and line number for each Frame in the stack

Format accepts flags that alter the printing of some verbs, as follows:


	%+v   Prints filename, function, and line number for each Frame in the stack.








- - -
Generated by [godoc2md](http://godoc.org/github.com/github.com/GlenDC)
