# 1.3 fmt — 格式化IO #

fmt 包实现了格式化I/O函数，类似于C的 printf 和 scanf. 格式“占位符”衍生自C，但比C更简单。

fmt 包的官方文档对 Printing 和 Scanning 有很详细的说明。这里就直接引用文档进行说明，同时附上额外的说明或例子，之后再介绍具体的函数使用。

以下例子中用到的类型或变量定义：
```go
	type Website struct {
	    Name string
	}
	
	// 定义结构体变量
	var site = Website{Name:"studygolang"}
```
## Printing ##
### Sample
```go

type user struct {
	name string
}

func main() {
	u := user{"tang"}
	//Printf 格式化输出
	fmt.Printf("% + v\n", u)     //格式化输出结构
	fmt.Printf("%#v\n", u)       //输出值的 Go 语言表示方法
	fmt.Printf("%T\n", u)        //输出值的类型的 Go 语言表示
	fmt.Printf("%t\n", true)     //输出值的 true 或 false
	fmt.Printf("%b\n", 1024)     //二进制表示
	fmt.Printf("%c\n", 11111111) //数值对应的 Unicode 编码字符
	fmt.Printf("%d\n", 10)       //十进制表示
	fmt.Printf("%o\n", 8)        //八进制表示
	fmt.Printf("%q\n", 22)       //转化为十六进制并附上单引号
	fmt.Printf("%x\n", 1223)     //十六进制表示，用a-f表示
	fmt.Printf("%X\n", 1223)     //十六进制表示，用A-F表示
	fmt.Printf("%U\n", 1233)     //Unicode表示
	fmt.Printf("%b\n", 12.34)    //无小数部分，两位指数的科学计数法6946802425218990p-49
	fmt.Printf("%e\n", 12.345)   //科学计数法，e表示
	fmt.Printf("%E\n", 12.34455) //科学计数法，E表示
	fmt.Printf("%f\n", 12.3456)  //有小数部分，无指数部分
	fmt.Printf("%g\n", 12.3456)  //根据实际情况采用%e或%f输出
	fmt.Printf("%G\n", 12.3456)  //根据实际情况采用%E或%f输出
	fmt.Printf("%s\n", "wqdew")  //直接输出字符串或者[]byte
	fmt.Printf("%q\n", "dedede") //双引号括起来的字符串
	fmt.Printf("%x\n", "abczxc") //每个字节用两字节十六进制表示，a-f表示
	fmt.Printf("%X\n", "asdzxc") //每个字节用两字节十六进制表示，A-F表示
	fmt.Printf("%p\n", 0x123)    //0x开头的十六进制数表示
}
```
### 占位符 ###

**普通占位符**
	
	占位符						说明						举例										输出
	%v		相应值的默认格式。								Printf("%v", site)，Printf("%+v", site)	{studygolang}，{Name:studygolang}
			在打印结构体时，“加号”标记（%+v）会添加字段名
	%#v		相应值的Go语法表示							Printf("#v", site)						main.Website{Name:"studygolang"}
	%T		相应值的类型的Go语法表示						Printf("%T", site)						main.Website
	%%		字面上的百分号，并非值的占位符					Printf("%%")							%

**布尔占位符**

	占位符						说明						举例										输出
	%t		单词 true 或 false。							Printf("%t", true)						true

**整数占位符**

	占位符						说明						举例									输出
	%b		二进制表示									Printf("%b", 5)						101
	%c		相应Unicode码点所表示的字符					Printf("%c", 0x4E2D)				中
	%d		十进制表示									Printf("%d", 0x12)					18
	%o		八进制表示									Printf("%o", 10)					12
	%q		单引号围绕的字符字面值，由Go语法安全地转义		Printf("%q", 0x4E2D)				'中'
	%x		十六进制表示，字母形式为小写 a-f				Printf("%x", 13)					d
	%X		十六进制表示，字母形式为大写 A-F				Printf("%x", 13)					D
	%U		Unicode格式：U+1234，等同于 "U+%04X"			Printf("%U", 0x4E2D)				U+4E2D

