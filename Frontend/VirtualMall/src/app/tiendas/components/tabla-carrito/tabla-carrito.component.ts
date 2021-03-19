import { Component } from '@angular/core';
import { Producto } from '../../interfaces/tienda.interface';
import { TiendaService } from '../../services/tienda.service';

@Component({
  selector: 'app-tabla-carrito',
  templateUrl: './tabla-carrito.component.html',
  styleUrls: ['./tabla-carrito.component.css']
})
export class TablaCarritoComponent  {

  productos : Producto[] = []

  constructor(private tiendasService : TiendaService) {

    this.tiendasService.obtenerCarrito().subscribe(
      product => {
        this.productos = product
      }, 
      err => {
        console.log(err)
      }
    );
  }

}
