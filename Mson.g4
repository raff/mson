grammar Mson;

mson
   : value*
   ;

obj
   : '{' pair (',' pair)* '}'
   | '{' '}'
   ;

pair
   : key=STRING ':' value
   ;

array
   : '[' value (',' value)* ']'
   | '[' ']'
   ;

value
   : STRING
   | NUMBER
   | binData
   | numberLong
   | isoDate
   | obj
   | array
   | 'true'
   | 'false'
   | 'null'
   ;

STRING
   : '"' (ESC | SAFECODEPOINT)* '"'
   ;

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;

fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;

fragment HEX
   : [0-9a-fA-F]
   ;

fragment SAFECODEPOINT
   : ~ ["\\\u0000-\u001F]
   ;

NUMBER
   : '-'? INT ('.' [0-9] +)? EXP?
   ;

fragment INT
   : '0' | [1-9] [0-9]*
   ;

// no leading zeros

fragment EXP
   : [Ee] [+\-]? INT
   ;

// extensions for Mson

binData
   : 'BinData' '(' typ=NUMBER ',' data=STRING ')'
   ;

numberLong
   : 'NumberLong' '(' num=NUMBER ')'
   | 'NumberLong' '(' str=STRING ')'
   ;

isoDate
   : 'ISODate' '(' date=STRING ')'
   ;

// \- since - means "range" inside [...]

WS
   : [ \t\n\r] + -> skip
   ;
