package util


import(
	"fmt"
	"os"
)

var (
	current_env_is string 
)

func init() {
	current_env_is = os.Getenv("is_dev")
}


func Print(values ...any) {
	if current_env_is == "prod" {
		return
	} 
	for _, val_of := range values {
		fmt.Println(val_of)
	}
}