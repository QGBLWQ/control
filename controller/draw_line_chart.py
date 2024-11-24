# draw_line_chart.py
import sys
import json
import matplotlib.pyplot as plt

def draw_line_chart(title, data):
    x = [item['x'] for item in data]
    y = [item['y'] for item in data]

    plt.plot(x, y, marker='o')
    plt.title(title)
    plt.xlabel('X-axis')
    plt.ylabel('Y-axis')
    plt.savefig('line_chart.png')

if __name__ == "__main__":
    params = json.loads(sys.argv[1])
    draw_line_chart(params['title'], params['data'])