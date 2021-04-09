package main

import "fmt"

func main(){
	A, B, K := 6,11,2
	fmt.Println(Countdiv(A,B,K))
}

func Countdiv(A, B, K int) int{
	remainderA := A % K;
    
    if remainderA!= 0 {
        A = A + (K - remainderA);
    }

    if (A > B) {
        return 0;
    }
    
    return ((B - A) / K) + 1;
}