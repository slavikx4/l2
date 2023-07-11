Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
nil
false

nil != nil, так как у первого nil os.PathError реализует интерфейс error

интерфейс в го состоит из
-указателя на на тип данных, который реализует его
-указатели на копию переданого объекта

пустой интерфейс имеет nil и nil



```