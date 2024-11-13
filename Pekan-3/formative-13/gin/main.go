package main

import (
	"github.com/kandlagifari/sb-go-batch61-kandla/Pekan-3/formative-13/gin/routers"
)

func main() {
	var PORT = ":30069"

	routers.StartServer().Run(PORT)
}
