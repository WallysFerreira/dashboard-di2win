import { Component, EventEmitter, Input, Output } from '@angular/core';
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
  @Input() title: any
  @Input() divWidth: any
  @Input() chartType: any = 'bar'
  @Input() data: any
  @Input() groupingButtons: any
  @Output() exchangedGroupby = new EventEmitter<string>()
  setData: any = {
    labels: [],
    datasets: [{
      label: 'Paginas processadas',
      data: [0]
    }]
  }
  imageUrl?: string

  async createChart(): Promise<void> {
    let canvas = <HTMLCanvasElement>document.getElementById(this.canvasID)
    this.ctx = canvas?.getContext('2d')
    let options = {
      animation: {
        onComplete: () => {
          this.saveImageUrl()
        }
      }
    }

    if (this.title != undefined) {
      options = Object.assign(options, {
        plugins: {
          title: {
            display: true,
            text: this.title
          }
        }
      })
    }

    if (this.chartType != 'doughnut') {
      options = Object.assign(options, {
        scales: {
          y: {
            beginAtZero: true,
          }
        }
      })
    }

    if (this.ctx) {
      this.chart = new Chart(this.ctx, {
        type: this.chartType,
        data: this.setData,
        options: options,
        plugins: [{
          id: 'custom_canvas_background_color',
          beforeDraw: (chart) => {
            const ctxA = chart.canvas.getContext('2d');
            ctxA!.save();
            ctxA!.globalCompositeOperation = 'destination-over';
            ctxA!.fillStyle = 'white';
            ctxA?.fillRect(0, 0, chart.canvas.width, chart.canvas.height);
            ctxA?.restore();
          }
        }]
      })
    }
  }

  ngOnInit(): void {
    let possibleID = 'chartCanvas' + Math.floor(Math.random() * 50)
    this.canvasID = possibleID
  }

  ngAfterContentChecked() {
    this.buttonsGroup = document.getElementById('groupDiv')?.getElementsByTagName('button')

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

  saveImageUrl() {
    let canvas = <HTMLCanvasElement>document.getElementById(this.canvasID)

    this.imageUrl = canvas.toDataURL()
  }
  groupbyClicked(value:any){
    this.exchangedGroupby.emit(value)
  }
}

Chart.register(...registerables)
