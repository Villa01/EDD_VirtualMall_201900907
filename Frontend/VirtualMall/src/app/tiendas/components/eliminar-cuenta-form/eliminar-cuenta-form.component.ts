import { Component, OnInit } from '@angular/core';
import { TiendaService } from '../../services/tienda.service';

@Component({
  selector: 'app-eliminar-cuenta-form',
  templateUrl: './eliminar-cuenta-form.component.html',
  styleUrls: ['./eliminar-cuenta-form.component.css']
})
export class EliminarCuentaFormComponent implements OnInit {

  password : string = ""
  error : boolean = false

  constructor(private servicio : TiendaService) { }

  ngOnInit(): void {
  }

  enviar(){
    this.servicio.eliminarUsuario(this.password).subscribe(
      (resp)=>{
        this.error = resp.eliminada
      }
    );
    };
  }

