import { Component, Input } from '@angular/core';
import { Producto } from '../../interfaces/tienda.interface';
import { TiendaService } from '../../services/tienda.service';

@Component({
  selector: 'app-producto',
  templateUrl: './producto.component.html',
  styleUrls: ['./producto.component.css']
})
export class ProductoComponent {

  @Input() producto : Producto = {
    Nombre: '',
    Descripcion: '',
    Codigo : -1,
    Precio : -1, 
    Cantidad: -1,
    Imagen : ''

  }
  stock : number  = 0
  constructor(private tiendasService: TiendaService) { 

  }

  agregarAlCarrito(){
    const temp: Producto = this.producto
    temp.Cantidad = this.stock

    let jsonProducto = JSON.stringify(temp)
    this.tiendasService.agregarAlCarrito(jsonProducto).subscribe(
      err => {
        console.log(err)
      }
    );
  }

  sumar(num : number){
    if (this.stock <= this.producto.Cantidad && this.stock >=0) { 
      this.stock++
    }
  }

  restar(num : number){
    if (this.stock <= this.producto.Cantidad && this.stock >0) {
      this.stock--
    }
  }

}
