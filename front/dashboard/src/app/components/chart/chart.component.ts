import { Component } from '@angular/core';
import { Chart } from 'chart.js';

@Component({
  selector: 'app-chart',
  templateUrl: './chart.component.html',
  styleUrls: ['./chart.component.css']
})
export class ChartComponent {
  public chart: any

  createChart(): void {
    new Chart('chartDiv', {
      type: 'bar',
      data: {
        labels: ['Red', 'Blue'],
        datasets: [{
          label: '# of votes',
          data: [12, 19],
          borderWidth: 1
        }]
      },
      options: {
        scales: {
          y: {
            beginAtZero: true
          }
        }
      }
    })
  }
}
