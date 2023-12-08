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

  buttons: any
  hiddenAttributeElement: any

  /*
  select(event: any) {
    this.buttons = document.getElementById(this.filterTitle + 'OptionsDiv')?.getElementsByClassName('nonDateOptions')[0].getElementsByTagName('button')

    const hasClass = event.target.classList.contains('selected')

    if (hasClass) {
      this.unselect(event)
    } else {
      this.setSelected(event)
    }
  }

  unselect (event: any) {
    event.target.classList.remove('selected')
  }

  setSelected(event: any) {
    this.clearSelectedButton()

    event.target.classList.add('selected')
  }

  clearSelectedButton() {
    for (let i = 0; i < this.buttons.length; i++) {
      this.buttons[i].classList.remove('selected')
    }
  }

  /*
  toggleOptionButtons(e: any) {
    let parentId = e.target.parentElement.id
    let optionsDiv = document.getElementById(parentId)?.getElementsByClassName('optionsDiv')[0]
    let attributeHidden = optionsDiv?.attributes.getNamedItem('hidden')

    console.log(attributeHidden)

    if (attributeHidden != null) {
      console.log("mostrando")
      this.showOptionButtons(e)
    }
    
    if (attributeHidden == null) {
      console.log("escondendo")
      this.hideOptionsButtons(e)
    }
  }

  showOptionButtons(e: any) {
    let parentId = e.target.parentElement.id
    let optionsDiv = document.getElementById(parentId)?.getElementsByClassName('optionsDiv')[0]
    
    this.hiddenAttributeElement = optionsDiv?.attributes.removeNamedItem('hidden')
  } 

  hideOptionsButtons(e: any) {
    let optionsDiv

    if (e.target.className == 'titleButton') {
      optionsDiv = e.target.parentElement.children[1]
    } else {
      optionsDiv = e.target.parentElement
    }

    optionsDiv.attributes.setNamedItem(this.hiddenAttributeElement)
  }
  */
}
