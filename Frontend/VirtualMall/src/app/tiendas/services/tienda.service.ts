import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { Observable } from 'rxjs';

import { Tienda, Producto, infoUsuario, RespuestaPassword, Cuenta, eliminarResponse, booleanResponse, ReportesResponse } from '../interfaces/tienda.interface';
import { flushMicrotasks } from '@angular/core/testing';

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

  cargarUsuarios(termino: string) {
    const url = `${this.apiUrl}/cargaUsuarios`
    return this.http.post<RespuestaPassword>(url, termino )
  }

  crearUsuario(nuevo : Cuenta){
    const url = `${this.apiUrl}/crearUsuario`
    let usuario = JSON.stringify(nuevo)
    return this.http.post(url, usuario)
  }

  eliminarUsuario(password : string){
    const url = `${this.apiUrl}/eliminarUsuario`
    let peticion = {
      password : password
    }
    let texto = JSON.stringify(peticion)
    console.log(texto)
    return this.http.post<booleanResponse>(url, peticion)
  }

  obtenerReportes(llave : string){
    const url = `${this.apiUrl}/obtenerReportes`
    let peticion = {
      "texto" : llave
    }
    return this.http.post<ReportesResponse>(url, peticion)
  }
}
