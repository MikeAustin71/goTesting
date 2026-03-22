[TOC]



# Mathematical Documentation for `sqrtBigFloat` (013Fix)**



## **Square Root Computation Using Newton’s Method**

### *Function: `bigFloatMathMacrobot.sqrtBigFloat()`*

------

## **1. Overview**

The `sqrtBigFloat` method computes
 $$ y = \sqrt{x} $$
 for a positive big.Float value $x$, using the **classic Newton–Raphson iteration**.
 This method is globally convergent for $x>0$, quadratically convergent near the root, and well‑suited for arbitrary‑precision arithmetic.

Your implementation enhances the classical method with:

- Precision‑aware working bit‑width
- Dynamic iteration scaling based on input size
- Stable initialization
- Strict error handling
- Final rounding to user‑requested precision

------

## **2. Mathematical Foundation**

To compute $\sqrt{x}$, Newton’s method applies to the function:

$$ f(y) = y^2 - x $$

The derivative is:

$$ f'(y) = 2y $$

Newton’s update rule is:

$$ y_{k+1} = y_k - \frac{f(y_k)}{f'(y_k)} $$

Substituting:

[ \begin{aligned} y_{k+1} &= y_k - \frac{y_k^2 - x}{2y_k} \ &= \frac{1}{2}\left(y_k + \frac{x}{y_k}\right) \end{aligned} ]

This is the **classic Babylonian iteration**:

$$ \boxed{ y_{k+1} = \frac{1}{2}\left(y_k + \frac{x}{y_k}\right) } $$

------

## **3. Algorithm Structure in Your Implementation**

### **3.1 Initial Guess**

Your code uses:

```go
y := (x / 2)
if y == 0 { y = 1 }
```

This is robust because:

- For large $x$, $x/2$ is a reasonable first approximation.
- For very small $x$, $x/2$ may underflow to zero, so you fall back to $1$.

------

### **3.2 Working Precision**

You compute:

```go
workingPrec := calculationPrecisionBits + 12%
```

This ensures:

- Newton’s intermediate values do not lose accuracy
- Final rounding to `calculationPrecisionBits` is correct
- Quadratic convergence is preserved

------

### **3.3 Iteration Count Scaling**

You compute:

- Number of integer digits in $x$
- Number of decimal digits implied by working precision

Then:

$$ \text{maxIter} = 100 \times (\text{integerDigits} + \text{decimalDigits}) $$

This is a **smart scaling rule**:

- Large integers → more iterations
- High precision → more iterations
- Ensures convergence even for 1000+ digit results

------

### **3.4 Convergence Check**

```go
if yNext.Cmp(y) == 0 {
    break
}
```

Because big.Float uses fixed precision, once two successive iterates are equal, the iteration has converged *at the current precision*.

------

### **3.5 Final Rounding**

After convergence:

```go
result := new(big.Float).SetPrec(calculationPrecisionBits)
result.Set(y)
```

This ensures:

- Output precision matches user request
- Internal guard bits do not leak into the final result

------

## **4. Summary of the Algorithm**

### **Newton Iteration Steps**

1. Choose initial guess $y_0$
2. Repeat until convergence:

$$ y_{k+1} = \frac{1}{2}\left(y_k + \frac{x}{y_k}\right) $$

1. Round to requested precision

------

## **5. Why This Method Works Well for big.Float**

- Quadratic convergence: digits double each iteration
- Stable for all $x>0$
- Easy to implement with arbitrary precision
- No transcendental functions required
- Works for extremely large integers (100–400 digits tested)
- Works for extremely high precision (1000+ digits tested)

------

# **End of Page 1**

------

------

# **Page 2 — Mathematical Documentation for `IntExpoBigFloat` (013Fix)**

### *A Typora‑Friendly Markdown Document*

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

# **End of Page 2**

------

If you'd like, I can now prepare:

- A combined “013Fix Technical Report”
- A developer‑oriented API reference
- A mathematical appendix with proofs of convergence
- A roadmap for integrating these into the AGM kernel

Just say the word, Mike.