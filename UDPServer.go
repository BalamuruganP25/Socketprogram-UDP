package main

 import (

     
     "net"
     "fmt"
     "log"
     "strings"
 )


 func handleUDPConnection(conn *net.UDPConn) {


 receviebuffer := make([]byte, 1024)
 n, addr, err := conn.ReadFromUDP(receviebuffer)
 fmt.Println("\t Buffer Data ===>",string(receviebuffer[:n]))
 fmt.Println(" \t UDP client : ", addr)
 if err != nil {
       
       log.Fatal(err)
    
    }else{

      maintokenvalidation(receviebuffer)
  }


 }


func Returnfourbytes(recbuffer [] byte,index int) int32{


  var temp int32 = 0

  var msgvalue int32 = 0

  msgvalue = int32 (recbuffer[index])
  index++

  temp = int32 (recbuffer[index])
  msgvalue |= (temp << 8)
  index++

  temp = int32 (recbuffer[index])
  msgvalue |= (temp << 16)
  index++

  temp = int32 (recbuffer[index])
  msgvalue |= (temp << 24)
  index++
 return msgvalue


}

 func maintokenvalidation(maintokenbuffer [] byte){

 rec_buffer_index := 0
 var getmaintoken int32 
 getmaintoken =  Returnfourbytes(maintokenbuffer,rec_buffer_index) 
 rec_buffer_index = rec_buffer_index +4
 fmt.Println("\t getmaintoken ==>",getmaintoken)
  switch getmaintoken {
    case 526:
        fmt.Println("\t Main token validation success")
         subtokenvalidation(maintokenbuffer,rec_buffer_index)
   
    default:
        fmt.Println("\t Main Token invalid")
    }



}


 func subtokenvalidation(subtokenbuffer [] byte,subbufferindex int){
 var getsubtoken int32
 getsubtoken  = Returnfourbytes(subtokenbuffer,subbufferindex)
 subbufferindex = subbufferindex +4
 fmt.Println("\t getsubtoken ==>",getsubtoken)

 switch getsubtoken {
    case 823:
        fmt.Println("\t sub token validation success")
         ProcessLoginMsg(subtokenbuffer,subbufferindex)
   
    default:
        fmt.Println("\t sub Token invalid")
    }


 }


func ProcessLoginMsg(loginbuffermsg [] byte,loginindex int){

	var usernamelength int
	var passwordlength int 
	var login_status string

	usernamelength = int (loginbuffermsg[loginindex])
	loginindex++
	var username string = string(loginbuffermsg[loginindex:(usernamelength+loginindex)])
    fmt.Println("\t username ==>",username)
    loginindex = loginindex + usernamelength

    passwordlength = int (loginbuffermsg[loginindex])
	loginindex++
	var password string = string(loginbuffermsg[loginindex:(passwordlength+loginindex)])
    fmt.Println("\t password ==>",password)

    if strings.TrimRight(username, "\n") == "balamurugan.p@XXX.in" && strings.TrimRight(username, "\n") == "bala12" {
  
  	login_status =  "success"
   }else{


   	login_status =  "falied"


   }

   fmt.Println("\t login_status ",login_status)




}



 func main(){

    addr := net.UDPAddr{
                   Port: 53552,
                   IP:   net.ParseIP("192.168.0.166"),
        }

         ln, err := net.ListenUDP("udp", &addr)

         if err != nil {
                 log.Fatal(err)
         }

         fmt.Println("UDP server up and listening on port 53552")

         defer ln.Close()

         for {
                 // wait for UDP client to connect
                 handleUDPConnection(ln)

         }





 }