import { Component, Input, OnInit } from '@angular/core';
import { Reporte } from '../../interfaces/tienda.interface';

@Component({
  selector: 'app-reporte',
  templateUrl: './reporte.component.html',
  styles: [
  ]
})
export class ReporteComponent implements OnInit {

  @Input() reporte : Reporte = {
    nombre : "",
    ruta : ""
  }

  constructor() { }

  ngOnInit(): void {
  }

}
