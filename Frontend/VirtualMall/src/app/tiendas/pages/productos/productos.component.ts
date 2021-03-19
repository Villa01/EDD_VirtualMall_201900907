import { Component } from '@angular/core';
import { Producto } from '../../interfaces/tienda.interface';
import { TiendaService } from '../../services/tienda.service';

@Component({
  selector: 'app-productos',
  templateUrl: './productos.component.html',
  styleUrls: ['./productos.component.css']
})
export class ProductosComponent {

  productos : Producto[] = []

  constructor(private tiendasService : TiendaService) { 

    tiendasService.obtenerProductos().subscribe(
      products => {
        console.log(products)
        this.productos = products
      }, 
      err => {
        console.log(err)
      }
    );
  }


}
