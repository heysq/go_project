package main

func mapTest() {
    m := make(map[int]int)
    // 开启一段并发代码
    go func() {
        // 不停地对map进行写入
        for {
            m[1] = 1
        }
    }()
    // 开启一段并发代码
    go func() {
        // 不停地对map进行读取
        for {
            _ = m[1]
        }
    }()
    // 无限循环, 让并发程序在后台执行

}
