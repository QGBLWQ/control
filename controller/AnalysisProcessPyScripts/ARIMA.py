import sys
import json
import warnings
from statsmodels.tsa.arima.model import ARIMA

# 忽略 warnings
warnings.filterwarnings("ignore")

try:
    input_data = json.loads(sys.argv[1])
    data = input_data.get("data", [])
    p = input_data.get("p", 1)
    d = input_data.get("d", 1)
    q = input_data.get("q", 1)
    forecast_steps = input_data.get("forecast_steps", 5)

    # 创建并拟合 ARIMA 模型
    model = ARIMA(data, order=(p, d, q))
    model_fit = model.fit()
    forecast = model_fit.forecast(steps=forecast_steps)
    predictions = forecast.tolist()
    
    # 输出预测结果
    print(json.dumps(predictions))
except Exception as e:
    # 如果出错，输出错误信息
    print(json.dumps({"error": str(e)}))
