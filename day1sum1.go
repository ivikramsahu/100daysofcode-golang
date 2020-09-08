package main

import ("fmt";"os";"strconv")

func main(){

    if (len(os.Args) < 5){
        fmt.Printf("Please enter in below format\nExample : 10 3 | -\n")
        os.Exit(1)
    }

    num1,_ := strconv.Atoi(os.Args[1])
    num2,_ := strconv.Atoi(os.Args[2])
    for i:=0;i<=num1;i++{
        if(os.Args[3] == "n"){
            fmt.Print(i)
        }else{
            fmt.Print(os.Args[3])
        }
        if(i == num1){
            fmt.Print("\n")
            break
        }
        for j:=0;j<num2;j++{
            fmt.Print(os.Args[4])
        }
    }
}
