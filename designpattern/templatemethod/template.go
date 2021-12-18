package templatemethod
//四个步骤 准备下载 下载 保存 完成
/*
说实话这个没怎么看懂
因为父类需要调用子类方法，所以子类需要匿名组合父类的同时，父类需要持有子类的引用。
*/
import "fmt"
type Downloader interface {
	Download(uri string)
}
type implement interface {
	download()
	save()
}
//组合了 《模板实现父类 implement》
type template struct {
	implement
	uri string
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}
//一个模板的下载步骤 四个步骤，相当于java种的那个抽象类，第一步和第四步有默认实现，中间两步交给子类实现
//整个过程的中间两步 由template持有的《模板实现父类 implement》来实现。
func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *template) save() {
	fmt.Print("default save\n")
}
//组合了template其实就是指继承了template父类。同时也实现了downloader接口，
type HTTPDownloader struct {
	*template
}
func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

/**/
type FTPDownloader struct {
	*template
}
func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}