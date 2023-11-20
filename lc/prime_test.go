package lc

import (
	"fmt"
	"testing"
)

// primeFactors 函數對一個整數 n 進行質因數分解
func primeFactors(n int) []int {
	var factors []int

	// 先處理所有的 2 因數
	// 如果 n 是偶數，2 會是它的一個因數
	for n%2 == 0 {
		factors = append(factors, 2) // 將 2 加入因數列表
		n = n / 2                    // 除以 2，減少 n 的值
	}

	// 處理奇數因數
	// 從 3 開始，只需檢查到 sqrt(n) 就可以了
	for i := 3; i*i <= n; i = i + 2 {
		// 檢查 i 是否是 n 的因數
		for n%i == 0 {
			factors = append(factors, i) // 將 i 加入因數列表
			n = n / i                    // 除以 i，減少 n 的值
		}
	}

	// 如果 n 此時仍大於 2，則它自身是一個質數且是自己的一個因數
	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}

// isPrime 函數檢查一個整數 n 是否為質數
func isPrime(n int) bool {
	// 處理小於 2 的情況
	if n < 2 {
		return false
	}

	// 只需檢查到 i*i <= n 即可
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false // 如果找到其他因數，則不是質數
		}
	}
	return true // 沒有找到其他因數，是質數
}

// TestPrimeFunctions 測試 primeFactors 和 isPrime 函數
func TestPrimeFunctions(t *testing.T) {
	for i := 2; i <= 1000; i++ {
		factors := primeFactors(i)
		isPrimeResult := isPrime(i)

		if isPrimeResult && len(factors) != 1 {
			t.Errorf("不一致: %d 被認為是質數，但 primeFactors 返回多個因數：%v", i, factors)
		}

		if !isPrimeResult && len(factors) == 1 {
			t.Errorf("不一致: %d 被認為不是質數，但 primeFactors 只返回一個因數：%v", i, factors)
		}

		if isPrimeResult && factors[0] != i {
			t.Errorf("錯誤: %d 被認為是質數，但 primeFactors 返回的因數不是自己：%v", i, factors)
		}

		if isPrimeResult && len(factors) == 1 && factors[0] == i {
			fmt.Println(i, factors)
		}
		if !isPrimeResult && len(factors) > 1 {
			fmt.Println(i, factors)
		}
	}
}
