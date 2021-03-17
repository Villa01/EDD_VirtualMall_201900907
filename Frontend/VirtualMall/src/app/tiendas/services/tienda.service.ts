import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { Observable } from 'rxjs';

import { Tienda } from "../interfaces/tienda.interface";

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

  obtenerTiendas(): Observable<Tienda[]>{
    const url= `${this.apiUrl}/getTiendas`;
    return this.http.get<Tienda[]>(url);
  }
}
