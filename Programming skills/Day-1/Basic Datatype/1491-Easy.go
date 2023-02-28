/*
1491. Average Salary Excluding the Minimum and Maximum Salary
Easy
1.2K
132
Companies
You are given an array of unique integers salary where salary[i] is the salary of the ith employee.

Return the average salary of employees excluding the minimum and maximum salary. Answers within 10-5 of the actual answer will be accepted.

*/









func average(salary []int) float64 {
    count := 0;
    var sum float64;
    sum = 0.0;
    sort.Ints(salary)
    slice := salary[1:len(salary)-1];
    for _,value:=range slice{
        sum+=float64(value)
        count++
    }
    return sum/float64(count)
}