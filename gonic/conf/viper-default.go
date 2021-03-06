package main

import (
	"fmt"

	"github.com/spf13/viper"
)


func overrideLoad(in string) {
    fmt.Println("=================")
    //viper.Set("env", "dy_new_value")
	viper.SetConfigName(in)
	viper.AddConfigPath("./")
    viper.SetDefault("env", "default_env")
	err := viper.ReadInConfig()
	if err != nil {
        fmt.Println(err)
	}
    keys := []string{"env", "app"}
    for _, key := range keys{
        fmt.Println(key+":",viper.GetString(key) )
    }

}


func main(){
    overrideLoad("conf")
    overrideLoad("conf_dev")
}
