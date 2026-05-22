**Sudoku Solver in Go**

Hi there! This is a small Go program that solves Sudoku puzzles for you.

You give it a Sudoku board, and it will fill in all the missing numbers—or let you know if something is wrong with your puzzle.

How It Works

1. **Takes Your Puzzle as Input**
   - You provide 9 rows of numbers.  
   - Use `.` for empty cells.  
   - Example row: `"53..7...."`  

2. **Checks Your Puzzle**  
   - Ensures there are exactly 9 rows and each row has 9 characters.  
   - Only numbers `1–9` or `.` are allowed.  
   - Validates that the pre-filled numbers do not break Sudoku rules.  

3. **Solves the Puzzle**  
   - Fills in empty cells one by one using recursive backtracking.  
   - Ensures each number follows Sudoku rules (no repeats in rows, columns, or 3x3 blocks).  
   - Backtracks if a placement leads to a dead end.  
   - Continues until the puzzle is solved or deemed impossible.  

4. **Shows You the Solution**  
   - Prints the solved Sudoku with spaces between numbers.  
   - If the puzzle is invalid or unsolvable, it prints:  
     ```
     Error
     ```

**How to Run It**

Open your terminal and run the program using `go run .` with 9 arguments (one for each row of the puzzle):

**Example 1: Solving a Valid Sudoku**

bash

go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e

**Expected Output**

3 9 6 2 4 5 7 8 1$
1 7 8 3 6 9 5 2 4$
5 2 4 8 1 7 3 9 6$
2 8 7 9 5 1 6 4 3$
9 3 1 4 8 6 2 7 5$
4 6 5 7 2 3 9 1 8$
7 1 2 6 3 8 4 5 9$
6 5 9 1 7 4 8 3 2$
8 4 3 5 9 2 1 6 7$

**Examples of expected outputs for invalid inputs or sudokus**

$ go run . 1 2 3 4 | cat -e
Error$

$ go run . | cat -e
Error$

$ go run . ".96.4...1" "1...6.1.4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
Error$
$