**浮点数和复数的组成部分（实部和虚部）**

	占位符						说明												举例									输出
	%b		无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat	
			的 'b' 转换格式一致。例如 -123456p-78
	%e		科学计数法，例如 -1234.456e+78									Printf("%e", 10.2)							1.020000e+01
	%E		科学计数法，例如 -1234.456E+78									Printf("%e", 10.2)							1.020000E+01
	%f		有小数点而无指数，例如 123.456									Printf("%f", 10.2)							10.200000
	%g		根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出				Printf("%g", 10.20)							10.2
	%G		根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出				Printf("%G", 10.20+2i)						(10.2+2i)

**字符串与字节切片**

	占位符						说明												举例									输出
	%s		输出字符串表示（string类型或[]byte)							Printf("%s", []byte("Go语言中文网"))		Go语言中文网
	%q		双引号围绕的字符串，由Go语法安全地转义							Printf("%q", "Go语言中文网")				"Go语言中文网"
	%x		十六进制，小写字母，每字节两个字符								Printf("%x", "golang")						676f6c616e67
	%X		十六进制，大写字母，每字节两个字符								Printf("%X", "golang")						676F6C616E67

**指针**

	占位符						说明												举例									输出
	%p		十六进制表示，前缀 0x											Printf("%p", &site)							0x4f57f0
	
这里没有 'u' 标记。若整数为无符号类型，他们就会被打印成无符号的。类似地，这里也不需要指定操作数的大小（int8，int64）。

宽度与精度的控制格式以 Unicode 码点为单位。（这点与C的 printf 不同，它以字节数为单位）二者或其中之一均可用字符 '*' 表示，此时它们的值会从下一个操作数中获取，该操作数的类型必须为 int。

对数值而言，宽度为该数值占用区域的最小宽度；精度为小数点之后的位数。 但对于 %g/%G 而言，精度为所有数字的总数。例如，对于123.45，格式 %6.2f 会打印123.45，而 %.4g 会打印123.5。%e 和 %f 的默认精度为6；但对于 %g 而言，它的默认精度为确定该值所必须的最小位数。

对大多数的值而言，宽度为输出的最小字符数，如果必要的话会为已格式化的形式填充空格。对字符串而言，精度为输出的最大字符数，如果必要的话会直接截断。

**其它标记**

	占位符						说明												举例									输出
	+		总打印数值的正负号；对于%q（%+q）保证只输出ASCII编码的字符。			Printf("%+q", "中文")					"\u4e2d\u6587"
	-		在右侧而非左侧填充空格（左对齐该区域）
	#		备用格式：为八进制添加前导 0（%#o），为十六进制添加前导 0x（%#x）或	Printf("%#U", '中')						U+4E2D '中'
			0X（%#X），为 %p（%#p）去掉前导 0x；如果可能的话，%q（%#q）会打印原始
			（即反引号围绕的）字符串；如果是可打印字符，%U（%#U）会写出该字符的
			Unicode 编码形式（如字符 x 会被打印成 U+0078 'x'）。
	' '		（空格）为数值中省略的正负号留出空白（% d）；
			以十六进制（% x, % X）打印字符串或切片时，在字节之间用空格隔开
	0		填充前导的0而非空格；对于数字，这会将填充移到正负号之后

标记有时会被占位符忽略，所以不要指望它们。例如十进制没有备用格式，因此 %#d 与 %d 的行为相同。

对于每一个 Printf 类的函数，都有一个 Print 函数，该函数不接受任何格式化，它等价于对每一个操作数都应用 %v。另一个变参函数 Println 会在操作数之间插入空白，并在末尾追加一个换行符。

不考虑占位符的话，如果操作数是接口值，就会使用其内部的具体值，而非接口本身。 因此：
```go
	var i interface{} = 23
	fmt.Printf("%v\n", i)
```
会打印 23。

若一个操作数实现了 Formatter 接口，该接口就能更好地用于控制格式化。

