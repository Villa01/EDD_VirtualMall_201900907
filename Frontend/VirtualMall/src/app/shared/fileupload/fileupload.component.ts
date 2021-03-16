import { Component, EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-fileupload',
  templateUrl: './fileupload.component.html',
  styleUrls: []
})
export class FileuploadComponent {

  
  @Output() onEnter : EventEmitter<string> = new EventEmitter();

  termino : string = ""

  constructor() { }

  buscar() {
    console.log(this.termino)
    this.onEnter.emit( this.termino );
  }

}
