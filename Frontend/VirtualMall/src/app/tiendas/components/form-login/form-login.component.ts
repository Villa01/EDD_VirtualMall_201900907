import { Component, Output } from '@angular/core';
import { TiendaService } from '../../services/tienda.service';
import { Cuenta } from '../../interfaces/tienda.interface';
import { Router } from '@angular/router';

@Component({
  selector: 'app-form-login',
  templateUrl: './form-login.component.html',
  styleUrls: [
    './form-login.component.css'
  ]
})
export class FormLoginComponent{
  @Output() acceso: boolean =  false
  @Output() cuenta:Cuenta = {
    Dpi:      0,
    Nombre:   "",
    Correo:   "",
    Password: "",
    Usuario:  "",
  }
  DPI : number = 0
  password : string = ""

  constructor(private service : TiendaService, private router: Router) { }

  enviar(){
    let info = {
      DPI: this.DPI,
      password: this.password
    }
    this.service.verificarPassword(info).subscribe(
      acceso => {

        console.log("La contaseña es ", acceso.correcta)
        this.acceso = acceso.correcta
        this.cuenta = acceso.cuenta

        if (this.acceso) {
          this.router.navigate(['home'])
        } else {
          
          alert("Las credenciales son incorrectas")
        }
        
      }, 
      err => {
        console.log(err)
      }
    );
  }

}
