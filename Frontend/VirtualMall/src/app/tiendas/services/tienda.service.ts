import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class TiendaService {

  private apiUrl : string = "http://localhost:3000";

  constructor(private http: HttpClient) { }

  cargarTiendas(termino:string){
    const url = `${this.apiUrl}/cargarTienda`;
    return this.http.post(url, termino)
  }
}
