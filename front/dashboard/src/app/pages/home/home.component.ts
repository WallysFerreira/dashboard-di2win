import { DatePipe } from '@angular/common';
import { Component } from '@angular/core';
import { getCount, getUsers } from 'src/app/app.component';
import { DestaqueComponent } from 'src/app/components/destaque/destaque.component';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {
  tituloDocumento = "O documento mais testado foi"
  tituloEmpresa = "A empresa que mais testou foi"
  tituloQntTotalPaginasTestadas = "Total de paginas testadas"
  documentoMes!: string 
  documentoSemana!: string
  empresaMes!: string
  empresaSemana!: string
  qntTotalPaginasTestadas: number = 0
  userChartData: any
  docChartData: any

  ngOnInit() {
    this.pegarInfoSemana()
    this.pegarInfoMes()
    this.pegarInfoGeral()
  }

  pegarInfoGeral() {
    this.pegarDadosUsuarios()
    this.pegarDadosDocumentos()
    this.pegarTotalPaginasTestadas()
  }

  async pegarTotalPaginasTestadas() {
    let res = await getCount("EXTRACT(month FROM created_at::date)", false, "0", null, null, null).then((res) => res.data.count)

    for (let month of res) {
      this.qntTotalPaginasTestadas += month.value
    }
  }

  pegarInfoMes() {
    this.documentoMes = "FATURAMENTO"
    this.empresaMes = "Inoa"
  }

  pegarInfoSemana() {
    this.documentoSemana = "CNH"
    this.empresaSemana = "Augusto"
  }

  async pegarDadosDocumentos() {
    let availableDocs = ["FATURAMENTO", "CNH", "POSICAO_CONSOLIDADA", "FATURA_ENERGIA", "CONTRATO_SOCIAL", "DECLARACAO_IR", "CAPA_SERASA", "ENDIVIDAMENTO", "COMPROVANTE_RESIDENCIA", "BALANCO_PATRIMONIAL", "RECIBOS"]
    let res = await getCount("EXTRACT(month FROM created_at::date)", false, "0", null, null, null)
    let labels = []
    let datasets = []

    for (let count of res.data.count) {
      let month = new Date(2000, count.name - 1, 1).toLocaleString('default', { month: 'long' })
      labels.push(month)
    }

    for (let doc of availableDocs) {
      let res = await getCount("EXTRACT(month FROM created_at::date)", false, "0", doc, null, null)
      let docData = []

      for (let count of res.data.count) {
        docData.push(count.value)
      }
      
      datasets.push({
        label: doc,
        data: docData
      })
    }

    this.docChartData = {
      labels: labels,
      datasets: datasets
    }
  }

  async pegarDadosUsuarios() {
    let availableUsers = await getUsers().then((res) => res.data.user)
    let res = await getCount("EXTRACT(month FROM created_at::date)", false, "0", null, null, null)
    let labels = []
    let datasets = []
    
    for (let count of res.data.count) {
      let month = new Date(2000, count.name - 1, 1).toLocaleString('default', { month: 'long' })
      labels.push(month)
    }

    for (let user of availableUsers) {
      let res = await getCount("EXTRACT(month FROM created_at::date)", false, user.id, null, null, null)
      let userData = []

      for (let count of res.data.count) {
        userData.push(count.value)
      } 

      datasets.push({
        label: user.name,
        data: userData
      })
    }

    this.userChartData = {
      labels: labels,
      datasets: datasets
    }
  }
}
