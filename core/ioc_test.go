package ioc

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestDefaultBeanFactory_AutoWire(t *testing.T) {
	factoryBean := NewDefaultBeanFactory()
	dao := &Dao{"my name is dao."}
	factoryBean.RegisterBean(dao)

	biz := &biz{}
	factoryBean.RegisterBeanWithName("biz.impl", biz)
	ctl := &Controll{}
	factoryBean.RegisterBean(ctl)
	factoryBean.AutoWire()

	assert.NotNil(t, biz.Dao)
	assert.NotNil(t, ctl.Biz)
	assert.Equal(t, dao, biz.Dao)
	assert.Equal(t, ctl.Biz, biz)
}

func TestDefaultFactoryBean(t *testing.T) {
	dao := Dao{}
	_type := reflect.TypeOf(dao)
	daoPtr := &Dao{}
	_typePtr := reflect.TypeOf(daoPtr)
	fmt.Printf("#%v ,kind:%v\n", _type, _type.Kind())
	fmt.Printf("#%v ,kind:%v\n", _typePtr, _typePtr.Kind())

}
