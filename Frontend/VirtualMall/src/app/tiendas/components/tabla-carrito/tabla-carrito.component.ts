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

  sumar(num : number){
    console.log(num)
    let indice = this.obtenerProducto(num)
    console.log(`El indice es ${indice}`)
    this.productos[indice].Cantidad++
  }

  restar(num : number){
    let indice = this.obtenerProducto(num)
    this.productos[indice].Cantidad--
  }

  eliminar(codigo : number){
    let indice = this.obtenerProducto(codigo)
    let producto = JSON.stringify(this.productos[indice])
    console.log(producto)
    this.tiendasService.eliminarDelCarrito(producto).subscribe(
      err => {
        console.log(err)
      }
    );

    this.tiendasService.obtenerCarrito().subscribe(
      product => {
        this.productos = product
      }, 
      err => {
        console.log(err)
      }
    );
  }

  obtenerProducto(codigo: number) :number{
    
    let numero = -1;
    this.productos.forEach(
      (product, i) => {
        if (product.Codigo === codigo){
            numero = i
        } 
      }
    );
      
    return numero
  }

}
