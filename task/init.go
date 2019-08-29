package task

import (
	"log"
)

func init() {
	log.Println("++++++++++++++++++ task init ++++++++++++++++++++++")
	go Task1()
}
