package test

type Expecter interface {
	NoArg() string
	NoReturn(str string)
	ManyArgsReturns(str string, i int) (strs []string, err error)
	Variadic(ints ...int) error
	VariadicMany(i int, a string, intfs ...interface{}) error
}

type VariadicNoReturnInterface interface {
	VariadicNoReturn(j int, is ...interface{})
}

type ExpecterNoArgArgs struct {
	StrAnything bool
}

type ExpecterNoArgReturns struct {
	Str string
}

type ExpecterNoArgExpectation struct {
	Args    ExpecterNoArgArgs
	Returns ExpecterNoArgReturns
}

type ExpecterNoReturnArgs struct {
	Str         string
	StrAnything bool
}

type ExpecterNoReturnReturns struct {
}

type ExpecterNoReturnExpectation struct {
	Args    ExpecterNoReturnArgs
	Returns ExpecterNoReturnReturns
}

type ExpecterManyArgsReturnsArgs struct {
	Str         string
	I           int
	StrAnything bool
	IAnything   bool
}

type ExpecterManyArgsReturnsReturns struct {
	Strs []string
	Err  error
}

type ExpecterManyArgsReturnsExpectation struct {
	Args    ExpecterManyArgsReturnsArgs
	Returns ExpecterManyArgsReturnsReturns
}

type ExpecterVariadicArgs struct {
	Ints       []int
	IntsAnything bool
}

type ExpecterVariadicReturns struct {
	Err error
}

type ExpecterVariadicExpectation struct {
	Args    ExpecterVariadicArgs
	Returns ExpecterVariadicReturns
}

type ExpecterVariadicManyArgs struct {
	I           int
	A           string
	Intfs       []interface{}
	IAnything   bool
	AAnything   bool
	IntfsAnything bool
}

type ExpecterVariadicManyReturns struct {
	Err error
}

type ExpecterVariadicManyExpectation struct {
	Args    ExpecterVariadicManyArgs
	Returns ExpecterVariadicManyReturns
}
