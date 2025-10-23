🧩 Common Format Verbs by Type:
✅ General Verbs
Verb	Description
%v	Default format (most commonly used)
%+v	Adds field names when printing structs
%#v	Go-syntax representation of the value
%T	Type of the value
%%	A literal percent sign (%)
📦 Boolean
Verb	Description
%t	The word true or false
🔢 Integer (signed and unsigned)
Verb	Description
%b	Base 2 (binary)
%c	Character represented by the corresponding Unicode code point
%d	Base 10
%o	Base 8
%O	Base 8 with 0o prefix
%x	Base 16, lowercase
%X	Base 16, uppercase
%U	Unicode format: U+1234
🧮 Floating-Point / Decimal
Verb	Description
%b	Scientific notation with exponent as power of two
%e / %E	Scientific notation (e.g. 1.23e+02)
%f	Decimal point but no exponent
%g / %G	Uses e or f format, whichever is more compact
🪄 String and Byte Slice
Verb	Description
%s	String or slice of bytes as string
%q	Double-quoted string (escapes special chars)
%x / %X	Hex encoding (lower/upper)
🧱 Pointer
Verb	Description
%p	Pointer address in base 16