# let-go-inject

It is a simple dependency injection library for golang.
Using it, you can make your code very simple and easy to maintain.
Compared with other libraries, this is simpler

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

	factoryBean := NewDefaultBeanFactory()
	dao := &Dao{"my name is dao."}
	factoryBean. RegisterBean(dao)

	biz := &biz{}
	factoryBean. RegisterBeanWithName("biz.impl", biz)
	ctl := &Controll{}
	factoryBean. RegisterBean(ctl)
	factoryBean. AutoWire()

	assert. NotNil(t, biz. Dao)
	assert. NotNil(t, ctl. Biz)
	assert. Equal(t, dao, biz. Dao)
	assert. Equal(t, ctl. Biz, biz)

}
```

## Other inject libraries

* [facebook inject](https://github.com/facebookarchive/inject)
* [google wire](https://github.com/google/wire)

