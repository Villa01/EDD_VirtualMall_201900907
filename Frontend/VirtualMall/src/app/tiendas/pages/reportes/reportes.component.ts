import { Component, OnInit } from '@angular/core';
import { TiendaService } from '../../services/tienda.service';
import { Reporte } from '../../interfaces/tienda.interface';

@Component({
  selector: 'app-reportes',
  templateUrl: './reportes.component.html',
  styleUrls: ['./reportes.component.css']
})
export class ReportesComponent {

  reportes : Reporte[] = []
  
  constructor() {
  }

  imprimirReportes(reportes : Reporte[]){
    this.reportes = reportes
    console.log(this.reportes)
  }
  
}
