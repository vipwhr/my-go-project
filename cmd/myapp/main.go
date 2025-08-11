package main

import(

	"strings"
	"fmt"
	"github.com/vipwhr/my-go-project/internal/pkg1"

)

func  main()  {
	
	fmt.Println(strings.TrimPrefix("http://zjuqx.cn:9014/drawings-backend/daily-paintings-full/22fe315f35ff692ea67d37f160c30601.jpg", "/"))
	fmt.Println("My go project Started!!!!")
	pkg1.HelloWorld()
}