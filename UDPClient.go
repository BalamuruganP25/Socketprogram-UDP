package main

 import (

     
     "net"
     "fmt"
 )

 func main(){

 Buffer := make([]byte, 1024)
 var Buffferindex int32 = 0
 var Validation_Main_Token int32 = 526
 var Validation_sub_Token  int32 = 823
 var Login_user_username   string = "balamurugan.p@XXX.in"
 var Login_user_password   string = "bala12"



 Buffer[Buffferindex] = byte (Validation_Main_Token)
 Buffferindex ++
 Buffer[Buffferindex] = byte (Validation_Main_Token >> 8)
 Buffferindex ++
 Buffer[Buffferindex] = byte (Validation_Main_Token >> 16)
 Buffferindex ++
 Buffer[Buffferindex] = byte (Validation_Main_Token >> 24)
 Buffferindex ++

 Buffer[Buffferindex] = byte (Validation_sub_Token)
 Buffferindex ++
 Buffer[Buffferindex] = byte (Validation_sub_Token >> 8)
 Buffferindex ++
 Buffer[Buffferindex] = byte (Validation_sub_Token >> 16)
 Buffferindex ++
 Buffer[Buffferindex] = byte (Validation_sub_Token >> 24)
 Buffferindex ++

 Buffer[Buffferindex] = byte(len(Login_user_username))
 fmt.Println(" \t Login_user_username length ===>",len(Login_user_username))
 Buffferindex++

 usernamebuffer := []byte (Login_user_username)

 for i:=0; i< len(Login_user_username); i++ {

          Buffer[Buffferindex] = usernamebuffer[i]
          Buffferindex++

}

 Buffer[Buffferindex] = byte(len(Login_user_password))
  fmt.Println(" \t Login_user_password length ===>",len(Login_user_password))
 Buffferindex++

 passwordbuffer := []byte (Login_user_password)

 for j:=0; j< len(Login_user_password); j++ {

          Buffer[Buffferindex] = passwordbuffer[j]
          Buffferindex++

}


   	 hostName := "192.168.0.166"
	 portNum := "53552"

	 service := hostName + ":" + portNum

	 RemoteAddr, err := net.ResolveUDPAddr("udp", service)

	 //LocalAddr := nil
	 // see https://golang.org/pkg/net/#DialUDP

	 conn, err := net.DialUDP("udp", nil, RemoteAddr)

	 // note : you can use net.ResolveUDPAddr for LocalAddr as well
	 // for this tutorial simplicity sake, we will just use nil

	 if err != nil {
	    fmt.Println(err)
	 }

	 fmt.Printf(" \t Established connection to %s \n", service)
	 fmt.Printf(" \t Remote UDP address : %s \n", conn.RemoteAddr().String())
	 fmt.Printf(" \t Local UDP client address : %s \n", conn.LocalAddr().String())

	      Sendmessage := Buffer
         _, err = conn.Write(Sendmessage)

         if err != nil {
             fmt.Println(err)

         }else{

         	fmt.Println("\t >>> Message has been send <<< ")
         }

     defer conn.Close()

     






 }