若其格式（它对于 Println 等函数是隐式的 %v）对于字符串是有效的 （%s %q %v %x %X），以下两条规则也适用：
```go
	1. 若一个操作数实现了 error 接口，Error 方法就能将该对象转换为字符串，随后会根据占位符的需要进行格式化。
	2. 若一个操作数实现了 String() string 方法，该方法能将该对象转换为字符串，随后会根据占位符的需要进行格式化。
```
为避免以下这类递归的情况：
```go
	type X string
	func (x X) String() string { return Sprintf("<%s>", x) }
```
需要在递归前转换该值：
```go
	func (x X) String() string { return Sprintf("<%s>", string(x)) }
```
**格式化错误**

如果给占位符提供了无效的实参（例如将一个字符串提供给 %d），所生成的字符串会包含该问题的描述，如下例所示：
```bash
	类型错误或占位符未知：%!verb(type=value)
		Printf("%d", hi):          %!d(string=hi)
	实参太多：%!(EXTRA type=value)
		Printf("hi", "guys"):      hi%!(EXTRA string=guys)
	实参太少： %!verb(MISSING)
		Printf("hi%d"):            hi %!d(MISSING)
	宽度或精度不是int类型: %!(BADWIDTH) 或 %!(BADPREC)
		Printf("%*s", 4.5, "hi"):  %!(BADWIDTH)hi
		Printf("%.*s", 4.5, "hi"): %!(BADPREC)hi
	所有错误都始于“%!”，有时紧跟着单个字符（占位符），并以小括号括住的描述结尾。
```
## Scanning ##

一组类似的函数通过扫描已格式化的文本来产生值。
Scan、Scanf 和 Scanln 从 os.Stdin 中读取；
Fscan、Fscanf 和 Fscanln 从指定的 io.Reader 中读取； 
Sscan、Sscanf 和 Sscanln 从实参字符串中读取。
Scanln、Fscanln 和 Sscanln 在换行符处停止扫描，且需要条目紧随换行符之后；
Scanf、Fscanf 和 Sscanf 需要输入换行符来匹配格式中的换行符；其它函数则将换行符视为空格。

Scanf、Fscanf 和 Sscanf 根据格式字符串解析实参，类似于 Printf。例如，%x 会将一个整数扫描为十六进制数，而 %v 则会扫描该值的默认表现格式。

格式化行为类似于 Printf，但也有如下例外：

	%p 没有实现
	%T 没有实现
	%e %E %f %F %g %G 都完全等价，且可扫描任何浮点数或复数数值
	%s 和 %v 在扫描字符串时会将其中的空格作为分隔符
	标记 # 和 + 没有实现

在使用 %v 占位符扫描整数时，可接受友好的进制前缀0（八进制）和0x（十六进制）。

宽度被解释为输入的文本（%5s 意为最多从输入中读取5个 rune 来扫描成字符串），而扫描函数则没有精度的语法（没有 %5.2f，只有 %5f）。

当以某种格式进行扫描时，无论在格式中还是在输入中，所有非空的连续空白字符 （除换行符外）都等价于单个空格。由于这种限制，格式字符串文本必须匹配输入的文本，如果不匹配，扫描过程就会停止，并返回已扫描的实参数。

在所有的扫描参数中，若一个操作数实现了 Scan 方法（即它实现了 Scanner 接口）， 该操作数将使用该方法扫描其文本。此外，若已扫描的实参数少于所提供的实参数，就会返回一个错误。

所有需要被扫描的实参都必须是基本类型或 Scanner 接口的实现。

注意：Fscan 等函数会从输入中多读取一个字符（rune），因此，如果循环调用扫描函数，可能会跳过输入中的某些数据。一般只有在输入的数据中没有空白符时该问题才会出现。若提供给 Fscan 的读取器实现了 ReadRune，就会用该方法读取字符。若此读取器还实现了 UnreadRune 方法，就会用该方法保存字符，而连续的调用将不会丢失数据。若要为没有 ReadRune 和 UnreadRune 方法的读取器加上这些功能，需使用 bufio.NewReader。

## Print 序列函数 ##

这里说的 Print 序列函数包括：Fprint/Fprintf/Fprintln/Sprint/Sprintf/Sprintln/Print/Printf/Println。之所以将放在一起介绍，是因为它们的使用方式类似、参数意思也类似。

