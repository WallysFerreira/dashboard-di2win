import { Component } from '@angular/core';

@Component({
  selector: 'app-grafico',
  templateUrl: './grafico.component.html',
  styleUrls: ['./grafico.component.css']
})

export class GraficoComponent {
  filtroButtons: any
  selectedDocType?: string
  selectedUserId?: string

  ngAfterContentChecked() {
    this.filtroButtons = document.getElementById('filtroDiv')?.getElementsByTagName('button')
    
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
  }
}
