import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-destaque',
  templateUrl: './destaque.component.html',
  styleUrls: ['./destaque.component.css']
})


export class DestaqueComponent {
  @Input() titulo!: string 
  @Input() destaque!: string
}
