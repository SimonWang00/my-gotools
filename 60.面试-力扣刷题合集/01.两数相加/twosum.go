package twosum

/**
 * @Author: SimonWang00
 * @Description:
 * @File:  twosum.go
 * @Version: 1.0.0
 * @Date: 2021/3/21 13:44
 */


/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/


func twoSum(arr []int, target int) []int {
	m := make(map[int]int)
	for k, v := range arr {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
