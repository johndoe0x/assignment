# assignment
## Scope of the presentation
### Please see PDF file name `KZG commitment scheme Fundamentals & build other schemes based on`.
You are expected to give a technical presentation on the KZG commitment scheme. The presentation should cover its use-cases, implementation aspect, security aspect and/or the underlying theory. In particular, explain how the KZG commitment scheme can be used as a building block to construct more complex cryptographic schemes.


## Coding assignment ; 
### Please take a look solution.go file. 

**WIP-solBySetCircle.go is trial for set up circle based on the formula, but not so much successful and could make finished.** 

Let 𝐶 be a circle and let 𝑃1, ..., 𝑃𝑛 be 𝑛 fixed points of 𝐶. As an input, you are given a list of two
intervals of 𝐶, 𝐼 and 𝐼', such they both satisfy either of the following:


● The interval is empty

● The interval is the entire circle 𝐶

● The interval is half-open of the form [𝑃𝑖, 𝑃𝑗) where 𝑖 ≠ 𝑗. i.e., the interval contains 𝑃𝑖 but
not 𝑃𝑗

The goal is to implement a function that detects if the union of 𝐼 and 𝐼' is an interval and output the union if so. You are free to implement circular interval as you wish.