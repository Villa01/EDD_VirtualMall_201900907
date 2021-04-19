import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-mostrar-error',
  templateUrl: './mostrar-error.component.html',
  styleUrls: ['./mostrar-error.component.css']
})
export class MostrarErrorComponent implements OnInit {
  @Input() error : boolean = false
  @Input() success : boolean = false
  @Input() mensajeErr : string = ""
  @Input() mensajeSucc : string = ""
  constructor() { }

  ngOnInit(): void {
  }

}
