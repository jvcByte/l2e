def is_palindrome(input_string):
    cleaned = "".join(c.lower() for c in input_string if c.isalnum())
    
    left = 0
    right = len(cleaned) - 1
    while left < right:
        if cleaned[left] != cleaned[right]:
            return ("Stopped Being a palindrom comparing index: " + str(left) + " to index: " + str(right))
        left += 1
        right -= 1
    return True


print(is_palindrome("racecar"))
print(is_palindrome("hello"))
print(is_palindrome("A man a plan a canal Panama"))
