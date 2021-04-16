import { Component, OnInit, Output } from '@angular/core';
import { Cuenta } from '../../interfaces/tienda.interface';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  @Output()acceso: boolean =  false
  @Output()cuenta:Cuenta = {
    Dpi:      0,
    Nombre:   "",
    Correo:   "",
    Password: "",
    Usuario:  "",
  }

  constructor() { }

  ngOnInit(): void {
  }

}