一般的，我们将 Fprint/Fprintf/Fprintln 归为一类；Sprint/Sprintf/Sprintln 归为一类；Print/Printf/Println 归为另一类。其中，Print/Printf/Println 会调用相应的F开头一类函数。如：
```go
	func Print(a ...interface{}) (n int, err error) {
		return Fprint(os.Stdout, a...)
	}
```
Fprint/Fprintf/Fprintln 函数的第一个参数接收一个io.Writer类型，会将内容输出到 io.Writer 中去。而 Print/Printf/Println 函数是将内容输出到标准输出中，因此，直接调用 F类函数 做这件事，并将 os.Stdout 作为第一个参数传入。

Sprint/Sprintf/Sprintln 是格式化内容为 string 类型，而并不输出到某处，需要格式化字符串并返回时，可以用这组函数。

在这三组函数中，`S/F/Printf`函数通过指定的格式输出或格式化内容；`S/F/Print`函数只是使用默认的格式输出或格式化内容；`S/F/Println`函数使用默认的格式输出或格式化内容，同时会在最后加上"换行符"。

Print 序列函数的最后一个参数都是 `a ...interface{}` 这种不定参数。对于`S/F/Printf`序列，这个不定参数的实参个数应该和`formt`参数的占位符个数一致，否则会出现格式化错误；而对于其他函数，当不定参数的实参个数为多个时，它们之间会直接（对于`S/F/Print`）或通过" "（空格）（对于`S/F/Println`）连接起来（注：对于`S/F/Print`，当两个参数都不是字符串时，会自动添加一个空格，否则不会加。感谢guoshanhe1983 反馈。[官方 effective_go](http://docs.studygolang.com/doc/effective_go.html#Printing) 也有说明）。利用这一点，我们可以做如下事情：

	result1 := fmt.Sprintln("studygolang.com", 2013)
	result2 := fmt.Sprint("studygolang.com", 2013)
	
result1的值是：`studygolang.com 2013`，result2的值是：`studygolang.com2013`。这起到了连接字符串的作用，而不需要通过`strconv.Itoa()`转换。

Print 序列函数用的较多，而且也易于使用（可能需要掌握一些常用的占位符用法），接下来我们结合 fmt 包中几个相关的接口来掌握更多关于 Print 的内容。

## Stringer 接口 ##

Stringer接口的定义如下：
```go
	type Stringer interface {
	    String() string
	}
```
根据 Go 语言中实现接口的定义，一个类型只要有 `String() string` 方法，我们就说它实现了 Stringer 接口。而在本节开始已经说到，如果格式化输出某种类型的值，只要它实现了 String() 方法，那么会调用 String() 方法进行处理。

我们定义如下struct：
```go
	type Person struct {
		Name string
		Age  int
		Sex  int
	}
```
我们给Person实现String方法，这个时候，我们输出Person的实例：
```go
	p := &Person{"polaris", 28, 0}
	fmt.Println(p)
```
输出：
```bash
	&{polaris 28 0}
```
接下来，为Person增加String方法。
```go
	func (this *Person) String() string {
		buffer := bytes.NewBufferString("This is ")
		buffer.WriteString(this.Name + ", ")
		if this.Sex == 0 {
			buffer.WriteString("He ")
		} else {
			buffer.WriteString("She ")
		}

		buffer.WriteString("is ")
		buffer.WriteString(strconv.Itoa(this.Age))
		buffer.WriteString(" years old.")
		return buffer.String()
	}
```
这个时候运行：
```go
	p := &Person{"polaris", 28, 0}
	fmt.Println(p)
```
输出变为：
```bash
	This is polaris, He is 28 years old
```
可见，Stringer接口和Java中的ToString方法类似。

## Formatter 接口 ##

Formatter 接口的定义如下：
```go
	type Formatter interface {
	    Format(f State, c rune)
	}
```
官方文档中关于该接口方法的说明：

> Formatter 接口由带有定制的格式化器的值所实现。 Format 的实现可调用 Sprintf 或 Fprintf(f) 等函数来生成其输出。

也就是说，通过实现 Formatter 接口可以做到自定义输出格式（自定义占位符）。

接着上面的例子，我们为 Person 增加一个方法：
```go
	func (this *Person) Format(f fmt.State, c rune) {
		if c == 'L' {
			f.Write([]byte(this.String()))
			f.Write([]byte(" Person has three fields."))
		} else {
			// 没有此句，会导致 fmt.Printf("%s", p) 啥也不输出
		    f.Write([]byte(fmt.Sprintln(this.String())))
		}
	}
```
这样，Person便实现了Formatter接口。这时再运行：
```go
	p := &Person{"polaris", 28, 0}
	fmt.Printf("%L", p)
```
输出为：
```bash
	This is polaris, He is 28 years old. Person has three fields.
```
这里需要解释以下几点：

1）fmt.State 是一个接口。由于 Format 方法是被 fmt 包调用的，它内部会实例化好一个 fmt.State 接口的实例，我们不需要关心该接口；

2）可以实现自定义占位符，同时 fmt 包中和类型相对应的预定义占位符会无效。因此例子中 Format 的实现加上了 else 子句；

3）实现了 Formatter 接口，相应的 Stringer 接口不起作用。但实现了 Formatter 接口的类型应该实现 Stringer 接口，这样方便在 Format 方法中调用 String() 方法。就像本例的做法；

4）Format 方法的第二个参数是占位符中%后的字母（有精度和宽度会被忽略，只保留字母）；

一般地，我们不需要实现 Formatter 接口。如果对 Formatter 接口的实现感兴趣，可以看看标准库 [math/big](http://docscn.studygolang.com/src/math/big/floatconv.go?s=7989:8041#L261) 包中 Int 类型的 Formatter 接口实现。

**小贴士**

State接口相关说明：
```go
	type State interface {
	    // Write is the function to call to emit formatted output to be printed.
	    // Write 函数用于打印出已格式化的输出。
	    Write(b []byte) (ret int, err error)
	    // Width returns the value of the width option and whether it has been set.
	    // Width 返回宽度选项的值以及它是否已被设置。
	    Width() (wid int, ok bool)
	    // Precision returns the value of the precision option and whether it has been set.
	    // Precision 返回精度选项的值以及它是否已被设置。
	    Precision() (prec int, ok bool)
	
	    // Flag returns whether the flag c, a character, has been set.
	    // Flag 返回标记 c（一个字符）是否已被设置。
	    Flag(c int) bool
	}
```
fmt 包中的 print.go 文件中的`type pp struct`实现了 State 接口。由于 State 接口有 Write 方法，因此，实现了 State 接口的类型必然实现了 io.Writer 接口。

## GoStringer 接口 ##

GoStringer 接口定义如下；
```go
	type GoStringer interface {
	    GoString() string
	}
```
该接口定义了类型的Go语法格式。用于打印(Printf)格式化占位符为 %#v 的值。

用前面的例子演示。执行：
```go
	p := &Person{"polaris", 28, 0}
	fmt.Printf("%#v", p)
```
输出：
```bash
	&main.Person{Name:"polaris", Age:28, Sex:0}
```
接着为Person增加方法：
```go
	func (this *Person) GoString() string {
		return "&Person{Name is "+this.Name+", Age is "+strconv.Itoa(this.Age)+", Sex is "+strconv.Itoa(this.Sex)+"}"
	}
```
这个时候再执行
```go
	p := &Person{"polaris", 28, 0}
	fmt.Printf("%#v", p)
```
输出：
```bash
	&Person{Name is polaris, Age is 28, Sex is 0}
```
一般的，我们不需要实现该接口。

## Scan 序列函数 ##

该序列函数和 Print 序列函数相对应，包括：Fscan/Fscanf/Fscanln/Sscan/Sscanf/Sscanln/Scan/Scanf/Scanln。

一般的，我们将Fscan/Fscanf/Fscanln归为一类；Sscan/Sscanf/Sscanln归为一类；Scan/Scanf/Scanln归为另一类。其中，Scan/Scanf/Scanln会调用相应的F开头一类函数。如：
```go
	func Scan(a ...interface{}) (n int, err error) {
		return Fscan(os.Stdin, a...)
	}
```
Fscan/Fscanf/Fscanln 函数的第一个参数接收一个 io.Reader 类型，从其读取内容并赋值给相应的实参。而 Scan/Scanf/Scanln 正是从标准输入获取内容，因此，直接调用 F类函数 做这件事，并将 os.Stdin 作为第一个参数传入。

Sscan/Sscanf/Sscanln 则直接从字符串中获取内容。

对于Scan/Scanf/Scanln三个函数的区别，我们通过例子来说明，为了方便讲解，我们使用Sscan/Sscanf/Sscanln这组函数。

1) Scan/FScan/Sscan
```go
	var (
		name string
		age  int
	)
	n, _ := fmt.Sscan("polaris 28", &name, &age)
	// 可以将"polaris 28"中的空格换成"\n"试试
	// n, _ := fmt.Sscan("polaris\n28", &name, &age)
	fmt.Println(n, name, age)
```
输出为：
```bash
	2 polaris 28
```
不管"polaris 28"是用空格分隔还是"\n"分隔，输出一样。也就是说，`Scan/FScan/Sscan` 这组函数将连续由空格分隔的值存储为连续的实参（换行符也记为空格）。

2) Scanf/FScanf/Sscanf
```go
	var (
		name string
		age  int
	)
	n, _ := fmt.Sscanf("polaris 28", "%s%d", &name, &age)
	// 可以将"polaris 28"中的空格换成"\n"试试
	// n, _ := fmt.Sscanf("polaris\n28", "%s%d", &name, &age)
	fmt.Println(n, name, age)
```
输出：
```bash
	2 polaris 28
```
如果将"空格"分隔改为"\n"分隔，则输出为：1 polaris 0。可见，`Scanf/FScanf/Sscanf` 这组函数将连续由空格分隔的值存储为连续的实参， 其格式由 `format` 决定，换行符处停止扫描(Scan)。

