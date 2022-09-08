# Jetgo

**dynamic-typing, script-like, rule-based language**  

## Concept
* Fact: user's input data structure.
* Rule: user's jetgo code.
* Vm: lang-runtime, uses FACT as data context to evalute the given RULE.

## Types

* **boolean**

* **number**

* **string**

* **array**

* **table**

* **nil**
```
    res = false
    elem = 123
    str = "hello world"
    arr = array(123, true, "123", 3.2)   
    res = in(elem, arr)  
```

## Value assign

```  
    boolean:
        &&, ||, !, ==, !=
    
    number:
        +, -, *, /, %, >>, <<, &, |, >, <, >=, <=, ==, !=

    string:
        +, ==, !=, >> (regexp match)

    nil:
        ==, !=
```

## Value assign

* ONLY support assign value to 'LOCAL VARIANT', i.e.:  
```
rule test {
    a = 12
    b = 2
    c = a + b
    return c
}
```
local varints: a,b,c 

* table declare assign
```
rule myTable {
    key = "test"
    val = 12.6

    t = table {
        "number": 123,
        "hello" : "world",
        key     : val,
    }
    return t
}

```

NOTE:
* ONLY support: =
* NOT support: ++, --, +=, -=, *=, /=, %=, &=, |=

## Built-in methods

### array  
```
    array(elems...) []any

    in(any, array) bool
    
    len(array) number

    append(array, any) array  
```

### table  
```
    table() table

    in(any, table) bool
    
    len(table) number

    tset(table_var, string_key, any)

    tget(table_var, string_key, default) any // default is optional, if no such key in table, return default

    tdel(table_val, string_key)

    tseti(table_var, string_key, any) // any would be integer when any is a number  
```

### string  
```
    string(any) string

    len(string) number

    match(string, pattern) bool  

    contains(str, substr) bool  
    
    supper(string) string // upper string

    slower(string) string // lower string

    sprintf(fmt_str, arg...) string

    ssplit(string_var, string_sp) array // split str  
```

### number  
```
    number(number or string) number

    integer(number or string) int  
```

### untils
```
    time() number // unix timestamp
```

## Comment
```
    // Single Line

    /*
        Multiple Lines
    */
```  

## Reserved keywords
* All keywords composed with lowercases
* All GO keywords

## Syntax
* Supported keywords: if / return / for / break / continue / fork
* Not support **else** branch keyword
* Not support **func**

## Loop

* Classic
```go
for i = 0; i < 99; i = i + 1 {
    if i % 2 == 0 {
        continue
    }
    if i == 97 {
        break
    }
}
```

* Infinite Loop
```go
i = 0
for  {
    if i >= 99 {
        break
    }
    if i % 2 == 0 {
        continue
    }
    if i == 97 {
        break
    }
    i = i + 1
}
```

## Concurrent
 
#### How-to
```cassandraql
results = [timeout_ms] fork {
    method_call1(args...),
    method_call2(args...),
    method_call3(args...),
}
```  

* keyword: ***fork***
* timeout: timeout_ms number类型 毫秒单位
* callee: MUST be method
* returns: array type, length equals to the number of methods

## How rule looks like?

```
rule hello_word {  
    pi = 3.1415926  
    arr = array(pi, 123, 20, "test", false)

    if in(pi, arr) {  
        return true  
    }

    if Fact.ScHeader == nil {
        return false
    }

    if Fact.ScHeader.Puin == 123 {
        return false
    }
    return false  
}  
```

## Rule Syntax Definition
```  
rule $NAME $ATTR... {
   ... ...
}  
```  

* rule: keyword [MUST]  
* $NAME: rule name [MUST] 
* $ATTR: attr (0 - N)  

## How-to call Go native method
+ Bridge interface: method.Method
```
func(args []interface{}) (interface{}, error)
```
- CAN NOT return args or args' slice
- If error returned, vm would stop immediately

* Register jetgo method  
```
    // NOTE: func_name should be lowercases， your_method type should be method.Method
    native.Register("func_name", your_method)
```

## Pointer
* NOT support pointer operators
* pointers in Fact would be deferenced, it means jetgo SUPPORT **protobuf2 struct as Fact data**

## Limitation
* Local variants \ method *MUST* be lowercases
* The scope of local variant *ONLY* be in its rules
* Max number of local variants must not exceed than *256* per rule
* Rule ONLY return <= 1 value
* Can not pass method as value
* Can not call Fact's Go method

