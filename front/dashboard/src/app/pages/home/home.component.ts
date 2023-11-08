import { Component } from '@angular/core';
import { DestaqueComponent } from 'src/app/components/destaque/destaque.component';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {
  tituloDocumento = "O documento mais testado foi"
  tituloEmpresa = "A empresa que mais testou foi"
  documentoMes!: string 
  documentoSemana!: string
  empresaMes!: string
  empresaSemana!: string

  ngOnInit() {
    this.pegarInfo()
  }

  pegarInfo() {
    this.documentoMes = "FATURAMENTO"
    this.documentoSemana = "CNH"

    this.empresaMes = "Inoa"
    this.empresaSemana = "Augusto"
  }
}
