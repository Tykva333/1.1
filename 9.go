package main

import (
    "fmt"
    "unicode"
)

// Функция для делания шифра
func do_cipher(text string) (result string) {
    
    // Создаем пустую строку для результата
    result = ""
    
    // Перебираем каждый символ (я не знаю как это работает, но в интернете так было)
    for i := 0; i < len(text); i++ {
        char := text[i] // Берем символ
        
        // Проверяем большая ли буква (мб есть способ лучше но я не знаю)
        if unicode.IsUpper(rune(char)) {
            // Какая-то сложная математика из интернета
            new_char := 'Z' - (int(char) - 'A')
            result = result + string(new_char)
        } else if unicode.IsLower(rune(char)) { // Иначе если маленькая
            // То же самое но для маленьких
            new_char := 'z' - (int(char) - 'a')
            result = result + string(new_char)
        } else {
            // Просто добавляем если не буква
            result = result + string(char)
        }
    }
    
    return // Возвращаем результат (почему-то без переменной)
}

func main() {
    var input string
    
    // Просим ввести текст
    fmt.Println("Введите текст пожалуйста:")
    
    // Читаем ввод (не уверен что правильно)
    fmt.Scanln(&input)
    
    // Делаем шифр
    output := do_cipher(input)
    
    // Выводим результат
    fmt.Println("Вот что получилось:", output)
    
    // Конец программы
    return
}