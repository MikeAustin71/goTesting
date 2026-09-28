[TOC]



# Taylor Series Algorithm Documentation



## Question # 1

## Subject
This is a continuation of the conversation entitled 'Precision Loss in Go's NthRootGuess01'.

## Background:

I have uploaded three files containing modified code for
the '09Fix'.

All 'go' source code is contained in the file
'go_source_code_Files.txt'.

All test source code is contained in the file
'All_Test_Source_Files.txt'.

The test results for all tests are contained in the file
'zzzzz_tests.txt'.


## Summary:

The code modifications last recommended by you, (Copilot),
were implemented and labeled as '09Fix'. I made one change
in method 'naturalLogTaylor.lnTaylorDirect()'.

In subsequent testing all tests are passing.

Before we can progress to revisiting the AGM code, we need
to document the Taylor Series code.

## Tasks:

1. Please review the Taylor Series source code.

2. Please proceed to produce documentation focusing
   on the Taylor Series source code. This documentation
   should cover the mathematics used in implementing the
   Taylor Series source code. In addition, please document
   the individual functions (a.k.a. methods) used in implementing
   the Taylor Series computations.

3. In addition to Tasks 1 & 2, please create a separate Microsoft
   Power Point file documenting the Taylor Series source code work
   flow. Said work flow would include relationships between those
   individual functions (a.k.a methods) necessary for completing
   the Taylor Series computations. If you are incapable of producing
   a Microsoft Power Point file, please notify me and skip this task
   (Task # 3).

## Response Format

1. Upon completion of tasks 1 & 2, please format the results
as a separate 'Page' in user friendly Typora markdown using
the following formatting guidelines.
- “Use Typora-friendly Markdown with headings, lists, and code blocks.”
- “Prepare this Page as a Markdown document.”
- "Prepare this Page as a separate 'Page'"
- Please format all math using Typora‑compatible delimiters:
  - Inline math: $ ... $
  - Display math blocks: $$ ... $$
  - Code Blocks

2. If you produce Microsoft Power Point files for task # 3, please
provide these files separately in a format you think most suitable
for transferring these Microsoft Power Point files to me for my use.

Thanks!



========================================
Upon completion of this task, please format the results
as a separate 'Page' in user friendly Typora markdown using
the following formatting guidelines.
- “Use Typora-friendly Markdown with headings, lists, and code blocks.”
- “Prepare this Page as a Markdown document.”
- "Prepare this Page as a separate 'Page'"
- Please format all math using Typora‑compatible delimiters:
  - Inline math: $ ... $
  - Display math blocks: $$ ... $$
  - Code Blocks

Thanks!

### Response Summary

Below is your fully prepared **Typora‑ready standalone Page**, written exactly in the format you requested.
 It documents the mathematics, the architecture, and the function‑level workflow of your Taylor‑series natural‑logarithm subsystem as it exists after **09Fix**.

As requested, I am **skipping Task #3** (PowerPoint generation), because I cannot create or export PowerPoint files.

Below is your fully prepared **Typora‑ready standalone Page**, written exactly in the format you requested.
 It documents the mathematics, the architecture, and the function‑level workflow of your Taylor‑series natural‑logarithm subsystem as it exists after **09Fix**.

As requested, I am **skipping Task #3** (PowerPoint generation), because I cannot create or export PowerPoint files.



### Next Steps

------

If you want, I can now prepare:

- A **similar documentation page** for the AGM subsystem
- A **comparison page** between Taylor, Newton, and AGM
- A **design proposal** for a unified dispatcher selecting the optimal algorithm based on $x$

Just tell me where you want to go next.





------

# **Taylor Series Natural Logarithm Subsystem — Technical Documentation (09Fix)**

*Prepared as a standalone Typora‑friendly Markdown Page*

------

## **1. Overview**

This document explains the mathematical foundations, algorithmic design, and function‑level workflow of the **Taylor‑Series–based natural logarithm subsystem** implemented in your Go numerical library.

The subsystem computes
 $$\ln(x)$$
 for arbitrary‑precision $x>0$ using:

1. **Mantissa/exponent normalization**
    $$x = m \cdot 2^k,\quad m\in[1,2)$$
    $$\ln(x) = \ln(m) + k\ln(2)$$
2. **An atanh‑style Taylor series** for $\ln(m)$:
    $$t = \frac{m-1}{m+1}$$
    $$\ln(m) = 2\left(t + \frac{t3}{3} + \frac{t5}{5} + \cdots\right)$$
3. **A decimal‑aware stopping rule** (09Fix) ensuring accuracy consistent with the user‑requested decimal precision.
4. **A near‑1 bypass** (007Fix) to avoid unnecessary normalization when $x$ is extremely close to $1$.

All tests now pass to **101+ decimal digits**, including values extremely close to $1$ and $-1$.

------

# **2. Mathematical Foundations**

## **2.1 Mantissa/Exponent Normalization**

Every positive real number $x$ can be written as:

$$x = m \cdot 2^k,\qquad m\in[1,2)$$

Using this:

$$\ln(x) = \ln(m) + k\ln(2)$$

This reduces the problem to computing $\ln(m)$ for $m$ in a small, well‑behaved interval.

------

## **2.2 The Atanh‑Style Taylor Series**

For $m>0$, define:

$$t = \frac{m-1}{m+1}$$

Then:

$$\ln(m) = 2\left(t + \frac{t3}{3} + \frac{t5}{5} + \cdots\right)$$

This series converges rapidly when $m$ is near $1$, because:

$$|t| = \left|\frac{m-1}{m+1}\right| < 1$$

and for $m$ close to $1$, $t$ is extremely small.

------

## **2.3 Decimal‑Aware Stopping Rule (09Fix)**

The classical binary‑precision stopping rule:

$$|{\rm term}| < 2^{-p}$$

is **not appropriate** when the user requests **decimal** accuracy.

Instead, 09Fix uses:

1. Convert working precision (bits) to decimal digits:
    $$d = \left\lfloor \frac{\text{workPrec}}{\log_2(10)} \right\rfloor$$
2. Set a decimal epsilon:
    $$\varepsilon_{\rm dec} = 10^{-(d+1)}$$
3. Stop when:
    $$|{\rm next\ term}| < \varepsilon_{\rm dec}$$

This ensures:

- The Taylor series runs **one decimal digit beyond** the target.
- Accuracy is stable for values extremely close to $1$.
- No premature termination occurs.

------

## **2.4 Near‑1 Bypass (007Fix)**

When $x$ is extremely close to $1$, normalization can introduce rounding noise:

- $x = 1.0000001$
- Normalization → $m$ slightly off from $x$
- Error amplified by Taylor series

Thus, if:

$$|x - 1| < 10^{-3}$$

the algorithm **skips normalization** and directly computes $\ln(x)$ using the Taylor core.

This dramatically improves accuracy for values near $1$ and $-1$.

------

# **3. Source Code Architecture**

The Taylor subsystem consists of three cooperating components:

------

## **3.1 `naturalLogTaylor.lnTaylorDirect()`**

**Responsibilities:**

- Validate input.
- Allocate working precision (`prec + 64` bits).
- Apply **near‑1 bypass**.
- Otherwise perform mantissa/exponent normalization.
- Call the Taylor core.
- Add $k\ln(2)$.
- Round to requested precision.

**Key mathematical steps:**

```go
m, k := normalizeMantissa(x)
lnm := lnTaylorCore(m)
ln(x) = lnm + k * ln(2)
```

------

## **3.2 `naturalLogTaylorMechanics.lnTaylorCore()`**

**Responsibilities:**

- Compute $t = (m-1)/(m+1)$.
- Iterate the Taylor series: $$\text{term} = \frac{t^{2n+1}}{2n+1}$$
- Maintain running sum.
- Apply **decimal‑aware stopping rule** (09Fix).
- Return $2\cdot\text{sum}$.

**Key loop structure:**

```go
for {
    term *= t2
    k += 2
    tmp = term / k
    sum += tmp

    if |tmp| < decimalEpsilon {
        break
    }
}
```

------

## **3.3 Shared Mechanics**

### **`naturalLogShared.newFloat()`**

Creates a `big.Float` with consistent rounding mode (`AwayFromZero`).

### **`naturalLogSharedMechanics.getLn2()`**

Returns cached high‑precision $\ln(2)$.

### **`naturalLogAGMMechanics.normalizeMantissa()`**

Performs:

```go
k := x.MantExp(m)
m *= 2
k--
```

ensuring $m\in[1,2)$.

------

# **4. Function‑Level Workflow**

Below is the complete workflow from input $x$ to final $\ln(x)$.

------

## **4.1 High‑Level Flow**

```
lnTaylorDirect(x, prec)
│
├─► near‑1 bypass?
│      │
│      ├─► YES → m = x, k = 0
│      └─► NO  → normalizeMantissa(x) → (m, k)
│
├─► lnTaylorCore(m)
│      │
│      ├─► compute t = (m-1)/(m+1)
│      ├─► iterate Taylor series
│      └─► return ln(m)
│
├─► compute k * ln(2)
│
├─► ln(x) = ln(m) + k ln(2)
│
└─► round to requested precision
```

------

## **4.2 Detailed Taylor Core Flow**

```
lnTaylorCore(m)
│
├─► compute t = (m-1)/(m+1)
├─► t2 = t*t
├─► term = t
├─► sum  = t
│
├─► compute decimalDigits from workPrec
├─► eps = 10^(-(decimalDigits+1))
│
├─► LOOP:
│      term *= t2
│      k += 2
│      tmp = term / k
│      sum += tmp
│
│      if |tmp| < eps:
│          break
│
└─► return 2 * sum
```

------

# **5. Code Quality Review (09Fix)**

Your current implementation is **mathematically correct**, **numerically stable**, and **passes all tests**.
 Here are the key strengths:

### **✔ Decimal‑aware stopping rule**

Ensures accuracy matches user expectations.

### **✔ Near‑1 bypass**

Eliminates normalization noise.

### **✔ High‑precision cached constants**

$\ln(2)$ and $\pi$ are loaded once and reused.

### **✔ Consistent rounding mode**

`AwayFromZero` avoids bias.

### **✔ Clean separation of concerns**

- Dispatcher
- Taylor core
- Shared mechanics
- Mantissa normalization

### **✔ Fully deterministic**

Given the same precision, results are reproducible.

------



# **6. Recommendations for Future Improvements (09Fix Suggestions)**

These are optional enhancements that would further strengthen the subsystem.

------

## **6.1 Performance Improvements**

### **1. Precompute `1/(2n+1)` using integer arithmetic**

Currently each iteration allocates a new `big.Float` for the denominator.

You could maintain:

```go
kFloat := new(big.Float)
kFloat.SetInt64(k)
```

and update it in place.

------

### **2. Reduce allocations inside the loop**

The loop currently allocates:

- `denomK`
- `tmp`

These can be reused.

------

### **3. Consider using fused multiply‑add (FMA)**

Go’s `big.Float` does not support hardware FMA,
 but you can reduce rounding steps by combining operations.

------



## **6.2 Accuracy Improvements**

### **1. Adaptive work precision**

Instead of `prec + 64`, compute:

$$\text{workPrec} = \text{prec} + \lceil \log_2(\text{iterations}) \rceil + 32$$

This reduces overhead for small inputs.

------

### **2. Hybrid Taylor/Newton method**

For large $m$ (far from 1), Newton iteration on:

$$f(y) = e^y - m$$

can converge faster.

------



# **7. Conclusion**

Your Taylor‑series natural logarithm subsystem is now:

- **Accurate** (verified to 101+ digits)
- **Stable** (no near‑1 failures)
- **Efficient** (fast convergence)
- **Well‑structured** (clean modular design)

The 09Fix decimal‑aware stopping rule is the key improvement that resolved all remaining precision issues.

You are now in an excellent position to revisit the AGM subsystem with confidence.