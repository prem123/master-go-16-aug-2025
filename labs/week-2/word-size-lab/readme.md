## Given the following string

```
Commodo Lorem enim dolore culpa minim ipsum ex excepteur in Commodo duis nulla ex laborum irure sunt incididunt Incididunt amet Lorem amet dolor sit consectetur culpa esse quis laborum pariatur laborum fugiat mollit Mollit voluptate aliquip Lorem incididunt mollit pariatur eu enim proident culpa esse laborum voluptate Nostrud aliqua magna ipsum qui duis euNisi pariatur sit do magna Lorem nostrud voluptate occaecat occaecat quis dolore irure Velit aliqua reprehenderit eu duis aliqua excepteur duis non enim Nostrud qui voluptate enim eiusmod dolor proident laboris nostrud commodo laborum aliquip sunt Exercitation do anim do ullamco Ipsum eiusmod aute qui ea consectetur Veniam laborum occaecat mollit pariatur commodo id ullamco dolore ipsum sit dolore elit fugiatUt ad aliqua dolor quis velit reprehenderit Dolore aliquip exercitation do fugiat Pariatur irure aliqua magna quis
```

Find the highest occurring "word size" in the given string and print the highest occurring "word size" & the number of occurrences

Output should look like:

```
The highest occurring word size is: 5
The number of occurrences for word size 5 is: 24
```

<details>
  <summary>Hints</summary>

- **Split the string into words:** This can be done using the strings.Split function.
- **Measure the length of each word:** Loop through the words and measure their lengths.
- **Store the frequency of each word size:** Use a map where the key is the word size, and the value is the frequency of that size.
- **Find the word size with the maximum frequency:** Iterate through the map to find the size with the most occurrences.
- **Print the result.**

</details>
<br>

<details>
  <summary>Not sure how?</summary>

  ```go
    func main() {
      str := `Commodo Lorem enim dolore culpa minim ipsum ex excepteur in Commodo duis nulla ex laborum irure sunt incididunt Incididunt amet Lorem amet dolor sit consectetur culpa esse quis laborum pariatur laborum fugiat mollit Mollit voluptate aliquip Lorem incididunt mollit pariatur eu enim proident culpa esse laborum voluptate Nostrud aliqua magna ipsum qui duis euNisi pariatur sit do magna Lorem nostrud voluptate occaecat occaecat quis dolore irure Velit aliqua reprehenderit eu duis aliqua excepteur duis non enim Nostrud qui voluptate enim eiusmod dolor proident laboris nostrud commodo laborum aliquip sunt Exercitation do anim do ullamco Ipsum eiusmod aute qui ea consectetur Veniam laborum occaecat mollit pariatur commodo id ullamco dolore ipsum sit dolore elit fugiatUt ad aliqua dolor quis velit reprehenderit Dolore aliquip exercitation do fugiat Pariatur irure aliqua magna quis`
      wordSizeCount := make(map[int]int)
      // Split the string into words
      words := strings.Split(str, " ")
      for _, word := range words {
        size := len(word)
        wordSizeCount[size]++
      }

      // Find the highest occurring word size and its frequency
      var maxSize, maxCount int
      for size, count := range wordSizeCount {
        if count > maxCount {
          maxSize = size
          maxCount = count
        }
      }

      // Output the result
      fmt.Printf("The highest occurring word size is: %d\n", maxSize)
      fmt.Printf("The number of occurrences for word size %d is: %d\n", maxSize, maxCount)
    }

  ```
</details>
<br>
