import sys
import json
import numpy as np
from scipy.stats import kurtosis, skew

def calculate_statistics(data):
    stats = {}
    data = np.array(data)

    # 计算描述性统计
    stats["count"] = len(data)
    stats["max"] = np.max(data)
    stats["min"] = np.min(data)
    stats["mean"] = np.mean(data)
    stats["std_dev"] = np.std(data, ddof=1)  # 样本标准差
    stats["median"] = np.median(data)
    stats["variance"] = np.var(data, ddof=1)  # 样本方差
    stats["kurtosis"] = kurtosis(data)
    stats["skewness"] = skew(data)
    stats["cv"] = stats["std_dev"] / stats["mean"] if stats["mean"] != 0 else float('nan')  # 变异系数

    return stats

try:
    # 从命令行接收 JSON 数据
    input_data = json.loads(sys.argv[1])

    # 获取输入数据
    data = input_data.get("data", [])

    # 计算统计信息
    stats = calculate_statistics(data)
    
    # 输出统计结果
    print(json.dumps(stats))
except Exception as e:
    # 如果出错，输出错误信息
    print(json.dumps({"error": str(e)}))
