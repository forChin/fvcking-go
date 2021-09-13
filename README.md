# fvcking-go
In this repository I will add Golang's wtf moments as I study it.

## Some good ones
### Declaration of variable `nil`
```
  nil := 12
```
### Call main recursively
```
func main() {
  main()
}
```
### Without package prefix
You can use exported function of package without package prefix if specify alias for imported package as dot.
```
import (
	. "fmt"
)

func main() {
	Println("Hello World!")
}
```
## Support
Share with me wtf Golang moments you found.
