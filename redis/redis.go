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

    //getData := client.HMGet(APIName, Formvalues)
    
   // if getData.Err() != nil{
     //   fmt.Println(getData.Err().Error())
    //}
    
    
}


func GetComponentURI(APIName string) string{
    
    client := utils.NewRedisClient()
    resultVar := client.HMGet(APIName,"ComponentURI").Val()
    return resultVar[0].(string)
}

