# pig-latin
CLI for converting text according to Pig Latin
## Rules :



* ### words that begin with consonant sounds
  
    `"sm-ile" > "ile" +"sm" + "ay"`  
    `"l-atin" = "atin" + "l" + ay"`  
    `"d-uck" = "uck" + "d" + "ay"`  

* ### words begin with consonant clusters 
  
    `"sm-ile" = "ile" + "sm" + "ay"`  
    `"str-ing" = "ing"+ "str" + "ay"`  
    `"gl-ove" = "ove" + "gl" + "ay"`  

* ### words that begin with vowel sounds

    `"explain" = "explain" + "yay"`  
    `"egg" = "egg" + "yay"`  
    `"omelet" = "omelet" + "yay"`   
    `"eat" = "eat" + "yay"`   

*Consonants:  
[B, C, D, F, G, H, J, K, L, M, N, P, Q, R, S, T, V, X, Z]*

*Vowels:  
[A, E, I, O, U, Y]*

---
How to install :
---

`$ go build`

---
How to use :
---

`$ ./pig-latin -h`  
will show help-menu

`$ ./pig-latin -translate "hello Dolly and Polly"`
<br/> will return `"ellohay ollyDay andyay ollyPay"`
