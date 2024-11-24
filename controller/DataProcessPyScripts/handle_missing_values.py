import sys
import json
import numpy as np

def handle_missing_values(data):
    # 将 None 转换为 NaN，以便计算均值
    cleaned_data = [x if x is not None else np.nan for x in data]

    # 计算均值（忽略 NaN）
    mean_value = np.nanmean(cleaned_data)

    # 用均值填充缺失值
    filled_data = [x if x is not None else mean_value for x in data]

    return filled_data

if __name__ == "__main__":
    # 从命令行参数获取输入数据
    input_data = sys.argv[1]  # 假设输入是一个 JSON 格式的字符串
    data = json.loads(input_data)

    # 处理缺失值
    result = handle_missing_values(data)

    # 输出处理后的数据（以 JSON 格式）
    print(json.dumps(result))
