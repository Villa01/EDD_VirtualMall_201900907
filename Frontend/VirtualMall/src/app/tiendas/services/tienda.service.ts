import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { Observable } from 'rxjs';

import { Tienda, Producto, infoUsuario, RespuestaPassword, Cuenta } from '../interfaces/tienda.interface';

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

  cargarInventarios(termino:string){
    const url = `${this.apiUrl}/cargarInventarios`;
    return this.http.post(url, termino)
  }

  cargarPedidos(termino:string){
    const url = `${this.apiUrl}/cargarPedidos`;
    return this.http.post(url, termino)
  }

  obtenerTiendas(): Observable<Tienda[]>{
    const url= `${this.apiUrl}/getTiendas`;
    return this.http.get<Tienda[]>(url);
  }

  obtenerProductos(): Observable<Producto[]>{
    const url= `${this.apiUrl}/getProductos`;
    return this.http.get<Producto[]>(url);
  }

  agregarAlCarrito(producto: string){
    const url = `${this.apiUrl}/agregarAlCarrito`;
    return this.http.post(url, producto)
  }

  obtenerCarrito():  Observable<Producto[]>{
    const url= `${this.apiUrl}/getCarrito`;
    return this.http.get<Producto[]>(url);
  }

  eliminarDelCarrito(producto: string){
    const url= `${this.apiUrl}/eliminarDelCarrito`;
    return this.http.post(url, producto)
  }
  
  hacerPedido(pedido : string){
    const url = `${this.apiUrl}/hacerPedido`;
    return this.http.post(url, pedido)
  }

  verificarPassword(info : infoUsuario) {
    const url = `${this.apiUrl}/verificacionLogIn`
    let texto = JSON.stringify(info)
    return this.http.post<RespuestaPassword>(url, info )
  }

  obtenerCuenta(){
    const url =  `${this.apiUrl}/cuentaActual`
    return this.http.get<Cuenta>(url)
  }
}
