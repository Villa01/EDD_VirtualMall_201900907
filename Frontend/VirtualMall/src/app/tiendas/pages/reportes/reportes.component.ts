import { Component, OnInit } from '@angular/core';
import { TiendaService } from '../../services/tienda.service';
import { Reporte } from '../../interfaces/tienda.interface';

@Component({
  selector: 'app-reportes',
  templateUrl: './reportes.component.html',
  styleUrls: ['./reportes.component.css']
})
export class ReportesComponent implements OnInit {

  password : string = ""
  reportes : Reporte[] = []

  constructor(private servicio : TiendaService) { }

  ngOnInit(): void {
  }

  enviar(){
    this.servicio.obtenerReportes(this.password).subscribe(
      resp => {
        this.reportes = resp.Reportes
        console.log(this.reportes)
      }
    );
  }
}
