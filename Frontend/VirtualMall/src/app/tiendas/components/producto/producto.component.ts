import { Component, Input } from '@angular/core';
import { Producto, Departamento } from '../../interfaces/tienda.interface';
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

  constructor(private tiendasService: TiendaService) { }

  agregarAlCarrito(){
    console.log(`Agregando ${this.producto.Nombre}` )
    let jsonProducto = JSON.stringify(this.producto)
    this.tiendasService.agregarAlCarrito(jsonProducto).subscribe(
      err => {
        console.log(err)
      }
    );
  }

}
