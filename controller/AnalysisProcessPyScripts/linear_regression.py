import sys
import json
from sklearn.linear_model import LinearRegression

# 从命令行接收 JSON 数据
input_data = json.loads(sys.argv[1])

# 获取 x_data, y_data 和 predict_data
x_data = input_data.get("x_data", [])
y_data = input_data.get("y_data", [])
predict_data = input_data.get("predict_data", [])

# 数据格式转换，将 x_data 转换为 [[x1], [x2], ...] 的二维格式
X = [[x] for x in x_data]
y = y_data
X_predict = [[x] for x in predict_data]

# 线性回归模型训练和预测
model = LinearRegression()
model.fit(X, y)
predictions = model.predict(X_predict).tolist()

# 输出预测结果
print(json.dumps(predictions))
