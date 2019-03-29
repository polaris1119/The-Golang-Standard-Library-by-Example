# 1.2 ioutil — 方便的IO操作函数集 #

虽然 io 包提供了不少类型、方法和函数，但有时候使用起来不是那么方便。比如读取一个文件中的所有内容。为此，标准库中提供了一些常用、方便的IO操作函数。

说明：这些函数使用都相对简单，一般就不举例子了。

## NopCloser 函数 ##

有时候我们需要传递一个 io.ReadCloser 的实例，而我们现在有一个 io.Reader 的实例，比如：strings.Reader ，这个时候 NopCloser 就派上用场了。它包装一个io.Reader，返回一个 io.ReadCloser ，而相应的 Close 方法啥也不做，只是返回 nil。

比如，在标准库 net/http 包中的 NewRequest，接收一个 io.Reader 的 body，而实际上，Request 的 Body 的类型是 io.ReadCloser，因此，代码内部进行了判断，如果传递的 io.Reader 也实现了 io.ReadCloser 接口，则转换，否则通过ioutil.NopCloser 包装转换一下。相关代码如下：
```go
    rc, ok := body.(io.ReadCloser)
    if !ok && body != nil {
		rc = ioutil.NopCloser(body)
    }
```
如果没有这个函数，我们得自己实现一个。当然，实现起来很简单，读者可以看看 [NopCloser](http://docscn.studygolang.com/src/io/ioutil/ioutil.go?s=5557:5598#L145) 的实现。

## ReadAll 函数 ##

很多时候，我们需要一次性读取 io.Reader 中的数据，通过上一节的讲解，我们知道有很多种实现方式。考虑到读取所有数据的需求比较多，Go 提供了 ReadAll 这个函数，用来从io.Reader 中一次读取所有数据。
```go
	func ReadAll(r io.Reader) ([]byte, error)
```
阅读该函数的源码发现，它是通过 bytes.Buffer 中的 [ReadFrom](http://docscn.studygolang.com/src/bytes/buffer.go?s=5385:5444#L144) 来实现读取所有数据的。该函数成功调用后会返回 err == nil 而不是 err == EOF。(成功读取完毕应该为 err == io.EOF，这里返回 nil 由于该函数成功期望 err == io.EOF，符合无错误不处理的理念)

## ReadDir 函数 ##

笔试题：编写程序输出某目录下的所有文件（包括子目录）

是否见过这样的笔试题？

在 Go 中如何输出目录下的所有文件呢？首先，我们会想到查 os 包，看 File 类型是否提供了相关方法（关于 os 包，后面会讲解）。

其实在 ioutil 中提供了一个方便的函数：ReadDir，它读取目录并返回排好序的文件和子目录名（ []os.FileInfo ）。通过这个方法，我们可以很容易的实现“面试题”。

```go	
func main() {
	dir := os.Args[1]
	listAll(dir,0)
}

func listAll(path string, curHier int){
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil{fmt.Println(err); return}

	for _, info := range fileInfos{
		if info.IsDir(){
			for tmpHier := curHier; tmpHier > 0; tmpHier--{
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(),"\\")
			listAll(path + "/" + info.Name(),curHier + 1)
		}else{
			for tmpHier := curHier; tmpHier > 0; tmpHier--{
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}
```

## ReadFile 和 WriteFile 函数 ##

ReadFile 读取整个文件的内容，在上一节我们自己实现了一个函数读取文件整个内容，由于这种需求很常见，因此 Go 提供了 ReadFile 函数，方便使用。ReadFile 的实现和ReadAll 类似，不过，ReadFile 会先判断文件的大小，给 bytes.Buffer 一个预定义容量，避免额外分配内存。

ReadFile 函数的签名如下:
```go
	func ReadFile(filename string) ([]byte, error)
```
函数文档：

> ReadFile 从 filename 指定的文件中读取数据并返回文件的内容。成功的调用返回的err 为 nil 而非 EOF。因为本函数定义为读取整个文件，它不会将读取返回的 EOF 视为应报告的错误。(同 ReadAll )

WriteFile 函数的签名如下：
```go 
	func WriteFile(filename string, data []byte, perm os.FileMode) error
```
函数文档:

> WriteFile 将data写入filename文件中，当文件不存在时会根据perm指定的权限进行创建一个,文件存在时会先清空文件内容。对于 perm 参数，我们一般可以指定为：0666，具体含义 os 包中讲解。

**小提示**

ReadFile 源码中先获取了文件的大小，当大小 < 1e9 时，才会用到文件的大小。按源码中注释的说法是 FileInfo 不会很精确地得到文件大小。

## TempDir 和 TempFile 函数 ##

操作系统中一般都会提供临时目录，比如 linux 下的 /tmp 目录（通过 os.TempDir() 可以获取到)。有时候，我们自己需要创建临时目录，比如 Go 工具链源码中（src/cmd/go/build.go），通过 TempDir 创建一个临时目录，用于存放编译过程的临时文件：
```go
	b.work, err = ioutil.TempDir("", "go-build")
```
第一个参数如果为空，表明在系统默认的临时目录（ os.TempDir ）中创建临时目录；第二个参数指定临时目录名的前缀，该函数返回临时目录的路径。

相应的，TempFile 用于创建临时文件。如 gofmt 命令的源码中创建临时文件：
```go
	f1, err := ioutil.TempFile("", "gofmt")
```
参数和 ioutil.TempDir 参数含义类似。

这里需要**注意**：创建者创建的临时文件和临时目录要负责删除这些临时目录和文件。如删除临时文件：
```go
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()
```
## Discard 变量 ##

Discard 对应的类型（`type devNull int`）实现了 io.Writer 接口，同时，为了优化 io.Copy 到 Discard，避免不必要的工作，实现了 io.ReaderFrom 接口。

devNull 在实现 io.Writer 接口时，只是简单的返回（标准库文件：[src/pkg/io/ioutil.go](http://docscn.studygolang.com/pkg/io/ioutil/#pkg-variables))。
```go
	func (devNull) Write(p []byte) (int, error) {
		return len(p), nil
	}
```
而 ReadFrom 的实现是读取内容到一个 buf 中，最大也就 8192 字节，其他的会丢弃（当然，这个也不会读取）。

# 导航 #

- [目录](/preface.md)
- 上一节：[io — 基本的IO接口](01.1.md)
- 下一节：[fmt — 格式化IO](01.3.md)
