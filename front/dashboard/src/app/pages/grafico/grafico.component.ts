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
  selectedGroupBy?: string
  selectedDocType?: string
  selectedUserId?: string = "0"
  selectedStartDate?: string
  selectedEndDate?: string

  async ngAfterContentChecked() {
    let changed = false
    this.filtroButtons = document.getElementById('filtroDiv')?.getElementsByTagName('button')
    this.dateInputs = document.getElementById('filtroDiv')?.getElementsByTagName('input')
    this.groupByButtons = document.getElementById('groupDiv')?.getElementsByTagName('button')

    for (let button of this.filtroButtons) {
      let parentId = button.parentElement.parentElement.id
      if (button.classList == 'selected') {
        if (parentId == 'docDiv') {
          if (button.value != this.selectedDocType) {
            this.selectedDocType = button.value
            changed = true;
          }
        } else if (parentId == 'userDiv') {
          if (button.value != this.selectedUserId) {
            this.selectedUserId = button.value
            changed = true;
          }
        }
      }
    }

    for (let date of this.dateInputs) {
      if (date.value != '') {
        if (date.id == 'dateStart') {
          if (date.value != this.selectedStartDate) {
            this.selectedStartDate = date.value
            changed = true
          }
        } else if (date.id == 'dateEnd') {
          if (date.value != this.selectedEndDate) {
            this.selectedEndDate = date.value
            changed = true
          }
        }
      }
    }

    for (let button of this.groupByButtons) {
      if (button.classList == 'selected') {
        if (button.value != this.selectedGroupBy) {
          this.selectedGroupBy = button.value
          changed = true
        }
      }
    }

    if (changed) {
      let apiRes = await getCount(this.selectedGroupBy || "user_id", this.selectedUserId || "0",  this.selectedDocType || null, this.selectedStartDate || null, this.selectedEndDate || null)
      console.log(apiRes)
    }
  }
}
