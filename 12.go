package main

// Импортируем нужные штуки
import (
    "bufio"
    "fmt"
    "math"
    "os"
)

// Главная функция
func main() {
    // Открываем файлик
    file, err := os.Open("input.txt")
    
    // Если не получилось открыть файл
    if err != nil {
        fmt.Println("Ой, файл не открылся! Ошибка:", err)
        return // Выходим из программы
    }
    // Не забываем закрыть файл
    defer file.Close()

    // Читаем файл построчно
    scanner := bufio.NewScanner(file)
    
    // Переменные для коэффициентов
    var a, b, c int
    
    // Читаем первую строку
    if scanner.Scan() {
        line := scanner.Text() // Берем текст строки
        // Парсим коэффициенты из строки
        _, err := fmt.Sscanf(line, "%d %d %d", &a, &b, &c)
        if err != nil {
            fmt.Println("Не смог прочитать числа из файла :(")
            return
        }
    }
    
    // Проверяем ошибки сканера
    if err := scanner.Err(); err != nil {
        fmt.Println("Ошибка при чтении файла:", err)
        return
    }

    // Выводим что прочитали (для проверки)
    fmt.Println("Коэффициенты:")
    fmt.Println("A =", a)
    fmt.Println("B =", b)
    fmt.Println("C =", c)

    // Считаем дискриминант по формуле
    D := b*b - 4*a*c
    fmt.Println("Дискриминант D =", D)

    // Проверяем разные случаи
    if D > 0 {
        // Если D положительный - два корня
        sqrtD := math.Sqrt(float64(D))
        x1 := (-float64(b) + sqrtD) / (2 * float64(a))
        x2 := (-float64(b) - sqrtD) / (2 * float64(a))
        
        // Выводим с округлением до двух знаков
        fmt.Printf("Два корня:\n")
        fmt.Printf("x1 = %.2f\n", x1)
        fmt.Printf("x2 = %.2f\n", x2)
    } else if D == 0 {
        // Если D ноль - один корень
        x := -float64(b) / (2 * float64(a))
        fmt.Printf("Один корень:\n")
        fmt.Printf("x = %.2f\n", x)
    } else {
        // Если D отрицательный - нет корней
        fmt.Println("Нет корней (D < 0)")
    }
}