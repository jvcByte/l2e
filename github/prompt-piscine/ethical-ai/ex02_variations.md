## Step 3 — Modify your palindrome function to:
- Ignore spaces and punctuation.
- Be case-insensitive.
- Return the position where the string stops being a palindrome (if not one). 
Function is_palindrome(input_string):
	Declare cleaned_string
	cleaned_string = clean the string
	Declare left
	Declare right = length of cleaned_string -1 
	While left < right:
		if cleaned_string[left] =! cleaned_string[right]
			Return  left
		Else:
			left = left + 1
			right = right - 1
	Return True
	End While
End Function