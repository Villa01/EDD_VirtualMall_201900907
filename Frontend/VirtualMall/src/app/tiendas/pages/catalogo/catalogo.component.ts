import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { Tienda } from '../../interfaces/tienda.interface';
import { TiendaService } from '../../services/tienda.service';

@Component({
  selector: 'app-catalogo',
  templateUrl: './catalogo.component.html',
  styleUrls: ['./catalogo.component.css']
})
export class CatalogoComponent{
  tiendas : Tienda[] = []

  constructor( private tiendasService: TiendaService) {

    this.tiendasService.obtenerTiendas().subscribe(tiendas => {
      console.log(tiendas)
      this.tiendas = tiendas;
    }, err => {
      console.log("No se han cargado las tiendas aún. ")
      this.tiendas = [];
    });
  }



}
