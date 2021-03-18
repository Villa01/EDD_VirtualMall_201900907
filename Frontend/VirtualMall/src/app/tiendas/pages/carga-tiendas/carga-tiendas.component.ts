import { Component } from '@angular/core';
import { TiendaService } from '../../services/tienda.service';

@Component({
  selector: 'app-carga-tiendas',
  templateUrl: './carga-tiendas.component.html',
  styleUrls: ['./carga-tiendas.component.css']
})
export class CargaTiendasComponent  {

  termino :string = ""

  opcion : string[] = ['tienda', 'inventario']

  constructor(private tiendaService : TiendaService) { }

  enviar(){
    console.log(this.termino);
    this.tiendaService.cargarTiendas(this.termino).subscribe(
      () => {
        console.log(this.termino)
        
      }, err => {
        console.log("No se pudo cargar la información")
      }
    );
  }

}
