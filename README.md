# README

The word_type_formatter provides a simple utility for identifying and formatting words by type.

```
$ echo "apple dog mango carrot horse foo broccoli potato mouse" | go run word_type_formatter.go
> APPLE Unknown word: dog MANGO [carrot] h*o*r*s*e Unknown word: foo [broccoli] Unknown word: potato m*o*u*s*e
```
