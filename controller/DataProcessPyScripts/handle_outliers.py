import sys
import json
import numpy as np

def handle_outliers(data):
    # 计算数据的均值和标准差
    mean = np.mean(data)
    std_dev = np.std(data)

    # 设置 Z-score 阈值，例如 3 标准差之外的值为异常值
    threshold = 3
    processed_data = [x if abs((x - mean) / std_dev) < threshold else None for x in data]

    return processed_data

if __name__ == "__main__":
    # 从命令行参数获取输入数据
    input_data = sys.argv[1]  # 假设输入是一个 JSON 格式的字符串
    data = json.loads(input_data)

    # 处理异常值
    result = handle_outliers(data)

    # 输出处理后的数据（以 JSON 格式）
    print(json.dumps(result))
