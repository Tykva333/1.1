package main

import (
	"fmt"
	"unicode"
)

func main() {
    // Объявляем переменные
    var text string
    var sdvig int
    
    // Просим пользователя ввести текст
    fmt.Println("Введите текст который надо зашифровать:")
    fmt.Scanln(&text) // Считываем текст
    
    // Просим ввести сдвиг
    fmt.Println("Теперь введите сдвиг (число):")
    fmt.Scanln(&sdvig) // Считываем сдвиг
    
    // Шифруем и выводим результат
    result := ""
    for _, bukva := range text {
        if unicode.IsUpper(bukva) {
            // Для больших букв
            newPos := (int(bukva) - 'А' + sdvig) % 33
            if newPos < 0 {
                newPos += 33
            }
            result += string(newPos + 'А')
        } else if unicode.IsLower(bukva) {
            // Для маленьких букв
            newPos := (int(bukva) - 'а' + sdvig) % 33
            if newPos < 0 {
                newPos += 33
            }
            result += string(newPos + 'а')
        } else {
            // Остальное оставляем как есть
            result += string(bukva)
        }
    }
    
    fmt.Println("Вот что получилось:", result)
}