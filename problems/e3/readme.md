# [Problem 3](https://projecteuler.net/problem=3)

Largest Prime Factor

> The prime factors of 13195 are 5, 7, 13 and 29.
> What is the largest prime factor of the number 600851475143 ?

---

## Thought Process

### Initial Thoughts

That's a really big number! Luckily this number fits in a 64bit uint or int so it shouldn't be too hard to deal with.

I'm first going to divide the number by 2 as having a factor larger than half of your number shouldn't be possible.
I already wrote a function in the helper library that pulls a list of all prime numbers up to a number. I can do that to get a list of all primes that we can check against for factorization. I'll then itterate through that list (starting at the back) and determine if the prime is a factor of half of the larger number. I can use % to determine if there's a remainder, and if so, it cannot be a factor.

### Thinking Clearly

The above method would have worked over enough time but was not needed at all. Since we're workign with factors, dividing the number by it's factors makes this process extremely fast. The new stratagey is to factor out the number searching for primes and adding them to a list. Once the last  factor is found record it and report back with the last value in the list.


