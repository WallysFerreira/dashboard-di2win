import { Component } from '@angular/core';

@Component({
  selector: 'app-grafico',
  templateUrl: './grafico.component.html',
  styleUrls: ['./grafico.component.css']
})

export class GraficoComponent {
  filtroButtons: any

  ngAfterContentChecked() {
    this.filtroButtons = document.getElementById('filtroDiv')?.getElementsByTagName('button')
    
    for (let button of this.filtroButtons) {
      let parentId = button.parentElement.parentElement.id
      if (button.classList == 'selected') {
        if (parentId == 'docDiv') {
          console.log('Documento selecionado:', button.value)
        } else if (parentId == 'userDiv') {
          console.log('Empresa selecionada:', button.value)
        }
      }
    }
  }
}
