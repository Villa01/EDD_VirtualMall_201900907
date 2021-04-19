import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { TiendaService } from '../../services/tienda.service';
import { Reporte } from '../../interfaces/tienda.interface';

@Component({
  selector: 'app-formulario-reporte',
  templateUrl: './formulario-reporte.component.html',
  styleUrls: ['./formulario-reporte.component.css']
})
export class FormularioReporteComponent implements OnInit {
  
  password : string = ""
  reportes : Reporte[] = []
  error : boolean = false
  success : boolean = false
  mensErro : string = "No se pudieron generar los reportes"
  mensSucc : string = "Reportes generados correctame"
  @Output() evento = new EventEmitter<Reporte[]>();


  constructor(private servicio: TiendaService) { }

  ngOnInit(): void {
    this.servicio.obtenerReportes().subscribe(
      resp => {
        this.reportes = resp
        this.evento.emit(this.reportes)
        
        this.error = false
        this.success = !this.error
      },
      err => {
        this.error = true
        this.success = !this.error
      }
    );
  }

  enviar(){
    
    this.servicio.generarReportes(this.password).subscribe(
      resp => {
        this.error = !resp.booleano
        this.success = !this.error
      },
      err => {
        this.error = true
        this.success = !this.error
      }
    );
  }
}
