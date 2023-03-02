/*//191. Number of 1 Bits
Easy
5.1K
1.1K
Companies
Write a function that takes the binary representation of an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).

/*
//initial solution/

func hammingWeight(num uint32) int {
    binaryStr := strconv.FormatInt(int64(num), 2)

    count:=0;
    for _, c := range binaryStr {
        if c == '1' {
            count++
        }
    }

    return count
} */

//better solution using bitwise operator 
func hammingWeight(num uint32) int {

    n:=int(num)
    count:=0;
    helper:=1;
    for i:=0;i<32;i++{
        if (helper & n) != 0{
            count++
        }
        helper=helper<<1;
    }
    return count
}