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
      let idPai = button.parentElement.parentElement.id
      if (button.classList == 'selected') {
        if (idPai == 'docDiv') {
          console.log('Documento selecionado:', button.value)
        } else if (idPai == 'userDiv') {
          console.log('Empresa selecionada:', button.value)
        }
      }
    }
  }
}
