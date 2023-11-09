import { Component } from '@angular/core';
import { getUsers } from 'src/app/app.component';

@Component({
  selector: 'app-filtro',
  templateUrl: './filtro.component.html',
  styleUrls: ['./filtro.component.css']
})

export class FiltroComponent {
  empresas: any

  async ngOnInit() {
    let res = await getUsers()
    this.empresas = res.data.user
    console.log(this.empresas)
  }
}
