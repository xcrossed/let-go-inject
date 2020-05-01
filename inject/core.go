package inject

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/xcrossed/let-go-inject/utils"
)

var NotExistBeanError = errors.New("Not exeis bean")
var AutoWireFinishedError = errors.New("autowire is finish,can not autowire")

type RegisteStatus uint

const (
	Initialize     RegisteStatus = iota //initialize status
	Registeing                          // regeisteing
	InjectFinished                      // Inject finished
)

// Bean
type Bean struct {
	UniqueName   string // bean global unique name
	BeanType     reflect.Type
	BeanValue    reflect.Value
	BeanInstance interface{} // instance
	Alias        string      //bean alias
}

//BeanFactory inteface
type BeanFactory interface {
	RegisterBean(instance interface{})
	CanAutoWire() bool
	RegisterBeanWithName(aliasName string, instance interface{})
	GetBeanByName(beanName string) (*Bean, error)
	AutoWire() error
}

type DefaultBeanFactory struct {
	beanMap       map[string]*Bean
	beanAliasMap  map[string]string
	registeStatus RegisteStatus
	mutx          sync.Mutex
}

var _ BeanFactory = &DefaultBeanFactory{}

// NewDefaultBeanFactory init method
func NewDefaultBeanFactory() *DefaultBeanFactory {
	return &DefaultBeanFactory{
		beanMap:      make(map[string]*Bean),
		beanAliasMap: make(map[string]string),
	}
}

// RegisterBean register a bean to factory
func (defaultBeanFactory *DefaultBeanFactory) RegisterBean(instance interface{}) {
	bean := defaultBeanFactory.createBean("", instance)
	defaultBeanFactory.mutx.Lock()
	defer defaultBeanFactory.mutx.Unlock()
	defaultBeanFactory.addToFactory(bean)
}

//
func (defaultBeanFactory *DefaultBeanFactory) createBean(aliasName string, instance interface{}) *Bean {
	if !utils.CanRegeiste(instance) {
		panic(fmt.Sprintf("%#v is not a ptr", instance))
	}

	bean := &Bean{}
	bean.BeanType = reflect.TypeOf(instance)
	bean.BeanValue = reflect.ValueOf(instance)
	bean.BeanInstance = instance
	if aliasName != "" {
		bean.Alias = aliasName
	}
	bean.UniqueName = utils.GetFullUniqueName(instance)
	return bean
}

// CanAutoWire check can autowire
func (defaultBeanFactory *DefaultBeanFactory) CanAutoWire() bool {
	if defaultBeanFactory.registeStatus == Initialize {
		return true
	}
	return false
}

//  RegisterBean register
func (defaultBeanFactory *DefaultBeanFactory) RegisterBeanWithName(aliasName string, instance interface{}) {
	bean := defaultBeanFactory.createBean(aliasName, instance)
	t := reflect.TypeOf(instance)
	if t.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("inject struct must be ptr.%v", instance))
	}
	defaultBeanFactory.addToFactory(bean)
}
func (defaultBeanFactory *DefaultBeanFactory) addToFactory(bean *Bean) {
	if _, ok := defaultBeanFactory.beanMap[bean.UniqueName]; ok {
		panic(fmt.Sprintf("can not repeat registe bean,alias:%v,instance:%#v", bean.Alias, bean.BeanInstance))
	}
	defaultBeanFactory.beanMap[bean.UniqueName] = bean
	aliasName := bean.Alias
	if aliasName != "" {
		if _, ok := defaultBeanFactory.beanAliasMap[aliasName]; !ok {
			defaultBeanFactory.beanAliasMap[aliasName] = bean.UniqueName
		}
	}
}

// GetBeanByName get bean by name or alias name
func (defaultBeanFactory *DefaultBeanFactory) GetBeanByName(beanName string) (*Bean, error) {
	// alias bean name
	if trueBeanName, ok := defaultBeanFactory.beanAliasMap[beanName]; ok {
		if _, ok := defaultBeanFactory.beanMap[trueBeanName]; ok {
			return defaultBeanFactory.beanMap[trueBeanName], nil
		}
	}
	if _, ok := defaultBeanFactory.beanMap[beanName]; ok {
		return defaultBeanFactory.beanMap[beanName], nil
	}
	return nil, NotExistBeanError
}

// AutoWire finish to inject
func (defaultBeanFactory *DefaultBeanFactory) AutoWire() error {
	defaultBeanFactory.mutx.Lock()
	defer defaultBeanFactory.mutx.Unlock()
	if !defaultBeanFactory.CanAutoWire() {
		return AutoWireFinishedError
	}

	defaultBeanFactory.registeStatus = Registeing

	for _, val := range defaultBeanFactory.beanMap {
		elemType := val.BeanType.Elem()
		eleVal := val.BeanValue.Elem()
		for i := 0; i < elemType.NumField(); i++ {
			fieldType := elemType.Field(i)
			vType := eleVal.Field(i)

			tag := fieldType.Tag
			if !utils.FieldNeedToInject(fieldType) {
				continue
			}

			alias := tag.Get(utils.InjectTagKey)
			var bean *Bean
			var fullUniqueName string
			var ok bool
			if alias != "" {
				alias = defaultBeanFactory.getInjectName(alias)
				fullUniqueName, ok = defaultBeanFactory.beanAliasMap[alias]
				if !ok {
					fullUniqueName = vType.Type().String()
				}
			} else {
				fullUniqueName = vType.Type().String()
			}
			bean, ok = defaultBeanFactory.beanMap[fullUniqueName]
			if !ok {
				panic(fmt.Sprintf("the bean is not exists:%v", fullUniqueName))
			}
			vType.Set(bean.BeanValue)
		}
	}
	defaultBeanFactory.registeStatus = InjectFinished
	return nil
}

func (defaultBeanFactory *DefaultBeanFactory) getInjectName(tag string) string {
	tags := strings.Split(tag, ",")
	if len(tags) == 0 {
		return ""
	}
	return tags[0]
}

func (defaultBeanFactory *DefaultBeanFactory) string() {
	for k, v := range defaultBeanFactory.beanMap {
		fmt.Printf("k:%#v,v:%#v \n", k, v)
	}
}
