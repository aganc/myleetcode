package main

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

// 20. 有效的括号
// https://leetcode.cn/problems/valid-parentheses/
func isValid(s string) bool {
	queue := make([]byte, 0)
	for i := range s {
		if len(queue) > 0 && s[i] == ')' {
			if queue[len(queue)-1] != '(' {
				return false
			}
			queue = queue[:len(queue)-1]
		} else if len(queue) > 0 && s[i] == ']' {
			if queue[len(queue)-1] != '[' {
				return false
			}
			queue = queue[:len(queue)-1]

		} else if len(queue) > 0 && s[i] == '}' {
			if queue[len(queue)-1] != '{' {
				return false
			}
			queue = queue[:len(queue)-1]
		} else {
			queue = append(queue, s[i])
		}
	}
	if len(queue) != 0 {
		return false
	}
	return true
}

// 21. 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val: 0,
	}
	cur := dummy
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}
		if list1.Val > list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

// 22. 括号生成
// https://leetcode.cn/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	res := make([]string, 0)

	var tracback func(string, int, int)
	tracback = func(path string, left, right int) {
		if len(path) == 2*n {
			res = append(res, path)
			return
		}

		if left < n {
			tracback(path+"(", left+1, right)
		}
		if right < left {
			tracback(path+")", left, right+1)
		}
	}
	tracback("", 0, 0)
	return res
}

// 23. 合并K个升序链表
// https://leetcode.cn/problems/merge-k-sorted-lists/submissions/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 0 {
		return nil
	}

	newlist := lists[0]
	for i := 1; i < len(lists); i++ {
		newlist = merge2Lists(newlist, lists[i])
	}
	return newlist
}
func merge2Lists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}
	if list1 == nil && list2 != nil {
		cur.Next = list2
	}
	if list2 == nil && list1 != nil {
		cur.Next = list1
	}
	return dummy.Next
}

// 31. 下一个排列
// https://leetcode.cn/problems/next-permutation/
func nextPermutation(nums []int) {
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		for k >= 0 && nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	i, j = i+1, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// 32. 最长有效括号
// https://leetcode.cn/problems/longest-valid-parentheses/?favorite=2cktkvj
func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}
	res := 0
	// 动态规划
	dp := make([]int, len(s))
	for i := range s {
		if i == 0 {
			continue
		}
		// s[i] = '(' dp[i] = 0
		// s[i] = ')'
		if s[i] == ')' && s[i-1] == '(' {
			if i-2 >= 0 {
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2
			}
		}

		if s[i] == ')' && s[i-1] == ')' {
			// 最左边对应'('，则dp[i] = dp[i-1] + 2
			if i-dp[i-1]-1 >= 0 {
				if s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2
					// 考虑...(())之前的
					if i-dp[i-1]-2 >= 0 {
						dp[i] += dp[i-dp[i-1]-2]
					}
				}
			}
		}

		if dp[i] > res {
			res = dp[i]
		}
	}
	return res

}

// 33. 搜索旋转排序数组
// https://leetcode.cn/problems/search-in-rotated-sorted-array/description/?favorite=2cktkvj&orderBy=most_relevant
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 二分法
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/?favorite=2cktkvj&orderBy=most_relevant
func searchRange(nums []int, target int) []int {
	// 二分找左边界
	start := searchLeftRange(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	// 二分找右边界
	last := searchRightRange(nums, target)
	return []int{start, last - 1}
}

func searchLeftRange(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if target <= nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
func searchRightRange(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target >= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// 39. 组合总和
// https://leetcode.cn/problems/combination-sum/?favorite=2cktkvj&orderBy=most_relevant
func combinationSum(candidates []int, target int) [][]int {

	res := make([][]int, 0)
	path := make([]int, 0)

	var traceback func(start int)
	traceback = func(start int) {
		if sum1(path) == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		if sum(path) > target {
			return
		}

		for k := start; k < len(candidates); k++ {
			path = append(path, candidates[k])
			traceback(k)
			path = path[:len(path)-1]
		}
	}
	traceback(0)
	return res
}
func sum1(path []int) int {
	s := 0
	for i := range path {
		s += path[i]
	}
	return s
}

// 40. 组合总和 II
// https://leetcode.cn/problems/combination-sum-ii/
func combinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	res := make([][]int, 0)
	path := make([]int, 0)
	history := make([]bool, len(candidates))

	var traceback func(start int)
	traceback = func(start int) {
		if sum2(path) == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		if sum(path) > target {
			return
		}

		for k := start; k < len(candidates); k++ {
			if k > 0 && candidates[k] == candidates[k-1] && history[k-1] == false {
				continue
			}
			path = append(path, candidates[k])
			history[k] = true
			traceback(k + 1)
			path = path[:len(path)-1]
			history[k] = false
		}
	}
	traceback(0)
	return res
}
func sum2(path []int) int {
	s := 0
	for i := range path {
		s += path[i]
	}
	return s
}
