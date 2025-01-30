package main

import "fmt"

func main() {

	bA := make([]byte, 0 , 100)

	for i:=65; i <91; i++ {

		if i== 70 {

			bA = append(bA, byte(0))

		}

		bA = append(bA, byte(i) )

	}

	fmt.Println("Byte Array Output!")
	fmt.Println("Searching for Zero Value")
	for k:=0; k < len(bA); k++ {

		if bA[k] == 0 {
			fmt.Println("--------------------------------")
			fmt.Printf("k=%v Found Zero byte=%d\n", k, bA[k])
			fmt.Println("--------------------------------")
			continue
		}

		fmt.Printf("%v. character= %c\n", k, bA[k])
	}

	/* Output
Byte Array Output!
Searching for Zero Value
0. character= A
1. character= B
2. character= C
3. character= D
4. character= E
--------------------------------
k=5 Found Zero byte=0
--------------------------------
6. character= F
7. character= G
8. character= H
9. character= I
10. character= J
11. character= K
12. character= L
13. character= M
14. character= N
15. character= O
16. character= P
17. character= Q
18. character= R
19. character= S
20. character= T
21. character= U
22. character= V
23. character= W
24. character= X
25. character= Y
26. character= Z
	 */

}
