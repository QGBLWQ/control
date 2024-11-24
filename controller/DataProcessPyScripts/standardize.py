import sys
import json
import numpy as np

def standardize_data(data):

    mean = np.mean(data)
    std_dev = np.std(data)


    standardized = [(x - mean) / std_dev for x in data]
    return standardized

if __name__ == "__main__":
    input_data = sys.argv[1]
    data = json.loads(input_data)

    # 标准化数据
    result = standardize_data(data)

    # 输出标准化后的数据（以 JSON 格式）
    print(json.dumps(result))
