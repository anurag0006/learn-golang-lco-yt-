package main

import "fmt"


func main()  {
	
	fmt.Println("hello")

	anurag := User{"Anurag","a@gmail.com",false,22}

	fmt.Println(anurag)

	anurag.GetStatus()
	anurag.NewMail()

	fmt.Println(anurag)

}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}


func (u User) GetStatus(){
   fmt.Println("Is user active?",u.Status)
}

func (u User) NewMail(){
	u.Email =  "test123@gmail.com"
	fmt.Println("Email of this :",u.Email)
}