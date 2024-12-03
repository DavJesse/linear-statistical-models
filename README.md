# Linear-Stats: Linear Regression Line and Pearson Correlation Coefficient Calculator

This program reads numerical data from a text file and computes two statistical metrics:
- **Linear Regression Line** equation.
- **Pearson Correlation Coefficient**.

The program is written in Go and processes data formatted as described below.

---

## Features

1. **Reads Data**: Reads a file containing numeric values.
2. **Linear Regression**: Calculates the equation of the linear regression line (`y = mx + b`) with 6 decimal precision.
3. **Pearson Coefficient**: Calculates the Pearson Correlation Coefficient with 10 decimal precision.

---

## Input Format

The program expects a file containing numeric values, one per line. Each line represents a `y` value, with the `x` values implicitly being the line numbers (starting from 0).  

### Example File: `data.txt`
```
189
113
121
114
145
110
```

For this example:
- `x` axis: 0, 1, 2, 3, 4, 5
- `y` axis: 189, 113, 121, 114, 145, 110

---

## Output Format

The program prints the following results:
1. **Linear Regression Line**: Equation in the form `y = mx + b` with `m` and `b` rounded to 6 decimal places.
2. **Pearson Correlation Coefficient**: A value rounded to 10 decimal places.

### Example Output
```
Linear Regression Line: y = 2.345678x + 10.123456 Pearson
Correlation Coefficient: 0.9876543210
```

---

## How to Run

1. Save your data in a text file (e.g., `data.txt`).
2. Run the program using the Go command:
   ```bash
   $ go run your-program.go data.txt

## Requirements

- **Go programming language** installed on your machine.
- **Data file** formatted as specified above.

---

## Calculation Details

- **Linear Regression Line**:
  - Equation: `y = mx + b`
  - `m` (slope) and `b` (intercept) are derived using standard regression formulas.

- **Pearson Correlation Coefficient**:
  - Measures the strength and direction of the linear relationship between `x` and `y`.

---

## License

This project is licensed under the **MIT License**. See `LICENSE` for more information.

Feel free to adapt the details to match the exact name of your program file or additional project-specific information.

---

## Author
[David Jesse Odhiambo](https://github.com/DavJesse)

Software Developer, Zone01 Kisumu