package patterns

import "fmt"

func Pattern1(n int) {
	/*
		 * * *
		* * *
		* * *

	*/

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%s", "* ")
		}
		fmt.Println()
	}

}

func Pattern2(n int) {
	/*
			Input Format: N = 6
		Result:
		*
		* *
		* * *
		* * * *
		* * * * *
		* * * * * *

	*/

	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%s", "* ")
		}
		fmt.Println()
	}

}

func Pattern3(n int) {
	/*
			Input Format: N = 6
		Result:
		1
		1 2
		1 2 3
		1 2 3 4
		1 2 3 4 5
		1 2 3 4 5 6

	*/

	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%d %s", j+1, " ")
		}
		fmt.Println()
	}

}

func Pattern4(n int) {
	/*
			Input Format: N = 6
		Result:
		1
		2 2
		3 3 3
		4 4 4 4
		5 5 5 5 5
		6 6 6 6 6 6

	*/

	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%d %s", i, " ")
		}
		fmt.Println()
	}

}

func Pattern5(n int) {
	/*
			Input Format: N = 6
		Result:
		* * * * * *
		* * * * *
		* * * *
		* * *
		* *
		*

	*/

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Printf("%s", "* ")
		}
		fmt.Println()
	}

}

func Pattern6(n int) {
	/*
				Input Format: N = 6
		Result:
		1 2 3 4 5 6
		1 2 3 4 5
		1 2 3 4
		1 2 3
		1 2
		1

	*/

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Printf("%d %s", j+1, " ")
		}
		fmt.Println()
	}

}

func Pattern7(n int) {
	/*
				Input Format: N = 6
			Result:
			     *
			    ***
			   *****
			  *******
			 *********
			***********

		considering 0 as index

		we can see that the first spaces are n-i-1 = 5 spaces
		then (2*i)+1  stars = 1 stars
		and the n again n-i-1 spaces

	*/

	for i := 0; i < n; i++ {

		//printing spaces
		for j := 0; j < n-i-1; j++ {
			fmt.Printf("%s", " ")
		}
		//printing stars
		for j := 0; j < (2*i)+1; j++ {
			fmt.Printf("%s", "*")
		}

		//printing spaces
		for j := 0; j < n-i-1; j++ {
			fmt.Printf("%s", " ")
		}
		fmt.Println()
	}

}
func Pattern8(n int) {
	/*
					Input Format: N = 6
		Result:
		***********
		 *********
		  *******
		   *****
		    ***
		     *
			considering 0 as index

			then ((n-i)*2)-1 or 2*n -( 2i +1)  stars = 9 stars
			and the i spaces

	*/

	for i := 0; i < n; i++ {

		//printing spaces
		for j := 0; j < i; j++ {
			fmt.Printf("%s", " ")
		}
		//printing stars
		for j := 0; j < 2*n-(2*i+1); j++ {
			fmt.Printf("%s", "*")
		}

		//printing spaces
		for j := 0; j < i; j++ {
			fmt.Printf("%s", " ")
		}
		fmt.Println()
	}

}

func Pattern9(n int) {
	/*
			Input Format: N = 6
		Result:
		     *
		    ***
		   *****
		  *******
		 *********
		***********
		***********
		 *********
		  *******
		   *****
		    ***
		     *
				considering 0 as index

				then ((n-i)*2)-1 or 2*n -( 2i +1)  stars = 9 stars
				and the i spaces

	*/
	for i := 0; i < n; i++ {

		//printing spaces
		for j := 0; j < n-i-1; j++ {
			fmt.Printf("%s", " ")
		}
		//printing stars
		for j := 0; j < (2*i)+1; j++ {
			fmt.Printf("%s", "*")
		}

		//printing spaces
		for j := 0; j < n-i-1; j++ {
			fmt.Printf("%s", " ")
		}
		fmt.Println()
	}

	for i := 0; i < n; i++ {

		//printing spaces
		for j := 0; j < i; j++ {
			fmt.Printf("%s", " ")
		}
		//printing stars
		for j := 0; j < 2*n-(2*i+1); j++ {
			fmt.Printf("%s", "*")
		}

		//printing spaces
		for j := 0; j < i; j++ {
			fmt.Printf("%s", " ")
		}
		fmt.Println()
	}

}

func Pattern10(n int) {
	/*
					Input Format: N = 6
		Result:
		     *
		     **
		     ***
		     ****
		     *****
		     ******
		     *****
		     ****
		     ***
		     **
		     *


	*/

	for i := 1; i <= (2*n)-1; i++ {
		stars := i
		if i > n {
			stars = (2 * n) - i
		}
		for j := 1; j <= stars; j++ {
			fmt.Printf("%s", "* ")
		}
		fmt.Println()
	}

}
