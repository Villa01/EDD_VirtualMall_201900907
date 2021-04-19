import { Component } from '@angular/core';
import { TiendaService } from '../../services/tienda.service';

@Component({
  selector: 'app-carga-tiendas',
  templateUrl: './carga-tiendas.component.html',
  styleUrls: ['./carga-tiendas.component.css']
})
export class CargaTiendasComponent  {

  termino :string = ""

  opcion : string[] = ['tienda', 'inventario', 'pedidos', 'usuarios', 'ubicaciones']

  constructor(private tiendaService : TiendaService) { }

  enviar(){
    this.tiendaService.cargarTiendas(this.termino).subscribe(
      () => {
        
      }, err => {
        console.log("No se pudo cargar la información")
      }
    );
  }

}
