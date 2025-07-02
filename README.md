PHP [serialize()](http://php.net/manual/en/function.serialize.php) and
[unserialize()](http://php.net/manual/en/function.unserialize.php) for Go.

This is a fork of github.com/elliotchance/phpserialize with all instances of map[any]any replaced with orderedmap.OrderedMap[any, any].

Since Associated Arrays in PHP are ordered, it is sometimes required to preserve that order when consuming such array.

# Install / Update

```bash
go get -u github.com/jamteacoffee/phpserialize
```

`phpserialize` requires Go 1.24+.

# Example

### Using struct field tags for marshalling

```go
package main

import (
	"fmt"
	"github.com/elliotchance/orderedmap/v3"
	"github.com/jamteacoffee/phpserialize"
)

func main() {
	input := []byte("a:2:{s:3:\"bar\";i:20;s:3:\"foo\";i:10;}")
	result := orderedmap.NewOrderedMap[any, any]()
	err := phpserialize.Unmarshal(input, &result)

	...

}
```
