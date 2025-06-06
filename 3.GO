 # Читаем лабиринт из файла
try:
    with open(input_file, "r") as f:
        maze = []
        for line in f:
            maze.append(list(line.strip()))
except:
    print("Ошибка! Файл не найден или не читается")
    return

# Ищем начало (S) и конец (E)
start = None
end = None
for i in range(len(maze)):
    for j in range(len(maze[i])):
        if maze[i][j] == "S":
            start = (i, j)
        elif maze[i][j] == "E":
            end = (i, j)

if not start or not end:
    print("В лабиринте нет начала или конца!")
    with open(output_file, "w") as f:
        f.write("В лабиринте нет начала или конца!\n")
    return

# Ищем путь
path = find_path(maze, start, end)

if path:
    # Помечаем путь в лабиринте
    steps = 0
    for i, j in path:
        if maze[i][j] not in ("S", "E"):
            maze[i][j] = "*"
            steps += 1

    # Записываем результат
    with open(output_file, "w") as f:
        f.write(f"Шагов: {steps-1}\n")  # -1 потому что начальная точка не считается
        for row in maze:
            f.write("".join(row) + "\n")

    # Выводим результат
    print(f"Нашли путь за {steps-1} шагов!")
    for row in maze:
        print("".join(row))
else:
    print("Путь не найден :(")
    with open(output_file, "w") as f:
        f.write("Путь не найден\n")# Куда можно ходить
directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]

while queue:
    pos, path = queue.popleft()
    if pos == end:
        return path

    # Проверяем соседние клетки
    for di, dj in directions:
        ni, nj = pos[0] + di, pos[1] + dj
        # Проверяем границы и стенки
        if 0 <= ni < len(maze) and 0 <= nj < len(maze[0]):
            if maze[ni][nj] != "#" and (ni, nj) not in visited:
                visited.add((ni, nj))
                queue.append(((ni, nj), path + [(ni, nj)]))

return None

