# Problem 1

Multiples of 3 and 5

`If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.`

---

## Thought Process

Simpilest way to approach this is to itterate over i and check if `i % 3 == 0` or `i % 3 == 5`.
If it is, add it to an int `sum` and return `sum` at the end.

The problem only has us going to 1000, but for a much larger data set, we might want to look at some fun shortcuts to accomplish it.

## Fun Times

We can use basic number theory to determine if larger numbers are divisible by 3 or 5 without dividing the entire number.


### Divisible by 3

A number is divisible by 3 if all of it's digits summed together is divisible by 3
1234 = 1 + 2 + 3 + 4 = 10 ... 10 is not divisible by 3 so 1234 is not divisible by 3. 
1233 = 1 + 2 + 3 + 3 = 9 ... 9 is divisible by 3 so 1233 is divisible by 3.

### Divisible by 5

A number is divisible by 5 if the last digit is 5 or 0 where x > 0;
