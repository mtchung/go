# Scripts to convert delimited columns into a vs b matrix

Given the following input file
```
user1,ad1,A
user2,ad2,A
user3,ad3,A
user4,ad4,A
user5,ad1,A
user6,ad5,A
user7,ad1,A
user8,ad3,A
```

When you execute the go script, you will get the following output
```
go run main.go -i input.csv -r 0 -c 1
        ad1     ad2     ad3     ad4     ad5
user1   Y       .       .       .       .
user2   .       Y       .       .       .
user3   .       .       Y       .       .
user4   .       .       .       Y       .
user5   Y       .       .       .       .
user6   .       .       .       .       Y
user7   Y       .       .       .       .
user8   .       .       Y       .       .


go run main.go -i input.csv -r 0 -c 1
        ad1     ad2     ad3     ad4     ad5
user1   Y       .       .       .       .
user2   .       Y       .       .       .
user3   .       .       Y       .       .
user4   .       .       .       Y       .
user5   Y       .       .       .       .
user6   .       .       .       .       Y
user7   Y       .       .       .       .
user8   .       .       Y       .       .
```