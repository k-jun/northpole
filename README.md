# northpole

simple game matching library in golang

## usage

```golang
package northpole

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/k-jun/northpole/room"
	"github.com/k-jun/northpole/user"
)

func main() {
	r := room.New(uuid.New().String(), 3, func(_ string) error { return nil })
	u1 := user.New(uuid.New().String())
	u2 := user.New(uuid.New().String())
	u3 := user.New(uuid.New().String())
	m := New()
	channel, err := m.CreateRoom(u1, r)
	if err != nil {
		fmt.Println(err)
	}
	channel, err = m.JoinRoom(u2, r)
	channel, err = m.JoinRoom(u3, r)
	fmt.Println(r.IsOpen()) // false

	r = <-channel
	fmt.Println(r) // latest state of the room
}
```

check examples directory for more information


## license
MIT
