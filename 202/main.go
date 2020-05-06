func isHappy(n int) bool {
    hash := make(map[int]bool)      // 开一个 map 判断是否循环
    for n!=1 {
        if _,ok := hash[n]; ok {    // 出现了循环，证明不是快乐数
            return false
        }
        hash[n] = true
        next := 0
        for n!=0 {                  // 计算下一个数字
            next += (n%10) * (n%10)
            n /= 10
        }
        n = next
    }
    return true
}
