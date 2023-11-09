import { Component } from '@angular/core';
import { getCount, getUsers } from 'src/app/app.component';

@Component({
  selector: 'app-filtro',
  templateUrl: './filtro.component.html',
  styleUrls: ['./filtro.component.css']
})

export class FiltroComponent {
  empresas: any
  tipos_documento: any

  async ngOnInit() {
    let res = await getUsers()
    this.empresas = res.data.user
    console.log("Empresas", this.empresas)

    this.tipos_documento = [
      "Faturamento",
      "CNH",
      "Posição Consolidada",
      "Fatura de Energia",
      "Contrato Social",
      "Declaração IR",
      "Capa Serasa",
      "Endividamento",
      "Comprovante de Residência",
      "Balanço Patrimonial",
      "Recibo",
    ]
  }
}
