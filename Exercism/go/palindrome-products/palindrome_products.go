package palindrome

import (
    "errors"
    "strconv"
)

// Define Product type here.
type Product struct {
    val int
    Factorizations [][2]int
}

func isPalindrome(n int) bool {
    sn := strconv.Itoa(n)
    if len(sn) == 1 { // always a palindrome number
        return true
    } else { // supposed length >= 2
        for i,j:=0,len(sn)-1; i<j; i,j=i+1,j-1 {
            if sn[i] != sn[j] {
                return false
            }
        }
        return true
    }
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
        return Product{}, Product{}, errors.New("fmin > fmax")
    }
    visited := map[int][][2]int{}
    for i:=fmin; i<=fmax; i++ {
        for j:=fmin; j<=fmax; j++ {
            p := i * j
            if !isPalindrome(p) {
                continue
            }
            if len(visited[p]) == 0 {
                visited[p] = append(visited[p], [2]int{i, j})
            } else {
                found := false
                for _,v := range visited[p] {
                    if v[0] == j && v[1] == i {
                        found = true
                        break
                    }
                }
                if !found {
                    visited[p] = append(visited[p], [2]int{i, j})
                }
            }
        }
    }
	if len(visited) == 0 {
        return Product{}, Product{}, errors.New("no palindromes...")
    }
    AllPalindromeProduct := []Product{}
    for k,v := range visited {
        AllPalindromeProduct = append(AllPalindromeProduct, Product{val:k, Factorizations: v})
    }
    pmin, pmax := AllPalindromeProduct[0], AllPalindromeProduct[0]
    for _,product := range AllPalindromeProduct {
        if pmin.val > product.val {
            pmin = product
        }
        if pmax.val < product.val {
            pmax = product
        }
    }
    return pmin, pmax, nil
}