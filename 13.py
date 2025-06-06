# Это программа для обработки чисел из двух файлов

# Сначала откроем первый файл
f1 = open('A.txt', 'r')
# Прочитаем все числа из файла
nums1 = f1.read().split()
# Закроем файл
f1.close()

# Сделаем список чисел из первого файла
list1 = []
for x in nums1:
    list1.append(int(x))

# Теперь второй файл
f2 = open('B.txt', 'r')
nums2 = f2.read().split()
f2.close()

list2 = []
for x in nums2:
    list2.append(int(x))

# Посчитаем сколько раз встречается каждое число в первом списке
count1 = {}
for num in list1:
    if num in count1:
        count1[num] += 1
    else:
        count1[num] = 1

# То же самое для второго списка
count2 = {}
for num in list2:
    if num in count2:
        count2[num] += 1
    else:
        count2[num] = 1

# Найдем числа, которые есть в обоих файлах
common_nums = []
for num in list1:
    if num in list2 and num not in common_nums:
        common_nums.append(num)

# Создадим итоговый список
final_list = []
for num in common_nums:
    # Найдем сколько раз число встречается в каждом файле
    cnt1 = count1[num]
    cnt2 = count2[num]
    # Возьмем большее значение
    if cnt1 > cnt2:
        max_cnt = cnt1
    else:
        max_cnt = cnt2
    # Добавим число столько раз
    for i in range(max_cnt):
        final_list.append(num)

# Отсортируем итоговый список
final_list.sort()

# Запишем результат в файл C.txt
f3 = open('C.txt', 'w')
for num in final_list:
    f3.write(str(num) + '\n')
f3.close()

print("Программа завершила работу! Результат в файле C.txt")