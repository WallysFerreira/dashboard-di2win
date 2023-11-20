import { Component } from '@angular/core';
import { getCount } from 'src/app/app.component';

@Component({
  selector: 'app-grafico',
  templateUrl: './grafico.component.html',
  styleUrls: ['./grafico.component.css']
})

export class GraficoComponent {
  filtroButtons: any
  dateInputs: any
  groupByButtons: any
  changed = false
  docWasSelected = false
  userIdWasSelected = false
  selectedGroupBy?: string
  selectedDocType?: string
  selectedUserId?: string = "0"
  selectedStartDate?: string
  selectedEndDate?: string
  labelData: any = []
  valueData: any = []
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

  async ngAfterContentChecked() {
    this.changed = false
    this.docWasSelected = false
    this.userIdWasSelected = false

    for (let button of this.filtroButtons) {
      let parentId = button.parentElement.parentElement.id
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
        }
      }
    }

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
      let apiRes = await getCount(this.selectedGroupBy || "user_id", false, this.selectedUserId || "0",  this.selectedDocType || null, this.selectedStartDate || null, this.selectedEndDate || null)
      this.labelData = []
      this.valueData = []

      for (let count of apiRes.data.count) {
        if (this.selectedGroupBy == "EXTRACT(month FROM created_at::date)") {
          let month = new Date(2000, count.name - 1, 1).toLocaleString('default', { month: 'long' })
          count.name = month
        }

        this.labelData.push(count.name)
        this.valueData.push(count.value)
      }
      
      this.entireDataset = {
        labels: this.labelData,
        datasets: []
      }
      this.entireDataset.datasets.push({label: 'Paginas processadas', data: this.valueData})

      console.log(this.entireDataset)
    }
  }
}
