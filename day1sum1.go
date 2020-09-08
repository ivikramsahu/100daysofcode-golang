/* Problem Statement 1

sample input 1 :
10 4 | *

output :
|****|****|****|****|****|****|****|****|****|****|

Sample input 2 :
10 4 n -

output :
0----1----2----3----4----5----6----7----8----9----10

Note : When the user enters the first major mark as n then print the numeric number instead of mark..
*/

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
