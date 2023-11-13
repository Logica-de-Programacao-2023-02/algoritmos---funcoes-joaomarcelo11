func concatenarStrings(slice []string) string {
    var resultado string
    for _, str := range slice {
        resultado += str
    }
    return resultado
}