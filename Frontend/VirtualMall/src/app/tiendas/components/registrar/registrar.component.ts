import { Component, OnInit } from '@angular/core';
import { Cuenta } from '../../interfaces/tienda.interface';
import { TiendaService } from '../../services/tienda.service';

@Component({
  selector: 'app-registrar',
  templateUrl: './registrar.component.html',
  styleUrls: ['./registrar.component.css']
})
export class RegistrarComponent{

  password: string = ""
  DPI: number = 0
  nombre: string = ""
  correo: string = ""

  constructor(private service : TiendaService) { }

  enviar(){
    let nuevo : Cuenta = {
      Dpi:      this.DPI,
      Nombre:   this.nombre,
      Correo:   this.correo,
      Password: this.password,
      Usuario:  "Usuario",
    }
    this.service.crearUsuario(nuevo).subscribe(
      ()=> {
        
      }
    );
  }
}
