package Sumarray

func Sum(arr []int) int {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}

func SumAll(input [][]int)[]int{
    sumArray := make([]int,0,len(input))
    for _,sub := range input{
        tempSum := 0
        for _, j:=range sub{
            tempSum += j
        }
        
        sumArray = append(sumArray, tempSum)
        
    }
    return sumArray
}