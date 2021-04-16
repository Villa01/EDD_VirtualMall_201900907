import { Component, OnInit } from '@angular/core';
import { Cuenta } from './tiendas/interfaces/tienda.interface';
import { TiendaService } from './tiendas/services/tienda.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{
  acceso : boolean = false
  login : boolean = false
  cuenta:Cuenta = {
    Dpi:      0,
    Nombre:   "",
    Correo:   "",
    Password: "",
    Usuario:  "",
  }

  ngOnInit(): void {
    console.log(this.cuenta)
  }

  constructor(private servicio: TiendaService) {
    this.servicio.obtenerCuenta().subscribe(

      cuenta => {
        
        this.cuenta = cuenta

        if (this.cuenta.Dpi == 0) {
          
          this.login = false
        } else {
          
          this.login = true
        }
        if (cuenta.Usuario == "Admin") {
          this.acceso = true
        } else {
          this.acceso = false
        }
      }
    );
  }

  title = 'VirtualMall';
}
