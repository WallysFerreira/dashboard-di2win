import { Component } from '@angular/core';
import { Chart, registerables } from 'chart.js';

@Component({
  selector: 'app-chart',
  templateUrl: './chart.component.html',
  styleUrls: ['./chart.component.css']
})

export class ChartComponent {
  public chart: any
  buttonsGroup: any

  createChart(): void {
    this.chart = new Chart('chartCanvas', {
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

  ngOnInit(): void {
    this.createChart()
  }

  ngAfterContentChecked() {
    this.buttonsGroup = document.getElementById('groupDiv')?.getElementsByTagName('button')
  }

  setSelected(id: any) {
    this.clearSelectedButton()
    
    document.getElementById(id)?.classList.add('selected')
  }

  clearSelectedButton() {
    for (let button of this.buttonsGroup) {
      button.classList.remove('selected')
    }
  }
}

Chart.register(...registerables)