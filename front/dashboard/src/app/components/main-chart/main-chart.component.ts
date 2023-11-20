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
  canvasID: any
  ctx: any
  @Input() divWidth: any
  @Input() chartType: any = 'bar'
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
    let canvas = <HTMLCanvasElement> document.getElementById(this.canvasID)
    this.ctx = canvas?.getContext('2d')

    if (this.ctx) {
      this.chart = new Chart(this.ctx, {
        type: this.chartType,
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
  }

  ngOnInit(): void {
    this.canvasID = 'chartCanvas' + Math.floor(Math.random() * 3)
  }

  ngAfterContentChecked() {
    this.buttonsGroup = document.getElementById('groupDiv')?.getElementsByTagName('button')
    console.log(this.ctx)

    if (this.data !== this.setData) {
      this.setData = this.data

      if (this.chart) this.chart.destroy()
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