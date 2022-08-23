package myleetcode

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
