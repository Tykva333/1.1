# Исходная матрица 3x3
matrix = [
    [1, 2, 3],
    [4, 5, 2],
    [7, 1, 9]
]

# Функция для транспонирования матрицы
def transpose_matrix(m):
    # Создаем новую матрицу, где строки становятся столбцами
    transposed = []
    for i in range(len(m[0])):
        new_row = []
        for row in m:
            new_row.append(row[i])
        transposed.append(new_row)
    return transposed

# Функция для вычисления следа матрицы (сумма элементов на главной диагонали)
def matrix_trace(m):
    trace = 0
    for i in range(len(m)):
        trace += m[i][i]
    return trace

# Функция для вычисления определителя матрицы 3x3
def matrix_determinant(m):
    # Расписываем формулу определителя для матрицы 3x3
    a = m[0][0] * (m[1][1] * m[2][2] - m[1][2] * m[2][1])
    b = m[0][1] * (m[1][0] * m[2][2] - m[1][2] * m[2][0])
    c = m[0][2] * (m[1][0] * m[2][1] - m[1][1] * m[2][0])
    return a - b + c

# Выводим исходную матрицу
print("Исходная матрица:")
for row in matrix:
    print(row)

# Транспонируем и выводим матрицу
print("\nТранспонированная матрица:")
transposed = transpose_matrix(matrix)
for row in transposed:
    print(row)

# Вычисляем и выводим след и определитель
print("\nСлед матрицы:", matrix_trace(matrix))
print("Определитель матрицы:", matrix_determinant(matrix))