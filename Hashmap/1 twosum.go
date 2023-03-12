/*1. Two Sum
Easy
44.2K
1.4K
Companies
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.*/


/*
initial solution without hashmap 

func twoSum(nums []int, target int) []int {
    var result []int
    for i:=0;i<len(nums);i++{
        for j:=i+1;j<len(nums);j++{
            if nums[i]+nums[j]==target{
                result=append(result,i,j) 
                return result
            }
        }
    }
    return nil
}


*/

func twoSum(nums []int, target int) []int {
    valueMap:=make(map[int]int)
    var result []int
    for i,num:=range nums{
        if complement,ok:=valueMap[target-num];ok{
              result = append(result, complement, i)
            return result
        }else{
            valueMap[num] = i
        }
    }
    return nil
}
