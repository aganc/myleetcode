package myleetcode

import "sort"

// 1. 两数之和
// https://leetcode.cn/problems/two-sum/
// 使用hashmap降低时间复杂度
func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, v := range nums {
		gap := target - v
		if index, ok := hashmap[gap]; ok {
			return []int{i, index}
		} else {
			hashmap[v] = i
		}
	}
	return []int{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2. 两数相加
// https://leetcode.cn/problems/add-two-numbers/
//
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10

		newNode := &ListNode{
			Val: sum,
		}
		cur.Next = newNode
		cur = cur.Next
	}
	if carry != 0 {
		cur.Next = &ListNode{
			Val: carry,
		}
	}
	return dummy.Next
}

// 3. 无重复字符的最长子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 熟悉map的用法 + 字符串遍历的是字节byte
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	dict := make(map[byte]bool)
	start, res := 0, 1
	for i := 0; i < len(s); i++ {
		for dict[s[i]] {
			delete(dict, s[start])
			start++
		}
		dict[s[i]] = true

		if i-start+1 > res {
			res = i - start + 1
		}
	}
	return res
}

// 4. 寻找两个正序数组的中位数
// https://leetcode.cn/problems/median-of-two-sorted-arrays/

// 5. 最长回文子串
// https://leetcode.cn/problems/longest-palindromic-substring/
// 动态规划思想
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	max := 1
	start := 0
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > max {
				max = j - i + 1
				start = i
			}
		}
	}
	return s[start : start+max]
}

// 10. 正则表达式匹配
// https://leetcode.cn/problems/regular-expression-matching/

// 11. 盛最多水的容器
// https://leetcode.cn/problems/container-with-most-water/
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		minh := height[l]
		if height[l] > height[r] {
			minh = height[r]
		}
		area := minh * (r - l)
		if area > res {
			res = area
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return res
}

// 15. 三数之和
// https://leetcode.cn/problems/3sum/
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return res
}

// 17. 电话号码的字母组合
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if digits == "" {
		return res
	}
	numsmap := [10]string{
		"", "", "abc", "def",
		"ghi", "jkl", "mno",
		"pqrs", "tuv", "wxyz",
	}
	var trackback func(string, int)
	trackback = func(tmp string, start int) {
		if len(tmp) == len(digits) {
			res = append(res, tmp)
			return
		}
		index := digits[start] - '0'
		nums := numsmap[index]
		for i := 0; i < len(nums); i++ {
			tmp += string(nums[i])
			trackback(tmp, start+1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	trackback("", 0)
	return res
}

// 19. 删除链表的倒数第 N 个结点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := &ListNode{
		Val:  0,
		Next: head,
	}
	tmp := cur
	dummy := cur
	l := 0
	for tmp.Next != nil {
		l++
		tmp = tmp.Next
	}
	index := l - n
	for index > 0 {
		cur = cur.Next
		index--
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}
