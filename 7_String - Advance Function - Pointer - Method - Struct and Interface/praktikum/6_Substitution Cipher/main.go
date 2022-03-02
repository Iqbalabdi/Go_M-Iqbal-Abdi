package main

import "fmt"

type student struct {
	name, nameEncode string
	score int
}

type Chiper interface {
   Encode() string
   Decode() string
}

func (s *student) Encode() string {
	var temp []byte
	for i:=0; i<len(s.name); i++ {
		var convert byte = (s.name[i] + 13 - 97)%26 + 97
		temp = append(temp, convert)
	}
	s.nameEncode = string(temp)
	return s.nameEncode
}

func (s *student) Decode() string {
	var temp []byte
	for i:=0; i<len(s.nameEncode); i++ {
		var convert byte = (s.nameEncode[i] + 13 - 97)%26 + 97
		temp = append(temp, convert)
	}
	s.name = string(temp)
	return s.name
}

func main() {
   var menu int
   var a = student{}
   var c Chiper = &a
   fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
   fmt.Scan(&menu)
   if menu == 1 {
       fmt.Print("\nInput Student's Name : ")
       fmt.Scan(&a.name)
       fmt.Print("\nEncode of Student’s Name " + a.name + " is : " + c.Encode())
   } else if menu == 2 {
       fmt.Print("\nInput Student's Encode Name : ")
       fmt.Scan(&a.nameEncode)
       fmt.Print("\nDecode of Student’s Name " + a.nameEncode + " is : " + c.Decode())

   } else {
       fmt.Println("Wrong input name menu")
   }
}