import sys
import json
import numpy as np
from sklearn.neural_network import MLPRegressor
from sklearn.preprocessing import MinMaxScaler

def bp_neural_network_predict(data, forecast_steps, hidden_layer_size, max_iter, learning_rate_init):
    # 数据预处理：归一化
    scaler = MinMaxScaler(feature_range=(0, 1))
    data = np.array(data).reshape(-1, 1)
    scaled_data = scaler.fit_transform(data)

    # 准备训练数据（使用前 n-1 个数据预测第 n 个数据）
    X_train = []
    y_train = []
    for i in range(len(scaled_data) - 1):
        X_train.append(scaled_data[i])
        y_train.append(scaled_data[i + 1])
    X_train = np.array(X_train)
    y_train = np.array(y_train).ravel()

    # 定义并训练 BP 神经网络模型
    model = MLPRegressor(
        hidden_layer_sizes=(hidden_layer_size, hidden_layer_size),  # 两层隐藏层，每层节点数相同
        max_iter=max_iter,
        learning_rate_init=learning_rate_init,
        random_state=42
    )
    model.fit(X_train, y_train)

    # 预测未来数据
    predictions = []
    last_input = scaled_data[-1]
    for _ in range(forecast_steps):
        next_pred = model.predict(last_input.reshape(1, -1))
        predictions.append(next_pred[0])
        last_input = np.array([next_pred[0]])  # 滚动预测使用新预测的值

    # 反归一化预测结果
    predictions = scaler.inverse_transform(np.array(predictions).reshape(-1, 1)).flatten()
    return predictions.tolist()

try:
    # 从命令行接收 JSON 数据
    input_data = json.loads(sys.argv[1])

    # 获取输入数据和预测步数
    data = input_data.get("data", [])
    forecast_steps = input_data.get("forecast_steps", 5)
    hidden_layer_size = input_data.get("hidden_layers", 10)  # 隐藏层节点数
    max_iter = input_data.get("max_iter", 1000)
    learning_rate_init = input_data.get("learning_rate_init", 0.001)

    # 执行 BP 神经网络预测
    predictions = bp_neural_network_predict(
        data, forecast_steps, hidden_layer_size, max_iter, learning_rate_init
    )
    
    # 输出预测结果
    print(json.dumps(predictions))
except Exception as e:
    # 如果出错，输出错误信息
    print(json.dumps({"error": str(e)}))
