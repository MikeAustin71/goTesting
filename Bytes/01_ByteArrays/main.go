package main

import "fmt"

func main() {

	bA := make([]byte, 0 , 100)

	for i:=65; i <91; i++ {

		bA = append(bA, byte(i) )

	}

	fmt.Println("Byte Array Output!")
	for k:=0; k < len(bA); k++ {
		fmt.Printf("%v. character= %c\n", k, bA[k])
	}

/* Output
Byte Array Output!
0. character= A
1. character= B
2. character= C
3. character= D
4. character= E
5. character= F
6. character= G
7. character= H
8. character= I
9. character= J
10. character= K
11. character= L
12. character= M
13. character= N
14. character= O
15. character= P
16. character= Q
17. character= R
18. character= S
19. character= T
20. character= U
21. character= V
22. character= W
23. character= X
24. character= Y
25. character= Z

 */
}
