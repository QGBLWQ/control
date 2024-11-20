import sys
import json
import numpy as np

def calculate_correlation(data):
    # 将数据转换为 numpy 数组
    data_array = np.array(data)

    # 计算相关系数矩阵
    correlation_matrix = np.corrcoef(data_array, rowvar=False)

    # 将结果转换为列表格式
    return correlation_matrix.tolist()

if __name__ == "__main__":
    # 从命令行参数获取输入数据
    input_data = sys.argv[1]  # 假设输入是一个 JSON 格式的字符串
    data = json.loads(input_data)

    # 计算相关系数
    result = calculate_correlation(data)

    # 输出相关系数矩阵（以 JSON 格式）
    print(json.dumps(result))
