import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-filtro',
  templateUrl: './filtro.component.html',
  styleUrls: ['./filtro.component.css']
})

export class FiltroComponent {
  @Input() filterTitle?: string
  @Input() filterType?: string
  @Input() filterData: any
}
