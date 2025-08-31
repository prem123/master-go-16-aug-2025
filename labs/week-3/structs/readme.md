# Exercise: Working with Structs and Receiver Functions in Go

### Objective
Practice defining a struct, attaching methods using receiver functions, and manipulating struct data.

---

### Problem Statement
Create a program to represent a **Bank Account** using a `struct`.  
The program should:

1. Define a `BankAccount` struct with fields:
   - `accountNumber` (string)  
   - `accountHolder` (string)  
   - `balance` (float64)

2. Implement the following **receiver functions**:
   - `Deposit(amount float64)` → Adds money to the balance  
   - `Withdraw(amount float64)` → Deducts money if balance is sufficient  
   - `Display()` → Prints account details  

3. In `main()`:
   - Create a new account  
   - Perform a deposit and withdrawal  
   - Display the final account details  
   - Prevent deposit/withdrawal of negative amounts.

---

### 🧱 Starter Code (students fill in some parts)
```go
package main

import "fmt"

// Step 1: Define struct here


// Step 2: Define receiver functions (Deposit, Withdraw, Display)

func main() {
    // Step 3: Create a BankAccount instance

    // Step 4: Call Deposit()

    // Step 5: Call Withdraw()

    // Step 6: Display final account details
}
```

### Example run output

```yaml
Deposited 500. New Balance: 1500
Withdrawn 200. New Balance: 1300
Account Number: 12345
Account Holder: Alice
Final Balance: 1300
```


### Bonus Challenges

- Add a method `ChangeHolderName(newName string)` to update the account holder.
- Create multiple accounts and store them in a slice, then loop through and display them.