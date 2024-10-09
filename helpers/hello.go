package helpers

import (
	"fmt"
	"os"
)

func SayHello(name string) {
	fmt.Println(fmt.Sprintf("Hello %s %s", name, GetEnvWithDefaultString("GREETING_SAY", "Selamat Pagi")))
}
