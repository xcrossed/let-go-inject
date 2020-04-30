# let-go-inject

It is a simple dependency injection library for golang. 
Using it, you can make your code very simple and easy to maintain. 
Compared with other libraries, this is simpler

## install

no go mod support , you can go get. 

``` 
go get -u -v  github.com/xcrossed/let-go-inject
```

go mdo support, you can use it direct. 

## feature

* support struct inject
* support interface inject
* support method inject(Coming soon)

## tag support

* inject tag

## demo code

``` go
package main

import (
	"fmt"

	"github.com/xcrossed/let-go-inject/inject"
)

type Dao struct {
	Name string
}

type BizInterface interface {
	SayHello(string)
}

type biz struct {
	Dao *Dao `inject:""` 
}

func (biz *biz) SayHello(name string) {

	fmt.Printf("say:%v\n", name)

}

type Controll struct {
	Biz BizInterface `inject:"biz.impl"` 
}

func main() {
	beanfactory := inject.NewDefaultBeanFactory()
	dao := &Dao{"my name is dao."}
	beanfactory.RegisterBean(dao)

	biz := &biz{}
	beanfactory.RegisterBeanWithName("biz.impl", biz)
	ctl := &Controll{}
	beanfactory.RegisterBean(ctl)
	beanfactory.AutoWire()

	fmt.Println(biz.Dao == nil)
	fmt.Println(ctl.Biz == nil)

	fmt.Printf("equals:%v\n", dao == biz.Dao)
	fmt.Printf("equals:%v\n", biz == ctl.Biz)
}
```

## Other inject libraries

* [facebook inject](https://github.com/facebookarchive/inject)
* [google wire](https://github.com/google/wire)
* [go spring](https://github.com/go-spring)


