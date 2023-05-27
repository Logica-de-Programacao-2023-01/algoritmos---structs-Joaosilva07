type Triangulo struct {
    base   float64
    altura float64
}

func calcularAreaTriangulo(triangulo Triangulo) float64 {
    return triangulo.base * triangulo.altura / 2
}
