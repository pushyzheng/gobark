# gobark

## Quick start

```go
import (
	"fmt"
	bark "github.com/pushyzheng/gobark"
)

func main() {
	// 1. create bark client
	client := bark.BarkClient{Domain: bark.DefaultDomain}
    
	// 2. build options
	options := bark.NewPushOptions("Hello", "To be or not to be")
	options.SetReceiver("your key")
	options.SetLevel(bark.Passive)
	
	// 3. execute push
	failed, err := client.Push(options)
	if err != nil {
		panic(err)
	}
	fmt.Println("failed:", failed)
}
```