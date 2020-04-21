# let-go-inject

It is a simple dependency injection library for golang.
Using it, you can make your code very simple and easy to maintain.
Compared with other libraries, this is simpler

## install

go get https://github.com/xcrossed/let-go-inject

## feature

* support struct inject
* support interface inject
* support method inject(Coming soon)

## tag support

* inject tag

``` go

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

	fmt. Printf("say:%v\n", name)

}

type Controll struct {

	Biz BizInterface `inject:"biz.impl"` 

}

func TestDefaultBeanFactory_AutoWire(t *testing. T) {

	beanfactory := NewDefaultBeanFactory()
	dao := &Dao{"my name is dao."}
	beanfactory. RegisterBean(dao)

	biz := &biz{}
	beanfactory. RegisterBeanWithName("biz.impl", biz)
	ctl := &Controll{}
	beanfactory. RegisterBean(ctl)
	beanfactory. AutoWire()

	assert. NotNil(t, biz. Dao)
	assert. NotNil(t, ctl. Biz)
	assert. Equal(t, dao, biz. Dao)
	assert. Equal(t, biz, ctl. Biz)

}
```

## Other inject libraries

* [go spring](https://github.com/go-spring)
* [facebook inject](https://github.com/facebookarchive/inject)
* [google wire](https://github.com/google/wire)