3) Scanln/FScanln/Sscanln
```go
	var (
		name string
		age  int
	)
	n, _ := fmt.Sscanln("polaris 28", &name, &age)
	// 可以将"polaris 28"中的空格换成"\n"试试
	// n, _ := fmt.Sscanln("polaris\n28", &name, &age)
	fmt.Println(n, name, age)
```
输出：
```bash	
	2 polaris 28
```
`Scanln/FScanln/Sscanln`表现和上一组一样，遇到"\n"停止（对于Scanln，表示从标准输入获取内容，最后需要回车）。

一般地，我们使用 `Scan/Scanf/Scanln` 这组函数。

**提示**

如果你是Windows系统，在使用 `Scanf` 时，有一个地方需要注意。看下面的代码：
```go
	for i := 0; i < 2; i++ {
		var name string
		fmt.Print("Input Name:")
		n, err := fmt.Scanf("%s", &name)
		fmt.Println(n, err, name)
	}
```
编译、运行（或直接 go run )，输入：polaris 回车。控制台内如下：
```bash
	Input Name:polaris
	1 <nil> polaris
	Input Name:0 unexpected newline
```
为什么不是让输入两次？第二次好像有默认值一样。

同样的代码在Linux下正常。个人认为这是go在Windows下的一个bug，已经向官方提出：[issue5391](https://code.google.com/p/go/issues/detail?id=5391)。

目前的解决方法是：换用Scanln或者改为Scanf("%s\n", &name)。

## Scanner 和 ScanState 接口 ##

基本上，我们不会去自己实现这两个接口，只需要使用上文中相应的 Scan 函数就可以了。这里只是简单的介绍一下这两个接口的作用。

任何实现了 Scan 方法的对象都实现了 Scanner 接口，Scan 方法会从输入读取数据并将处理结果存入接收端，接收端必须是有效的指针。Scan 方法会被任何 Scan、Scanf、Scanln 等函数调用，只要对应的参数实现了该方法。Scan 方法接收的第一个参数为`ScanState`接口类型。

ScanState 是一个交给用户定制的 Scanner 接口的参数的接口。Scanner 接口可能会进行一次一个字符的扫描或者要求 ScanState 去探测下一个空白分隔的 token。该接口的方法基本上在 io 包中都有讲解，这里不赘述。

在fmt包中，scan.go 文件中的 ss 结构实现了 ScanState 接口。
## fmt/print.go 阅读

### Fprint
```go
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	p := newPrinter()           // 实际工作结构
	p.doPrint(a)
	n, err = w.Write(p.buf)
	p.free()
	return
}
```
### newPrinter
```go
// printer 状态结构
type pp struct {
	buf buffer

	arg interface{}

	value reflect.Value

	fmt fmt

	reordered bool
	
	goodArgNum bool
	
	panicking bool
	
	erroring bool
}

// 通过 sync.Pool 复用，避免回收造成 GC
var ppFree = sync.Pool{
	New: func() interface{} { return new(pp) },
}

// 分配或重用 pp 结构
func newPrinter() *pp {
	p := ppFree.Get().(*pp)
	p.panicking = false
	p.erroring = false
	p.fmt.init(&p.buf)
	return p
}
```

### doPrint
```go
func (p *pp) doPrint(a []interface{}) {
	prevString := false
	
	// 获取可变参数索引及参数
	for argNum, arg := range a {
	    // reflect.TypeOf.Kind
		isString := arg != nil && reflect.TypeOf(arg).Kind() == reflect.String
		
	    // 判断是否需要一个空格
		if argNum > 0 && !isString && !prevString {
			p.buf.WriteByte(' ')
		}
		p.printArg(arg, 'v')
		prevString = isString
	}
}
```

### printArg
```go
func (p *pp) printArg(arg interface{}, verb rune) {
	p.arg = arg
	p.value = reflect.Value{}

	if arg == nil {
		switch verb {
		case 'T', 'v':
			p.fmt.padString(nilAngleString)
		default:
			p.badVerb(verb)
		}
		return
	}

	switch verb {
	case 'T':
		p.fmt.fmt_s(reflect.TypeOf(arg).String())
		return
	case 'p':
		p.fmtPointer(reflect.ValueOf(arg), 'p')
		return
	}

	// 类型判断
	switch f := arg.(type) {
	case bool:
		p.fmtBool(f, verb)
	case float32:
		p.fmtFloat(float64(f), 32, verb)
	case float64:
		p.fmtFloat(f, 64, verb)
	case complex64:
		p.fmtComplex(complex128(f), 64, verb)
	case complex128:
		p.fmtComplex(f, 128, verb)
	case int:
		p.fmtInteger(uint64(f), signed, verb)
	case int8:
		p.fmtInteger(uint64(f), signed, verb)
	case int16:
		p.fmtInteger(uint64(f), signed, verb)
	case int32:
		p.fmtInteger(uint64(f), signed, verb)
	case int64:
		p.fmtInteger(uint64(f), signed, verb)
	case uint:
		p.fmtInteger(uint64(f), unsigned, verb)
	case uint8:
		p.fmtInteger(uint64(f), unsigned, verb)
	case uint16:
		p.fmtInteger(uint64(f), unsigned, verb)
	case uint32:
		p.fmtInteger(uint64(f), unsigned, verb)
	case uint64:
		p.fmtInteger(f, unsigned, verb)
	case uintptr:
		p.fmtInteger(uint64(f), unsigned, verb)
	case string:
		p.fmtString(f, verb)
	case []byte:
		p.fmtBytes(f, verb, "[]byte")
	case reflect.Value:
		// Handle extractable values with special methods
		// since printValue does not handle them at depth 0.
		if f.IsValid() && f.CanInterface() {
			p.arg = f.Interface()
			if p.handleMethods(verb) {
				return
			}
		}
		p.printValue(f, verb, 0)
	default:
		// If the type is not simple, it might have methods.
		if !p.handleMethods(verb) {
			// Need to use reflection, since the type had no
			// interface methods that could be used for formatting.
			p.printValue(reflect.ValueOf(f), verb, 0)
		}
	}
}
```

# 导航 #

- [目录](/preface.md)
- 上一节：[ioutil — 方便的IO操作函数集](01.2.md)
- 下一节：[bufio — 缓存IO](01.4.md)