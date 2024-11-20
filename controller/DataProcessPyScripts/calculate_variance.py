import sys
import json
import numpy as np

def calculate_variance(data):
    # 将数据转换为 numpy 数组
    data_array = np.array(data)

    # 计算每列的方差
    variance = np.var(data_array, axis=0)

    # 将结果转换为列表格式
    return variance.tolist()

if __name__ == "__main__":
    # 从命令行参数获取输入数据
    input_data = sys.argv[1]  # 假设输入是一个 JSON 格式的字符串
    data = json.loads(input_data)

    # 计算方差
    result = calculate_variance(data)

    # 输出方差列表（以 JSON 格式）
    print(json.dumps(result))
