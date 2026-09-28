[TOC]



# **Mathematical Documentation for `IntExpoBigFloat` (013Fix)**

------

# **Integer Exponentiation Using Exponentiation‑by‑Squaring**

### *Function: `bigFloatMathMacrobot.intExpoBigFloat()`*

------

## **1. Overview**

The `IntExpoBigFloat` method computes:

$$ y = x^n $$

where:

- $x$ is a big.Float
- $n$ is a signed integer
- Precision is user‑controlled

Your implementation uses the **Exponentiation‑by‑Squaring** algorithm, which is the optimal method for integer exponents in arbitrary‑precision arithmetic.

------

## **2. Mathematical Foundation**

### **2.1 Basic Rules**

For integer $n$:

- If $n = 0$:
   $$ x^0 = 1 $$
- If $n < 0$:
   $$ xn = \frac{1}{x{-n}} $$
- If $n > 0$:
   use exponentiation‑by‑squaring

------

## **3. Exponentiation‑by‑Squaring**

This method reduces the number of multiplications from $O(n)$ to $O(\log n)$.

### **3.1 Recurrence**

For integer $n$:

- If $n$ is even:

$$ xn = (x2)^{n/2} $$

- If $n$ is odd:

$$ xn = x \cdot x{n-1} $$

------

### **3.2 Binary Interpretation**

Let $n$ be written in binary:

$$ n = (b_k b_{k-1} \dots b_1 b_0)_2 $$

Then:

$$ xn = \prod_{i=0}{k} x^{2^i b_i} $$

This means:

- Square the base each step
- Multiply into result only when the current bit is 1

------

## **4. Algorithm Structure in Your Implementation**

### **4.1 Working Precision**

```go
workingPrec := calculationPrecisionBits + 22
```

This ensures:

- Intermediate squaring does not lose precision
- Final rounding is accurate

------

### **4.2 Handling Negative Exponents**

```go
if exponent < 0 {
    exponent = -exponent
    negExponent = true
}
```

After computing $x^{|n|}$:

```go
if negExponent {
    result = 1 / result
}
```

------

### **4.3 Main Loop**

Your code:

```go
for exponent > 0 {
    if exponent & 1 == 1 {
        result *= base
    }
    base *= base
    exponent >>= 1
}
```

This is the canonical exponentiation‑by‑squaring loop.

### **Why it’s optimal**

- Each loop iteration halves the exponent
- Only $\log_2(n)$ squarings
- Only as many multiplications as there are 1‑bits in $n$

------

### **4.4 Final Rounding**

```go
out := new(big.Float).SetPrec(calculationPrecisionBits)
out.Set(result)
```

This ensures:

- Output precision matches user request
- Guard bits do not leak into final value

------

## **5. Summary of the Algorithm**

### **Exponentiation‑by‑Squaring Steps**

1. If $n = 0$: return $1$
2. If $n < 0$: compute $x^{-n}$ and take reciprocal
3. Initialize:
    $$ \text{result} = 1 $$
4. While $n > 0$:
   - If $n$ is odd:
      $$ \text{result} \leftarrow \text{result} \cdot x $$
   - Square the base:
      $$ x \leftarrow x^2 $$
   - Shift exponent right:
      $$ n \leftarrow n / 2 $$
5. Round to requested precision

------

## **6. Why This Method Works Well for big.Float**

- Minimal number of multiplications
- Excellent performance for large exponents
- Stable for high‑precision arithmetic
- Works seamlessly with negative exponents
- Fully deterministic and reproducible
- Validated against naive multiplication for 200+ digit precision
- Validated against known high‑precision expected values (456–855 digits)

------



