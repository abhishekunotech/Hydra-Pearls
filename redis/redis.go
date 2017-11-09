package redis

import(
    "fmt"
    //"github.com/go-redis/redis"
    "github.com/Unotechsoftware/pearly-gates-go/utils"
    //"log"
    //"github.com/mediocregopher/radix.v2/redis"
)

func DBConnect(APIName string, FormValues map[string]interface{}) {
    
    client := utils.NewRedisClient()
    
    resultVar := client.HMSet(APIName, FormValues)
    
    if resultVar.Err() != nil {
        fmt.Println(resultVar.Err().Error())
    }    
}



