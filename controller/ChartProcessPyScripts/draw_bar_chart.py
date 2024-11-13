# draw_bar_chart.py
import sys
import json
import matplotlib.pyplot as plt

def draw_bar_chart(title, data):
    labels = [item['label'] for item in data]
    values = [item['value'] for item in data]

    plt.bar(labels, values)
    plt.title(title)
    plt.xlabel('Labels')
    plt.ylabel('Values')
    plt.savefig('bar_chart.png')

if __name__ == "__main__":
    params = json.loads(sys.argv[1])
    draw_bar_chart(params['title'], params['data'])