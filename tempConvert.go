package main

import (
    "fmt"
)

// Função para converter Celsius para Fahrenheit
func celsiusParaFahrenheit(celsius float64) float64 {
    return (celsius * 9 / 5) + 32
}

// Função para converter Fahrenheit para Celsius
func fahrenheitParaCelsius(fahrenheit float64) float64 {
    return (fahrenheit - 32) * 5 / 9
}

func main() {
    // Exemplo de conversão de Celsius para Fahrenheit
    temperaturaCelsius := 25.0
    temperaturaFahrenheit := celsiusParaFahrenheit(temperaturaCelsius)
    fmt.Printf("%.2f graus Celsius é equivalente a %.2f graus Fahrenheit\n", temperaturaCelsius, temperaturaFahrenheit)

    // Exemplo de conversão de Fahrenheit para Celsius
    temperaturaFahrenheit = 77.0
    temperaturaCelsius = fahrenheitParaCelsius(temperaturaFahrenheit)
    fmt.Printf("%.2f graus Fahrenheit é equivalente a %.2f graus Celsius\n", temperaturaFahrenheit, temperaturaCelsius)
}
