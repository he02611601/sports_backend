package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//// connect to firebase
	//opt := option.WithCredentialsFile("./sport-backend-firebase-adminsdk-8282e-58cf8bb531.json")
	//app, err := firebase.NewApp(context.Background(), nil, opt)
	//if err != nil {
	//	log.Fatalf("error initializing app: %v\n", err)
	//}
	//
	//client, err := app.Firestore(context.Background())
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//// testData struct
	//type dataStruct struct {
	//	TestString  string `json:"testString"`
	//	TestInteger int64  `json:"testInteger"`
	//}
	//
	//// testData variable and init
	//var testData dataStruct = dataStruct{
	//	TestString:  "testtesttestString",
	//	TestInteger: 3,
	//}
	//log.Println(testData)
	//
	//// insert to firebase
	//result, err := client.Collection("record").Doc("Paul").Set(context.Background(), testData)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(result)
	//
	//// select from firebase
	//dsnap, err := client.Collection("record").Doc("Paul").Get(context.Background())
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//m := dsnap.Data()
	//log.Println(m)
	//
	//// close connect
	//defer client.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
