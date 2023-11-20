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
  @Input() label: any
  @Input() data: any
  setLabel: any = []
  setData: any = []

  createChart(): void {
    this.chart = new Chart('chartCanvas', {
      type: 'bar',
      data: {
        labels: this.setLabel,
        datasets: [{
          label: '# of votes',
          data: this.setData,
          borderWidth: 1
        }]
      },
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

    if (JSON.stringify(this.data) !== JSON.stringify(this.setData)) {
      this.setLabel = this.label
      this.setData = this.data

      this.chart.destroy()
      this.createChart()
    }
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