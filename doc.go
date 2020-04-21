/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package letGoInject

/*
// demo code
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
*/
