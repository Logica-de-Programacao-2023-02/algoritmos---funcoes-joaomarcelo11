func encontraIndice(slice []int, valor int) int {
    for i, v := range slice {
        if v == valor {
            return i
        }
    }
    return -1
}