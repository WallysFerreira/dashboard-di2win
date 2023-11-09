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

    res = await getCount("doc_type", 0, null, null, null)
    this.tipos_documento = res.data.count
    console.log("Documentos", this.tipos_documento)
  }
}
