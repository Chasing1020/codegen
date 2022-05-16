/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/5/7-18:40
File: sd.go
*/

package consul

import "fmt"

func Discovery() {
	passingOnly := true
	addrs, meta, err := Client.Health().Service("crud", "", passingOnly, nil)
	fmt.Println(addrs)
	fmt.Println(meta)
	fmt.Println(err)
	//if len(addrs) == 0 && err == nil {
	//	return nil, fmt.Errorf("service ( %s ) was not found", service)
	//}
	//if err != nil {
	//	return nil, err
	//}
	//return addrs, meta, nil
}
