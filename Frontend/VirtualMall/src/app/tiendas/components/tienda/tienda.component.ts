import { Component, Input, OnInit } from '@angular/core';
import { Tienda } from '../../interfaces/tienda.interface';

@Component({
  selector: 'app-tienda',
  templateUrl: './tienda.component.html',
  styleUrls: ['./tienda.component.css']
})
export class TiendaComponent implements OnInit {
  @Input() tienda : Tienda = {
    Nombre:'',
    Descripcion: '',
    Contacto: '',
    Calificacion: 0,
    Logo : '',

  };
  constructor() { }

  ngOnInit(): void {
  }

}
