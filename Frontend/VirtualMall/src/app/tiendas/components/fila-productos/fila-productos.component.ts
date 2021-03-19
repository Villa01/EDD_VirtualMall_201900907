import { Component, Input } from '@angular/core';
import { Producto } from '../../interfaces/tienda.interface';

@Component({
  selector: 'app-fila-productos',
  templateUrl: './fila-productos.component.html',
  styleUrls: ['./fila-productos.component.css']
})
export class FilaProductosComponent {
  @Input() producto: Producto = {
    Nombre: '',
    Descripcion: '',
    Codigo : -1,
    Precio : -1, 
    Cantidad: -1,
    Imagen : ''

  }
  constructor() { }

  sumar(){
    this.producto.Cantidad++
  }

  restar(){
    
    this.producto.Cantidad--
  }

  eliminar(){

  }

}
