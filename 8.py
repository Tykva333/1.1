from string import ascii_uppercase as big_letters

file = open("input.txt")
text_from_file = file.readline()
file.close()

big_text = text_from_file.upper()

print("Введите число для сдвига (шифр Цезаря):")
shift = input()
shift_num = int(shift)

def caesar(text, shift):
    result = ""
    for char in text:
        if char in big_letters:
            index = big_letters.index(char)
            new_index = index + shift
            if new_index >= len(big_letters):
                new_index = new_index - len(big_letters)
            result += big_letters[new_index]
        else:
            result += char
    return result

def atbash(text):
    result = ""
    reversed_alphabet = "ZYXWVUTSRQPONMLKJIHGFEDCBA"
    for char in text:
        if char in big_letters:
            index = big_letters.index(char)
            result += reversed_alphabet[index]
        else:
            result += char
    return result

print("Шифр Атбаш:")
print(atbash(big_text))

print("Шифр Цезаря со сдвигом", shift_num)
print(caesar(big_text, shift_num))