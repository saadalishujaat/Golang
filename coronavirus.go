package main
import "fmt"

func main() {

   var check1,check2,check3,check4=0,0,0,0
   var mess string = ""

   for (check4==0){

     fmt.Println("Please select an option:")
     fmt.Println("1 – Print Covid19 cases in Pakistan ")
     fmt.Println("2 – Print Covid19 cases in SouthKorea ")
     fmt.Println("3 – Print Covid19 cases in France ")
     fmt.Println("4 – Print my personalized message to Coronavirus, optional")
     fmt.Println("0 – Exit")

     var option int
     fmt.Scan(&option)

       if (option == 1)  {
           fmt.Printf("1500 cases are in pakistan\n\n")
           check1=1
       }
       if (option==2) {
          fmt.Printf("9500 cases are in southkorea\n\n")
          check2=1
       }
       if (option==3) {
          fmt.Printf("37000 cases are in southkorea\n\n")
          check3=1
       }
       if (option==4) {
          fmt.Printf("Write the message to coronavirus : ")

          fmt.Scan(&mess)
       }
       if (option == 0) {
         if(check1==1){
           if(check2==1){
             if(check3==1){
               fmt.Println("Terminated")
                fmt.Printf("Your message to coronavirus is : ")
                 fmt.Println(mess)

                check4=1
                }else{fmt.Println("please see option 3 first")}
              }else{fmt.Println("please see option 2 first")}
            }else{fmt.Println("please see option 1 first")}
        }
    }
}
