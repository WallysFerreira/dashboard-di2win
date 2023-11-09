import { Component } from '@angular/core';

@Component({
  selector: 'app-grafico',
  templateUrl: './grafico.component.html',
  styleUrls: ['./grafico.component.css']
})

export class GraficoComponent {
  filtroButtons: any
  dateInputs: any
  selectedDocType?: string
  selectedUserId?: string
  selectedStartDate?: string
  selectedEndDate?: string

  ngAfterContentChecked() {
    this.filtroButtons = document.getElementById('filtroDiv')?.getElementsByTagName('button')
    this.dateInputs = document.getElementById('filtroDiv')?.getElementsByTagName('input')

    for (let button of this.filtroButtons) {
      let parentId = button.parentElement.parentElement.id
      if (button.classList == 'selected') {
        if (parentId == 'docDiv') {
          this.selectedDocType = button.value
          console.log('Documento selecionado:', this.selectedDocType)
        } else if (parentId == 'userDiv') {
          this.selectedUserId = button.value
          console.log('Empresa selecionada:', this.selectedUserId)
        }
      }
    }

    for (let date of this.dateInputs) {
      if (date.value != '') {
        if (date.id == 'dateStart') {
          this.selectedStartDate = date.value
          console.log('Data inicio', this.selectedStartDate)
        } else if (date.id == 'dateEnd') {
          this.selectedEndDate = date.value
          console.log('Data fim', this.selectedEndDate)
        }
      }
    }
  }
}
