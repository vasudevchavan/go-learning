
# String


|**Category**|	**Function / Operation**|	**Example**	|**Output / Description**|
|---------|---------------------|-----------|--------------------|
|**Basic**	|Declaration	|s := "Hello"	|Create a string|
||Length	|len(s)	|5|
||Access char (byte)	|s[0]|	'H' (byte)|
||Iterate runes	for i, r := range s {}	|Iterates Unicode characters|
|**Concatenation**	+	|"Hi " + "Go"	|"Hi Go"|
|**Comparison**	==, !=	|s1 == s2|	true/false|
||Lexical	|s1 < s2	|true/false|
|**Case & Trim**	|strings.ToUpper(s)|	"HELLO"	|Uppercase|
||strings.ToLower(s)	|"hello"	|Lowercase|
||strings.TrimSpace(s)	|"Hello " → "Hello"	|Remove spaces|
||strings.Trim(s, "H")	|"Hello" → "ello"	|Trim chars|
|**Search**	|strings.Contains(s, "ll")	|true|	Substring exists?|
||strings.HasPrefix(s, "He")	|true	|Starts with?|
||strings.HasSuffix(s, "lo")	|true	|Ends with?|
||strings.Index(s, "l")	|2	|First index|
||strings.LastIndex(s, "l")	|3|	Last index|
|**Replace**	|strings.Replace(s, "l", "L", -1)	|"HeLLo"	|Replace all occurrences|
|**Split & Join**	|strings.Split("a,b,c", ",")	|["a","b","c"]	|Split string|
||strings.Join([]string{"a","b"}, "-")	|"a-b"	|Join array|
|**Repeat**	|strings.Repeat("Go", 3)	|"GoGoGo"	|Repeat string|
|**Count**	|strings.Count(s, "l")	|2	|Count occurrences|
|**Conversion**	|strconv.Atoi("123")	|123	|String → int|
||strconv.Itoa(123)	|"123"	|Int → string|
||strconv.ParseFloat("3.14", 64)	|3.14	|String → float|
||[]byte(s)	|[]byte("Hi")	|String → byte slice|
||string([]byte{72,105})	|"Hi"	|Byte slice → string|
|**Formatting**	|fmt.Sprintf("Hi %s", "Go")	|"Hi Go"	|Format string|
|**Unicode / Runes**	|[]rune(s)	|[]rune("Hi")	|Convert string to runes|
|**Sorting**	|sort.Strings(slice)	|Sort array of strings	|["a","b","c"]|
|**Check empty**	|s == ""	|true/false	|Is string empty?|





# Go: String vs Byte vs Rune 

| **Operation / Task**       | **Use String**                         | **Use Byte (`[]byte`)**             | **Use Rune (`[]rune`)**                  |
|----------------------------|---------------------------------------|------------------------------------|----------------------------------------|
| **Length**                 | `len(s)` → bytes                       | `len(b)` → bytes                   | `len(runes)` → characters              |
| **Access / Index**         | `s[i]` → byte                          | `b[i]` → byte                      | `runes[i]` → Unicode character        |
| **Iterate**                | `for i := 0; i < len(s); i++ {}` → bytes | `for _, b := range b {}` → bytes  | `for i, r := range runes {}` → Unicode characters |
| **Concatenate**            | `s1 + s2`                              | `append(b1, b2...)`                | `append(runes1, runes2...)`           |
| **Substring / Slice**      | `s[a:b]` → bytes                        | `b[a:b]` → bytes                   | `string(runes[a:b])` → Unicode-safe slice |
| **Convert**                | `[]byte(s)`                             | `string(b)`                         | `string(runes)` → back to string       |
| **Count characters**       | ❌ counts bytes                         | ❌ counts bytes                      | ✅ counts actual characters             |
| **Modify characters**      | ❌ immutable                             | ✅ modify bytes                     | ✅ modify Unicode characters safely    |
| **Unicode-safe iteration** | ❌ may split multi-byte chars           | ❌ may split multi-byte chars        | ✅ iterates full characters            |
| **Emoji / Accents**        | ❌ risky                                | ❌ risky                             | ✅ correct                             |

---

