import { Component, ViewChild } from '@angular/core';
import { getCount } from 'src/app/app.component';
import { MainChartComponent } from 'src/app/components/main-chart/main-chart.component';

@Component({
  selector: 'app-grafico',
  templateUrl: './grafico.component.html',
  styleUrls: ['./grafico.component.css']
})

export class GraficoComponent {
  @ViewChild(MainChartComponent) chartChild?: MainChartComponent;
  filtroButtons: any
  dateInputs: any
  groupByButtons: any
  changed = false
  docWasSelected = false
  userIdWasSelected = false
  segmentWasSelected = false
  selectedChartType?: string = 'bar'
  selectedGroupBy?: string
  selectedDocType?: string
  selectedUserId?: string = "0"
  selectedSegment?: string
  selectedStartDate?: string
  selectedEndDate?: string
  labelData: any = []
  valueData: any = []
  groupingButtons: any = [
    {
      name: "userGroup",
      text: "Empresa",
      value: "users.name"
    },
    {
      name: "docGroup",
      text: "Documento",
      value: "doc_type"
    },
    {
      name: "monthGroup",
      text: "Mês",
      value: "EXTRACT(month FROM created_at::date)"
    },
    {
      name: "segmentGroup",
      text: "Segmento",
      value: "users.segment"
    }
  ]
  entireDataset: any = {
    labels: [],
    datasets: [{
      label: 'Paginas processadas',
      data: [0]
    }]
  }

  ngOnInit() {
    this.filtroButtons = document.getElementById('filtroDiv')?.getElementsByTagName('button')
    this.dateInputs = document.getElementById('filtroDiv')?.getElementsByTagName('input')
    this.groupByButtons = document.getElementById('groupDiv')?.getElementsByTagName('button')

    for (let date of this.dateInputs) {
      date.addEventListener('change', () => {
        if (date.value[0] != '0') {
          this.changeDate()
        }
      })
    }
  }

  changeChartType(event: any) {
    if (event.target.value != this.selectedChartType) {
      this.changed = true
      this.selectedChartType = event.target.value
    }

    this.updateData()
  }

  async ngAfterContentChecked() {
    this.changed = false
    this.docWasSelected = false
    this.userIdWasSelected = false
    this.segmentWasSelected = false

    for (let button of this.filtroButtons) {
      let parentId = button.parentElement.parentElement.id

      // This part checks if a button is selected and if it's value is different from the one previously selected
      if (button.classList == 'selected') {
        if (parentId == 'docDiv') {
          this.docWasSelected = true

          if (button.value != this.selectedDocType) {
            this.selectedDocType = button.value
            this.changed = true;
          }
        } else if (parentId == 'userDiv') {
          this.userIdWasSelected = true

          if (button.value != this.selectedUserId) {
            this.selectedUserId = button.value
            this.changed = true;
          }
        } else if (parentId == 'segmentoDiv') {
          this.segmentWasSelected = true

          if (button.value != this.selectedSegment) {
            this.selectedSegment = button.value
            this.changed = true
          }
        }
      }
    }

    // This part checks if a filter that was previously selected has been unselected
    if (this.selectedDocType != undefined) {
      if (!this.docWasSelected) {
        this.selectedDocType = undefined
        this.changed = true
      }
    }

    if (this.selectedUserId != undefined) {
      if (!this.userIdWasSelected) {
        this.selectedUserId = undefined
        this.changed = true
      }
    }

    if (this.selectedSegment != undefined) {
      if (!this.segmentWasSelected) {
        this.selectedSegment = undefined
        this.changed = true
      }
    }

    for (let button of this.groupByButtons) {
      if (button.classList == 'selected') {
        if (button.value != this.selectedGroupBy) {
          this.selectedGroupBy = button.value
          this.changed = true
        }
      }
    }

    this.updateData()
  }

  changeDate() {
    for (let date of this.dateInputs) {
      if (date.value != '') {
        if (date.id == 'dateStart') {
          if (date.value != this.selectedStartDate) {
            this.selectedStartDate = date.value
            this.changed = true
            this.updateData()
          }
        } else if (date.id == 'dateEnd') {
          if (date.value != this.selectedEndDate) {
            this.selectedEndDate = date.value
            this.changed = true
            this.updateData()
          }
        }
      }
    }
  }

  async updateData() {
    if (this.changed) {
      let apiRes = await getCount(this.selectedGroupBy || "users.name", false, this.selectedUserId || "0", this.selectedSegment || null, this.selectedDocType || null, this.selectedStartDate || null, this.selectedEndDate || null)
      this.labelData = []
      this.valueData = []
      let counts = []

      for (let count of apiRes.data.count) {
        if (this.selectedGroupBy == "EXTRACT(month FROM created_at::date)") {
          let month = new Date(2000, count.name - 1, 1).toLocaleString('default', { month: 'long' })
          count.name = month
        }

        counts.push(count)
      }

      if (this.selectedChartType == 'doughnut') {
        counts.sort((a: any, b: any) => {
          let nameA = a.name.toUpperCase()
          let nameB = b.name.toUpperCase()

          if (nameA < nameB) return -1

          if (nameA > nameB) return 1

          return 0
        })
      }

      for (let count of counts) {
        this.labelData.push(count.name)
        this.valueData.push(count.value)
      }

      this.entireDataset = {
        labels: this.labelData,
        datasets: []
      }

      this.entireDataset.datasets.push({ label: 'Paginas processadas', data: this.valueData })
    }
  }
}
