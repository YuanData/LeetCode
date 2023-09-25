def solution(N: int):
    # 第一步：將數字 N 轉換為二進制表示
    binary_representation = bin(N)[2:]

    # 初始化兩個變量，一個用於存儲當前的間隙長度，另一個用於存儲最大的間隙長度
    current_gap_length = 0
    max_gap_length = 0

    # 第二步：遍歷二進制表示中的每一位，找出所有的二進制間隙
    for bit in binary_representation:
        if bit == '0':
            # 如果當前位是 0，則增加當前間隙長度
            current_gap_length += 1
        else:
            # 如果當前位是 1，則檢查當前間隙長度是否大於最大間隙長度，如果是，則更新最大間隙長度
            if current_gap_length > max_gap_length:
                max_gap_length = current_gap_length
            # 重置當前間隙長度為 0
            current_gap_length = 0

    # 第三步：返回最大間隙長度
    return max_gap_length


if __name__ == '__main__':
    assert solution(1041) == 5
    assert solution(15) == 0
    assert solution(32) == 0
    assert solution(888) == 1
