import "math"

type Circulo struct {
    raio float64
}

func calcularAreaCirculo(circulo Circulo) float64 {
    return math.Pi * circulo.raio * circulo.raio
}
