package sort


func Sort(Array []int)[]int {
    if(len(Array) <= 0) {
        return Array
    }
    i := 0
    j := len(Array)-1
    target := Array[i]
    
        
        
    for i < j {
        for(i<j && Array[j] >= target)  {
            j--
        }
        Array[i] = Array[j];
        for(i<j && Array[i] <= target)  {
            i++
        }
        Array[j] = Array[i]
        Array[i] = target
        
    }
    
    Sort(Array[:i])
    Sort(Array[i+1:])

    return Array
}
