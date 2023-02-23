package main

import (
        "fmt"
)

type Matrix struct{
        rowlen int
        columnlen int
        list []int
}

func main() {
        var first,second,result Matrix
        Matrix_initialization(&first,2,4)
        Matrix_initialization(&second,4,3)
        result=matrix_multiplication(first,second)
        formate_print(result)
}


func Matrix_initialization (object *Matrix,row int,col int)  {
        var a int
        object.rowlen=row
        object.columnlen=col
        for i:=0;i<object.rowlen*object.columnlen;i++{
                fmt.Scan(&a)
                object.list=append(object.list,a)
        }
}



func formate_print (j Matrix){
        fmt.Println("the result  matrix :")
        for i:=0;i<j.columnlen;i++{
                for p:=0;p<j.rowlen;p++{
                        fmt.Print(j.list[i*j.rowlen+p]," ")
                }
                fmt.Println()
        }
}

func matrix_multiplication (first,second Matrix) Matrix{
        var result Matrix
        var tmp int
        result.rowlen=second.rowlen
        result.columnlen=first.columnlen
        for i:=0;i<first.columnlen;i++{
                for j:=0;j<second.rowlen;j++{
                        tmp=0
                        for k:=0;k<second.columnlen;k++{
                                tmp+=first.list[first.rowlen*i+k]*second.list[second.columnlen*k+j]
                        }
                        result.list=append(result.list,tmp)
                }
        }
        return result

}
