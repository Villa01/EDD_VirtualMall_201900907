import { Component, Output, EventEmitter, Input } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { TiendaService } from '../../services/tienda.service';

@Component({
  selector: 'app-input-json',
  templateUrl: './input-json.component.html',
  styleUrls: []
})
export class InputJSONComponent {
  @Input() opcion : string = ""
  @Output() onEnter : EventEmitter<string> = new EventEmitter();

  termino : string = ""
  error : boolean = false

  constructor(private tiendasService : TiendaService) { }
  // Obtiene el contenido del archivo JSON
  onFileChange(event: any){

      let files = event.target.files;
    
      var lector = new FileReader();
      lector.readAsText(files[0]);
      lector.onload = () => {
        let texto:any = lector.result;
        if(texto){
          this.termino = texto;
        }

      }
    }
    //Envia el archivo al padre
    enviar(){
      
      switch (this.opcion){
        case 'tienda':
          this.tiendasService.cargarTiendas(this.termino).subscribe(
            resp =>{
              //console.log(resp.toString);
            }, err => {
              this.error = true
            }
          );
          break
        case 'inventario':
          this.tiendasService.cargarInventarios(this.termino).subscribe(
            resp =>{
              //console.log(resp.toString);
            }, err => {
              this.error = true
            }
          );
          break
        
          case 'pedidos':
            this.tiendasService.cargarPedidos(this.termino).subscribe(
              resp =>{
                //console.log(resp.toString);
              }, err => {
                this.error = true
                
              }
            );
            break
      }
      //console.log(this.termino)

      

      //this.onEnter.emit(this.termino);

    }

}
