def main():
    original_text = input()
    key = input()
    action = input()

    key_long_text = key
    for i in range(len(original_text) - len(key)):
        key_long_text += key[i % len(key)]

    if action == "encode":
        shofer_text = ""
        for i in range(len(key_long_text)):
            new_digit = ord(original_text[i]) + ord(key_long_text[i]) % ord('A')
            if new_digit > 90:
                new_digit = new_digit % 90 + 64
            shofer_text += chr(new_digit)

        print(shofer_text)

    elif action == "decode":
        deshofer_text = ""
        for i in range(len(key_long_text)):
            new_digit = ord(original_text[i]) - ord(key_long_text[i]) % ord('A')
            if new_digit < 65:
                new_digit += 26
            deshofer_text += chr(new_digit)

        print(deshofer_text)

if __name__ == "__main__":
    main()