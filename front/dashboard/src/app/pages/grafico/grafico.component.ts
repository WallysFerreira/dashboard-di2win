import { Component, ViewChild } from '@angular/core';
import { ajeitarNome, getCount, getUsers } from 'src/app/app.component';
import { MainChartComponent } from 'src/app/components/main-chart/main-chart.component';

@Component({
  selector: 'app-grafico',
  templateUrl: './grafico.component.html',
  styleUrls: ['./grafico.component.css']
})

export class GraficoComponent {
  @ViewChild(MainChartComponent) chartChild?: MainChartComponent;
  filtroButtons: any
  filtroSelects: any
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
  userFilterData: any = []
  docFilterData: any = []
  segmentFilterData: any = []

  async ngOnInit() {
    this.filtroButtons = document.getElementById('filtrosDiv')?.getElementsByTagName('button')
    this.filtroSelects = document.getElementById('filtrosDiv')?.getElementsByTagName('select')
    this.dateInputs = document.getElementById('filtrosDiv')?.getElementsByTagName('input')
    this.groupByButtons = document.getElementById('groupDiv')?.getElementsByTagName('button')

    let users = await getUsers().then((res) => res.data.user)
    users.map((user: any) => {
      this.userFilterData.push({
        name: user.name,
        value: user.id
      })
    })

    let docs = await getCount("doc_type", true, "0", null, null, null, null).then((res) => res.data.count)
    docs.map((doc: any) => {
      this.docFilterData.push({
        name: ajeitarNome(doc.name),
        value: doc.name
      })
    })
    
    let segments = await getCount("users.segment", true, "0", null, null, null, null).then((res) => res.data.count)
    segments.map((segment: any) => {
      this.segmentFilterData.push({
        name: ajeitarNome(segment.name),
        value: segment.name
      })
    })
  }

  ngAfterViewInit() {
    for (let date of this.dateInputs) {
      date.addEventListener('change', () => {
        if (date.value[0] != '0') {
          this.changeDate()
        }
      })
    }
  }

  changeFilter(event: any) {
    let targetParentId = event.target.parentElement.parentElement.id
    let targetValue = event.target.value
    
    if (targetParentId.includes('Documento')) {
      if (this.selectedDocType != targetValue) {
        this.selectedDocType = targetValue
        this.changed = true
      }
    }

    if (targetParentId.includes('Cliente')) {
      if (this.selectedUserId != targetValue) {
        this.selectedUserId = targetValue
        this.changed = true
      }
    }

    if (targetParentId.includes('Segmento')) {
      if (this.selectedSegment != targetValue) {
        this.selectedSegment = targetValue
        this.changed = true
      }
    }

    this.updateData()
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

      this.entireDataset.datasets.push({ label: 'Paginas processadas', data: this.valueData, backgroundColor:['#e71c9a','#f52997','#ee47ac','#f682c5','#f5acd8','#f8b5dc']})
    }
  }
}
