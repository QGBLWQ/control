import sys
import json
import numpy as np
from scipy.stats import chi2_contingency

def calculate_chi_square(data):
    num_features = len(data[0])
    chi_square_matrix = np.zeros((num_features, num_features))

    # 逐对计算卡方检验值
    for i in range(num_features):
        for j in range(i + 1, num_features):
            # 构造列联表
            contingency_table = np.histogram2d(
                [row[i] for row in data],
                [row[j] for row in data],
                bins=(np.unique([row[i] for row in data]).size,
                      np.unique([row[j] for row in data]).size)
            )[0]
            # 计算卡方检验
            chi2, _, _, _ = chi2_contingency(contingency_table)
            chi_square_matrix[i, j] = chi2
            chi_square_matrix[j, i] = chi2  # 对称矩阵

    return chi_square_matrix.tolist()

if __name__ == "__main__":
    # 从命令行参数获取输入数据
    input_data = sys.argv[1]  # 假设输入是一个 JSON 格式的字符串
    data = json.loads(input_data)

    # 计算卡方检验矩阵
    result = calculate_chi_square(data)

    # 输出卡方检验矩阵（以 JSON 格式）
    print(json.dumps(result))
