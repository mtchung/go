# CSV to Matrix Converter

A lightweight Go CLI utility that transforms a delimited CSV file into a visual cross-tabulation matrix (pivot table). 

It dynamically maps unique values from a **row column** and a **column column**, marking intersections with a `Y` (present) or a `.` (absent).

---

## 🚀 Features

* **Dynamic Columns:** Automatically detects and sorts unique headers.
* **Flexible Mapping:** Configure which CSV columns represent your rows and columns via CLI flags.
* **Memory Efficient:** Built using optimized Go slices and maps for handling larger datasets.

---

## 🛠️ Usage

### Prerequisites
Ensure you have [Go](https://go.dev) installed (version 1.16 or higher recommended).

### Syntax
```bash
go run main.go -i <input-csv> -d <input-delimiter> -o <output-delimiter>  -r <row column_index> -c <column column_index> 
```

### Flags

| Flag | Description | Type | Mandatory | Default |
| :--- | :--- | :--- | :--- | :--- |
| `-i` | Path to the input CSV file | `string` | Yes | input.csv |
| `-d` | Input Delimiter | `string` | Yes | , |
| `-o` | Output Delimiter | `string` | Yes | \t |
| `-r` | Zero-based index for the **row** data | `int` | Yes | 0 |
| `-c` | Zero-based index for the **column** data | `int` | Yes | 1 |
| `-h` | Help | `string` | No | N/A |
| `-v` | Version | `string` | No | N/A |
---

## 📋 Example

### 1. Input Data (`input.csv`)
```csv
user1,ad1,A
user2,ad2,A
user3,ad3,A
user4,ad4,A
user5,ad1,A
user6,ad5,A
user7,ad1,A
user8,ad3,A
```

### 2. Help Menu
To pivot the data using **Users** (index `0`) as rows and **Ads** (index `1`) as columns:

```bash
go run main.go -h
```

### 3. Help Menu Output
```text
Usage: main [-hV] [-c value] [-d value] [-i value] [-r value] [parameters ...]
 -c, --column=value
                  Column Field2 (column) [1]
 -d, --delimiter=value
                  Delimiter Value Default to comma(,) [,]
 -h, --help       Display help menu
 -i, --inputCSV=value
                  Path to the output file [input.csv]
 -r, --row=value  Column Field1 (row)
 -V, --version    Display version info
```


### 4. Execution Command
To pivot the data using **Users** (index `0`) as rows and **Ads** (index `1`) as columns:

```bash
go run main.go -i input.csv -r 0 -c 1
```

### 5. Output Matrix
```text
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

### 6. Execution Command
To pivot the data using **Users** (index `1`) as rows and **Ads** (index `0`) as columns:

```bash
go run main.go -i input.csv -r 1 -c 0
```

### 7. Output Matrix
```text
        user1   user2   user3   user4   user5   user6   user7   user8
ad1     Y       .       .       .       Y       .       Y       .
ad2     .       Y       .       .       .       .       .       .
ad3     .       .       Y       .       .       .       .       Y
ad4     .       .       .       Y       .       .       .       .
ad5     .       .       .       .       .       Y       .       .
```
