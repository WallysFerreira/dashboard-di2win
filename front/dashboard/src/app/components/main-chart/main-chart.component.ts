import { Component, Input } from '@angular/core';
import { Chart, registerables } from 'chart.js';

@Component({
  selector: 'app-main-chart',
  templateUrl: './main-chart.component.html',
  styleUrls: ['./main-chart.component.css']
})

export class MainChartComponent {
  public chart: any
  buttonsGroup: any
  @Input() data: any
  @Input() groupingButtons: any
  setData: any = {
    labels: [],
    datasets: [{
      label: 'Paginas processadas',
      data: [0]
    }]
  }

  createChart(): void {
    this.chart = new Chart('chartCanvas', {
      type: 'bar',
      data: this.setData,
      options: {
        scales: {
          y: {
            beginAtZero: true,
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

    if (this.data !== this.setData) {
      this.setData = this.data

      this.chart.destroy()
      this.createChart()
    }
  }


  setSelected(id: any) {
    this.clearSelectedButton()
    
    document.getElementById(id + 'Button')?.classList.add('selected')
  }

  clearSelectedButton() {
    for (let button of this.buttonsGroup) {
      button.classList.remove('selected')
    }
  }
}

Chart.register(...registerables)