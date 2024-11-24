# draw_line_bar_mixed_chart.py
import sys
import json
import matplotlib.pyplot as plt

def draw_line_bar_mixed_chart(title, line_data, bar_data):
    fig, ax1 = plt.subplots()

    # Plot bar chart
    labels = [item['label'] for item in bar_data]
    values = [item['value'] for item in bar_data]
    ax1.bar(labels, values, color='b', alpha=0.6)
    ax1.set_xlabel('Labels')
    ax1.set_ylabel('Bar Values', color='b')

    # Plot line chart
    ax2 = ax1.twinx()
    x = [item['x'] for item in line_data]
    y = [item['y'] for item in line_data]
    ax2.plot(x, y, color='r', marker='o')
    ax2.set_ylabel('Line Values', color='r')

    plt.title(title)
    plt.savefig('line_bar_mixed_chart.png')

if __name__ == "__main__":
    params = json.loads(sys.argv[1])
    draw_line_bar_mixed_chart(params['title'], params['line_data'], params['bar_data'])