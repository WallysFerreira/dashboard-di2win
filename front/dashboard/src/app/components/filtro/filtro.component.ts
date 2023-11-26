import { Component, ViewChild, ElementRef } from '@angular/core';
import { getCount, getUsers } from 'src/app/app.component';

@Component({
  selector: 'app-filtro',
  templateUrl: './filtro.component.html',
  styleUrls: ['./filtro.component.css']
})

export class FiltroComponent {
  empresas: any
  tipos_documento: any
  buttonsDoc: any
  buttonsEmp: any
  segmentos: any

  async getSegmentos(){
    let res = await getUsers();
    this.empresas = res.data.user;
  
    this.segmentos = []
    
    for (let empresa of this.empresas) {
        this.segmentos.push(empresa.segment)
    }
  }
  
  async ngOnInit() {
    await this.getSegmentos();
    let res = await getUsers()
    this.empresas = res.data.user

    this.tipos_documento = [
      {
        name: "Faturamento",
        value: "FATURAMENTO"
      },
      {
        name: "CNH",
        value: "CNH"
      },
      {
        name: "Posição Consolidada",
        value: "POSICAO_CONSOLIDADA"
      },
      {
        name: "Fatura de Energia",
        value: "FATURA_ENERGIA"
      },
      {
        name: "Contrato Social",
        value: "CONTRATO_SOCIAL"
      },
      {
        name: "Declaração IR",
        value: "DECLARACAO_IR"
      },
      {
        name: "Capa Serasa",
        value: "CAPA_SERASA"
      },
      {
        name: "Endividamento",
        value: "ENDIVIDAMENTO"
      },
      {
        name: "Comprovante de Residência",
        value: "COMPROVANTE_RESIDENCIA"
      },
      {
        name: "Balanço Patrimonial",
        value: "BALANCO_PATRIMONIAL"
      },
      {
        name: "Recibo",
        value: "RECIBOS"
      }
    ]
  }

  ngAfterContentChecked() {
    this.buttonsDoc = document.getElementById('docDiv')?.getElementsByTagName('button')
    this.buttonsEmp = document.getElementById('userDiv')?.getElementsByTagName('button')
  }

  select(event: any, div: any) {
    const hasClass = event.target.classList.contains('selected')

    if (hasClass) {
      this.unselect(event)
    } else {
      this.setSelected(event, div)
    }
  }

  unselect (event: any) {
    event.target.classList.remove('selected')
  }

  setSelected(event: any, div: any) {
    this.clearSelectedButton(div)

    event.target.classList.add('selected')
  }

  clearSelectedButton(div: any) {
    let buttons

    if (div == 'docDiv') {
      buttons = this.buttonsDoc
    } else if (div == 'userDiv') {
      buttons = this.buttonsEmp
    }

    for (let i = 0; i < buttons.length; i++) {
      buttons[i].classList.remove('selected')
    }
  }
}
